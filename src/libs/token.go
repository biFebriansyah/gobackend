package libs

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecrets = []byte(os.Getenv("JWT_KEYS"))

type claims struct {
	User_id string
	Role    string
	jwt.StandardClaims
}

func NewToken(uid string, role string) *claims {
	return &claims{
		User_id: uid,
		Role:    role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}
}

func (c *claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return tokens.SignedString(mySecrets)
}

func CheckToken(token string) (*claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecrets), nil
	})

	if err != nil {
		return nil, err
	}

	claims := tokens.Claims.(*claims)
	return claims, err
}
