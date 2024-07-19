package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
	runtime.LockOSThread()
	flag.Parse()

	stub := flag.Arg(0)
	if gopp.Empty(stub) {
		return
	}

	var nowt = time.Now()

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
	log.Println(gopp.Lenof(signtx), time.Since(nowt)) // about 1.1s
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
	{
		bcc, err := json.Marshal(Classes)
		gopp.ErrPrint(err)
		gopp.SafeWriteFile("Classes.json", bcc, 0644)
		bcc = nil

		nowt := time.Now()
		bcc, err = os.ReadFile("Classes.json")
		gopp.ErrPrint(err)
		Classes = nil
		err = json.Unmarshal(bcc, &Classes)
		gopp.ErrPrint(err)
		log.Println("decode big json", time.Since(nowt)) // about 400ms
		bcc = nil
	}

	testcall()

	Libman.Open()

	app := NewQGuiApplication(1, []string{"./heh.exe"}, 0)
	ape := NewQQmlApplicationEngine(nil)
	ape.Load("hh.qml")
	app.Exec()

	log.Println("top -pid", os.Getpid())
	time.Sleep(gopp.DurandSec(23, 3))
	Classes = nil
	dedups = nil
	log.Println("clean vars")

	gopp.Forever()
}
