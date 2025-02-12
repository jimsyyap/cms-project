package utils

import (
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-jwt-secret-key") // Replace with your secret key

// GenerateJWT generates a JWT for a user
func GenerateJWT(userID uint, role string) (string, error) { // userID is uint
    claims := jwt.MapClaims{
        "user_id": userID,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// VerifyJWT verifies a JWT and returns the claims
func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}
