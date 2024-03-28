package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abhinandkakkadi/rampnow/pkg/order_svc/pb"
	"github.com/gin-gonic/gin"
)

// @Summary Fetch Order
// @ID Fetchorder
// @Tags Order-service
// @Produce json
// @Security BearerAuth
// @Param        status   query      string  false  "Status : "
// @Param        mintotal   query      string  false  "Min Total : "
// @Param        maxtolat   query      string  false  "Max Total : "
// @Param        sortby   query      string  false  "Sort By : "
// @Param        sortorder   query      string  false  "Sort Order : "
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /order [get]
func FetchOrder(ctx *gin.Context, c pb.OrderServiceClient) {

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

	ctx.JSON(http.StatusOK, &res)

}
