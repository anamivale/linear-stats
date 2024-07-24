package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

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

    
}