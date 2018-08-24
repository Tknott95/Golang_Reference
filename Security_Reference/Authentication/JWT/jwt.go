package ctrlJwt

import (
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Init() initializes the ctrlJwt
// func Init(uid uuid.UUID) {
// 	InitKeys()
// 	GenerateJWT(uid, "login")
// }

func InitKeys() {

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}

// GenerateJWT generates a new JWT token
func GenerateJWT(sub string, role string) (string, error) {
	// Create the Claims
	_subject := "fsff" // sub.String()
	_timeTwenty := time.Now().Add(time.Minute * 20).Unix()
	_timeNow := time.Now().Unix()
	claims := AppClaims{
		role,
		jwt.StandardClaims{
			ExpiresAt: _timeTwenty,
			IssuedAt:  _timeNow,
			Subject:   _subject,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	print(token)

	ss, err := token.SignedString(token)
	if err != nil {
		return "", err
	}

	print(ss)

	// print("\n JWT")

	// fmt.Printf("%s", ss)

	return ss, nil
}
