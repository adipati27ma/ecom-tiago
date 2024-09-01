package auth

import (
	"ecom-tiago/configs"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds); // docs: 15 minutes
	
	// docs: claims is a struct that will be encoded to a JWT (contains metadata)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userID),
		"exp": time.Now().Add(expiration).Unix(),
	});

	tokenString, err := token.SignedString(secret);
	if err != nil {
		return "", err;
	}

	return tokenString, nil;
}