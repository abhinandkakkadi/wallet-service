package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abhinandkakkadi/rampnow/pkg/payment/pb"
	"github.com/abhinandkakkadi/rampnow/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Get Wallet Balance By user id
// @ID Find wallet by user id
// @Tags Payment-service
// @Produce json
// @Security BearerAuth
// @param id path string true "Find wallet by user id"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Failure 502 {object} response.Response{}
// @Router /wallet_balance/{id} [get]
func GetWalletBalance(ctx *gin.Context, c pb.PaymentServiceClient) {

	paramsID := ctx.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, pb.GetWalletBalanceResponse{Error: fmt.Sprint(errors.New("id not found"))})
		return
	}

	res, err := c.GetWalletBalance(context.Background(), &pb.GetWalletBalanceRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), pb.GetWalletBalanceResponse{Error: fmt.Sprint(res.Error)})
		return
	}

	resp := response.SuccessResponse(true, "SUCCESS", res)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	response.ResponseJSON(*ctx, resp)

}
