package utils

import "github.com/golang-jwt/jwt/v5"

func GenerateJWT(claims jwt.Claims, methods jwt.SigningMethod, jwtSecret string) (string, error) {
	return jwt.NewWithClaims(methods, claims).SignedString([]byte(jwtSecret))
}
