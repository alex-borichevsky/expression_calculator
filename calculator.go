package expression_calculator
import (
	"errors"
	"fmt"
)

type token rune

func isFunc(token token) bool {
	return (token == operatorPlus() || token == operatorMinus())
}
func digit(t token) (int, bool){
	if t == '0' {return 0, true}
	if t == '1' {return 1, true}
	if t == '2' {return 2, true}
	if t == '3' {return 3, true}
	if t == '4' {return 4, true}
	if t == '5' {return 5, true}
	if t == '6' {return 6, true}
	if t == '7' {return 7, true}
	if t == '8' {return 8, true}
	if t == '9' {return 9, true}
	return ' ',false
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
	val, ok := digit(t)
	if ok{
		expression.value = val
		expression.result = expression.calculate()
	}
	return expression.result, nil
}
func Evaluate(input string) (int, error) {
	expression := Expression{}
	p := parser{runes: []rune(input)}
	return expression.process(&p)
}