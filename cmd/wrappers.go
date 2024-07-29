package main

import (
	"github.com/qtui/qtrt"
)

// deprecated file

// ///
type QObject struct {
	Cthis voidptr
}

func (me *QObject) GetCthis() voidptr { return me.Cthis }

func (me *QObject) Connect(args ...any) {
	qtrt.Callany[int](me, args...)
}
func testcall() {
	me := &QObject{}
	me.Connect(voidptr(usize(0)), 123, "aiewjff", 456.78, 999)
}

type QGuiApplication struct {
	Cthis voidptr
}

func (me *QGuiApplication) GetCthis() voidptr { return me.Cthis }

func NewQGuiApplication(argc int, argv []string, flags int) *QGuiApplication {
	// ptr := qtrt.Callany[voidptr](nil, argc, argv, flags)
	// return &QGuiApplication{gopp.FatptrAs[voidptr](ptr)}
	return nil
}

func (me *QGuiApplication) Exec() int {
	qtrt.Callany[int](me)
	return 0
}

type QQmlApplicationEngine struct {
	Cthis voidptr
}

func (me *QQmlApplicationEngine) GetCthis() voidptr { return me.Cthis }

func NewQQmlApplicationEngine(obj *QObject) *QQmlApplicationEngine {
	// ptr := qtrt.Callany[voidptr](nil, obj)
	// cthis := gopp.FatptrAs[voidptr](ptr)
	// return &QQmlApplicationEngine{cthis}
	return nil
}

func (me *QQmlApplicationEngine) Load(qmlfile string) {
	qtrt.Callany[int](me, qmlfile)
}

/*
	symname := "__ZN15QGuiApplicationC2ERiPPci"
	fnsym := Libman.Dlsym(symname)
	gopp.NilPrint(fnsym, symname)
	log.Println(symname, fnsym)

	argv := []string{"./hehehexe"}
	argv4c := voidptr(cgopp.CStrArrFromStrs(argv))
	argc := new(int)
	*argc = len(argv)
	argc4c := voidptr(argc)
	cthis := cgopp.Mallocgc(128)
	qapp := cgopp.FfiCall[voidptr](fnsym, cthis, argc4c, argv4c, 0)

	{
		symname := "__ZN15QGuiApplication4execEv"
		fnsym := Libman.Dlsym(symname)
		gopp.NilPrint(fnsym, symname)
		log.Println(symname, fnsym)

		rv := cgopp.FfiCall[int](fnsym, qapp)
		log.Println(rv)
	}
*/
