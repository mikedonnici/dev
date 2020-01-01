package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"time"

	"github.com/mkedonnici/grpc/project/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	// create(client)
	read(client)
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

func read(client blogpb.BlogServiceClient) {

	const id = "5df8a709e7ad6cd6b1bda0a8"
	req := blogpb.ReadPostRequest{
		PostId: id,
	}

	res, err := readBlogPost(client, req)
	if err != nil {
		statusErr, exists := status.FromError(err)
		if exists {
			if statusErr.Code() == codes.NotFound {
				log.Fatalf("Post with id %s not found, err = %s", id, statusErr)
			}
			log.Fatalf("statusErr = %s", statusErr)
		}
		log.Fatalf(".readBlogPost() err = %s", err)
	}
	fmt.Println(res)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {

}

func readPostHandler(w http.ResponseWriter, r *http.Request) {

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
