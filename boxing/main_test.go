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

// We use this interface to generate pointer and slice values that can't be
// determined at compile time.
var gen VarGen = &VarGenImpl{}

// sendOnChan is a helper that repeatedly calls a function which pushes an
// interface{} to a channel. The point is to use an interface{} value in a way
// that doesn't require allocations but that cannot be optimized out.
func sendOnChan(b *testing.B, f func(chan<- interface{})) {
	ch := make(chan interface{})
	go func() {
		for _ = range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		f(ch)
	}
	close(ch)
}

func BenchmarkInt(b *testing.B) {
	sendOnChan(b, func(ch chan<- interface{}) {
		x := int(5)
		ch <- x
	})
}

func BenchmarkString(b *testing.B) {
	sendOnChan(b, func(ch chan<- interface{}) {
		x := "test"
		ch <- x
	})
}

func BenchmarkEmptyStruct(b *testing.B) {
	sendOnChan(b, func(ch chan<- interface{}) {
		x := struct{}{}
		ch <- x
	})
}

func BenchmarkStruct(b *testing.B) {
	sendOnChan(b, func(ch chan<- interface{}) {
		x := VarGenStruct{a: 1}
		ch <- x
	})
}

func BenchmarkSlice(b *testing.B) {
	sendOnChan(b, func(ch chan<- interface{}) {
		x := gen.GenSlice()
		ch <- x
	})
}

func BenchmarkIntPtr(b *testing.B) {
	sendOnChan(b, func(ch chan<- interface{}) {
		x := gen.GenIntPtr()
		ch <- x
	})
}

func BenchmarkStructPtr(b *testing.B) {
	sendOnChan(b, func(ch chan<- interface{}) {
		x := gen.GenStructPtr()
		ch <- x
	})
}

func BenchmarkSlicePtr(b *testing.B) {
	sendOnChan(b, func(ch chan<- interface{}) {
		x := gen.GenStructPtr()
		ch <- x
	})
}

func BenchmarkUnsafePointer(b *testing.B) {
	sendOnChan(b, func(ch chan<- interface{}) {
		x := unsafe.Pointer(uintptr(5))
		ch <- x
	})
}
