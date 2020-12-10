package utils

import (
	"math/rand"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func GenerateOrderedArray(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = i
	}
	return ints
}

func GenerateRandomArray(n, scope int) []int {
	ints := make([]int, n)
	if scope == 1 {
		for i := 0; i < n; i++ {
			ints[i] = 1
		}
	} else {
		for i := 0; i < n; i++ {
			ints[i] = rand.Intn(scope)
		}
	}
	return ints
}
