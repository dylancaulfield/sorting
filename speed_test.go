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

var data = generateIntegers(100000)

func intGreater(a, b interface{}) bool {
	return a.(int) > b.(int)
}

func BenchmarkBubble(b *testing.B) {

	localData := data

	for i := 0; i < b.N; i++ {
		Bubble(&localData, intGreater)
	}

}
func BenchmarkMerge(b *testing.B) {

	localData := data

	for i := 0; i < b.N; i++ {
		Merge(&localData, intGreater)
	}

}
