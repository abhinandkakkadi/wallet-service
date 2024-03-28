package routes

import (
	"context"
	"net/http"

	"github.com/abhinandkakkadi/rampnow/pkg/order_svc/pb"
	"github.com/gin-gonic/gin"
)

// @Summary Update Order
// @ID Updateorder
// @Tags Order-service
// @Produce json
// @Security BearerAuth
// @Param updateorderdetials body domain.UpdateOrder{} true "Update Order Detials"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /order [put]
func UpdateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	
	res, err := c.GetTransactions(context.Background(), &pb.GetTransactionRequest{})

	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}

	ctx.JSON(http.StatusOK, &res)

}
