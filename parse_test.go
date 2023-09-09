package main

import (
	"reflect"
	"testing"
)

func CreateConvertToRPNTest(input, expected []string) func(*testing.T) {
    cache := NewCache()
    return func(t *testing.T) {
        tokens, err := convertToRPN(input, cache)
        if err != nil {
            t.Errorf("Error: %s", err)
        }
        if !reflect.DeepEqual(tokens, expected) {
            t.Errorf("Expected tokens %v, but got %v", expected, tokens)
        }
    }
}

func CreateConvertToRPNWithCacheTest(input, expected []string, cache *Cache) func(*testing.T) {
    return func(t *testing.T) {
        tokens, err := convertToRPN(input, cache)
        if err != nil {
            t.Errorf("Error: %s", err)
        }
        if !reflect.DeepEqual(tokens, expected) {
            t.Errorf("Expected tokens %v, but got %v", expected, tokens)
        }
    }
}


// Test the convertToRPN function with a valid input containing multiple operators and operands
func TestValidInputWithMultipleOperatorsAndOperands(t *testing.T) {
    input := []string{"2", "+", "3", "*", "4", "-", "5"}
    expected := []string{"2", "3", "4", "*", "+", "5", "-"}
    CreateConvertToRPNTest(input, expected)(t)
}

// Test the behavior of the convertToRPN function when the input contains only one operand
func test_input_with_only_one_operand(t *testing.T) {
    input := []string{"5"}
    expected := []string{"5"}
    CreateConvertToRPNTest(input, expected)(t)
}


// Test the behavior of the convertToRPN function when the input contains only one operator.
func TestInputWithOnlyOneInfixOperator(t *testing.T) {
    for _, operator := range []string{"+", "-", "*", "/", "%", "^"} {
        CreateConvertToRPNTest([]string{"6",operator,"2"}, []string{"6", "2", operator})(t)
    }
}


// Test if the function handles input with decimal numbers correctly
func test_input_with_decimal_numbers(t *testing.T) {
    input := []string{"2.5", "+", "3.7", "*", "4.2"}
    expected := []string{"2.5", "3.7", "4.2", "*", "+"}
    CreateConvertToRPNTest(input, expected)(t)
}


// Test the behavior of the convertToRPN function when given an input with multiple parentheses
func TestInputWithMultipleParentheses(t *testing.T) {
    input := []string{"(", "1", "+", "2", ")", "*", "(", "3", "-", "4", ")"}
    expected := []string{"1", "2", "+", "3", "4", "-", "*"}
    CreateConvertToRPNTest(input, expected)(t)
}


// Test if the function returns an error when there are unmatched parentheses in the input
func test_input_with_unmatched_parentheses(t *testing.T) {
    input := []string{"(", "2", "+", "3", ")"}
    cache := NewCache()
    _, err := convertToRPN(input, cache)
    if err == nil {
        t.Errorf("Expected an error, but got nil")
    }
}


// Test if the convertToRPN function returns an error for invalid tokens
func test_invalid_tokens(t *testing.T) {
    input := []string{"2", "..", "3"}
    cache := NewCache()
    _, err := convertToRPN(input, cache)
    if err == nil {
        t.Errorf("Expected an error, but got nil")
    }
}


// Test if the function handles input with consecutive operators correctly
func test_input_with_consecutive_operators(t *testing.T) {
    input := []string{"2", "+", "+", "3"}
    cache := NewCache()
    result, err := convertToRPN(input, cache)
    if err == nil {
        t.Errorf("Expected error, got nil")
    }
    expectedError := "Invalid token: +"
    if err.Error() != expectedError {
        t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
    }
    if len(result) != 0 {
        t.Errorf("Expected empty result, got %v", result)
    }
}


// Test the behaviour of the convertToRPN function when given input with negative numbers
func test_input_with_negative_numbers(t *testing.T) {
    input := []string{"2", "+", "-3"}
    expected := []string{"2", "-3", "+"}
    CreateConvertToRPNTest(input, expected)(t)
}


// Test if the function correctly handles input with constants
func test_input_with_constants(t *testing.T) {
    input := []string{"3.14", "pi", "e"}
    expected := []string{"3.14", "pi", "e"}
    CreateConvertToRPNTest(input, expected)(t)
}


// Test the behavior of the convertToRPN function when the input contains the exponentiation operator (^)
func test_input_with_exponentiation_operator(t *testing.T) {
    input := []string{"2", "^", "3"}
    expected := []string{"2", "3", "^"}
    CreateConvertToRPNTest(input, expected)(t)
}


// Test the behaviour of the convertToRPN function when given input with the modulo operator (%)
func TestInputWithModuloOperator(t *testing.T) {
    input := []string{"5", "+", "4", "%", "2"}
    expected := []string{"5", "4", "2", "%", "+"}
    CreateConvertToRPNTest(input, expected)(t)
}


// Test the behavior of the convertToRPN function when given an input with a right-associative operator
func test_input_with_right_associative_operator(t *testing.T) {
    input := []string{"2", "^", "3", "^", "4"}
    expected := []string{"2", "3", "4", "^", "^"}
    CreateConvertToRPNTest(input, expected)(t)
}


// Test the behavior of the convertToRPN function when given input with cache hit and miss
func test_input_with_cache_hit_and_miss(t *testing.T) {
    // Create a new cache
    cache := NewCache()
    cache.Set("x", 10)
    input := []string{"x", "+", "y"}
    expected := []string{"x", "y", "+"}
    CreateConvertToRPNWithCacheTest(input, expected, cache)(t)
}


