package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/mikedonnici/dev/grpc/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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

	c, err := s.collection("posts")
	if err != nil {
		return nil, err
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
	data.ID = oid

	return &blogpb.CreatePostResponse{
		Post: blogpbPost(data),
	}, nil
}

func (s *service) ReadPost(ctx context.Context, req *blogpb.ReadPostRequest) (*blogpb.ReadPostResponse, error) {

	log.Println("Reading post...")

	c, err := s.collection("posts")
	if err != nil {
		return nil, err
	}

	oid, err := primitive.ObjectIDFromHex(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf(".ObjectIDFromHex() err = %s", err))
	}

	f := bson.M{"_id": oid}
	r := c.FindOne(ctx, f)
	if r.Err() == mongo.ErrNoDocuments {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("no post found with id %s", oid))
	}
	if err != nil {
		return nil, fmt.Errorf(".FindOne() err = %w", err)
	}

	data := post{}
	if err := r.Decode(&data); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Decode() err = %s", err))
	}

	return &blogpb.ReadPostResponse{
		Post: blogpbPost(data),
	}, nil
}

func (s *service) UpdatePost(ctx context.Context, req *blogpb.UpdatePostRequest) (*blogpb.UpdatePostResponse, error) {

	log.Println("Updating post...")

	c, err := s.collection("posts")
	if err != nil {
		return nil, err
	}

	p := req.GetPost()
	oid, err := primitive.ObjectIDFromHex(p.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf(".ObjectIDFromHex() err = %s", err))
	}

	filter := bson.M{"_id": oid}
	data := post{
		ID:       oid,
		AuthorID: p.GetAuthorId(),
		Title:    p.GetTitle(),
		Content:  p.GetContent(),
	}

	_, err = c.ReplaceOne(ctx, filter, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("UpdateOne() err = %s", err))
	}

	return &blogpb.UpdatePostResponse{
		Post: blogpbPost(data),
	}, nil
}

func (s *service) DeletePost(ctx context.Context, req *blogpb.DeletePostRequest) (*blogpb.DeletePostResponse, error) {

	log.Println("Deleting post...")

	c, err := s.collection("posts")
	if err != nil {
		return nil, err
	}

	oid, err := primitive.ObjectIDFromHex(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("could not convert hex to object id, err = %s", err))
	}

	// ensure exists before deleting
	_, err = s.ReadPost(ctx, &blogpb.ReadPostRequest{
		PostId: oid.Hex(),
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("could not find post with id %q to delete", oid.Hex()))
	}

	filter := bson.M{"_id": oid}

	_, err = c.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("DeleteOne() err = %s", err))
	}

	return &blogpb.DeletePostResponse{
		PostId: oid.Hex(),
	}, nil
}

func (s *service) ListPost(req *blogpb.ListPostRequest, stream blogpb.BlogService_ListPostServer) error {

	log.Println("List posts...")

	c, err := s.collection("posts")
	if err != nil {
		return err
	}

	filter := bson.M{} // empty
	crsr, err := c.Find(context.Background(), filter)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Find() err = %s", err))
	}
	defer crsr.Close(context.Background())

	for crsr.Next(context.Background()) {

		data := post{}
		if err := crsr.Decode(&data); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Decode() err = %s", err))
		}

		stream.Send(&blogpb.ListPostResponse{
			Post: blogpbPost(data),
		})
	}

	// returns last error seen by cursor... ie if it leaps out of the .Next loop for some reason (?)
	if err := crsr.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("cursor.Next() err = %s", err))
	}

	return nil
}

// collection returns a pointer to a mongo collection associated with the service, or an an error
func (s *service) collection(name string) (*mongo.Collection, error) {
	c, ok := s.collections["posts"]
	if !ok {
		return nil, fmt.Errorf(errMissingCollection, "posts")
	}
	return c, nil
}

// blogpbPost is a convenience func that returns a pointer to a blogpb.Post from the provided post value
func blogpbPost(data post) *blogpb.Post {
	return &blogpb.Post{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}
