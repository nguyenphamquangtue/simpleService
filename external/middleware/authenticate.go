package middleware

import (
	"os"
	"simpleService/external/dto"
	"time"

	"github.com/golang-jwt/jwt"
)

// This method generates a JWT
func GenerateJWT(user *dto.User) (string, error) {
	// Create a claims object with user information
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // JWT expiration time
	}

	// Create a JWT with claims and a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(os.Getenv("SECRET_KEY")) // Secret key, replace with an actual and secure key

	// Sign and convert the JWT to a string
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
