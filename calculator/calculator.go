package calculator

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/lecrank/calc/editor"
)

// Returns the result of given expression (without brackets)
func calcExpression(exp string) (float64, error) {
	var index []int

	exp = editor.InputFilter(exp)

	if !correctExp(exp) {
		return 0, &IncorrectInput{}
	}

	for !numberOnly(exp) {
		res, err := calcUnit(exp, &index)
		if err != nil {
			return 0, err
		}
		exp = editor.ExpEdit(exp, res, index)
	}
	result, err := strconv.ParseFloat(exp, 64)
	if err != nil {
		return result, &IncorrectInput{}
	}
	return result, err
}

// Looks for "*" and "+" operations called "prime". If not found, looks for "+" and "-". Returns [<unit>, <sign>]
func extractOperation(exp string) ([]string, error) {
	rePrime := regexp.MustCompile(`-?[0-9]*\.?[0-9]+([/\*])-?[0-9]*\.?[0-9]+`)
	re := regexp.MustCompile(`-?[0-9]*\.?[0-9]+([-\+])-?[0-9]*\.?[0-9]+`)

	if primeResult := rePrime.FindAllStringSubmatch(exp, 1); len(primeResult) == 0 {
		if result := re.FindAllStringSubmatch(exp, 1); len(result) == 0 {
			return []string{"", ""}, &IncorrectInput{}
		} else {
			return result[0], nil
		}
	} else {
		return primeResult[0], nil
	}
}

// Parses given expression and returns found unit and the sign
func parseUnit(exp string) (string, string, error) {

	operation, err := extractOperation(exp)
	if err != nil {
		return exp + "+0", "+", err
	}

	unitExp := operation[0]
	sign := operation[1]

	return unitExp, sign, nil
}

// Calculates the result of the operation between 2 floats according to given sign
func calc(n1 float64, n2 float64, sign string) (float64, error) {
	switch sign {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		return n1 / n2, nil
	case "end":
		return n1, nil
	}
	return -1, &IncorrectInput{}
}

// Finds the position of unit in whole expression
func findIndex(content string, piece string) []int {
	re := regexp.MustCompile(editor.MakeRegExp(piece))
	loc := re.FindStringIndex(content)
	return loc
}

// Processes []string{"float", "float"} -> float, float
func findNumbers(strNumbers []string, fNeg bool, sNeg bool) (float64, float64, error) {
	num1, err := strconv.ParseFloat(strNumbers[0], 64)
	if err != nil {
		return 0, 0, &IncorrectInput{}
	}
	if fNeg {
		num1 *= -1
	}
	num2, err := strconv.ParseFloat(strNumbers[1], 64)
	if sNeg {
		num2 *= -1
	}
	if err != nil {
		return 0, 0, &IncorrectInput{}
	}
	return num1, num2, nil
}

// Calculates one operation
func calcUnit(exp string, pieceIndex *[]int) (float64, error) {
	var firstNeg, secondNeg bool

	unit, sign, err := parseUnit(exp)
	*pieceIndex = findIndex(exp, unit)

	if IsNegative(unit) {
		firstNeg = true
		unit = unit[1:]
	}

	if err != nil {
		return 0, err
	}
	strNumbers := strings.Split(unit, sign)

	if IsNegative(strNumbers[1]) {
		secondNeg = true
		strNumbers[1] = strNumbers[1][1:]
	}
	num1, num2, err := findNumbers(strNumbers, firstNeg, secondNeg)

	result, err := calc(num1, num2, sign)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// Returnes the string between brackets
func parseBrackets(exp string) (string, error) {
	var found bool

	ind1, ind2 := -1, -1
	l := len(exp)
	for i := 0; i < l; i++ {
		if string(exp[i]) == "(" {
			ind1 = i
		}
	}
	for i := ind1 + 1; i < l; i++ {
		if string(exp[i]) == ")" && !found {
			ind2 = i
			break
		}
	}
	if ind1 < 0 || ind2 < 0 {
		return "", &IncorrectInput{}
	}
	return exp[(ind1 + 1):ind2], nil
}

// The main func that returns actual result
func GetResult(exp string) (float64, error) {
	for bracketsIn(exp) {
		innerExp, err := parseBrackets(exp)
		if err != nil {
			return 0, err
		}
		index := findIndex(exp, "("+innerExp+")")
		tempResult, err := calcExpression(innerExp)
		if err != nil {
			return 0, err
		}
		exp = editor.ExpEdit(exp, tempResult, index)
		exp = editor.InputFilter(exp)
	}
	return calcExpression(exp)
}
