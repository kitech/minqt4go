package main

import (
	"log"
	"reflect"
	"strings"

	"github.com/kitech/gopp"
)

var Classes = map[string][]ccMethod{} // class name => struct type
type ccMethod struct {
	reflect.Method
	Static bool
	CCSym  string
	CCCrc  uint64
}

var dedups = map[uint64]int{} // sym crc =>
var dedupedcnt = 0

func addsymrawline(qtmodename string, line string) {
	flds := strings.Split(line, " ")
	// log.Println(line, flds)
	sym := gopp.LastofGv(flds)
	// log.Println(sym)
	addsym(sym)
}
func addsym(name string) {
	// log.Println("demangle...", len(name), name)
	sgnt, ok := demangle(name)
	if strings.HasPrefix(name, "GCC_except") {
	} else if strings.HasPrefix(name, "_OBJC_") {
	} else if strings.Contains(name, "QtPrivate") {
	} else {
		// gopp.FalsePrint(ok, "demangle failed", name)
	}
	// log.Println(ok, len(name), "=>", len(sgnt), sgnt, ok)
	if !ok {
		return
	}
	if strings.HasPrefix(sgnt, "typeinfo") {
		return
	}
	if strings.HasPrefix(sgnt, "vtable") {
		return
	}
	if strings.HasPrefix(sgnt, "operator") {
		return
	}
	if strings.Contains(sgnt, "operator+=") {
		return
	}
	if strings.Contains(sgnt, "anonymous namespace") {
		return
	}

	clzname, mthname := SplitMethod(sgnt)
	if clzname == "" || mthname == "" {
		if clzname == "" && mthname != "" {
			// maybe global function
		} else {
			gopp.Warn("somerr", clzname, mthname, sgnt)
		}
		return
	}
	// log.Println(clzname, mthname, sgnt)
	mths, ok := Classes[clzname]
	if ok {
	} else {
		Classes[clzname] = nil
	}
	mtho := ccMethod{}
	mtho.CCSym = name
	mtho.CCCrc = gopp.Crc64Str(name)

	if _, ok := dedups[mtho.CCCrc]; ok {
		// log.Println("already have", sgnt, len(dedups))
		dedupedcnt++
		return
	}
	defer func() { dedups[mtho.CCCrc] = 1 }()

	mtho.Name = strings.Title(mthname)
	mtho.Index = len(mths)
	mtho.Type = nil
	mtho.Func = reflect.Value{}
	mtho.PkgPath = "hhpkg"

	mths = append(mths, mtho)
	Classes[clzname] = mths
}

func SplitMethod(s string) (string, string) {
	idx := strings.LastIndexAny(s, " )")
	if idx != -1 {
		s = s[:idx]
	}

	flds := strings.Split(s, "::")
	for i, fld := range flds {
		idx := strings.Index(fld, "(")
		if idx != -1 {
			flds[i] = fld[:idx]
		}
	}
	if len(flds) < 2 {
		return "", flds[0]
	}
	return flds[0], flds[1]
}

func SplitArgs(s string) (rets []string) {
	pos1 := strings.Index(s, "(")
	pos2 := strings.LastIndex(s, ")")

	mid := s[pos1+1 : pos2]
	log.Println(mid)
	if mid == "" {
		return
	}

	rets = strings.Split(mid, ", ")

	return
}
