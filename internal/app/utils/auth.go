package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId int, userRole string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"user_role": userRole,
		"expires": time.Now().Add(3 * time.Hour),
	});

	tokenString, err := token.SignedString([]byte("my-super-private-secret-key"));
	if err != nil {
		return "", err
	};

	return tokenString, nil;
}

func GetToken(token string) string {
	splittedToken := strings.Split(token, " ");

	if len(splittedToken) == 2 && splittedToken[0] == "Bearer" {
		return splittedToken[1];
	}

	return ""
}

func ValidateJWT(t string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(t, func (token *jwt.Token) (interface{}, error) {
		return []byte("my-super-private-secret-key"), nil
	});

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Invalid token");
	}

	claims, ok := token.Claims.(jwt.MapClaims);
	if !ok {
		return nil, fmt.Errorf("Invalid token claims");
	}

	return claims, nil;
}