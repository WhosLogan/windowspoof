package main

import "testing"

func BenchmarkGetWindows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetWindows()
	}
}
