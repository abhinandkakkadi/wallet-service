package routes

import (
	"context"
	"net/http"

	"github.com/abhinandkakkadi/rampnow/pkg/payment/pb"
	"github.com/abhinandkakkadi/rampnow/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Get Transactions
// @ID Get Transactions
// @Tags Payment-service
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Failure 502 {object} response.Response{}
// @Router /transactions [get]
func GetAllTransactions(ctx *gin.Context, c pb.PaymentServiceClient) {
	
	res, err := c.GetTransactions(context.Background(), &pb.GetTransactionRequest{})

	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}

	resp := response.SuccessResponse(true, "SUCCESS", res)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	response.ResponseJSON(*ctx, resp)
}
