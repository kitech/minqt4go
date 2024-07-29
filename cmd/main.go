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
	_ "github.com/qtui/qtrt"
	"github.com/qtui/qtsyms"
	sitter "github.com/smacker/go-tree-sitter"
)

/*

 */
import "C"

var _ = sitter.ErrNoLanguage

// 按照 minqt 的格式规范，生成对应函数/类(?)的封装
// 包括能够取到的符号表
// TODO inline 的方法的处理
// TODO 获取不到 static
// usage:
func main() {
	runtime.LockOSThread()
	flag.Parse()

	if true {
		aaa()
		return
	}

	stub := flag.Arg(0)
	if gopp.Empty(stub) {
		log.Println("input some keywords like engine")
		return
	}

	// pr := sitter.NewParser()
	// sitter.NewLanguage()

	var nowt = time.Now()
	signt := qtsyms.LoadAllQtSymbols()
	log.Println(gopp.Lenof(signt), time.Since(nowt)) // about 1.1s
	// os.Exit(0)
	gopp.Mapdo(qtsyms.QtSymbols, func(idx int, clz string, mths []qtsyms.QtMethod) {
		// log.Println(idx, clz, mths)
		for _, mtho := range mths {
			if gopp.StrHaveNocase(mtho.CCSym, stub) {
				decsym, ok := qtsyms.Demangle(mtho.CCSym)
				gopp.FalsePrint(ok, mtho.CCSym)
				signt = append(signt, mtho.CCSym, decsym)
			}
		}
	})
	log.Println(gopp.Lenof(signt), time.Since(nowt)) // about 1.1s

	cp := gopp.NewCodePager()
	for i := 0; i < len(signt); i += 2 {
		// oname := signt[i]
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
		clz, mth := qtsyms.SplitMethod(dname)
		log.Println(dname, clz, mth)
		cp.APf("", "// %s", dname)
		if clz == mth {
			cp.APf("", "func  %sNew%d() {", clz)
			cp.AP("", "  rv := qtrt.Callany(nil)")
			cp.AP("", "  _ = rv")
			cp.AP("", "}\n")
		} else {
			cp.APf("", "func (me *%s) %s() {", clz, strings.Title(mth))
			cp.AP("", "  rv := qtrt.Callany(nil)")
			cp.AP("", "  _ = rv")
			cp.AP("", "}\n")
		}
		// cp.APf("", "func (me *%s) %s() {", clz, strings.Title(mth))
		// cp.APf("", "  name := \"%s\"", oname)
		// cp.AP("", "  fnsym := qtrt.GetQtSymAddr(name)")
		// cp.AP("", "  rv := cgopp.FfiCall[gopp.Fatptr](fnsym)")
		// cp.AP("", "}\n")

	}

	codesnip := cp.ExportAll()
	log.Println("codesnip", codesnip)

	log.Println(len(qtsyms.QtSymbols), "mthcnt", len(qtsyms.Symdedups), "deduped", qtsyms.Symdedupedcnt, gopp.Bytes2Hum(gopp.DeepSizeof(qtsyms.QtSymbols, 0)))

	if true {
		return
	}
	testcall()

	log.Println("top -pid", os.Getpid(), "lsof -p", os.Getpid())
	// gopp.PauseAk() // 到这儿，内存24M

	app := NewQGuiApplication(1, []string{"./heh.exe"}, 0)
	// gopp.PauseAk() // 到这儿，内存26M
	ape := NewQQmlApplicationEngine(nil)
	// gopp.PauseAk() // 到这儿，内存27M
	ape.Load("hh.qml")
	log.Println("top -pid", os.Getpid())

	// gopp.PauseAk() // 到这儿，内存28M
	log.Println("app.Exec ...", "top -pid", os.Getpid(), "lsof -p", os.Getpid())
	app.Exec() // 到这儿，内存32M

	log.Println("top -pid", os.Getpid())
	time.Sleep(gopp.DurandSec(23, 3))
	qtsyms.QtSymbols = nil
	qtsyms.Symdedups = nil
	log.Println("clean vars")

	gopp.Forever()
}
