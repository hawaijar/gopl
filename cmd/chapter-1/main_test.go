package main

import (
	"fmt"
	"testing"
)

func BenchmarkExercise11(b *testing.B) {
	fmt.Println("Exercise11 benchmark...")
	for i := 0; i < b.N; i++ {
		Exercise11()
	}
	fmt.Println(" ===================== ")

	for i := 0; i < b.N; i++ {
		Exercise12()
	}
}
