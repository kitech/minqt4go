package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/kitech/gopp"

	_ "github.com/kitech/minqt/qtinline" // import这个包，链接还是慢
	_ "github.com/qtui/qtmeta"
	"github.com/qtui/qtrt"
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
		log.Println("input some keywords like engine")
		return
	}

	var nowt = time.Now()
	signt := qtrt.LoadAllQtSymbols(stub)
	log.Println(gopp.Lenof(signt), time.Since(nowt)) // about 1.1s
	// os.Exit(0)

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
		clz, mth := qtrt.SplitMethod(dname)
		cp.APf("", "// %s", dname)
		cp.APf("", "func (me *%s) %s() {", clz, strings.Title(mth))
		cp.APf("", "  name := \"%s\"", oname)
		cp.AP("", "  sym := dlsym(name)")
		cp.AP("", "  rv := cgopp.FfiCall(sym)")
		cp.AP("", "}\n")
	}

	codesnip := cp.ExportAll()
	log.Println("codesnip", codesnip)

	log.Println(len(qtrt.Classes), "mthcnt", len(qtrt.Symdedups), "deduped", qtrt.Symdedupedcnt, gopp.Bytes2Hum(gopp.DeepSizeof(qtrt.Classes, 0)))

	testcall()

	log.Println("top -pid", os.Getpid())
	// gopp.PauseAk() // 到这儿，内存24M

	app := NewQGuiApplication(1, []string{"./heh.exe"}, 0)
	// gopp.PauseAk() // 到这儿，内存26M
	ape := NewQQmlApplicationEngine(nil)
	// gopp.PauseAk() // 到这儿，内存27M
	ape.Load("hh.qml")
	log.Println("top -pid", os.Getpid())
	// gopp.PauseAk() // 到这儿，内存28M
	log.Println("app.Exec ...")
	app.Exec() // 到这儿，内存32M

	log.Println("top -pid", os.Getpid())
	time.Sleep(gopp.DurandSec(23, 3))
	qtrt.Classes = nil
	qtrt.Symdedups = nil
	log.Println("clean vars")

	gopp.Forever()
}
