package main

import "testing"

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse([]byte("This is the Ý test data."))
	}
}
