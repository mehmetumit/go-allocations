package main

import (
	"testing"
)

func BenchmarkAllocs(b *testing.B) {
	n := 100
	var testByte byte

	b.Run("LotsOfAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buff []byte

			for i := 0; i < n; i++ {
				buff = append(buff, testByte)
			}
		}

	})
	b.Run("AppendAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buff := make([]byte, n)

			for i := 0; i < n; i++ {
				buff = append(buff, testByte)
			}
		}

	})
	b.Run("IndexAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buff := make([]byte, n)

			for i := 0; i < n; i++ {
				buff[i] = testByte
			}
		}

	})
	b.Run("ZeroAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			//n == 100
			var buff [100]byte

			for i := 0; i < n; i++ {
				buff[i] = testByte
			}
		}

	})

}
