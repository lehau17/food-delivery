package random

import (
	"math/rand"
	"strconv"
	"time"
)

type ramdom struct {
}

func NewRandom() *ramdom {
	return &ramdom{}
}

func (r *ramdom) Generate() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return strconv.Itoa(r1.Intn(999999999999))
}
