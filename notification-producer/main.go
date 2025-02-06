package main

import (
	"log"
	"net"

	"github.com/dhruvgupta7733/notification-service/database"
	handler "github.com/dhruvgupta7733/notification-service/handlers"
	model "github.com/dhruvgupta7733/notification-service/model"
	"github.com/dhruvgupta7733/notification-service/services"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5051")

	if err != nil {
		log.Fatalf("Failed to listen on port 5051! : %v", err)
		}

	db := database.DbConnect()
	rdb := services.MakeRedisConn()

	s := grpc.NewServer()
	model.RegisterNotifyServer(s, &handler.Server{
		DB: db,
		RDB: rdb,
	})
	
	log.Printf("gRPC server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
	log.Fatalf("Failed to serve! : %v", err)
	}
}
