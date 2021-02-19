package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "pingpong/pingpongpb"
)

const (
	waitMinSeconds = 5
	waitMaxSeconds = 10
	port           = ":50051"
)

type server struct {
	pb.UnimplementedGameServiceServer
}

func main() {
	fmt.Println("Ping Pong!")
	fmt.Println("Waiting for a new game request, anyone?")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGameServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (*server) NewGame(context.Context, *pb.NewGameRequest) (*pb.NewGameResponse, error) {
	log.Println("Received a request for a game!")
	delay(1, 2)
	return &pb.NewGameResponse{
		Hello: "Yes, I'll play!",
	}, nil
}

func (*server) Hit(context.Context, *pb.HitRequest) (*pb.HitResponse, error) {
	return &pb.HitResponse{
		Num: int32(randNum(10)),
	}, nil
}

func randNum(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

func delay(minSeconds, maxSeconds int) {
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//delaySeconds := r.Intn(maxSeconds-minSeconds) + minSeconds
	//fmt.Printf("Delay %d seconds", delaySeconds)
	time.Sleep(time.Duration(1) * time.Second)
}
