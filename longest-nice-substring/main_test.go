package main

import "testing"

func BenchmarkLongestNiceSubstring(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		longestNiceSubstring("hiIWIllOlRbOhXZbZZlNXxIIIOBIoioBCXlzbbioIIOOiwHhIozhblolIxoiXoOziLo")
	}
}
