package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pabloCode010/database_programming_project/config"
	"github.com/pabloCode010/database_programming_project/models"
)

func New(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":      "my-auth-server",
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
			"sub":      user.ID,
			"username": user.Username,
			"role":     user.Role,
		})

	return token.SignedString([]byte(config.JwtKey))
}
