package http

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"

	"github.com/abhinandkakkadi/rampnow-auth-service/pkg/pb"
	"google.golang.org/grpc"

	handler "github.com/abhinandkakkadi/rampnow-auth-service/pkg/api/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func StartGRPCServer(userHandler *handler.UserHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	fmt.Println("grpcPort/////", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, userHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewServerHTTP(userHandler *handler.UserHandler)  {
	// engine := gin.New()
	StartGRPCServer(userHandler, "50056")	
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3002")
}
