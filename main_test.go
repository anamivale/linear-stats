package main

import (
	"math"
	"testing"
)

func TestCalculateLinearRegression(t *testing.T) {
	// Test case 1: Basic case with positive numbers
	xs1 := []float64{0, 1, 2, 3, 4}
	ys1 := []float64{1, 3, 2, 5, 4}
	m1, b1 := calculateLinearRegression(xs1, ys1)
	verifyLinearRegressionResult(t, m1, b1, 0.8, 1.4)

	// Test case 2: Case with negative numbers
	xs2 := []float64{-2, -1, 0, 1, 2}
	ys2 := []float64{-1, -3, -2, -5, -4}
	m2, b2 := calculateLinearRegression(xs2, ys2)
	verifyLinearRegressionResult(t, m2, b2, -0.8, -3.0)

	// Test case 3: Case with floating-point numbers
	xs3 := []float64{0.5, 1.5, 2.5, 3.5, 4.5}
	ys3 := []float64{1.2, 3.4, 2.8, 5.6, 4.3}
	m3, b3 := calculateLinearRegression(xs3, ys3)
	verifyLinearRegressionResult(t, m3, b3, 0.84, 1.36)

	// Test case 4: Case with identical x values
	xs4 := []float64{0, 0, 0, 0, 0}
	ys4 := []float64{1, 2, 3, 4, 5}
	m4, b4 := calculateLinearRegression(xs4, ys4)
	verifyLinearRegressionResult(t, m4, b4, 0.0, 3.0)
}

func verifyLinearRegressionResult(t *testing.T, m, b, expectedM, expectedB float64) {
	epsilon := 1e-6 // Small value to handle floating-point precision issues
	if math.Abs(m-expectedM) > epsilon || math.Abs(b-expectedB) > epsilon {
		t.Errorf("Expected m = %.6f, b = %.6f, but got m = %.6f, b = %.6f", expectedM, expectedB, m, b)
	}
}

func TestCalculatePearsonCorrelation(t *testing.T) {
	// Test case 1: Basic case with positive numbers
	xs1 := []float64{0, 1, 2, 3, 4}
	ys1 := []float64{1, 3, 2, 5, 4}
	r1 := calculatePearsonCorrelation(xs1, ys1)
	verifyPearsonCorrelationResult(t, r1, 0.8000000000)

	// Test case 2: Case with negative numbers
	xs2 := []float64{-2, -1, 0, 1, 2}
	ys2 := []float64{-1, -3, -2, -5, -4}
	r2 := calculatePearsonCorrelation(xs2, ys2)
	verifyPearsonCorrelationResult(t, r2, -0.8000000000)

	// Test case 3: Case with floating-point numbers
	xs3 := []float64{0.5, 1.5, 2.5, 3.5, 4.5}
	ys3 := []float64{1.2, 3.4, 2.8, 5.6, 4.3}
	r3 := calculatePearsonCorrelation(xs3, ys3)
	verifyPearsonCorrelationResult(t, r3, 0.8070955641)

	// Test case 4: Case with identical x values
	xs4 := []float64{0, 0, 0, 0, 0}
	ys4 := []float64{1, 2, 3, 4, 5}
	r4 := calculatePearsonCorrelation(xs4, ys4)
	verifyPearsonCorrelationResult(t, r4, 0.0)
}

func verifyPearsonCorrelationResult(t *testing.T, r, expectedR float64) {
	epsilon := 1e-6 // Small value to handle floating-point precision issues
	if math.Abs(r-expectedR) > epsilon {
		t.Errorf("Expected r = %.10f, but got r = %.10f", expectedR, r)
	}
}


