package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/mkedonnici/grpc/project/greet/greetpb"
)

const (
	useTLS = true
	certFile = "ssl/ca.crt"
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

	client := greetpb.NewGreetServiceClient(conn)

	req := greetpb.GreetRequest{
		Person: &greetpb.Person{
			FirstName: "Mike",
			LastName:  "Donnici",
		},
	}

	//unaryRequest(client, req)
	//serverStreamingRequest(client, req)
	//clientStreamingRequest(client)
	//biDirectionalRequest(client)
	unaryRequestWithDeadline(client, req, 5 * time.Second)
	unaryRequestWithDeadline(client, req, 1 * time.Second)
}

func unaryRequest(client greetpb.GreetServiceClient, req greetpb.GreetRequest) {
	log.Println("Unary request...")
	res, err := client.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("client.Greet() err = %s", err)
	}
	fmt.Println(res.Greeting)
}

func serverStreamingRequest(client greetpb.GreetServiceClient, req greetpb.GreetRequest) {
	fmt.Println("Server-streaming request...")
	resStream, err := client.GreetManyTimes(context.Background(), &req)
	if err != nil {
		log.Fatalf("client.Greet() err = %s", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("resStream.Recv() err = %s", err)
		}
		fmt.Println(msg.GetGreeting())
	}
}

func clientStreamingRequest(client greetpb.GreetServiceClient) {
	fmt.Println("Client streaming request...")

	reqStream, err := client.GreetAfterManyTimes(context.Background())
	if err != nil {
		log.Fatalf("client.GreetAfterManyTimes() err = %s", err)
	}

	people := []*greetpb.Person{
		{
			FirstName: "Mike",
			LastName:  "Donnici",
		},
		{
			FirstName: "Christie",
			LastName:  "Wood",
		},
		{
			FirstName: "Maia",
			LastName:  "Donnici",
		},
		{
			FirstName: "Leo",
			LastName:  "Donnici",
		},
		{
			FirstName: "Milo",
			LastName:  "Beach",
		},
	}

	for _, p := range people {
		err := reqStream.Send(&greetpb.GreetRequest{Person: p})
		if err != nil {
			log.Fatalf("reqStream.Send() err = %s", err)
		}
	}

	res, err := reqStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("reqStream.CloseAndRecv() err = %s", err)
	}

	fmt.Println(res.Greeting)
}

func biDirectionalRequest(client greetpb.GreetServiceClient) {

	fmt.Println("Bi-directional request...")

	biStream, err := client.GreetAsync(context.Background())
	if err != nil {
		log.Fatalf("client.GreetAsync() err = %s", err)
	}

	waitChan := make(chan struct{})

	people := []*greetpb.Person{
		{
			FirstName: "Mike",
			LastName:  "Donnici",
		},
		{
			FirstName: "Christie",
			LastName:  "Wood",
		},
		{
			FirstName: "Maia",
			LastName:  "Donnici",
		},
		{
			FirstName: "Leo",
			LastName:  "Donnici",
		},
		{
			FirstName: "Milo",
			LastName:  "Beach",
		},
	}

	// send
	go func() {
		for _, p := range people {
			err := biStream.Send(&greetpb.GreetRequest{Person: p})
			if err != nil {
				log.Fatalf("biStream.Send() err = %s", err)
			}
		}
		// finished!
		err := biStream.CloseSend()
		if err != nil {
			log.Fatalf("biStream.CloseSend() err = %s", err)
		}
	}()

	// rec
	go func() {
		for {
			res, err := biStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("biStream.Recv() err = %s", err)
			}
			fmt.Println(res.Greeting)
		}
		close(waitChan)
	}()

	// wait here
	<-waitChan
}

func unaryRequestWithDeadline(client greetpb.GreetServiceClient, req greetpb.GreetRequest, timeout time.Duration) {

	log.Println("Unary request with deadline...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := client.GreetWithDeadline(ctx, &req)
	if err != nil {

		statusErr, exists := status.FromError(err)
		if exists {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalln("Deadline exceeded")
			}
			log.Fatalf("statusErr = %s", statusErr)
		}
		log.Fatalf("client.GreetWithDeadline() err = %s", err)
	}
	fmt.Println(res.Greeting)
}
