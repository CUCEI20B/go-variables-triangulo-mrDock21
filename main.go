package main

import (
	"fmt"
)

func psedoTernary(expression bool, trueRes string, falseRes string) string {
	res := ""
	if expression {
		res = trueRes
	} else {
		res = falseRes
	}
	return res
}

func main() {
	var base, height int = 1, 1
	// Read from console
	num_args, err1 := fmt.Scanf("%d", &base)
	_, err2 := fmt.Scanf("%d", &height)

	if err1 == nil && err2 == nil {
		fmt.Printf("%f\n", float32(base)*float32(height)/2.0)
	}
}