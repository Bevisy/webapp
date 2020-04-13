package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	_, err := Sqrt(-1)
	fmt.Println(err)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	} else if f == 0 {
		return 0, errors.New("")
	} else {
		return math.Sqrt(f), errors.New("")
	}
}
