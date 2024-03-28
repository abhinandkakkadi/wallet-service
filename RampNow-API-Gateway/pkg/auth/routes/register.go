package routes

import (
	"context"
	"fmt"
	"net/http"

	_ "github.com/abhinandkakkadi/rampnow/cmd/docs"
	"github.com/abhinandkakkadi/rampnow/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	RampId string `json:"ramp_id"`
	Password  string `json:"password"`
}

// @Summary Register new user
// @ID User Registration
// @Tags Authentication-Service
// @Produce json
// @param RegisterUser body RegisterRequestBody{} true "User registration"
// @Success 200 {object} pb.RegisterResponse{}
// @Failure 422 {object} pb.RegisterResponse{}
// @Failure 502 {object} pb.RegisterResponse{}
// @Router /auth/register [post]
func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:     body.Email,
		Password:  body.Password,
		FullName:  body.FullName,
		RampId:    body.RampId,
	})
	if err != nil {
		fmt.Println("er////r", err)
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
