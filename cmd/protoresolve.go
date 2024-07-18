package main

import (
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/kitech/gopp"
)

func callany(obj voidptr, args ...any) {
	pc, _, _, _ := runtime.Caller(1)
	fno := runtime.FuncForPC(pc)
	fnname := fno.Name()
	log.Println(fno, fnname, gopp.Retn(fno.FileLine(pc)))

	funcname := "main.(*QObject).Dummy"
	funcname = fnname
	pos1 := strings.Index(funcname, "(")
	pos2 := strings.Index(funcname, ")")
	clzname := funcname[pos1+2 : pos2]
	mthname := funcname[pos2+2:]

	log.Println(clzname, mthname, obj, args)

	mths, ok := Classes[clzname]
	gopp.FalsePrint(ok, "not found???", clzname)

	//
	var rcmths = mths
	rcmths = resolvebyname(mthname, rcmths)
	// 根据参数个数
	rcmths = resolvebyargc(len(args), rcmths)

	//
	argtys := reflecttypes(args...)
	rcmths = resolvebyargty(argtys, rcmths)

	log.Println("final rcmths:", len(rcmths))
	switch len(rcmths) {
	case 0:
		// so fuck
	case 1:
		// good
	default:

	}
}

func qttypemathch(tystr string, tyo reflect.Type) bool {
	log.Println("tymat", tystr, "?<=", tyo.String())
	return false
}

func resolvebyargty(tys []reflect.Type, mths []ccMethod) (rets []ccMethod) {
	for _, mtho := range mths {
		// log.Println(mtho.Name)

		sgnt, _ := demangle(mtho.CCSym)
		// log.Println(sgnt, mtho.CCSym)
		vec := SplitArgs(sgnt)
		// log.Println(vec)

		allmat := true
		for j := 0; j < len(vec); j++ {
			mat := qttypemathch(vec[j], tys[j])
			if !mat {
				allmat = false
			}
		}
		if allmat {
			log.Println(gopp.MyFuncName(), "rc", mtho.CCCrc, mtho.CCSym)
			rets = append(rets, mtho)
		}
	}
	return
}

func resolvebyname(mthname string, mths []ccMethod) (rets []ccMethod) {

	for _, mtho := range mths {
		log.Println(mtho.Name)
		if mtho.Name == mthname {
			log.Println(gopp.MyFuncName(), "rc", mthname, mtho.CCCrc, mtho.CCSym)
			rets = append(rets, mtho)
		}
	}
	return
}
func resolvebyargc(argc int, mths []ccMethod) (rets []ccMethod) {
	for _, mtho := range mths {
		// log.Println(mtho.Name)
		// if mtho.Type.NumIn() == argc {
		sgnt, _ := demangle(mtho.CCSym)
		// log.Println(sgnt, mtho.CCSym)
		vec := SplitArgs(sgnt)
		// log.Println(vec)
		if true || len(vec) == argc {
			log.Println(gopp.MyFuncName(), "rc", mtho.CCCrc, mtho.CCSym)
			rets = append(rets, mtho)
		}
		// }
	}
	return
}

func reflecttypes(args ...any) (rets []reflect.Type) {
	for _, argx := range args {
		ty := reflect.TypeOf(argx)
		rets = append(rets, ty)
	}

	return
}

// ///
type QObject struct {
	Cthis voidptr
}

func (me *QObject) Connect(args ...any) {
	callany(me.Cthis, args...)
}
func testcall() {
	me := &QObject{}
	me.Connect(voidptr(usize(0)), 123, "aiewjff", 456.78, 999)
}
