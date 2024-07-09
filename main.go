package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/magiclandpipi/bluecara/handlers"
    pb "github.com/magiclandpipi/bluecara/proto"
    "github.com/joho/godotenv"
    "google.golang.org/grpc"
)

const (
    port = ":50051"
)

func main() {

    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterExperienceServiceServer(s, &handlers.ExperienceServiceServer{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
