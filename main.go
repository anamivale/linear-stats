package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// calculatePearsonCorrelation computes the Pearson correlation coefficient
func calculatePearsonCorrelation(xs, ys []float64) float64 {
    n := float64(len(xs))
    sumX, sumY, sumXY, sumX2, sumY2 := 0.0, 0.0, 0.0, 0.0, 0.0

    for i := 0; i < len(xs); i++ {
        sumX += xs[i]
        sumY += ys[i]
        sumXY += xs[i] * ys[i]
        sumX2 += xs[i] * xs[i]
        sumY2 += ys[i] * ys[i]
    }

    numerator := n*sumXY - sumX*sumY
    denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))
    return numerator / denominator
}

func calculateLinearRegression(xs, ys []float64) (float64, float64) {
    n := float64(len(xs))
    sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0

    for i := 0; i < len(xs); i++ {
        sumX += xs[i]
        sumY += ys[i]
        sumXY += xs[i] * ys[i]
        sumX2 += xs[i] * xs[i]
    }

    m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
    b := (sumY - m*sumX) / n
    return m, b
}
func main() {
    if len(os.Args) != 2 {
        log.Fatal("Usage: go run your-program.go data.txt")
    }

    filename := os.Args[1]

    file, err := os.Open(filename)
    if err != nil {
        log.Fatalf("failed to open file: %v", err)
    }
    defer file.Close()

    var ys []float64
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        y, err := strconv.ParseFloat(line, 64)
        if err != nil {
            log.Fatalf("failed to parse line: %v", err)
        }
        ys = append(ys, y)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("error reading file: %v", err)
    }
    if err := scanner.Err(); err != nil {
        fmt.Println(err.Error())
        return
    }

    xs := make([]float64, len(ys))

    for i := range xs{
        xs[i] = float64(i)
    }

    slope, yintercept := calculateLinearRegression(xs, ys)
    fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, yintercept)

    coefcorr := calculatePearsonCorrelation(xs,ys)
    fmt.Printf("Pearson Correlation Coefficient: %.10f\n", coefcorr)
}