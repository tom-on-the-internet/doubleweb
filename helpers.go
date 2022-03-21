package main

import "math/rand"

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	bytes := make([]rune, n)

	for i := range bytes {
		bytes[i] = letters[rand.Intn(len(letters))]
	}

	return string(bytes)
}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}

	return num
}
