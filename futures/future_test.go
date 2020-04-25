package futures

import (
	"testing"
	"time"
)

func BenchmarkIntFuture(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		f := CreateIntFuture(func() int {
			time.Sleep(5 * time.Millisecond)
			return 5
		})
		f.Get()
	}
}

func BenchmarkGenericFuture(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		f := NewFuture(func() int {
			time.Sleep(5 * time.Millisecond)
			return 5
		})
		f.Get()
	}
}