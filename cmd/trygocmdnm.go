package main

// import "cmd/internal/objfile"
// import "objfile"

import (
	"cmd/goinct"
	"log"
	"time"

	"github.com/kitech/gopp"
)

func aaa() {
	// objfile.Open()
	goinct.Keep()

	nowt := time.Now()
	// goinct.NM("main")
	syms := goinct.NMget("../srcc/libQtQuickInline.dylib")
	log.Println(len(syms), gopp.DeepSizeof(syms, 0), time.Since(nowt))
	for i, symo := range syms {
		log.Println(i, string(symo.Code), symo.Name, symo.Size, symo.Type)
	}
}
