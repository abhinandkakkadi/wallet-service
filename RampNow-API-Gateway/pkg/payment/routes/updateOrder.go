package routes

import (
	"context"
	"net/http"

	"github.com/abhinandkakkadi/rampnow/pkg/payment/pb"
	"github.com/gin-gonic/gin"
)

// @Summary Get Transactions
// @ID Get Transactions
// @Tags Payment-Service
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.GetTransactionResponse{}
// @Failure 422 {object} pb.GetTransactionResponse{}
// @Failure 502 {object} pb.GetTransactionResponse{}
// @Router /transactions [get]
func GetAllTransactions(ctx *gin.Context, c pb.OrderServiceClient) {
	
	res, err := c.GetTransactions(context.Background(), &pb.GetTransactionRequest{})

	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
