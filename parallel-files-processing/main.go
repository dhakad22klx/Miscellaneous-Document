package main

import (
	"fmt"
	concurrent "parallelfileprocessing/concurrent"
	sequential "parallelfileprocessing/sequential"
)

// I want to find duplicate files based on their content.

// Use secure hash, because names /dates may differ.

func main() {
	fmt.Println("Hello, World!")
	// Sequential
	sequential.Sequential()
	// Concurrent - like a map reduce
	concurrent.Concurrent()
}
