package editor

import (
	"fmt"
	"strings"
)

// Edits the expression by removing excess symbols
func InputFilter(input string) string {
	input = strings.TrimLeft(input, "*/+ ")
	input = strings.TrimRight(input, "*/+- ")
	input = strings.ReplaceAll(input, "--", "+")
	input = strings.ReplaceAll(input, "+-", "-")
	return input
}

// If piece contains "*" or "+", puts "\"" before them
func MakeRegExp(piece string) string {
	for _, symb := range []string{"*", "+", "(", ")"} {
		piece = strings.ReplaceAll(piece, symb, `\`+symb)
	}
	return piece
}

// Checks if given symbol is "*" or "+" or "/" or "-"
func isSign(symbol string) bool {
	var pool = []string{"+", "-", "*", "/"}
	for _, elem := range pool {
		if elem == symbol {
			return true
		}
	}
	return false
}

// Replaces unit with the result of unit operation
func ExpEdit(exp string, result float64, index []int) string {
	prevExp := exp
	var tempResult string

	tempResult = fmt.Sprintf("%f", result)
	exp = prevExp[:index[0]]
	if index[0] > 0 && !isSign(string(prevExp[index[0]-1])) {
		exp += "+"
	}
	exp += tempResult
	if len(prevExp) >= index[1] {
		exp = exp + prevExp[index[1]:]
	}
	return exp
}
