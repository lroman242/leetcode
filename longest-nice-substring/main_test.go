package main

import "testing"

func BenchmarkLongestNiceSubstring(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		longestNiceSubstring("wWOExoVhvXebB")
		longestNiceSubstring("ijIJwuUnW")
		longestNiceSubstring("YazaAay")
		longestNiceSubstring("hiIWIllOlRbOhXZbZZlNXxIIIOBIoioBCXlzbbioIIOOiwHhIozhblolIxoiXoOziLo")
	}
}

func BenchmarkLongestNiceSubstring1(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		longestNiceSubstring1("wWOExoVhvXebB")
		longestNiceSubstring1("ijIJwuUnW")
		longestNiceSubstring1("YazaAay")
		longestNiceSubstring1("hiIWIllOlRbOhXZbZZlNXxIIIOBIoioBCXlzbbioIIOOiwHhIozhblolIxoiXoOziLo")
	}
}
