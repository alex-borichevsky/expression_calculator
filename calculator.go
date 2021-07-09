package expression_calculator
import (
	"errors"
	"fmt"
)

var digits = map[rune]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

type token rune

func isFunc(token token) bool {
	return (token == operatorPlus() || token == operatorMinus())
}
func isDigit(token token) bool {
	_, ok := digits[rune(token)]
	return ok
}
func operatorPlus() token {
	return '+'
}
func operatorMinus() token {
	return '-'
}

type nextTokener interface {
	nextToken() (token, error)
}
type parser struct {
	runes    []rune
	position int
}

func (parser parser) length() int {
	return len(parser.runes)
}
func (parser *parser) nextToken() (token, error) {
	if parser.position == parser.length() {
		return ' ', errors.New("EOF")
	}
	var current = parser.runes[parser.position]
	parser.position++
	return token(current), nil
}

type Expression struct {
	operator token
	value    int
	result   int
}

func (expression Expression) calculate() int {
	if expression.operator == operatorPlus() {
		expression.result += expression.value
	} else if expression.operator == operatorMinus() {
		expression.result -= expression.value
	}
	return expression.result
}

func (expression *Expression) process(n nextTokener) (int, error) {
	var previousOperator token
	expression.operator = operatorPlus()
	var err error
	var t token
	for i := 0; err == nil; i++ {
		t, err = n.nextToken()
		if err != nil {
			break
		}
		if isFunc(t) {
			if t == previousOperator {
				return 0, fmt.Errorf("unexpected token at pos: %v", i)
			}
		}
		previousOperator = t
		expression.apply(t)
	}
	return expression.result, nil
}

func (expression *Expression) apply(t token) (int, error) {
	if isFunc(t) {
		expression.operator = t
	}
	if isDigit(t) {
		expression.value = digits[rune(t)]
		expression.result = expression.calculate()
	}
	return expression.result, nil
}
func Evaluate(input string) (int, error) {
	expression := Expression{}
	p := parser{runes: []rune(input)}
	return expression.process(&p)
}