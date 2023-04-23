package parser

import "os"

type NoArgsError struct{}

func (m *NoArgsError) Error() string {
	return "\n\nWrong usage!\nCorrect usage: go run calc.go \"(1+2)*3\"\n\n"
}

func ParseExp() (string, error) {
	exp := os.Args

	if len(exp) != 2 {
		return "", &NoArgsError{}
	}
	return exp[1], nil
}
