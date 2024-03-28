package http

import (
	"fmt"
	"log"
	"net"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/api/services"
	pb "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewGRPCServer(orderService *services.OrderService, grpcPort string) *grpc.Server {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	fmt.Println("grpcPort/////", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, orderService)
	// pb.RegisterOrderServiceServer(grpcServer, orderService)

	if err := grpcServer.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err).
	}
	return grpcServer
}

func NewServerHTTP(orderService *services.OrderService) *ServerHTTP {
	engine := gin.New()
	go NewGRPCServer(orderService, "50057")
	// Use logger from Gin
	engine.Use(gin.Logger())

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8888")
}
