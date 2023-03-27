package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Println("Enter the formula to calculate or '#' to exit:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		if line == "#" {
			break
		}

		elements := strings.Split(line, " ")
		if len(elements) != 3 {
			log.Println(`not enough arguments: use " " between operands and operator`)
			continue
		}

		a, err := strconv.Atoi(elements[0])
		if err != nil {
			log.Println("first operand error string to int", err)
			continue
		}

		b, err := strconv.Atoi(elements[2])
		if err != nil {
			log.Println("second operand error string to int", err)
			continue
		}

		operation := elements[1]

		var result int
		switch operation {
		case "+":
			result = Add(a, b)
		case "-":
			result = Subtract(a, b)
		case "*":
			result = Multiply(a, b)
		case "/":
			result = Divide(a, b)
		default:
			log.Println("invalid operation", operation)
			continue
		}
		fmt.Println(result)
	}
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	if b == 0 {
		log.Fatal("integer divide by zero error")
		return 0
	}
	return a / b
}
