package handlers

import (
	"context"

	"github.com/dhruvgupta7733/notification-service/database"
	model "github.com/dhruvgupta7733/notification-service/model"
	"github.com/dhruvgupta7733/notification-service/services"
	"github.com/redis/go-redis/v9"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct{
	model.UnimplementedNotifyServer
	DB *sql.DB
	RDB *redis.Client
}

func (s *Server) SendNotificationInfo(ctx context.Context, in *model.EmailRequest) (*model.GenericResponse, error){
	database.CreateEmailRequest(s.DB, in)
	services.SendKafka(int64(in.Id))
	
	return &model.GenericResponse{Message: "Info Recieved"}, nil
}

func (s *Server) AddBlacklisted(ctx context.Context, in *model.Email) (*model.GenericResponse, error){
	msg  := services.AddBlacklistEmail(s.RDB, in.EmailId)
	return &model.GenericResponse{Message: msg}, nil
}

func (s *Server) RemoveBlacklisted(ctx context.Context, in *model.Email) (*model.GenericResponse, error){
	msg := services.RemoveBlacklistEmail(s.RDB, in.EmailId)
	return &model.GenericResponse{Message: msg}, nil
}