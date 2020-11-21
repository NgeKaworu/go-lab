package main

import (
	"fmt"
	"math"
)

func main() {
	for subLength := 1; subLength < math.MaxInt8; subLength <<= 1 {
		fmt.Println(subLength)
	}

}
