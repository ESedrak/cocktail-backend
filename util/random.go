package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName generates a random name string
func RandomNameString() string {
	return RandomString(15)
}

// RandomQty generates a random quantity amount
func RandomQty() int64 {
	return RandomInt(0, 500)
}

// RandomMeasurement generates a random measurement type
func RandomMeasurement() string {
	measurements := []string{"mls", "oz", "spoons", "cup", "tbsp", "tsp"}
	n := len(measurements)
	return measurements[rand.Intn(n)]
}
