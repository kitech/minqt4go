package main

import (
	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
)

/*
// #include <cxxabi.h>

#cgo LDFLAGS: -lc++

extern void* _Z20cxxabi__cxa_demanglePcS_PmPi(void*, void*, void*, void*);
*/
import "C"

// /// \see ../srcc/demangle.cpp
func demangle(name string) (string, bool) {
	name4c := cgopp.CStringgc(name)

	// var reslen usize = 300
	var resok int
	var rv voidptr
	if gopp.RandNum(0, 2) == 1 {
		// res4c := cgopp.Mallocgc(int(reslen))
		// 如果长度不够会realloc，传递go分配的内存则crash
		rv = C._Z20cxxabi__cxa_demanglePcS_PmPi(name4c, nil, nil, (voidptr)(&resok))
	} else {
		fnsym := cgopp.Dlsym0("__cxa_demangle")
		rv = cgopp.Litfficallg(fnsym, name4c, nil, nil, (voidptr)(&resok))
	}
	// log.Println(name, resok, reslen, rv, cgopp.GoString(rv), res4c)
	gopp.TruePrint(rv != nil && resok != 0, "wow", rv, resok)
	if resok == 0 {
		defer cgopp.Cfree(rv)
		return cgopp.GoString(rv), resok == 0
	}

	return "", resok == 0
}
