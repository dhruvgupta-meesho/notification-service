package main

import (
	"log"
	"net"

	"github.com/dhruvgupta7733/notification-service/database"
	handler "github.com/dhruvgupta7733/notification-service/handlers"
	model "github.com/dhruvgupta7733/notification-service/model"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5051")

	if err != nil {
		log.Fatalf("failed to listen on port 5051: %v", err)
		}

	db := database.DbConnect()

	s := grpc.NewServer()
	model.RegisterNotifyServer(s, &handler.Server{DB: db})
	log.Printf("gRPC server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
	log.Fatalf("failed to serve: %v", err)
	}
}
