package helpers

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type randomGenerator struct {
	rand *rand.Rand
}

func NewRandomGenerator() *randomGenerator {
	return &randomGenerator{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (rg *randomGenerator) RandomInt(min, max int64) int64 {
	return min + rg.rand.Int63n(max-min+1)
}

func (rg *randomGenerator) RandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = alphabet[rg.rand.Intn(len(alphabet))]
	}
	return string(result)
}

func (rg *randomGenerator) RandomUsername() string {
	return rg.RandomString(8)
}

func (rg *randomGenerator) RandomEmail() string {
	username := rg.RandomString(8)
	domain := rg.RandomString(6) + ".com" // Example domain
	return username + "@" + domain
}

func (rg *randomGenerator) RandomPhone() string {
	// Generating a random phone number with a length between 10 and 13 characters
	length := rg.RandomInt(10, 13)
	return rg.RandomString(int(length))
}

func (rg *randomGenerator) RandomName() string {
	return rg.RandomString(10)
}

func (rg *randomGenerator) RandomAddress() string {
	return rg.RandomString(10)
}

func (rg *randomGenerator) RandomAge() uint {
	// Generating a random age between 1 and 100
	return uint(rg.RandomInt(1, 100))
}

func (rg *randomGenerator) RandomGender() string {
	// Assuming gender can be "Male", "Female", or "Other"
	genders := []string{"Male", "Female", "Other"}
	return genders[rg.rand.Intn(len(genders))]
}
