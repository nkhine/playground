package main

import (
	"fmt"

	"github.com/nkhine/go-tutorials/go-tests/math"
)

func main() {
	sum := math.Sum([]int{10, -2, 3})
	if sum != 11 {
		msg := fmt.Sprintf("Fail: Wanted 11 but received %d", sum)
		panic(msg)
	}
	fmt.Println("PASS")
}
