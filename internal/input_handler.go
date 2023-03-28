package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputHandler() (int, int, string, string, error) {
	fmt.Println("Enter the formula to calculate or '#' to exit:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	if line == "#" {
		return 0, 0, "#", "", nil
	}

	elements := strings.Split(line, " ")
	if len(elements) != 3 {
		return 0, 0, "", "", errors.New("not enough argumnets error")
	}

	err := ValidateFormula(elements[0], elements[2], elements[1])
	if err != nil {
		return 0, 0, "", "", err
	}
	var a, b int
	var operandsType string
	if isRoman(elements[0]) {
		a = ConvertRomanToArab(elements[0])
		b = ConvertRomanToArab(elements[2])
		operandsType = "roman"
	} else {
		a, err = strconv.Atoi(elements[0])
		if err != nil {
			return 0, 0, "", "", errors.New(fmt.Sprintf("first operand convertation from string to int error: %s", err))
		}

		b, err = strconv.Atoi(elements[2])
		if err != nil {
			return 0, 0, "", "", errors.New(fmt.Sprintf("second operand convertation from string to int error: %s", err))
		}
		operandsType = "arabic"
	}

	operation := elements[1]

	return a, b, operation, operandsType, nil
}

func ValidateFormula(a, b, oper string) error {
	if !ValidOperation(oper) { // validate operator
		return errors.New(fmt.Sprintf("invalid operation %s, only '+', '-', '/' and '*' are available", oper))
	}
	if isRoman(a) != isRoman(b) {
		return errors.New("numbers must be both romans or both arabic")
	}
	return nil
}

func ValidOperation(oper string) bool {
	return oper == "+" || oper == "-" || oper == "*" || oper == "/"
}
