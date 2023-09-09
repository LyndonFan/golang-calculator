package main

import (
	"reflect"
	"testing"
)

func TestIsDigit(t *testing.T) {
    testCases := []struct {
        input byte
        expected bool
    }{
        {byte('0'), true},
        {byte('5'), true},
        {byte('9'), true},
        {byte('a'), false},
        {byte('A'), false},
        {byte(' '), false},
        {byte('+'), false},
        {byte('-'), false},
    }

    for _, tc := range testCases {
        result := isDigit(tc.input)
        if result != tc.expected {
            t.Errorf("Expected isDigit(%c) to be %v, but got %v", tc.input, tc.expected, result)
        }
    }
}

func CreateTokenizeTest(input string, expectedTokens []string) func(*testing.T) {
	cache := NewCache()
	return func(t *testing.T) {
		tokens, err := tokenize(input, cache)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if !reflect.DeepEqual(tokens, expectedTokens) {
			t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
		}
	}
}

func CreateTokenizeTestExpectError(input string, expectedError error) func(*testing.T) {
	cache := NewCache()
	return func(t *testing.T) {
		tokens, err := tokenize(input, cache)
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
