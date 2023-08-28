package main

import (
	"testing"
)

func CreateTokenizeTest(input string, expectedTokens []string) func(*testing.T) {
	return func(t *testing.T) {
		tokens, err := tokenize(input)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(tokens) != len(expectedTokens) {
			t.Errorf("Expected %d tokens, got %s", len(expectedTokens), tokens)
		}
		for i, token := range tokens {
			if token != expectedTokens[i] {
				t.Errorf("Expected token %d to be %s, got %s", i, expectedTokens[i], token)
			}
		}
	}
}

func CreateTokenizeTestExpectError(input string, expectedError error) func(*testing.T) {
	return func(t *testing.T) {
		tokens, err := tokenize(input)
		if err == nil {
			t.Errorf("Did not get error, but got tokens %s", tokens)
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("Expected error %s, got %s", expectedError, err)
		}
	}
}

func TestTokenizeNormal(t *testing.T) {
	CreateTokenizeTest("", []string{})(t)
	CreateTokenizeTest("1", []string{"1"})(t)
	CreateTokenizeTest("1+2", []string{"1", "+", "2"})(t)
	// not do subtraction as it has its own test
	CreateTokenizeTest("1*2", []string{"1", "*", "2"})(t)
	CreateTokenizeTest("1/2", []string{"1", "/", "2"})(t)
	CreateTokenizeTest("1^2", []string{"1", "^", "2"})(t)
	CreateTokenizeTest("1%2", []string{"1", "%", "2"})(t)
}

func TestTokenizeWithLongNumbers(t *testing.T){
	CreateTokenizeTest("123", []string{"123"})(t)
	CreateTokenizeTest("123+456", []string{"123", "+", "456"})(t)
}

func TestTokenizeWithMinusSign(t *testing.T){
	// negative number
	CreateTokenizeTest("-1", []string{"-1"})(t)
	CreateTokenizeTest("-123", []string{"-123"})(t)
	CreateTokenizeTest("-123+456", []string{"-123", "+", "456"})(t)

	CreateTokenizeTest("-1+-2", []string{"-1", "+", "-2"})(t)
	CreateTokenizeTest("-123+-456", []string{"-123", "+", "-456"})(t)
	CreateTokenizeTest("-123^-2", []string{"-123", "^", "-2"})(t)

	// minus sign
	CreateTokenizeTest("1-2", []string{"1", "-", "2"})(t)
	CreateTokenizeTest("1-2-3", []string{"1", "-", "2", "-", "3"})(t)
	CreateTokenizeTest("123-456", []string{"123", "-", "456"})(t)

	// both
	CreateTokenizeTest("-2-3", []string{"-2", "-", "3"})(t)
	CreateTokenizeTest("-123-456", []string{"-123", "-", "456"})(t)
	CreateTokenizeTest("-78--90", []string{"-78", "-", "-90"})(t)
}

func TestTokenizeWithDecimalPoint(t *testing.T){
	// full numbers
	CreateTokenizeTest("1.2", []string{"1.2"})(t)
	CreateTokenizeTest("12.345", []string{"12.345"})(t)

	// missing leading zero
	CreateTokenizeTest(".1", []string{".1"})(t)
	CreateTokenizeTest(".123", []string{".123"})(t)
	CreateTokenizeTest("-.1", []string{"-.1"})(t)
	CreateTokenizeTest("-.123", []string{"-.123"})(t)

	// missing trailing zero
	CreateTokenizeTest("1.", []string{"1."})(t)
	CreateTokenizeTest("12.", []string{"12."})(t)
	CreateTokenizeTest("-1.", []string{"-1."})(t)
	CreateTokenizeTest("-123.", []string{"-123."})(t)

	// no need to test both
	// as it will error when in rpn
}

func TestTokenizeWithMinusAndPoint(t *testing.T){
	CreateTokenizeTest("1.2-3", []string{"1.2", "-", "3"})(t)
	CreateTokenizeTest("-1.2-3", []string{"-1.2", "-", "3"})(t)
	CreateTokenizeTest("-1.2-3.4", []string{"-1.2", "-", "3.4"})(t)
	CreateTokenizeTest("1.2-.4", []string{"1.2", "-", ".4"})(t)
}
