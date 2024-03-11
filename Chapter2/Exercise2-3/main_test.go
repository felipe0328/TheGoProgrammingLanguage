package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(255)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(255)
	}
}

func BenchmarkPopCountComplete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoopComplete(255)
	}
}
