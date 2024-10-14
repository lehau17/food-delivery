package common

import (
	"math/rand"
	"strconv"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomSequence(n int) string {
	b := make([]rune, n)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range b {
		b[i] = letters[r1.Intn(999999)%len(letters)]
	}
	return string(b)
}

func GetSalt(len int) string {
	if len <= 0 {
		len = 50
	}
	return randomSequence(len)
}

func randomOTP() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return strconv.Itoa(r1.Intn(999999) + 100000)
}

func GetOtp() string {
	return randomOTP()
}
