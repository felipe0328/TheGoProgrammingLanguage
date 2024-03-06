package main

import "testing"

var testArray []string = []string{"1", "2", "3", "4", "5"}

func BenchmarkStringsImplementation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		joinsImplementation(testArray)
	}
}

func BenchmarkCustomImplementation(b *testing.B) {
	customImplementation(testArray)
}

func BenchmarkCustomImplementationWithBytes(b *testing.B) {
	customImplementationWithBytes(testArray)
}
