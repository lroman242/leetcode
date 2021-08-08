package main

import (
	"testing"
)

func BenchmarkTest1(b *testing.B) {
	var input1 = [][]int{{1,0,0,0},{0,0,0,0},{0,0,2,-1}}

	for i := 0; i < b.N; i++ {
		result1 := uniquePathsIII(input1)
		if result1 != 2 {
			b.Error("unexpected result1")
		}
	}
}

func BenchmarkTest2(b *testing.B) {
	var input2 = [][]int{{1,0,0,0},{0,0,0,0},{0,0,0,2}}

	for i := 0; i < b.N; i++ {
		result2 := uniquePathsIII(input2)
		if result2 != 4 {
			b.Error("unexpected result2")
		}
	}
}

func BenchmarkTest3(b *testing.B) {
	var input3 = [][]int{{0,1},{2,0}}

	for i := 0; i < b.N; i++ {
		result3 := uniquePathsIII(input3)
		if result3 != 0 {
			b.Error("unexpected result3")
		}
	}
}
