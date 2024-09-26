package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lehau17/food_delivery/components/jwtprovider"
)

type jwtProvider struct {
	secretkey string
}

func NewJwtProvider(secretkey string) *jwtProvider {
	return &jwtProvider{secretkey: secretkey}
}

type myClaims struct {
	Payload *jwtprovider.TokenPayload
	jwt.RegisteredClaims
}

func NewMyClaims() *myClaims {
	return &myClaims{}
}

func (j *jwtProvider) SignToken(payload *jwtprovider.TokenPayload, expiry int) (*jwtprovider.Token, error) {
	claims := myClaims{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiry) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Server GO_Lang : Le Trung Hau",
			// Audience:  []string{"somebody_else"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(j.secretkey))
	if err != nil {
		return nil, nil // handle error
	}
	return &jwtprovider.Token{Token: ss, Created: time.Now(), Expiry: expiry}, nil
}

func (j *jwtProvider) Validate(token string) (*jwtprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretkey), nil
	})
	if err != nil {
		return nil, jwtprovider.ErrinvalidToken
	}
	if !res.Valid {
		return nil, jwtprovider.ErrinvalidToken

	}
	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, jwtprovider.ErrinvalidToken

	}
	return claims.Payload, nil
}
