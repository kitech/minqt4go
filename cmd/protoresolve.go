package main

import (
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
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

// like jit, name jitqt
func callany(cobj voidptr, args ...any) voidptr {
	clzname, mthname := getclzmthbycaller()
	log.Println(clzname, mthname, cobj, args)
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
	var ccret voidptr
	switch len(rcmths) {
	case 0:
		// sowtfuck
		gopp.Warn("No match, rcmths", len(namercmths), namercmths, len(namercmths))
	case 1:
		// good
		mtho := rcmths[0]
		convedargs := argsconvert(mtho, argtys, args...)
		log.Println("oriargs", len(args), args)
		log.Println("convedargs", len(convedargs), convedargs)
		fnsym := Libman.Dlsym(mtho.CCSym)
		if isctor {
			cthis := cgopp.Mallocgc(123)
			ccargs := append([]any{cthis}, convedargs...)
			// log.Println("fficall info", mthname, fnsym, len(args), len(ccargs), ccargs)
			// cpp ctor 函数是没有返回值的
			cgopp.FfiCall[int](fnsym, ccargs...)
			ccret = cthis
		} else {
			// todo
			ccargs := append([]any{cobj}, convedargs...)
			// log.Println("fficall info", clzname, mthname, fnsym, len(args), len(ccargs), ccargs)
			ccret = cgopp.FfiCall[voidptr](fnsym, ccargs...)
		}
	default:
		// sowtfuck
	}
	return ccret
}

func resolvebyctorno(mths []ccMethod) (rets []ccMethod) {
	// bye C1E, C2E, C3E?
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
func qttypemathch(idx int, tystr string, tyo reflect.Type, conv bool, argx any) (any, bool) {
	// log.Println("tymat", idx, tystr, "?<=", tyo.String())
	// goty := tyo.String()

	var rvx = argx
	tymat := false

	mcdata := &TMCData{}
	mcdata.idx = idx
	mcdata.ctys = tystr
	mcdata.gotyo = tyo
	mcdata.goargx = argx

	for _, mater := range typemcers {
		tymat = mater.Match(mcdata, conv)
		if tymat {
			if conv {
				rvx = mcdata.ffiargx
			}
			log.Println("matched", reflect.TypeOf(mater), conv)
			break
		}
	}

	// if goty == tystr {
	// 	tymat = true
	// } else if goty+"&" == tystr {
	// 	tymat = true
	// 	if conv {
	// 		// 只对primitive type可以
	// 		refval := reflect.New(tyo)
	// 		refval.Elem().Set(reflect.ValueOf(argx))
	// 		rvx = refval.Interface()
	// 	}
	// } else if goty == "[]string" && tystr == "char**" {
	// 	tymat = true
	// 	if conv {
	// 		// todo how freeit
	// 		ptr := cgopp.CStrArrFromStrs(argx.([]string))
	// 		rvx = ptr
	// 	}

	// 	// QObject* ?<= *main.QObject
	// } else if isqtptrtymat(tystr, tyo) {
	// 	tymat = true
	// 	if conv {
	// 		tvx := reflect.ValueOf(argx)
	// 		if tvx.IsNil() {

	// 		} else {
	// 			// .Elem().FieldByName("Cthis")
	// 		}
	// 		log.Println(tvx)
	// 	}
	// }
	gopp.FalsePrint(tymat, "tymat", idx, tystr, "?<=", tyo.String(), tymat)

	return rvx, tymat
}

func argsconvert(mtho ccMethod, tys []reflect.Type, args ...any) (rets []any) {
	sgnt, _ := demangle(mtho.CCSym)
	// log.Println(sgnt, mtho.CCSym)
	vec := SplitArgs(sgnt)
	log.Println(len(vec), vec, sgnt)

	for j := 0; j < len(vec); j++ {
		argx, mat := qttypemathch(j, vec[j], tys[j], true, args[j])
		if !mat {
			// wtf???
		}
		rets = append(rets, argx)
	}
	return
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
			_, mat := qttypemathch(j, vec[j], tys[j], false, nil)
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
