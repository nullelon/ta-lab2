package lab2

import (
	"fmt"
	"math/rand"
	"testing"
)

var reversionData = map[string]string{
	"12340":  "04321",
	"5550":   "0555",
	"8760":   "0678",
	"976240": "042679",
}

func TestRecursiveReversion_ReversedNums(t *testing.T) {
	a := RecursiveReversion{}
	for input, r := range reversionData {
		result := a.ReversedNums(input)
		if result != r {
			t.Fatalf("Got %s, expected %s\n", result, r)
		}
		fmt.Printf("%s -> %s\n", input, result)
	}
}

func TestIteratedReversion_ReversedNums(t *testing.T) {
	a := IteratedReversion{}
	for input, r := range reversionData {
		result := a.ReversedNums(input)
		if result != r {
			t.Fatalf("Got %s, expected %s\n", result, r)
		}
		fmt.Printf("%s -> %s\n", input, result)
	}
}

func TestRandomImplementation(t *testing.T) {
	var reverse IReverse
	if rand.Intn(2) == 0 {
		reverse = RecursiveReversion{}
		fmt.Println("Using RecursiveReversion implementation...")
	} else {
		reverse = IteratedReversion{}
		fmt.Println("Using IteratedReversion implementation...")
	}

	for input, r := range reversionData {
		result := reverse.ReversedNums(input)
		if result != r {
			t.Fatalf("Got %s, expected %s\n", result, r)
		}
		fmt.Printf("%s -> %s\n", input, result)
	}
}

func TestTaskTwoFunction(t *testing.T) {
	m := []int{2, 2, 3, 4, 5}
	n := []int{4, 20, 15, 40, 50}
	result := []int{6, 190, 455, 91390, 2118760}

	for i := range m {
		r := C(m[i], n[i])
		if r != result[i] {
			t.Fatalf("Expected %d, got %d", result[i], r)
		}
		fmt.Printf("C(%d, %d) = %d\n", m[i], n[i], r)
	}
}
