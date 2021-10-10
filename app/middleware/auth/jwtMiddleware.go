package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

/*func (jwtConfig *ConfigJWT) Init() jwtware.JWTConfig {

}*/

//generate Token
func (jwtConfig *ConfigJWT) GenerateToken(userID int) string {
	claims := JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConfig.ExpiresDuration))).Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConfig.SecretJWT))
	return token
}
