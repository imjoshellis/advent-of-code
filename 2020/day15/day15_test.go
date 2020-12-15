package main

import "testing"

func benchmarkGetNth(target int, b *testing.B) {
	initial := []int{0, 12, 6, 13, 20, 1, 17}
	for n := 0; n < b.N; n++ {
		getNth(initial, target)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkGetNth(10, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkGetNth(100, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkGetNth(1000, b) }
func BenchmarkFib10(b *testing.B) { benchmarkGetNth(10000, b) }
func BenchmarkFib20(b *testing.B) { benchmarkGetNth(1000000, b) }
func BenchmarkFib40(b *testing.B) { benchmarkGetNth(10000000, b) }
