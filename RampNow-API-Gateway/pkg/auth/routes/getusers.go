package routes

import (
	"context"
	"net/http"

	"github.com/abhinandkakkadi/rampnow/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

// @Summary Get users
// @ID Get users
// @Tags User-Section
// @Produce json
// @Success 200 {object} pb.GetUsersResponse{}
// @Failure 422 {object} pb.GetUsersResponse{}
// @Failure 502 {object} pb.GetUsersResponse{}
// @Router /user/getusers [get]
func GetUsers(ctx *gin.Context, c pb.AuthServiceClient) {
	res, err := c.GetUsers(context.Background(), &pb.GetUsersRequest{})
	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
