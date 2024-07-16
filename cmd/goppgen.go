package main

// dont modify main

import "unsafe"

// begin cgo

/*
#include <stdint.h>
*/
import "C"

// end cgo

type i32 = int32
type i64 = int64
type u32 = uint32
type u64 = uint64
type f32 = float32
type f64 = float64
type usize = uintptr
type vptr = unsafe.Pointer
type voidptr = unsafe.Pointer

// begin cgo
type cuptr = C.uintptr_t
type cvptr = *C.void
type charptr = *C.char
type ucharptr = *C.uchar
type scharptr = *C.schar
type cint = C.int
type cshort = C.short
type clong = C.long
type cfloat = C.float
type cdouble = C.double
type cuintptr = C.uintptr_t
type ci64 = C.int64_t
type cu64 = C.uint64_t
type ci32 = C.int32_t
type cu32 = C.uint32_t
type ci16 = C.int16_t
type cu16 = C.uint16_t

// end cgo

func anyptr2uptr[T any](p *T) usize {
	var pp = usize(vptr(p))
	return pp
}

// begin cgo

func anyptr2uptrc[T any](p *T) cuptr {
	var pp = uintptr(vptr(p))
	return cuptr(pp)
}

// end cgo
