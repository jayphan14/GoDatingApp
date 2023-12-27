package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt(minVal, maxVal int64) int64 {
	return minVal + rand.Int63n(maxVal-minVal+1)
}

func randString(length int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var builder strings.Builder

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(alphabet))
		builder.WriteByte(alphabet[randomIndex])
	}

	return builder.String()
}

func randName() string {
	availableName := []string{"Jay", "John", "David", "Leon"}
	randomIndex := rand.Intn(len(availableName))
	return availableName[randomIndex]
}

func randEmail() string {
	return randName() + string(randInt(100, 999)) + "@gmail.com"
}

func randPassword() string {
	return randString(15)
}

func randGender() string {
	availableGender := []string{"Male", "Female", "Other"}
	randomIndex := rand.Intn(len(availableGender))
	return availableGender[randomIndex]
}

func randUniversity() string {
	availableUniversity := []string{"University of Waterloo", "Wilfrid Laurier University"}
	randomIndex := rand.Intn(len(availableUniversity))
	return availableUniversity[randomIndex]
}


