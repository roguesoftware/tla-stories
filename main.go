package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/roguesoftware/tla-proto"
)

const port = ":50506"

var initialStories []*pb.StoryItem

type server struct {
	pb.UnimplementedStoryServiceServer
}

func (s *server) GetStories(ctx context.Context, in *pb.StoryRequest) (*pb.StoryReply, error) {
	ids := in.GetIds()
	locationIds := in.GetLocationIds()
	userIds := in.GetUserIds()

	log.Printf("Received: %v %v %v", ids, locationIds, userIds)

	var stories []*pb.StoryItem
	stories = initialStories[1:2]

	return &pb.StoryReply{Stories: stories}, nil
}

func main() {
	// load initial votes
	fileName := "stories.json"
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening %v: %v", fileName, err)
	}
	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(jsonBytes, &initialStories)

	// create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStoryServiceServer(s, &server{})
	log.Printf("Registered story server with %v stories", len(initialStories))
	if err := s.Serve(lis); err != nil {
		log.Fatal(s.Serve(lis))
	}
}
