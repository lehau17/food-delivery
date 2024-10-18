package jwtprovider

import (
	"errors"
	"net/http"
	"time"

	"github.com/lehau17/food_delivery/common"
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	Uid  int    `json:"uid"`
	Role string `json:"role"`
}

type TokenProvider interface {
	SignToken(payload *TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrinvalidToken = common.NewFullErrorResponse(int(http.StatusUnauthorized), errors.New("Invalid token"), "Invalid token", "Invalid token", "ErrTokenInvalid")
	ErrCreateToken  = common.NewFullErrorResponse(int(http.StatusUnauthorized), errors.New("error create token"), "error create token", "error create token", "ErrCreateToken")
)
