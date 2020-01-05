package main

import (
	"context"
	"fmt"
	"github.com/mikedonnici/dev/grpc/greet/greetpb"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

// needs to implement the GreetServer interface
type server struct{}

const (
	useTLS = true
	certFile = "ssl/server.crt"
	keyFile = "ssl/server.pem"
)

func main() {

	tls := "WITHOUT"
	if useTLS {
		tls = "WITH"
	}
	fmt.Printf("Server is firing up %s TLS\n", tls)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("net.Listen() err = %s", err)
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
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf(".Serve() err = %s", err)
	}
}


func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greeting %s %s", req.Person.FirstName, req.Person.LastName)
	return &greetpb.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s %s", req.Person.FirstName, req.Person.LastName),
	}, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("Greeting many times %s %s", req.Person.FirstName, req.Person.LastName)
	for i := 0; i < 3; i++ {
		res := &greetpb.GreetResponse{
			Greeting: fmt.Sprintf("Hello, %s %s (%d)", req.Person.FirstName, req.Person.LastName, i),
		}
		stream.Send(res)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (*server) GreetAfterManyTimes(stream greetpb.GreetService_GreetAfterManyTimesServer) error {
	log.Println("Greeting *after* many times...")

	// Server will respond *after* the last client request, however it can respond at any time
	var reqs []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		reqs = append(reqs, req.String())
	}

	allResponse := &greetpb.GreetResponse{
		Greeting: fmt.Sprintf("All requests: %v", reqs),
	}

	return stream.SendAndClose(allResponse)
}

func (*server) GreetAsync(stream greetpb.GreetService_GreetAsyncServer) error {

	log.Println("Greeting bi-directionally...")

	// server will response 1 second after getting the request
	for {

		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv() err = %s", err)
		}

		time.Sleep(1 * time.Second)
		stream.Send(&greetpb.GreetResponse{
			Greeting: fmt.Sprintf("Asyncly hello %s", req.Person.FirstName),
		})
	}

	return nil
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	log.Printf("Greeting with deadline %s %s", req.Person.FirstName, req.Person.LastName)

	// fake delay
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Client cancelled the request")
			return nil, status.Error(codes.Canceled, "Client cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}

	return &greetpb.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s %s", req.Person.FirstName, req.Person.LastName),
	}, nil
}


