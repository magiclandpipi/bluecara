// handlers/experience.go
package handlers

import (
    "context"

    pb "github.com/magiclandpipi/bluecara/proto"
)

type ExperienceServiceServer struct {
    pb.UnimplementedExperienceServiceServer
}

func (s *ExperienceServiceServer) GetHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
    return &pb.HelloResponse{Message: "Hello, World!"}, nil
}
