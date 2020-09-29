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
	_, err1 := fmt.Scanf("%d\n%d", &base, &height)

	if err1 == nil {
		fmt.Printf("%f\n", float32(base)*float32(height)/2.0)
	}
}