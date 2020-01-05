package main

import (
	"context"
	"fmt"
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
	//create(client)
	id := "5e1114d999b4f9b9c48dfbae"
	read(client, id)

	updatedPost := blogpb.Post{
		Id:       id,
		AuthorId: "Michael Peter Donnici",
		Title:    "The new way to handle things... with care and tests!",
		Content:  "The complexity of arranging code and corresponding tests is a constant battle between...",
	}
	update(client, updatedPost)
}

func create(client blogpb.BlogServiceClient) {
	post := blogpb.CreatePostRequest{
		Post: &blogpb.Post{
			AuthorId: "mike",
			Title:    "The way I like it...",
			Content:  "One day I was walking on the beach when I had a thought.",
		},
	}
	res, err := createBlogPost(client, post)
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
	fmt.Println(res.Post)
}

// read handles reading a post ... printing here but would manage the response from here based on
// status code / error types etc.
func read(client blogpb.BlogServiceClient, postID string) {

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
	fmt.Println(res)
}

func update(client blogpb.BlogServiceClient, updatedPost blogpb.Post) {

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
	fmt.Println(res)
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
