package randomGenerator

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomInt() int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	return rand.Intn(MaxInt)
}
