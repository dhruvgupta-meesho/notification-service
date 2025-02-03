package handlers

import (
	"context"

	"github.com/dhruvgupta7733/notification-service/database"
	model "github.com/dhruvgupta7733/notification-service/model"
	"github.com/dhruvgupta7733/notification-service/services"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct{
	model.UnimplementedNotifyServer
	DB *sql.DB
}

func (s *Server) GetNotificationInfo(ctx context.Context, in *model.EmailRequest) (*model.GenericResponse, error){
	database.CreateEmailRequest(s.DB, in)
	services.SendKafka(int64(in.Id))

	return &model.GenericResponse{Message: "Info Recieved"}, nil
}
