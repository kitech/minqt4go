package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kitech/gopp"
)

/*

 */
import "C"

// 按照 minqt 的格式规范，生成对应函数/类(?)的封装
// 包括能够取到的符号表
// TODO inline 的方法的处理
// TODO 获取不到 static
// usage:
func main() {
	flag.Parse()

	stub := flag.Arg(0)
	if gopp.Empty(stub) {
		return
	}

	libpfx := gopp.Mustify(os.UserHomeDir())[0].Str() + "/.nix-profile/lib"
	globtmpl := fmt.Sprintf("%s/Qt*.framework/Qt*", libpfx)
	libs, err := filepath.Glob(globtmpl)
	gopp.ErrPrint(err, libs)
	log.Println(libs, len(libs))
	signtx := gopp.Mapdo(libs, func(idx int, vx any) (rets []any) {
		log.Println(idx, vx, gopp.Bytes2Humz(gopp.FileSize(vx.(string))))
		tmpfile := "symfiles/" + filepath.Base(vx.(string)) + ".sym"
		var lines []string
		if !gopp.FileExist2(tmpfile) {
			lines, err := gopp.RunCmd(".", "nm", vx.(string))
			gopp.ErrPrint(err, vx)
			log.Println(idx, vx, len(lines))
			// save cache
			gopp.SafeWriteFile(tmpfile, []byte(strings.Join(lines, "\n")), 0644)
		} else {
			bcc, err := os.ReadFile(tmpfile)
			gopp.ErrPrint(err, tmpfile)
			lines = strings.Split(string(bcc), "\n")
		}

		for _, line := range lines {
			if strings.Contains(line, stub) && !strings.Contains(line, "Private") {
				// log.Println(line)
				name := gopp.Lastof(strings.Split(line, " ")).Str()
				signt, ok := demangle(name)
				log.Println(name, ok, signt)
				rets = append(rets, name, signt)
			}
			addsymrawline(filepath.Base(vx.(string)), line)
		}
		return
	})
	log.Println(gopp.Lenof(signtx))
	signt := gopp.IV2Strings(signtx.([]any))
	cp := gopp.NewCodePager()
	for i := 0; i < len(signt); i += 2 {
		oname := signt[i]
		dname := signt[i+1]
		if strings.HasPrefix(dname, "typeinfo for") ||
			strings.HasPrefix(dname, "non-virtual thunk to") ||
			strings.HasPrefix(dname, "typeinfo name for") ||
			strings.HasPrefix(dname, "guard variable for") ||
			strings.Count(dname, "<") > 0 {
			continue
		}

		// txt := fmt.Sprintf("// %s\nfunc () {\nsymname=\"%s\"\n}\n", dname, oname)
		// fmt.Println(txt)
		clz, mth := SplitMethod(dname)
		cp.APf("", "// %s", dname)
		cp.APf("", "func (me *%s) %s() {", clz, strings.Title(mth))
		cp.APf("", "  name := \"%s\"", oname)
		cp.AP("", "  sym := dlsym(name)")
		cp.AP("", "  rv := cgopp.FfiCall(sym)")
		cp.AP("", "}\n")
	}

	codesnip := cp.ExportAll()
	log.Println(codesnip)

	log.Println(len(Classes), "mthcnt", len(dedups), "deduped", dedupedcnt, gopp.Bytes2Hum(gopp.DeepSizeof(Classes, 0)))

	testcall()

	log.Println("top -pid", os.Getpid())
	time.Sleep(gopp.DurandSec(23, 3))
	Classes = nil
	dedups = nil
	log.Println("clean vars")

	gopp.Forever()
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

	rets = strings.Split(mid, ", ")

	return
}
