package calculator

import (
	"regexp"
)

type IncorrectInput struct{}

func (m *IncorrectInput) Error() string {
	return "\nIncorrect input!\n"
}

// Checks if the expression is correct
func correctExp(input string) bool {
	re := regexp.MustCompile(`[0-9]+|[[0-9]+[\*/\+-\.]?[0-9]+]*`)
	return re.MatchString(input)
}

// Checks if no operation in the expression
func numberOnly(str string) bool {
	for i, symbol := range str {
		if (string(symbol) == "+" || string(symbol) == "-" || string(symbol) == "/" ||
			string(symbol) == "*") && i != 0 && i != len(str)-1 {
			return false
		}
	}
	return true
}

// Checks if expression contains brackets
func bracketsIn(str string) bool {
	re := regexp.MustCompile(`[)()]+`)
	return re.MatchString(str)
}

// Checks if there is "-" before the number
func IsNegative(exp string) bool {
	return string(exp[0]) == "-"
}
