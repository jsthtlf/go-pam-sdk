package utils

import "math/rand/v2"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	if n < 1 {
		n = 1
	}
	length := len(letterRunes)
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.IntN(length)]
	}
	return string(b)
}
