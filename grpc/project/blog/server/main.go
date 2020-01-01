package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/mkedonnici/grpc/project/blog/blogpb"
)

const (
	useTLS   = true
	certFile = "ssl/server.crt"
	keyFile  = "ssl/server.pem"
	mongoDSN = "mongodb://localhost:27017"
)

const (
	errMissingCollection = "the %s collection does not seem to be attached"
)

type service struct {
	collections map[string]*mongo.Collection
}

type post struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	AuthorID string             `json:"authorId" bson:"author_id"`
	Title    string             `json:"title" bson:"title"`
	Content  string             `json:"content" bson:"content"`
}

func main() {

	// adds some addition guff to log lines
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// create the service and attach mongo collection(s)
	srvc := NewService()
	if err := srvc.AttachMongoCollection(mongoDSN, "blog", "posts"); err != nil {
		log.Fatalf(".AttachMongoCollection() err = %s", err)
	}
	log.Println("Attached mongodb")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("net.Listen() err = %srvc", err)
	}

	var serverOptions []grpc.ServerOption
	if useTLS {
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("credentials.NewServerTLSFromFile() err = %s", err)
		}
		serverOptions = append(serverOptions, grpc.Creds(creds))
	}

	s := grpc.NewServer(serverOptions...)
	blogpb.RegisterBlogServiceServer(s, srvc)
	go func() {
		tls := "WITHOUT"
		if useTLS {
			tls = "WITH"
		}
		log.Printf("Starting service %s TLS\n", tls)
		if err := s.Serve(lis); err != nil {
			log.Fatalf(".Serve() err = %s", err)
		}
	}()

	// block until ctl+c
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("Stopping service...")
	s.Stop()
	log.Println("Done.")
}

func NewService() *service {
	cols := make(map[string]*mongo.Collection)
	return &service{
		collections: cols,
	}
}

func (s *service) AttachMongoCollection(mongoDSN string, dbName, collectionName string) error {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDSN))
	if err != nil {
		return fmt.Errorf("mongo.NewClient() err = %w", err)
	}
	log.Println("Created mongo client...")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		return fmt.Errorf("mongo.Client.Connect() err = %w", err)
	}
	log.Println("Created mongo connection...")

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("mongo.Client.Ping() err = %w", err)
	}
	log.Println("Successfully pinged mongo database...")

	// attach db collection keyed by collection name
	s.collections[collectionName] = client.Database(dbName).Collection(collectionName)
	log.Printf("Attached collection: %s", collectionName)

	return nil
}

func (s *service) CreatePost(ctx context.Context, req *blogpb.CreatePostRequest) (*blogpb.CreatePostResponse, error) {

	log.Println("Creating post...")

	c, ok := s.collections["posts"]
	if !ok {
		return nil, fmt.Errorf(errMissingCollection, "posts")
	}

	p := req.GetPost()
	data := post{
		AuthorID: p.GetAuthorId(),
		Title:    p.GetTitle(),
		Content:  p.GetContent(),
	}
	r, err := c.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("InsertOne() err = %s", err))
	}
	oid, ok := r.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "could not assert ObjectID")
	}

	return &blogpb.CreatePostResponse{
		Post: &blogpb.Post{
			Id:       oid.Hex(),
			AuthorId: p.GetAuthorId(),
			Title:    p.GetTitle(),
			Content:  p.GetContent(),
		},
	}, nil
}

func (s *service) ReadPost(ctx context.Context, req *blogpb.ReadPostRequest) (*blogpb.ReadPostResponse, error) {

	log.Println("Reading post...")

	c, ok := s.collections["posts"]
	if !ok {
		return nil, fmt.Errorf(errMissingCollection, "posts")
	}

	oid, err := primitive.ObjectIDFromHex(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf(".ObjectIDFromHex() err = %s", err))
	}

	// f := post{ID: oid} // tags don't convert this to BSON?
	f := bson.M{"_id": oid}
	r := c.FindOne(ctx, f)
	if r.Err() == mongo.ErrNoDocuments {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("no post found with id %s", oid))
	}
	if err != nil {
		return nil, fmt.Errorf(".FindOne() err = %w", err)
	}

	p := post{}
	if err := r.Decode(&p); err != nil {
		return nil, fmt.Errorf(".Decode() err = %w", err)
	}

	return &blogpb.ReadPostResponse{
		Post: &blogpb.Post{
			Id:       p.ID.Hex(),
			AuthorId: p.AuthorID,
			Title:    p.Title,
			Content:  p.Content,
		},
	}, nil
}

func (s *service) UpdatePost(ctx context.Context, req *blogpb.UpdatePostRequest) (*blogpb.UpdatePostResponse, error) {

	log.Println("Updating post...")

	c, ok := s.collections["posts"]
	if !ok {
		return nil, fmt.Errorf(errMissingCollection, "posts")
	}

	p := req.GetPost()
	oid, err := primitive.ObjectIDFromHex(p.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf(".ObjectIDFromHex() err = %s", err))
	}
	data := post{
		ID:       oid,
		AuthorID: p.GetAuthorId(),
		Title:    p.GetTitle(),
		Content:  p.GetContent(),
	}

	_, err = c.UpdateOne(ctx, bson.M{"_id": oid}, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("UpdateOne() err = %s", err))
	}

	return &blogpb.UpdatePostResponse{
		Post: &blogpb.Post{
			Id:       oid.Hex(),
			AuthorId: data.GetAuthorId(),
			Title:    data.GetTitle(),
			Content:  p.GetContent(),
		},
	}, nil
}
