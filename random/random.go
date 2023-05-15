package random

import (
	"math/rand"
)

type Random struct {
}

func New() *Random {
	return &Random{}
}

func (random *Random) RandomInt() int {
	return rand.Intn(101)
}