package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/roguesoftware/tla-proto"
)

const (
	port = ":50506"
)

type server struct {
	pb.UnimplementedStoryServiceServer
}

func (s *server) GetStories(ctx context.Context, in *pb.StoryRequest) (*pb.StoryReply, error) {
	ids := in.GetIds()
	locationIds := in.GetLocationIds()
	userIds := in.GetUserIds()

	log.Printf("Received: %v %v %v", ids, locationIds, userIds)

	var stories []*pb.StoryItem
	var story pb.StoryItem

	story.Id = "123456-abcdef"
	story.LocationId = "loc-123"
	story.Title = "City Hall"
	story.Story = "This is a story about a city with a hall"

	stories = append(stories, &story)

	story.Id = "923456-kbcdef"
	story.LocationId = "loc-123"
	story.Title = "City Hall"
	story.Story = "An interesting debate occurred here in 2006"

	stories = append(stories, &story)

	return &pb.StoryReply{Stories: stories}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStoryServiceServer(s, &server{})
	log.Printf("Registered story server")
	if err := s.Serve(lis); err != nil {
		log.Fatal(s.Serve(lis))
	}
}
