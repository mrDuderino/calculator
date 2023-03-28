package main

import (
	"fmt"
	"github.com/mrDuderino/calculator/internal"
)

func main() {
	a, b, operation, operandsType, err := internal.InputHandler()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	var result int
	switch operation {
	case "+":
		result = internal.Add(a, b)
	case "-":
		result = internal.Subtract(a, b)
	case "*":
		result = internal.Multiply(a, b)
	case "/":
		result = internal.Divide(a, b)
	default:
		return
	}

	if operandsType == "roman" {
		if result < 1 {
			fmt.Println("unknown result: roman numbers can be only positive (greater or equal to one)")
		} else {
			fmt.Println(internal.ConvertArabToRoman(result))
		}
	} else {
		fmt.Println(result)
	}
}
