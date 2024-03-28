package routes

import (
	"context"
	"net/http"

	"github.com/abhinandkakkadi/rampnow/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary Login user
// @ID User Login
// @Tags Authentication-Service
// @Produce json
// @param LoginUser body LoginRequestBody{} true "User Login"
// @Success 200 {object} pb.LoginResponse{}
// @Failure 422 {object} pb.LoginResponse{}
// @Failure 502 {object} pb.LoginResponse{}
// @Router /auth/login [post]
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}
	
	ctx.Writer.Header().Set("accesstoken", res.AccessToken)
	ctx.JSON(http.StatusCreated, &res)
}
