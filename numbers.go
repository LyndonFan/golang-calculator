package main

import "math"

func getConstants() map[string]float64 {
	consts := map[string]float64{
		"pi": math.Pi,
		"e": math.E,
	}
	return consts
}