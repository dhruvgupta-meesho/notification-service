package handlers

import (
	"context"

	"github.com/dhruvgupta7733/notification-service/database"
	model "github.com/dhruvgupta7733/notification-service/model"
	"github.com/dhruvgupta7733/notification-service/services"
	"github.com/hashicorp/go-uuid"
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
	uid, err:= uuid.GenerateUUID()
	if err != nil{
		return &model.GenericResponse{Message: "Error creating UUID"}, nil
	}
	database.CreateEmailRequest(uid, s.DB, in)
	services.SendKafka(uid)
	
	return &model.GenericResponse{Message: "Info Recieved"}, nil
}

// func (s *Server) AddBlacklisted(ctx context.Context, in *model.Email) (*model.GenericResponse, error){
// 	msg  := services.AddBlacklistEmail(s.RDB, in.EmailId)
// 	return &model.GenericResponse{Message: msg}, nil
// }

// func (s *Server) RemoveBlacklisted(ctx context.Context, in *model.Email) (*model.GenericResponse, error){
// 	msg := services.RemoveBlacklistEmail(s.RDB, in.EmailId)
// 	return &model.GenericResponse{Message: msg}, nil
// }

func (s *Server) AddBlacklisted(ctx context.Context, in *model.EmailList) (*model.GenericResponse, error) {
	msg := services.AddBlacklistEmails(s.RDB, in.EmailIds)
	return &model.GenericResponse{Message: msg}, nil
}

func (s *Server) RemoveBlacklisted(ctx context.Context, in *model.EmailList) (*model.GenericResponse, error) {
	msg := services.RemoveBlacklistEmails(s.RDB, in.EmailIds)
	return &model.GenericResponse{Message: msg}, nil
}


func (s *Server) GetRequestStatus(ctx context.Context, in *model.RequestID) (*model.RequestStatusResponse, error){
	req, _ := database.GetEmailRequest(s.DB, in.Id)
	return &model.RequestStatusResponse{
		EmailId : req.EmailId,
		Message : req.Message,
		FailureCode : req.Failure_code,
		FailureComment: req.Failure_comments,
	}, nil
}