package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func CreateConvertToRPNTest(input, expected []string, cache *Cache) func(*testing.T) {
    if cache == nil {
        cache = NewCache()
    }
    return func(t *testing.T) {
        tokens, err := convertToRPN(input, cache)
        if err != nil {
            t.Errorf("Error: %s", err)
        } else if !reflect.DeepEqual(tokens, expected) {
            t.Errorf("Expected tokens %v, but got %v", expected, tokens)
        }
    }
}

func CreateConvertToRPNExpectErrorTest(input []string, expectedError error, cache *Cache) func(*testing.T) {
    if cache == nil {
        cache = NewCache()
    }
    return func(t *testing.T) {
        tokens, err := convertToRPN(input, cache)
        if err == nil {
            t.Errorf("Expected an error, but got: %s", tokens)
        }
        if errors.Is(err, expectedError) {
            t.Errorf("Expected error %v, but got %v", expectedError, err)
        }
    }
}

// Test the convertToRPN function with a valid input containing multiple operators and operands
func TestValidInputWithMultipleOperatorsAndOperands(t *testing.T) {
    input := []string{"2", "+", "3", "*", "4", "-", "5"}
    expected := []string{"2", "3", "4", "*", "+", "5", "-"}
    CreateConvertToRPNTest(input, expected, nil)(t)
}

// Test the behavior of the convertToRPN function when the input contains only one operand
func TestInputWithOnlyOneOperand(t *testing.T) {
    input := []string{"5"}
    expected := []string{"5"}
    CreateConvertToRPNTest(input, expected, nil)(t)
}


// Test the behavior of the convertToRPN function when the input contains only one operator.
func TestInputWithOnlyOneInfixOperator(t *testing.T) {
    for _, op := range []string{"+", "-", "*", "/", "%", "^"} {
        CreateConvertToRPNTest([]string{"6",op,"2"}, []string{"6", "2", op}, nil)(t)
    }
}


// Test if the function handles input with decimal numbers correctly
func TestInputWithDecimalNumbers(t *testing.T) {
    input := []string{"2.5", "+", "3.7", "*", "4.2"}
    expected := []string{"2.5", "3.7", "4.2", "*", "+"}
    CreateConvertToRPNTest(input, expected, nil)(t)
}


// Test the behavior of the convertToRPN function when given an input with multiple parentheses
func TestInputWithMultipleParentheses(t *testing.T) {
    input := []string{"(", "1", "+", "2", ")", "*", "(", "3", "-", "4", ")"}
    expected := []string{"1", "2", "+", "3", "4", "-", "*"}
    CreateConvertToRPNTest(input, expected, nil)(t)
}


// Test if the function returns an error when there are unmatched parentheses in the input
func TestInputWithUnmatchedParentheses(t *testing.T) {
    input := []string{"(", "2", "+", "3"}
    cache := NewCache()
    _, err := convertToRPN(input, cache)
    if err == nil {
        t.Errorf("Expected an error, but got nil")
    }
}


// Test if the convertToRPN function returns an error for invalid tokens
func TestInvalidTokens(t *testing.T) {
    input := []string{"2", "..", "3"}
    cache := NewCache()
    _, err := convertToRPN(input, cache)
    if err == nil {
        t.Errorf("Expected an error, but got nil")
    }
}


// Test if the function handles input with consecutive operators correctly
func TestInputWithConsecutiveOperators(t *testing.T) {
    input := []string{"2", "+", "+", "3"}
    expectedError := fmt.Errorf("Invalid token: +")
    CreateConvertToRPNExpectErrorTest(input, expectedError, nil)(t)
}


// Test the behaviour of the convertToRPN function when given input with negative numbers
func TestInputWithNegativeNumbers(t *testing.T) {
    input := []string{"2", "+", "-3"}
    expected := []string{"2", "-3", "+"}
    CreateConvertToRPNTest(input, expected, nil)(t)
}


// Test if the function correctly handles input with constants
func TestInputWithConstants(t *testing.T) {
    input := []string{"pi", "+", "3", "*", "e"}
    expected := []string{"pi", "3", "e", "*", "+"}
    CreateConvertToRPNTest(input, expected, nil)(t)
}


// Test the behavior of the convertToRPN function when the input contains the exponentiation operator (^)
func TestInputWithExponentiationOperator(t *testing.T) {
    input := []string{"2", "^", "3"}
    expected := []string{"2", "3", "^"}
    CreateConvertToRPNTest(input, expected, nil)(t)
}


// Test the behaviour of the convertToRPN function when given input with the modulo operator (%)
func TestInputWithModuloOperator(t *testing.T) {
    input := []string{"5", "+", "4", "%", "2"}
    expected := []string{"5", "4", "2", "%", "+"}
    CreateConvertToRPNTest(input, expected, nil)(t)
}


// Test the behavior of the convertToRPN function when given an input with a right-associative operator
func TestInputWithRightAssociativeOperator(t *testing.T) {
    input := []string{"2", "^", "3", "^", "4"}
    expected := []string{"2", "3", "4", "^", "^"}
    CreateConvertToRPNTest(input, expected, nil)(t)
}


// Test the behavior of the convertToRPN function when given input with cache hit and miss
func TestInputWithCacheHitAndMiss(t *testing.T) {
    // Create a new cache
    cache := NewCache()
    cache.Set("x", 10)
    input := []string{"x", "+", "y"}
    expectedError := fmt.Errorf("Invalid token: y")
    CreateConvertToRPNExpectErrorTest(input, expectedError, cache)(t)
}


