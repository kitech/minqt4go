package main

import "github.com/qtui/qtrt"

// ///
type QObject struct {
	Cthis voidptr
}

func (me *QObject) Connect(args ...any) {
	qtrt.Callany(me.Cthis, args...)
}
func testcall() {
	me := &QObject{}
	me.Connect(voidptr(usize(0)), 123, "aiewjff", 456.78, 999)
}

type QGuiApplication struct {
	Cthis voidptr
}

func NewQGuiApplication(argc int, argv []string, flags int) *QGuiApplication {
	cthis := qtrt.Callany(nil, argc, argv, flags)
	return &QGuiApplication{cthis}
}

func (me *QGuiApplication) Exec() int {
	qtrt.Callany(me.Cthis)
	return 0
}

type QQmlApplicationEngine struct {
	Cthis voidptr
}

func NewQQmlApplicationEngine(obj *QObject) *QQmlApplicationEngine {
	cthis := qtrt.Callany(nil, obj)
	return &QQmlApplicationEngine{cthis}
}

func (me *QQmlApplicationEngine) Load(qmlfile string) {
	qtrt.Callany(me.Cthis, qmlfile)
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
