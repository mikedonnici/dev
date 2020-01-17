package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/mikedonnici/dev/grpc/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

const (
	useTLS         = true
	certFile       = "ssl/ca.crt"
	timeoutSeconds = 5
)

func main() {

	var options = grpc.WithInsecure()
	if useTLS {
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("credentials.NewClientTLSFromFile() err = %s", err)
		}
		options = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalf("grpc.Dial() err = %s", err)
	}
	defer conn.Close()

	client := blogpb.NewBlogServiceClient(conn)

	// test runs
	newPost1 := blogpb.Post{
		AuthorId: "Leo Anthony Donnici",
		Title:    "How to build a transformer with Lego.",
		Content:  "Lorem ipsum....",
	}
	newPost2 := blogpb.Post{
		AuthorId: "Maia Grace Donnici",
		Title:    "On the benefits of Fairies and Spells.",
		Content:  "Lorem ipsum....",
	}

	// Create the posts
	p1 := create(client, &newPost1)
	log.Printf("Created post with id %s: %q", p1.GetId(), p1.GetTitle())
	p2 := create(client, &newPost2)
	log.Printf("Created post with id %s: %q", p2.GetId(), p2.GetTitle())

	// Read a post
	id := p1.Id
	p3 := read(client, id)
	log.Printf("Read post with id %s: %q", id, p3.GetTitle())

	// Update a post
	updatedPost := p3
	updatedPost.Title = "Lego transformers - advanced course"
	updatedPost.Content = "The complexity of arranging code and corresponding tests is a constant battle between..."
	p4 := update(client, updatedPost)
	log.Printf("Updated post with id %s", p3.Id)
	log.Printf("---> from: %v", p3)
	log.Printf("---> to: %v", p4)

	// Delete a post
	deletedId := delete(client, p4.Id)
	log.Printf("Deleted post with id %q", deletedId)

	// List the posts
	posts := list(client)
	for _, p := range posts {
		fmt.Println(p)
	}
}

func create(client blogpb.BlogServiceClient, post *blogpb.Post) blogpb.Post {

	req := blogpb.CreatePostRequest{
		Post: post,
	}

	res, err := createBlogPost(client, req)
	if err != nil {
		statusErr, exists := status.FromError(err)
		if exists {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalln("Deadline exceeded")
			}
			log.Fatalf("statusErr = %s", statusErr)
		}
		log.Fatalf("createBlogPost() err = %s", err)
	}

	return *res.Post
}

func read(client blogpb.BlogServiceClient, postID string) blogpb.Post {

	req := blogpb.ReadPostRequest{
		PostId: postID,
	}

	res, err := readBlogPost(client, req)
	if err != nil {
		statusErr, exists := status.FromError(err)
		if exists {
			if statusErr.Code() == codes.NotFound {
				log.Fatalf("Post with id %q not found, err = %s", postID, statusErr.Err())
			}
			log.Fatalf("readBlogPost() statusErr = %s", statusErr.Err())
		}
		log.Fatalf("readBlogPost() err = %s", err)
	}

	return *res.Post
}

func update(client blogpb.BlogServiceClient, updatedPost blogpb.Post) blogpb.Post {

	req := blogpb.UpdatePostRequest{
		Post: &updatedPost,
	}

	res, err := updateBlogPost(client, req)
	if err != nil {
		statusErr, exists := status.FromError(err)
		if exists {
			log.Fatalf("updateBlogPost() statusErr = %s", statusErr.Err())
		}
		log.Fatalf("updateBlogPost() err = %s", err)
	}

	return *res.Post
}

func delete (client blogpb.BlogServiceClient, postID string) string {

	req := blogpb.DeletePostRequest{
		PostId:               postID,
	}

	res, err := deleteBlogPost(client, req)
	if err != nil {
		statusErr, exists := status.FromError(err)
		if exists {
			log.Fatalf("deleteBlogPost() statusErr = %s", statusErr.Err())
		}
		log.Fatalf("deleteBlogPost() err = %s", err)
	}

	return res.PostId
}

func list(client blogpb.BlogServiceClient) []*blogpb.ListPostResponse {

	req := blogpb.ListPostRequest{}

	posts, err := listBlogPosts(client, req)
	if err != nil {
		statusErr, exists := status.FromError(err)
		if exists {
			log.Fatalf("listBlogPosts() statusErr = %s", statusErr.Err())
		}
		log.Fatalf("listBlogPosts() err = %s", err)
	}

	return posts
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {

}

func readPostHandler(w http.ResponseWriter, r *http.Request) {

}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {

}

func createBlogPost(client blogpb.BlogServiceClient, req blogpb.CreatePostRequest) (*blogpb.CreatePostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()
	return client.CreatePost(ctx, &req)
}

func readBlogPost(client blogpb.BlogServiceClient, req blogpb.ReadPostRequest) (*blogpb.ReadPostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()
	return client.ReadPost(ctx, &req)
}

func updateBlogPost(client blogpb.BlogServiceClient, req blogpb.UpdatePostRequest) (*blogpb.UpdatePostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()
	return client.UpdatePost(ctx, &req)
}

func deleteBlogPost(client blogpb.BlogServiceClient, req blogpb.DeletePostRequest) (*blogpb.DeletePostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()
	return client.DeletePost(ctx, &req)
}

func listBlogPosts(client blogpb.BlogServiceClient, req blogpb.ListPostRequest) ([]*blogpb.ListPostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()

	var posts []*blogpb.ListPostResponse
	stream, err := client.ListPost(ctx, &req)
	if err != nil {
		return nil, err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("resStream.Recv() err = %s", err)
		}
		posts = append(posts, res)
	}

	return posts, nil
}
