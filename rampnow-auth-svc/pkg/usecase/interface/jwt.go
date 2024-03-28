package interfaces

import (
	domain "github.com/abhinandkakkadi/rampnow-auth-service/pkg/domain"
	"github.com/golang-jwt/jwt"
)

type JWTUsecase interface {
	GenerateAccessToken(userid uint, userName string) (string, error)
	VerifyToken(token string) (bool, *domain.SignedDetails)
	GetTokenFromString(signedToken string, claims *domain.SignedDetails) (*jwt.Token, error)
	GenerateRefreshToken(userid uint, userName string) (string, error)
}
