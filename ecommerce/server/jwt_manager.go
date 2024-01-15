package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

// JWT token manager
type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

// user's information
type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

// returns a new JWT manager
func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}


// return user claim
func (manager *JWTManager) Verify(accessToken string) ([]string, error) {

	token, _ := jwt.Parse(accessToken, nil)
	claims := token.Claims.(jwt.MapClaims) 
	rolesVar := claims["realm_access"].(map[string]interface{})["roles"]
	roles := rolesVar.([]interface {})
	log.Println("--> Verify. User roles: \n",  roles)
	rolesArr := make([]string, 0, len(roles))
	for _, v := range roles {
        rolesArr = append(rolesArr, fmt.Sprintf("%v", v))
    }
	return rolesArr, nil
}