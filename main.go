package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	l := len(arr)
	fmt.Println(arr[l/2:])
	fmt.Println(arr[:l/2])

}
