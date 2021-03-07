// Package main implements a client for Game service.
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"

	pb "pingpong/pingpongpb"
)

const (
	waitMaxSeconds = 10
	address = "localhost:50051"
)

type gameScore struct {
	serverScore int
	clientScore int
}

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGameServiceClient(conn)

	// Contact the server and print out its response.
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	ctx, cancel := context.WithTimeout(context.Background(), waitMaxSeconds * time.Second)
	defer cancel()
	r, err := c.NewGame(ctx, &pb.NewGameRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", r.GetHello())
	log.Printf("Ready!")

	score := gameScore{}
	for {
		ctx, cancel = context.WithTimeout(context.Background(), 3 * time.Second)
		h, err := c.Hit(ctx, &pb.HitRequest{Num: int32(randNum(10))})
		if err != nil {
			cancel()
			log.Fatalf("could not hit: %v", err)
		}
		log.Printf("Response: %d", h.GetNum())
		if h.GetNum() > 6 {
			score.serverScore += 1
		}
		if h.GetNum() < 4 {
			score.clientScore += 1
		}
		if score.serverScore > 20 || score.clientScore > 20 {
			break
		}
	}
	fmt.Println("Game!")
	fmt.Printf("Final score: Server %d, Client %d", score.serverScore, score.clientScore)
}

func randNum(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}
