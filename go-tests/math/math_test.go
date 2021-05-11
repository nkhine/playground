package math

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	fmt.Println("running tests...")
	// sum := math.Sum()
	sum := Sum([]int{10, -2, 3})
	if sum != 11 {
		msg := fmt.Sprintf("Wanted 11 but received %d", sum)
		t.Errorf(msg)
	}
}
