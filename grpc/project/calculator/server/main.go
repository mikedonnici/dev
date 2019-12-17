package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"math"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"

	"github.com/mkedonnici/grpc/project/calculator/calculatorpb"
)

// implement the SumServiceServer interface
type server struct{}

func main() {

	// set up listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("net.Listen() err = %s", err)
	}

	// Set up server
	s := grpc.NewServer()

	// register service(s)
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	// listen...
	log.Println("Server listening...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc.Server.Server() err = %s", err)
	}
}

func (*server) Sum(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	log.Printf("Server request for Sum(%v, %v)...", req.Num1, req.Num2)
	return &calculatorpb.AddResponse{
		Sum: req.Num1 + req.Num2,
	}, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {

	log.Printf("PrimeNumberDecomposition(%v)\n", req.Number)

	k := int32(2)
	n := req.Number
	for {
		if n <= 1 {
			break
		}

		// if k evenly divides into N ... this is a factor
		if n%k == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				Number: k,
			})
			n = n / k // divide N by k so that we have the rest of the number left.
		} else {
			k++
		}
	}
	return nil
}

func (*server) Average(stream calculatorpb.CalculatorService_AverageServer) error {

	log.Printf("Average()\n")

	var nums []int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv() err = %s", err)
		}
		nums = append(nums, req.Number)
	}

	var sum int32
	for _, n := range nums {
		sum += n
	}

	ave := float32(sum) / float32(len(nums))
	return stream.SendAndClose(&calculatorpb.AverageResponse{Average: ave})
}

func (*server) BiggestInSet(stream calculatorpb.CalculatorService_BiggestInSetServer) error {

	log.Printf("BiggestInSet()\n")

	var biggest int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv() err = %s", err)
		}
		if req.Number > biggest {
			biggest = req.Number
			stream.Send(&calculatorpb.BiggestInSetResponse{
				Number: req.Number,
			})
		}
	}

	return nil
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {

	log.Printf("Server request for SquareRoot(%v)...", req.Number)

	if req.Number < 0 {
		msg := fmt.Sprintf("Argument must be a positive integer, received: %v", req.Number)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	result := math.Sqrt(float64(req.Number))
	return &calculatorpb.SquareRootResponse{
		Number: float32(result),
	}, nil
}
