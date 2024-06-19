package minqt

import (
	"log"
	"reflect"

	"github.com/ebitengine/purego"
	_ "github.com/ebitengine/purego"

	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
)

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

// current process
func dlsym(name string) voidptr {
	sym, err := purego.Dlsym(purego.RTLD_DEFAULT, name)
	gopp.ErrPrint(err, name)
	if gopp.ErrHave(err, "symbol not found") {
		// sym, err = purego.Dlsym(purego.RTLD_DEFAULT, "_"+name)
	}
	return voidptr(sym)
}

type QObject struct {
	Cthis voidptr
}

func QObjectof(ptr voidptr) QObject {
	return QObject{ptr}
}

func (me QObject) SetProperty(name string, valuex any) bool {

	var value = QVarintNew2(valuex)

	symname := "__ZN7QObject11setPropertyEPKcRK8QVariant"
	sym := dlsym(symname)
	rv := cgopp.Litfficallg(sym, me.Cthis, cgopp.StrtoVptrRef(&name), value.Cthis)
	log.Println(rv)
	return true
}

// int/str/list???
func (me QObject) Property(name string) gopp.Any {
	symname := "__ZNK7QObject8propertyEPKc"
	sym := dlsym(symname)
	name4c := cgopp.CString(name)
	defer cgopp.Cfree(name4c)
	rv := cgopp.Litfficallg(sym, me.Cthis, name4c)
	// log.Println(rv)
	return gopp.ToAny(rv)
}
func (me QObject) FindChild(objname string) QObject {

	sym := dlsym("QObjectFindChild1")
	on4c := cgopp.StrtoVptrRef(&objname)
	rv := cgopp.Litfficallg(sym, me.Cthis, on4c)
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
	log.Println(reflect.TypeOf(any(vx)), vx)
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
		v4cc := cgopp.CString(v)
		defer cgopp.Cfree(v4cc)
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

type QQmlApplicationEngine struct{ Cthis voidptr }

func QQmlApplicationEngineof(ptr voidptr) QQmlApplicationEngine {
	return QQmlApplicationEngine{ptr}
}
func QQmlApplicationEngineNew() QQmlApplicationEngine {
	sym := dlsym("__ZN21QQmlApplicationEngineC1EP7QObject")
	rv := cgopp.Litfficallg(sym, nil)
	return QQmlApplicationEngineof(rv)
}
func (me QQmlApplicationEngine) Load(u string) {
	sym := dlsym("QQmlApplicationEngineLoad1")
	v4cc := cgopp.CString(u)
	defer cgopp.Cfree(v4cc)
	cgopp.Litfficallg(sym, v4cc)
}

func (me QQmlApplicationEngine) RootObject() QObject {
	sym := dlsym("QQmlApplicationEngineRootObject1")
	rv := cgopp.Litfficallg(sym, me.Cthis)
	return QObjectof(rv)
}
