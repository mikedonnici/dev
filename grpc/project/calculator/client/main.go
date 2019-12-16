package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/mkedonnici/grpc/project/calculator/calculatorpb"
)

func main() {

	ctx := context.Background()

	// dial server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial() err = %s", err)
	}
	client := calculatorpb.NewCalculatorServiceClient(conn)

	//if err := unaryRequest(ctx, client); err != nil {
	//	log.Printf("unaryRequest() err = %s", err)
	//}
	//
	//if err := serverStreamRequest(ctx, client); err != nil {
	//	log.Printf("serverStreamRequest() err = %s", err)
	//}
	//
	//if err := clientStreamRequest(ctx, client); err != nil {
	//	log.Printf("clientStreamRequest() err = %s", err)
	//}
	//
	//if err := biDirectionalRequest(ctx, client); err != nil {
	//	log.Printf("biDirectionalRequest() err = %s", err)
	//}

	if err := unaryRequestWithCodes(ctx, client); err != nil {
		log.Printf("unaryRequestWithCodes() err = %s", err)
	}
}

func unaryRequest(ctx context.Context, client calculatorpb.CalculatorServiceClient) error {
	// set up a client

	// make a request
	req := calculatorpb.AddRequest{
		Num1: 7,
		Num2: 32,
	}
	res, err := client.Sum(ctx, &req)
	if err != nil {
		return err
	}
	fmt.Println("Sum is", res.Sum)
	return nil
}

func serverStreamRequest(ctx context.Context, client calculatorpb.CalculatorServiceClient) error {

	req := calculatorpb.PrimeNumberDecompositionRequest{
		Number: 120,
	}
	resStream, err := client.PrimeNumberDecomposition(ctx, &req)
	if err != nil {
		return err
	}

	// process stream
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(res.Number)
	}
	return nil
}

func clientStreamRequest(ctx context.Context, client calculatorpb.CalculatorServiceClient) error {

	reqs := []calculatorpb.AverageRequest{
		{Number: 43},
		{Number: 59},
		{Number: 453},
		{Number: 345},
		{Number: 5594},
		{Number: 21},
	}

	stream, err := client.Average(ctx)
	if err != nil {
		log.Fatalf("client.Average() err = %s", err)
	}
	for _, r := range reqs {
		err := stream.Send(&r)
		if err != nil {
			log.Fatalf("stream.Send() err = %s", err)
		}
	}

	ave, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("stream.CloseAndRecv() err = %s", err)
	}
	fmt.Printf("Average = %v\n", ave.Average)
	return nil
}

func biDirectionalRequest(ctx context.Context, client calculatorpb.CalculatorServiceClient) error {

	nums := []int32{
		1, 3, 4, 6, 5, 234, 345, 66, 2, 34, 5, 6, 456, 4, 234, 23, 4,
		34, 234, 345, 546, 345, 63, 56, 456, 4256, 24, 56, 2436, 2, 46,
	}

	stream, err := client.BiggestInSet(ctx)
	if err != nil {
		log.Fatalf("client.BiggestInSet() err = %s", err)
	}

	waitHere := make(chan struct{})

	// sending channel
	go func() {
		for _, n := range nums {
			time.Sleep(500 * time.Millisecond)
			if err := stream.Send(&calculatorpb.BiggestInSetRequest{Number: n}); err != nil {
				log.Fatalf("stream.Send() err = %s", err)
			}
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("stream.CloseSend() err = %s", err)
		}
	}()

	// rec chan
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("stream.Recv() err = %s", err)
			}
			// anything we receive is now the biggest number
			fmt.Printf("Biggest inset is now %d\n", res.Number)
		}
		close(waitHere)
	}()

	<-waitHere
	return nil
}

func unaryRequestWithCodes(ctx context.Context, client calculatorpb.CalculatorServiceClient) error {
	nums := []int32{64, 9, -3, -100, 100, 1000, 623}
	for _, n := range nums {
		err := doSquareRoot(client, ctx, n)
		if err != nil {
			return err
		}
	}
	return nil
}

func doSquareRoot(c calculatorpb.CalculatorServiceClient, ctx context.Context, num int32) error {

	req := calculatorpb.SquareRootRequest{
		Number: num,
	}

	res, err := c.SquareRoot(ctx, &req)
	if err != nil {
		// first check if it a gRPC error...
		grpcStatus, ok := status.FromError(err)
		if ok {
			log.Printf(grpcStatus.Message())
			return nil
		}
		// ...otherwise its a code-level error
		return err
	}

	fmt.Printf("Square root of %v is %v\n", req.Number, res.Number)
	return nil
}
