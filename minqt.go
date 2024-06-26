package minqt

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"sync/atomic"
	"time"

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

func QCompVersion() string {
	rv := call0("QCompileVersion")
	return cgopp.GoString(rv)
}

func QRuntimeVersion() string {
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
	// gopp.Debug(typ, file, funcname, msg)
	// QOV qt output via
	var rfmsg = msg
	if strings.HasPrefix(rfmsg, "file://") {
		pos := strings.Index(rfmsg, " ")
		path := rfmsg[:pos]
		bname := gopp.Lastof(strings.Split(path, "/")).Str()
		rfmsg = fmt.Sprintf("%s%s", bname, rfmsg[pos:])
	}
	fmt.Printf("QOVGoLog: %s\n", rfmsg)

	if strings.Contains(msg, "ReferenceError") &&
		strings.HasSuffix(msg, "is not defined") {
		// missing function/slot
		//ReferenceError: neslot1 is not defined
		reg := regexp.MustCompile(`ReferenceError: ([^ ]+) is not defined`)
		mats := reg.FindAllStringSubmatch(msg, -1)
		// log.Println(mats)
		if len(mats) > 0 && len(mats[0]) > 0 {
			callee := mats[0][1]
			OnMissingCall(callee) //
		}
	}

	return true
}

//export qtMessageOutputGoimpl
func qtMessageOutputGoimpl(typex C.int, filex *C.char, funcx *C.char, msgx *C.char) {
	// gopp.Debug(typex, filex, funcx, msgx)
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

type QObjectst struct {
	Cthis voidptr
}
type QObject = *QObjectst

func QObjectof(ptr voidptr) QObject {
	return &QObjectst{ptr}
}
func (me QObject) SetCthis(ptr voidptr) QObject {
	me.Cthis = ptr
	return me
}

// why slow, 1ms?
func (me QObject) SetProperty(name string, valuex any) bool {

	var value = QVarintNew2(valuex)
	// defer value.Dtor()

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

// not need caller free property value, auto free now
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
	cnt   int
	ci    string
}

func QVariantof(ptr voidptr) QVariant {
	// gopp.NilPrint(ptr)

	me := QVariant{ptr, 0, gopp.ZeroStr}
	if false {
		ci := gopp.GetCallerInfo(5)
		me.ci = ci
	}

	time.AfterFunc(gopp.DurandSec(3, 5), me.Dtor)
	return me
}

func (me QVariant) Dtor() {
	sym := dlsym("QVariantDtor")
	// log.Println(sym, me.Cthis)
	cgopp.Litfficallg(sym, me.Cthis)
	me.cnt++
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
	case bool:
		vp = QVarintNew(value)
	default:
		log.Println("unimpl", reflect.TypeOf(vx), value)
	}
	// time.AfterFunc(gopp.DurandSec(3, 5), vp.Dtor)
	return vp
}
func QVarintNew[T int | int64 | string | voidptr | bool](vx T) QVariant {
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
	case bool:
		sym := dlsym("QVariantNewBool")
		// rv := cgopp.Litfficallg(sym, v)
		rv := cgopp.FfiCall[voidptr](sym, v)
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
func (me QVariant) Tobool() bool {
	sym := dlsym("QVariantTobool")
	rv := cgopp.Litfficallg(sym, me.Cthis)
	return usize(rv) != 0
}

// ////
type QQuickItemst struct {
	QObject
}
type QQuickItem = *QQuickItemst

func QQuickItemof(ptr voidptr) QQuickItem {
	me := &QQuickItemst{QObjectof(ptr)}
	return me
}

// ////
// 没有C++类型?
type QStackViewst struct {
	QObject
}
type QStackView = *QStackViewst

func QStackViewof(ptr voidptr) QStackView {
	me := &QStackViewst{QObjectof(ptr)}
	return me
}

func (me QStackView) ReplaceCurrentItem(nextitem QQuickItem) QQuickItem {
	// gopp.Info("todododooooo", nextitem)
	sym := dlsym("QQuickStackView_replaceCurrentItem")
	rv := cgopp.FfiCall[voidptr](sym, me.Cthis, nextitem.Cthis)
	return QQuickItemof(rv)
}

func (me QStackView) Get(idx int) QQuickItem {
	sym := dlsym("QQuickStackView_get")
	// rv := cgopp.Litfficallg(sym, me.Cthis, idx)
	log.Println(sym, me, me.Cthis, idx)
	rv := cgopp.FfiCall[voidptr](sym, me.Cthis, idx)
	// gopp.Info("todododooooo", curritem, nextitem)
	return QQuickItemof(rv)
}

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

// //////
type QQmlComponentst struct {
	QObject
}
type QQmlComponent = *QQmlComponentst

func QQmlComponentof(ptr voidptr) QQmlComponent {
	me := &QQmlComponentst{}
	me.SetCthis(ptr)
	return me
}

// todo
func QQmlComponentNew() {

}

type QtObjectst struct {
	QObject
}
type QtObject = *QtObjectst

// how get QtObject instance??? just call QtObject::create
func QtObjectof(ptr voidptr) QtObject {
	me := &QtObjectst{QObjectof(ptr)}
	return me
}

func QtObjectCreate(e QQmlApplicationEngine) QtObject {
	symname := "QtObjectCreate"
	sym := dlsym(symname)

	rv := cgopp.Litfficallg(sym, e.Cthis)
	return QtObjectof(rv)
}

func (me QtObject) CreateQmlObject(qmltxt string, parent QObject) QObject {
	symname := "QtObjectCreateQmlObject"
	sym := dlsym(symname)

	qmltxt4c := cgopp.CStringaf(qmltxt)
	rv := cgopp.Litfficallg(sym, me.Cthis, qmltxt4c, parent.Cthis)

	return QObjectof(rv)
}

//////////

type QArgument struct {
	Data   voidptr
	Tyname voidptr // type string, like QVariant/int/double
}

type QMetaObjectst struct {
}
type QMetaObject = *QMetaObjectst

// 用于调用静态方法
func QMetaObjectof0() QMetaObject {
	return &QMetaObjectst{}
}

/*
# define QMETHOD_CODE  0                        // member type codes
# define QSLOT_CODE    1
# define QSIGNAL_CODE  2
*/
func QMethodof(name string) string { return fmt.Sprintf("0%s", name) }
func QSlotof(name string) string   { return fmt.Sprintf("1%s", name) }
func QSignalof(name string) string { return fmt.Sprintf("2%s", name) }

// todo
func (me QMetaObject) Invoke2(obj QObject, slotname string, args ...any) {
	var argv [3]voidptr

	a0 := &QArgument{}
	a0.Data = QVarintNew2(123).Cthis
	a0.Tyname = cgopp.CStringaf("QVariant")
	argv[0] = (voidptr)(a0)

	for i := 0; i < len(argv) && i < len(args); i++ {
		// todo
	}

	// gopp.Println(argv)
	symname := "QMetaObjectInvokeMethod2"
	sym := dlsym(symname)
	name4c := cgopp.CStringaf(slotname)
	rv := cgopp.Litfficallg(sym, obj.Cthis, name4c, argv[0], argv[1], argv[2])
	// gopp.Println(rv, sym, slotname)
	gopp.GOUSED(rv)
}

// 单独提取出来，因为它要转换参数为QVariant
// slotname 不需要参数类型，foo(int), 那么直接传递foo
func (me QMetaObject) InvokeQmlmf(obj QObject, slotname string, args ...any) {
	var argv [3]voidptr

	a0 := &QArgument{}
	if false {
		a0.Data = QVarintNew2(123).Cthis
		a0.Tyname = cgopp.CStringaf("QVariant")
		argv[0] = (voidptr)(a0)
	}

	for i := 0; i < len(argv) && i < len(args); i++ {
		v := QVarintNew2(args[i])

		a := &QArgument{}
		a.Data = v.Cthis
		// a.Tyname = cgopp.StrtoRefc("QVariant")
		a.Tyname = cgopp.CStringaf("QVariant")
		argv[i] = (voidptr)(a)

	}

	// gopp.Println(argv)
	symname := "QMetaObjectInvokeMethod2"
	sym := dlsym(symname)
	name4c := cgopp.StrtoRefc(&slotname)
	// log.Println(args, slotname, argv)
	rv := cgopp.Litfficallg(sym, obj.Cthis, name4c, argv[0], argv[1], argv[2])
	// gopp.Println(rv, sym, slotname)
	gopp.GOUSED(rv)
}

// todo how simple get root object
func Qmljsgc2(robj QObject) {
	me := QMetaObjectof0()
	me.InvokeQmlmf(robj, "jsgc")
}
