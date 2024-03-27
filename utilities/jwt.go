package utilities

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "SuperSecretKey123!"

func GenerateToken(email string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 16).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (string, error) {
	ParsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check if the token was signed with the correct method type (using go type checking)
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return "", errors.New("could not parse token")
	}

	tokenIsValid := ParsedToken.Valid

	if !tokenIsValid {
		return "", errors.New("invalid token")
	}

	userId, err := getUserID(ParsedToken)

	if err != nil {
		return "", errors.New("failed to get user id from token")
	}

	return userId, nil
}

func getUserID(parsedToken *jwt.Token) (string, error) {
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("invalid token claims")
	}

	userId := claims["userId"].(string)

	return userId, nil
}
