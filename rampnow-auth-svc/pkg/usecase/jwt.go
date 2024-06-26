package usecase

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abhinandkakkadi/rampnow-auth-service/pkg/domain"
	usecase "github.com/abhinandkakkadi/rampnow-auth-service/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
)

type jwtUsecase struct {
	SecretKey string
}

// GenerateRefreshToken implements interfaces.JWTUsecase
func (j *jwtUsecase) GenerateRefreshToken(userid uint, username string) (string, error) {
	claims := &domain.SignedDetails{
		UserId:   userid,
		UserName: username,
		Source:   "refreshtoken",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(150)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}

	return signedToken, err
}

// GenerateRefreshToken implements interfaces.JWTUsecase
func (j *jwtUsecase) GenerateAccessToken(userid uint, username string) (string, error) {

	claims := &domain.SignedDetails{
		UserId:   userid,
		UserName: username,
		Source:   "accesstoken",

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}

	return signedToken, err
}

// GetTokenFromString implements interfaces.JWTUsecase
func (j *jwtUsecase) GetTokenFromString(signedToken string, claims *domain.SignedDetails) (*jwt.Token, error) {
	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.SecretKey), nil
	})

}

// VerifyToken implements interfaces.JWTUsecase
func (j *jwtUsecase) VerifyToken(signedToken string) (bool, *domain.SignedDetails) {
	claims := &domain.SignedDetails{}
	token, _ := j.GetTokenFromString(signedToken, claims)

	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func NewJWTUsecase() usecase.JWTUsecase {
	return &jwtUsecase{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
