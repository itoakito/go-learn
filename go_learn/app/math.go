package main

import (
	"fmt"

	"math"
)

func main() {
	var a float64 = 55
	var b float64 = 3
	c := math.Pow(a,b)
	fmt.Println(a,b,c)
}