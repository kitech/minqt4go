package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/ebitengine/purego"
	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"

	_ "github.com/kitech/minqt/qtinline" // import这个包，链接还是慢
	_ "github.com/qtui/qtmeta"
	"github.com/qtui/qtrt"
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

	stub := flag.Arg(0)
	if gopp.Empty(stub) {
		log.Println("input some keywords like engine")
		return
	}

	// pr := sitter.NewParser()
	// sitter.NewLanguage()

	var nowt = time.Now()
	signt := qtrt.LoadAllQtSymbols(stub)
	log.Println(gopp.Lenof(signt), time.Since(nowt)) // about 1.1s
	// os.Exit(0)
	if true {
		genqtclzszcpp()
		// make cxlib
		rungenclzsz()
		return
	}

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

	log.Println(len(qtrt.QtSymbols), "mthcnt", len(qtrt.Symdedups), "deduped", qtrt.Symdedupedcnt, gopp.Bytes2Hum(gopp.DeepSizeof(qtrt.QtSymbols, 0)))

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
	qtrt.QtSymbols = nil
	qtrt.Symdedups = nil
	log.Println("clean vars")

	gopp.Forever()
}

func genqtclzszcpp() {

	classes := gopp.MapKeys(qtrt.QtSymbols)
	sort.Strings(classes)

	cp := gopp.NewCodePager()
	cp.AP("", "#include <QtCore>")
	cp.AP("", "#include <QtGui>")
	cp.AP("", "#include <QtWidgets>")
	cp.AP("", "#include <QtQml>")
	cp.AP("", "#include <QtNetwork>")
	cp.AP("", "#include <QtOpenGL>")
	cp.AP("", "#include <QtQuick>")
	cp.AP("", "#include <QtQuickControls2>")
	cp.AP("", "#include <QtQuickTemplates2>")
	cp.AP("", "#include <QtQuickWidgets>")

	cp.AP("", "int genqtclzsz(quint64 crc){")
	cp.AP("", "switch (crc){")
	gopp.Mapdo(classes, func(idx int, v string) {
		// log.Println(idx, v)
		// if strings.HasPrefix(v, "QAbstract") {
		// 	return
		// }
		// if strings.HasPrefix(v, "QAccessible") {
		// 	return
		// }
		// if strings.HasPrefix(v, "QApple") {
		// 	return
		// }
		// if strings.HasPrefix(v, "QBacking") {
		// 	return
		// }
		// if strings.HasPrefix(v, "QBenchmark") {
		// 	return
		// }

		// if slices.Contains(gopp.Sliceof("QArgumentType", "QActionAnimation", "QAdoptedThread", "QAlphaPaintEngine", "QAlphaWidget", "QAnimationGroupJob", "QAnimationTimer", "QBasicDrag"), v) {
		// 	return
		// }
		if slices.Contains(gopp.Sliceof("QNativeInterface", "QPasswordDigestor", "QQuickOpenGLUtils", "QSsl"), v) {
			return
		}
		// if slices.Contains(gopp.Sliceof("QPain", "QRa", "QLockFile", "QLoggingC", "QP"), v) {
		// 	return
		// }
		crc := gopp.Crc64Str(v)
		// cp.APf("", "#include <%s> // %d", v, idx)
		// todo todo replace #ifdef
		cp.APf("", "#ifdef %s_H", strings.ToUpper(v))
		cp.APf("", "case %vUL: return int(sizeof(%s));  // %d", crc, v, idx)
		cp.APf("", "#endif")
	})
	cp.AP("", "}")
	cp.AP("", "return -1;")
	cp.AP("", "}")
	codestr := cp.ExportAll()
	// log.Println(cp.ExportAll())

	savefile := "../qtallcc/genqtclzsz.cpp"
	gopp.SafeWriteFile(savefile, []byte(codestr), 0644)
}

func rungenclzsz() {

	dlh, err := purego.Dlopen("../qtallcc/libQtAllInline.dylib", purego.RTLD_LAZY)
	gopp.ErrPrint(err)
	name := "_Z10genqtclzszy"
	sym, err := purego.Dlsym(dlh, name)
	gopp.ErrPrint(err)

	classes := gopp.MapKeys(qtrt.QtSymbols)
	sort.Strings(classes)

	validcnt := 0
	cp := gopp.NewCodePager()
	cp.AP("", "package qtclzsz\n")
	cp.AP("", "// for amd64")
	cp.AP("", "func getfromgened(class string) int {")
	cp.AP("", "switch class {")
	gopp.Mapdo(classes, func(idx int, v string) {
		crc := gopp.Crc64Str(v)
		sz := cgopp.FfiCall[int](sym, crc)
		// log.Println(idx, v, sz)
		comment := ""
		if sz == -1 || sz == 4294967295 {
			comment = "// "
		} else {
			validcnt++
		}
		cp.APf("", "%scase \"%s\": return %v", comment, v, sz)
	})
	cp.AP("", "}")
	cp.AP("", "return -1")
	cp.AP("", "}")
	codestr := cp.ExportAll()
	// log.Println(codestr)

	savefile := "../../qtui/qtclzsz/sizes_gened.go"
	gopp.SafeWriteFile(savefile, []byte(codestr), 0644)
	log.Println("gen valid", validcnt, "of", len(classes))
}
