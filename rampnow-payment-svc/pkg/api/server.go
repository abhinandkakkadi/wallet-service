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

func NewGRPCServer(paymentService *services.PaymentService, grpcPort string) *grpc.Server {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	fmt.Println("grpcPort/////", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPaymentServiceServer(grpcServer, paymentService)
	// pb.RegisterOrderServiceServer(grpcServer, orderService)

	if err := grpcServer.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err).
	}
	return grpcServer
}

func NewServerHTTP(paymentService *services.PaymentService) {

	NewGRPCServer(paymentService, "50057")

}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8888")
}
