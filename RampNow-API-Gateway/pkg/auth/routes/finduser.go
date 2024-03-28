package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abhinandkakkadi/rampnow/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

// @Summary Find user by id
// @ID Find user by id
// @Tags User-Section
// @Produce json
// @param id path string true "Find user by id"
// @Success 200 {object} pb.FindUserResponse{}
// @Failure 422 {object} pb.FindUserResponse{}
// @Failure 502 {object} pb.FindUserResponse{}
// @Router /user/finduser/{id} [get]
func FindUser(ctx *gin.Context, c pb.AuthServiceClient) {
	paramsID := ctx.Param("id")

	id, err := strconv.Atoi(paramsID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, pb.FindUserResponse{Error: fmt.Sprint(errors.New("id not found"))})
		return
	}

	res, err := c.FindUser(context.Background(), &pb.FindUserRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), pb.FindUserResponse{Error: fmt.Sprint(res.Error)})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
