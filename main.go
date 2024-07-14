package main

import (
    "log"
    "net"
    "net/http"

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
    // Initialize Gorilla Mux router
    router := mux.NewRouter()

    // Test endpoint
    router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    }).Methods("GET")

    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterExperienceServiceServer(s, &handlers.ExperienceServiceServer{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }

    // Start HTTP server
    log.Fatal(http.ListenAndServe(":8080", router))
}
