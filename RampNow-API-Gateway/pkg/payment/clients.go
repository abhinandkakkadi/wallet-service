package payment

import (
	"fmt"

	"github.com/abhinandkakkadi/rampnow/pkg/config"
	"github.com/abhinandkakkadi/rampnow/pkg/payment/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.PaymentServiceClient
}

func InitServiceClient(c *config.Config) pb.PaymentServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewPaymentServiceClient(cc)
}
