package minqt

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"sync/atomic"

	"github.com/ebitengine/purego"
	_ "github.com/ebitengine/purego"

	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"

	cmap "github.com/orcaman/concurrent-map/v2"
)

/*
 */
import "C"

// inline 的函数/方法就没法搞了。。。

func QVersion() string {
	rv := call0("qVersion")
	return cgopp.GoString(rv)
}

func call0(name string) voidptr {
	sym, err := purego.Dlsym(purego.RTLD_DEFAULT, name)
	gopp.ErrPrint(err, name)
	rv := cgopp.Litfficallg(voidptr(sym))
	return rv
}

// todo
type seqfnpair struct {
	np *int64
	f  func()
}

var runuithfns = cmap.New[seqfnpair]()
var runuithseq int64 = 10000

//export qtuithcbfningo
func qtuithcbfningo(n *int64) {
	key := fmt.Sprintf("%d", *n)
	// log.Println(*n, key)
	pair, ok := runuithfns.Get(key)
	if ok {
		pair.f()
		runuithfns.Remove(key)
	}
}

func RunonUithread(f func()) {

	const name = "QMetaObjectInvokeMethod1"
	sym := dlsym(name)
	sym2 := dlsym("qtuithcbfningo")
	// log.Println(sym, name, sym2)

	seq := new(int64)
	*seq = atomic.AddInt64(&runuithseq, 3)

	key := fmt.Sprintf("%d", *seq)
	runuithfns.Set(key, seqfnpair{seq, f})

	cgopp.Litfficallg(sym, sym2, seq)
}

// 应该放在minqt包里面
func SetQtmsgout(f func(typ int, file, funcname, msg string) bool) {
	if f == nil {
		log.Println("set nil func is not allowed")
		return
	}
	qtmsgoutfn = f
}

// 不过没法传参数... 暂时先放弃了
var OnMissingCall = func(callee string) {
	gopp.Debug("caller missing", callee)
}
var qtmsgoutfn = qtMsgoutput

func qtMsgoutput(typ int, file, funcname, msg string) bool {
	gopp.Debug(typ, file, funcname, msg)

	if strings.Contains(msg, "ReferenceError") &&
		strings.HasSuffix(msg, "is not defined") {
		// missing function/slot
		//ReferenceError: neslot1 is not defined
		reg := regexp.MustCompile(`ReferenceError: ([^ ]+) is not defined`)
		mats := reg.FindAllStringSubmatch(msg, -1)
		log.Println(mats)
		if len(mats) > 0 && len(mats[0]) > 0 {
			callee := mats[0][1]
			OnMissingCall(callee) //
		}
	}

	return true
}

//export qtMessageOutputGoimpl
func qtMessageOutputGoimpl(typex C.int, filex *C.char, funcx *C.char, msgx *C.char) {
	gopp.Debug(typex, filex, funcx, msgx)
	typ := int(typex)
	file := cgopp.GoString(voidptr(filex))
	funcname := cgopp.GoString(voidptr(funcx))
	msg := cgopp.GoString(voidptr(msgx))

	ok := qtmsgoutfn(typ, file, funcname, msg)
	if !ok {
		qtMsgoutput(typ, file, funcname, msg)
	}
}

// ///
var symcache = cmap.New[voidptr]()

// 这个函数很快，50ns
// current process
func dlsym(name string) voidptr {
	// if sym, ok := symcache.Get(name); ok {
	// 	return sym
	// }
	symi, err := purego.Dlsym(purego.RTLD_DEFAULT, name)
	gopp.ErrPrint(err, name)
	if gopp.ErrHave(err, "symbol not found") {
		// sym, err = purego.Dlsym(purego.RTLD_DEFAULT, "_"+name)
	}
	sym := voidptr(symi)
	// if sym != nil {
	// 	symcache.Set(name, sym)
	// }
	return sym
}
func Dlsym0(name string) voidptr { return dlsym(name) }

type QObject struct {
	Cthis voidptr
}

func QObjectof(ptr voidptr) QObject {
	return QObject{ptr}
}

// why slow, 1ms?
func (me QObject) SetProperty(name string, valuex any) bool {

	var value = QVarintNew2(valuex)

	const symname = "_ZN7QObject11setPropertyEPKcRK8QVariant"

	sym := dlsym(symname)
	name4c := cgopp.StrtoRefc(&name)
	// name4c := cgopp.CString(name)
	// defer cgopp.Cfree(name4c)
	rv := cgopp.Litfficallg(sym, me.Cthis, name4c, value.Cthis)
	// log.Println(rv)
	gopp.GOUSED(rv)
	return true
}

// int/str/list???
func (me QObject) Property(name string) QVariant {
	const symname = "QObjectProperty1"
	sym := dlsym(symname)
	// name4c := cgopp.CString(name)
	// defer cgopp.Cfree(name4c)
	name4c := cgopp.StrtoRefc(&name)
	rv := cgopp.Litfficallg(sym, me.Cthis, name4c)
	// log.Println(rv)
	return QVariantof(rv)
}
func (me QObject) FindChild(objname string) QObject {

	sym := dlsym("QObjectFindChild1")
	on4c := cgopp.StrtoRefc(&objname)
	// on4c := cgopp.CString(objname)
	// defer cgopp.Cfree(on4c)
	rv := cgopp.Litfficallg(sym, me.Cthis, on4c)
	// log.Println(rv)
	return QObjectof(rv)
}

type QVariant struct {
	Cthis voidptr
}

func QVariantof(ptr voidptr) QVariant { return QVariant{ptr} }

func (me QVariant) Dtor() {
	sym := dlsym("QVariantDtor")
	cgopp.Litfficallg(sym, me.Cthis)
}

func QVarintNew2(vx any) QVariant {
	var vp QVariant
	switch value := vx.(type) {
	case int:
		vp = QVarintNew(value)
	case int64:
		vp = QVarintNew(value)
	case string:
		vp = QVarintNew(value)
	case voidptr:
		vp = QVarintNew(value)
	default:
		log.Println("unimpl", reflect.TypeOf(vx), value)
	}
	return vp
}
func QVarintNew[T int | int64 | string | voidptr](vx T) QVariant {
	// log.Println(reflect.TypeOf(any(vx)), vx)
	switch v := any(vx).(type) {
	case int:
		sym := dlsym("QVariantNewInt")
		v4cc := voidptr(usize(v))
		rv := cgopp.Litfficallg(sym, v4cc)
		return QVariantof(rv)
	case int64:
		sym := dlsym("QVariantNewInt64")
		v4cc := voidptr(usize(v))
		rv := cgopp.Litfficallg(sym, v4cc)
		return QVariantof(rv)

	case string:
		sym := dlsym("QVariantNewStr")
		// v4cc := cgopp.CString(v)
		// defer cgopp.Cfree(v4cc)
		v4cc := cgopp.StrtoRefc(&v)
		rv := cgopp.Litfficallg(sym, v4cc)
		return QVariantof(rv)
	case voidptr:
		sym := dlsym("QVariantNewPtr")
		rv := cgopp.Litfficallg(sym, v)
		return QVariantof(rv)
	}
	return QVariant{}
}
func (me QVariant) Toint() int {
	sym := dlsym("QVariantToint")
	rv := cgopp.Litfficallg(sym, me.Cthis)
	return int(usize(rv))
}
func (me QVariant) Toint64() int64 {
	sym := dlsym("QVariantToint64")
	rv := cgopp.Litfficallg(sym, me.Cthis)
	return int64(usize(rv))
}
func (me QVariant) Tostr() string {
	sym := dlsym("QVariantTostr")
	rv := cgopp.Litfficallg(sym, me.Cthis)
	return cgopp.GoString(rv)
}
func (me QVariant) Toptr() voidptr {
	sym := dlsym("QVariantToptr")
	rv := cgopp.Litfficallg(sym, me.Cthis)
	return rv
}

// ////
type QQuickItem struct {
	Cthis voidptr
}

func QQuickItemof(ptr voidptr) QQuickItem { return QQuickItem{ptr} }

// ////
// 没有C++类型?
type QStackView struct {
	Cthis voidptr
}

func (me QStackView) Replace(curritem, nextitem QQuickItem) {
	gopp.Info("todododooooo", curritem, nextitem)
}
func QStackViewof(ptr voidptr) QStackView { return QStackView{ptr} }

// ////
type QQmlApplicationEngine struct{ Cthis voidptr }

func QQmlApplicationEngineof(ptr voidptr) QQmlApplicationEngine {
	return QQmlApplicationEngine{ptr}
}
func QQmlApplicationEngineNew() QQmlApplicationEngine {
	// sym := dlsym("_ZN21QQmlApplicationEngineC1EP7QObject")
	sym := dlsym("QQmlApplicationEngineNew")
	// log.Println(sym)
	rv := cgopp.Litfficallg(sym)
	return QQmlApplicationEngineof(rv)
}
func (me QQmlApplicationEngine) Load(u string) {
	sym := dlsym("QQmlApplicationEngineLoad1")
	// v4cc := cgopp.CString(u)
	// defer cgopp.Cfree(v4cc)
	v4cc := cgopp.StrtoRefc(&u)
	cgopp.Litfficallg(sym, v4cc)
}

func (me QQmlApplicationEngine) RootObject() QObject {
	sym := dlsym("QQmlApplicationEngineRootObject1")
	rv := cgopp.Litfficallg(sym, me.Cthis)
	return QObjectof(rv)
}
