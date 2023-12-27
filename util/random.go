package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(minVal, maxVal int64) int64 {
	return minVal + rand.Int63n(maxVal-minVal+1)
}

func RandString(length int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var builder strings.Builder

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(alphabet))
		builder.WriteByte(alphabet[randomIndex])
	}

	return builder.String()
}

func RandName() string {
	availableName := []string{"Jay", "John", "David", "Leon"}
	randomIndex := rand.Intn(len(availableName))
	return availableName[randomIndex]
}

func RandEmail() string {
	return RandName() + strconv.Itoa(int(RandInt(100, 999))) + "@gmail.com"
}

func RandPassword() string {
	return RandString(15)
}

func RandGender() string {
	availableGender := []string{"Male", "Female", "Other"}
	randomIndex := rand.Intn(len(availableGender))
	return availableGender[randomIndex]
}

func RandUniversity() string {
	availableUniversity := []string{"University of Waterloo", "Wilfrid Laurier University"}
	randomIndex := rand.Intn(len(availableUniversity))
	return availableUniversity[randomIndex]
}
