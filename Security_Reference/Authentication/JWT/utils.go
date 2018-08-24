package ctrlJwt

import (
	"crypto/rsa"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/tk.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/tk.rsa.pub"
)

type AppClaims struct {
	// Sub       uuid.UUID `json:"Sub"`
	// IssuedAt  int64     `json:"iss"`
	// ExpiresAt int64     `json:"exp"`
	Role string `json:"role"`
	jwt.StandardClaims
}
