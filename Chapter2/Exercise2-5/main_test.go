package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(18446744073709551615)
	}
}

func BenchmarkPopCountClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearing(18446744073709551615)
	}
}
