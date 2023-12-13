// auth/auth.go
package auth

import (
	"errors"
	"main/model"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var secretKey = []byte(os.Getenv("SECRET_KEY")) // Ganti dengan kunci rahasia yang lebih aman

func GenerateToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku selama 1 hari

	claims := &Claims{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(c *gin.Context, db *gorm.DB) (*Claims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header is missing")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	// Cek apakah token ada di database dan belum expired
	var user model.User
	if err := db.Where("username = ? AND token = ?", claims.Username, tokenString).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("user not found or token mismatch")
		}
		return nil, errors.New("internal server error")
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
