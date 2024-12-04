package random

import "math/rand"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charsetLen = len(charset)

func GetRandomString(length int) string {
	chars := make([]byte, length)
	for i := 0; i < length; i++ {
		chars[i] = charset[rand.Intn(charsetLen)]
	}
	return string(chars)
}
