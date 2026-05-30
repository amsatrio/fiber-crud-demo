package util

import "math/rand/v2"

func GenerateWord(min, max int, isNumeric bool) string {
	length := min + rand.IntN(max-min+1)
	b := make([]byte, length)

	var charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if isNumeric {
		charset = "0123456789"
	}

	charsetLen := len(charset)
	for i := range b {
		b[i] = charset[rand.IntN(charsetLen)]
	}

	return string(b)
}

func GenerateNumber(min, max int) int {
	return rand.IntN(max-min+1) + min
}
