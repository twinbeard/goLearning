package random

import (
	"math/rand"
	"time"
)

func GenerateSixDigitOtp() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator
	otp := 100000 + rng.Intn(900000)                       // Generate a random 6-digit number
	return otp
}
