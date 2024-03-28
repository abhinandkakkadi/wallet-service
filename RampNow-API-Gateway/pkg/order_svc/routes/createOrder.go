package routes

import (
	"fmt"
	"net/http"

	"github.com/abhinandkakkadi/rampnow/pkg/domain"
	"github.com/abhinandkakkadi/rampnow/pkg/order_svc/pb"
	"github.com/abhinandkakkadi/rampnow/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Create Order
// @ID createorder
// @Tags Order-service
// @Produce json
// @Security BearerAuth
// @Param orderdetials body domain.Order{} true "Order Detials"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /order [post]
func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := domain.Transaction{}
	

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Printf("trnasaction %+v", body)

	res, err := c.CreateTransaction(ctx, &pb.CreateTransactionRequest{
		PayerRampId: body.PayerRampId,
		PayeeRampId: body.PayeeRampId,
		PaymentAmount: float32(body.PaymentAmount),
	})

	if err != nil {
		responses := response.ErrorResponse("Failed to Create Order", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		response.ResponseJSON(*ctx, responses)
		return
	}

	responses := response.SuccessResponse(true, "SUCCESS", res)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	response.ResponseJSON(*ctx, responses)
	return

}
