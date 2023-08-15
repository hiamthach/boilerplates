package router

import (
	"context"
	"encoding/json"
	"go-microservices/app/config"
	"go-microservices/app/controller"
	"go-microservices/app/middleware"
	pb "go-microservices/pd"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func StartGRPC() {
	grpcServer := grpc.NewServer()
	server := controller.GetServer()
	port := config.Get().Server.GRPCServerUrl
	pb.RegisterMyServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("grpc server start at port:", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Start() {
	server := controller.GetServer()
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	err := pb.RegisterMyServiceHandlerServer(context.Background(), grpcMux, server)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	port := config.Get().Server.HTTPServerUrl
	log.Printf("Listening on port %s\n", port)
	_ = http.ListenAndServe(port, before(grpcMux))
}

func before(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := middleware.Auth.AuthToken(w, r); err != nil {
			data, _ := json.Marshal(&pb.BaseReply{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			contentType := getRequestContentType(r)
			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write(data)

			return
		}
		middleware.CORS(w, r)
		h.ServeHTTP(w, r)
	})
}

func getRequestContentType(r *http.Request) string {
	contentType := r.Header.Get("Content-Type")
	if len(contentType) == 0 {
		contentType = "application/json"
	}
	return contentType
}
