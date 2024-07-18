package main

import (
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/kitech/gopp"
)

// only call by callany
func getclzmthbycaller() (clz string, mth string) {
	pc, _, _, _ := runtime.Caller(2)
	fno := runtime.FuncForPC(pc)
	fnname := fno.Name()
	log.Println(fno, fnname, gopp.Retn(fno.FileLine(pc)))

	// main.NewQxx
	if pos := strings.LastIndex(fnname, ".NewQ"); pos > 0 {
		clz = fnname[pos+4:]
		mth = clz
		return
	}
	// main.(*QObject).Dummy
	funcname := "main.(*QObject).Dummy"
	funcname = fnname
	pos1 := strings.Index(funcname, "(")
	pos2 := strings.Index(funcname, ")")
	clzname := funcname[pos1+2 : pos2]
	mthname := funcname[pos2+2:]

	log.Println(clzname, mthname)

	clz = clzname
	mth = mthname
	return
}

func callany(obj voidptr, args ...any) {
	clzname, mthname := getclzmthbycaller()
	log.Println(clzname, mthname, obj, args)
	isctor := clzname == mthname

	mths, ok := Classes[clzname]
	gopp.FalsePrint(ok, "not found???", clzname)

	//
	var namercmths []ccMethod // 备份
	var rcmths = mths

	rcmths = resolvebyname(mthname, rcmths)
	namercmths = rcmths

	// 根据参数个数
	rcmths = resolvebyargc(len(args), rcmths)

	//
	argtys := reflecttypes(args...)
	rcmths = resolvebyargty(argtys, rcmths)

	if isctor {
		rcmths = resolvebyctorno(rcmths)
	}

	log.Println("final rcmths:", len(rcmths), rcmths)
	switch len(rcmths) {
	case 0:
		// so fuck
		gopp.Warn("No match, rcmths", namercmths)
	case 1:
		// good
	default:

	}
}

func resolvebyctorno(mths []ccMethod) (rets []ccMethod) {
	// bye C1E, C2E
	c2idx := -1
	c1idx := -1
	for idx, mtho := range mths {
		if strings.Contains(mtho.CCSym, mtho.Name+"C1E") {
			c1idx = idx
		} else if strings.Contains(mtho.CCSym, mtho.Name+"C2E") {
			c2idx = idx
		}
	}
	if c2idx >= 0 {
		rets = append(rets, mths[c2idx])
	} else if c2idx >= 0 {
		rets = append(rets, mths[c1idx])
	} else {
	}
	return
}

// todo todo todo
func qttypemathch(idx int, tystr string, tyo reflect.Type) bool {
	// log.Println("tymat", idx, tystr, "?<=", tyo.String())
	goty := tyo.String()

	tymat := false
	if goty == tystr {
		tymat = true
	} else if goty+"&" == tystr {
		tymat = true
	} else if goty == "[]string" && tystr == "char**" {
		tymat = true
	}
	log.Println("tymat", idx, tystr, "?<=", tyo.String(), tymat)

	return tymat
}

func resolvebyargty(tys []reflect.Type, mths []ccMethod) (rets []ccMethod) {
	for _, mtho := range mths {
		// log.Println(mtho.Name)

		sgnt, _ := demangle(mtho.CCSym)
		// log.Println(sgnt, mtho.CCSym)
		vec := SplitArgs(sgnt)
		log.Println(len(vec), vec, sgnt)

		allmat := true
		for j := 0; j < len(vec); j++ {
			mat := qttypemathch(j, vec[j], tys[j])
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
		// log.Println(mtho.Name)
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
		if len(vec) == argc {
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
