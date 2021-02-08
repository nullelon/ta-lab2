package lab2

import (
	"fmt"
	"testing"
)

func TestReverseDigitsSequence(t *testing.T) {
	data := map[string]string{
		"12340":  "04321",
		"5550":   "0555",
		"8760":   "0678",
		"976240": "042679",
	}
	for input, r := range data {
		result := ReverseDigitsSequence(input)
		if result != r {
			t.Fatalf("Got %s, expected %s\n", result, r)
		}
		fmt.Printf("%s -> %s\n", input, result)
	}
}

func TestTaskTwoFunction(t *testing.T) {
	m := []int{1, 2, 3, 4, 7}
	n := []int{4, 20, 15, 40, 50}

	for i := range m {
		fmt.Printf("C(%d, %d) = %d\n", m[i], n[i], TaskTwoFunction(m[i], n[i]))
	}
}
