package sorting

import (
	"math/rand"
	"testing"
)

func generateIntegers(n int) []interface{} {

	ints := make([]interface{}, n)

	for i := 0; i < n; i++ {
		ints[i] = (rand.Int() % 1000) + 1
	}

	return ints

}

func intGreater(a, b interface{}) bool {
	return a.(int) > b.(int)
}

func BenchmarkBubble(b *testing.B) {

	data := generateIntegers(1000)

	for i := 0; i < 1; i++ {
		Bubble(&data, intGreater)
	}

}
