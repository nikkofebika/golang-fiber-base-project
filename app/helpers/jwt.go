package helpers

import (
	"golang-fiber-base-project/app/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenDetail struct {
	AccessToken       string    `json:"access_token"`
	AccessTokenExpiry time.Time `json:"access_token_expiry"`
}

func GenerateToken(user *models.User, jwtSecret string) (*JwtTokenDetail, error) {
	jwtTokenDetail := &JwtTokenDetail{}

	jwtTokenDetail.AccessTokenExpiry = time.Now().Add(time.Hour * 24)

	accessClaims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   jwtTokenDetail.AccessTokenExpiry.Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	var err error

	jwtTokenDetail.AccessToken, err = accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	return jwtTokenDetail, nil
}

func ValidateToken(tokenString, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	})
}

func ExtractUserID(token *jwt.Token) (uint, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, jwt.ErrInvalidKey
	}

	return uint(claims["id"].(float64)), nil
}
