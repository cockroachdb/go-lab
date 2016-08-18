// Test programs used to determine in what cases assigning an interface to a
// variable results in Go making an allocation.
//
// Sample output:
//   BenchmarkInt-4             	  500000	       242 ns/op	       8 B/op	       1 allocs/op
//   BenchmarkString-4          	  500000	       265 ns/op	      16 B/op	       1 allocs/op
//   BenchmarkEmptyStruct-4     	 1000000	       228 ns/op	       0 B/op	       0 allocs/op
//   BenchmarkStruct-4          	  500000	       258 ns/op	      32 B/op	       1 allocs/op
//   BenchmarkSlice-4           	  500000	       273 ns/op	      32 B/op	       1 allocs/op
//   BenchmarkIntPtr-4          	 1000000	       217 ns/op	       0 B/op	       0 allocs/op
//   BenchmarkStructPtr-4       	 1000000	       217 ns/op	       0 B/op	       0 allocs/op
//   BenchmarkSlicePtr-4        	 1000000	       223 ns/op	       0 B/op	       0 allocs/op
//   BenchmarkUnsafePointer-4   	 1000000	       213 ns/op	       0 B/op	       0 allocs/op
//   PASS

package main

import (
	"testing"
	"unsafe"
)

var gen VarGen

func init() {
	// We use this interface to generate pointer and slice values that can't be
	// determined at compile time.
	gen = &VarGenImpl{}
}

func BenchmarkInt(b *testing.B) {
	ch := make(chan interface{})
	go func() {
		for _ = range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		x := int(i * i)
		ch <- x
	}
	close(ch)
}

func BenchmarkString(b *testing.B) {
	ch := make(chan interface{})
	go func() {
		for _ = range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		x := "test"
		ch <- x
	}
	close(ch)
}

func BenchmarkEmptyStruct(b *testing.B) {
	ch := make(chan interface{})
	go func() {
		for _ = range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		x := struct{}{}
		ch <- x
	}
	close(ch)
}

func BenchmarkStruct(b *testing.B) {
	ch := make(chan interface{})
	go func() {
		for _ = range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		x := VarGenStruct{a: 1}
		ch <- x
	}
	close(ch)
}

func BenchmarkSlice(b *testing.B) {
	ch := make(chan interface{})
	go func() {
		for _ = range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		x := gen.GenSlice()
		ch <- x
	}
	close(ch)
}

func BenchmarkIntPtr(b *testing.B) {
	ch := make(chan interface{})
	go func() {
		for _ = range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		x := gen.GenIntPtr()
		ch <- x
	}
	close(ch)
}

func BenchmarkStructPtr(b *testing.B) {
	ch := make(chan interface{})
	go func() {
		for _ = range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		x := gen.GenStructPtr()
		ch <- x
	}
	close(ch)
}

func BenchmarkSlicePtr(b *testing.B) {
	ch := make(chan interface{})
	go func() {
		for _ = range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		x := gen.GenSlicePtr()
		ch <- x
	}
	close(ch)
}

func BenchmarkUnsafePointer(b *testing.B) {
	ch := make(chan interface{})
	go func() {
		for _ = range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		x := unsafe.Pointer(uintptr(i))
		ch <- x
	}
	close(ch)
}
