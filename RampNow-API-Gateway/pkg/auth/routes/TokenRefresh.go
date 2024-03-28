package routes

import (
	"context"
	"net/http"
	"strings"

	"github.com/abhinandkakkadi/rampnow/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

// @Summary Refresh token for users
// @ID User RefreshToken
// @Tags Authentication-Service
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.TokenRefreshResponse{}
// @Failure 422 {object} pb.TokenRefreshResponse{}
// @Router /auth/token-refresh [post]
func TokenRefresh(ctx *gin.Context, c pb.AuthServiceClient) {

	// Get access token
	autheader := ctx.Request.Header["Authorization"]
	auth := strings.Join(autheader, " ")
	bearerToken := strings.Split(auth, " ")
	token := bearerToken[1]

	// Check if token is empty
	if token == "" {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Generate Refresh token
	res, err := c.TokenRefresh(context.Background(), &pb.TokenRefreshRequest{
		Token: token,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}

	// Set access token
	ctx.Writer.Header().Set("accesstoken", res.Token)
	ctx.JSON(int(res.Status), &res)
}
