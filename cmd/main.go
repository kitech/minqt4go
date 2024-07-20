package main

import (
	"encoding/json"
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
		return
	}

	var nowt = time.Now()
	signt := qtrt.LoadAllQtSymbols(stub)
	log.Println(gopp.Lenof(signt), time.Since(nowt)) // about 1.1s
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
	log.Println(codesnip)

	log.Println(len(qtrt.Classes), "mthcnt", len(qtrt.Symdedups), "deduped", qtrt.Symdedupedcnt, gopp.Bytes2Hum(gopp.DeepSizeof(qtrt.Classes, 0)))
	{
		bcc, err := json.Marshal(qtrt.Classes)
		gopp.ErrPrint(err)
		gopp.SafeWriteFile("Classes.json", bcc, 0644)
		bcc = nil

		nowt := time.Now()
		bcc, err = os.ReadFile("Classes.json")
		gopp.ErrPrint(err)
		qtrt.Classes = nil
		err = json.Unmarshal(bcc, &qtrt.Classes)
		gopp.ErrPrint(err)
		log.Println("decode big json", time.Since(nowt)) // about 400ms
		bcc = nil
	}

	testcall()

	// Libman.Open()

	app := NewQGuiApplication(1, []string{"./heh.exe"}, 0)
	ape := NewQQmlApplicationEngine(nil)
	ape.Load("hh.qml")
	log.Println("top -pid", os.Getpid())
	app.Exec()

	log.Println("top -pid", os.Getpid())
	time.Sleep(gopp.DurandSec(23, 3))
	qtrt.Classes = nil
	qtrt.Symdedups = nil
	log.Println("clean vars")

	gopp.Forever()
}
