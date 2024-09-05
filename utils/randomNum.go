package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber() uint64 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a new random generator
	return 1000000000 + rnd.Uint64()%9000000000            // ensures it's a 10-digit number
}
