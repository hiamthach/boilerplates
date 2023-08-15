package controller

import (
	log_helper "go-microservices/cores/logs"
	pb "go-microservices/pd"
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
