package controller

import (
	"context"
	"go-microservices/app/config"
	log_helper "go-microservices/core/logs"
	pb "go-microservices/pd"
	"net/http"
)

type Server struct {
	pb.UnimplementedMyServiceServer
	logger log_helper.ILogHelper
}

func GetServer() *Server {
	return &Server{
		logger: log_helper.GetLogHelper(),
	}
}

func (s *Server) HealthCheck(ctx context.Context, req *pb.PingReq) (*pb.PingResp, error) {
	return &pb.PingResp{
		Code:    http.StatusOK,
		Message: "OK",
		Version: config.Get().App.Version,
	}, nil
}
