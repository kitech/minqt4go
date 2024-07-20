package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/kitech/minqt"

	"github.com/kitech/gopp/cgopp"
)

type TypeMatcher interface {
	Match(d *TMCData, conv bool) bool
}

type TypeConver interface {
	Conv() any
}

type TMCData struct {
	idx    int
	ctys   string
	gotyo  reflect.Type
	goargx any

	// results
	ffiargx any

	// tmps
	// gotystr string
}

func (me *TMCData) Dbgstr() string {
	return fmt.Sprintf("idx: %v cty: %v goty: %v", me.idx, me.ctys, me.gotyo.String())
}

var typemcers = []TypeMatcher{
	&TMCEQ{}, &TMCTocxref{}, &TMCTocxCharpp{},
	&TMCQtptr{}, &TMCToQStrref{},
}

// ///
type TMCEQ struct{}

func (me *TMCEQ) Match(d *TMCData, conv bool) bool {
	if d.ctys == d.gotyo.String() {
		if conv {
			d.ffiargx = d.goargx
		}
		return true
	}
	return false
}

type TMCTocxref struct{}

func (me *TMCTocxref) Match(d *TMCData, conv bool) bool {
	if d.gotyo.String()+"&" == d.ctys {
		if conv {
			// 只对primitive type可以
			refval := reflect.New(d.gotyo)
			refval.Elem().Set(reflect.ValueOf(d.goargx))
			d.ffiargx = refval.Interface()
		}
		return true
	}
	return false
}

type TMCTocxCharpp struct{}

func (me *TMCTocxCharpp) Match(d *TMCData, conv bool) bool {
	if d.gotyo.String() == "[]string" && d.ctys == "char**" {
		if conv {
			// todo how freeit
			ptr := cgopp.CStrArrFromStrs(d.goargx.([]string))
			d.ffiargx = ptr
		}
		return true
	}
	return false
}

type TMCQtptr struct{}

func isqtptrtymat(tystr string, tyo reflect.Type) bool {
	// QObject* ?<= *main.QObject
	// goty := tyo.String()
	if tyo.Kind() == reflect.Pointer {
		ety := tyo.Elem()
		log.Println(ety, ety.Name())
		if ety.Name()+"*" == tystr {
			return true
		}
	}

	return false
}
func (me *TMCQtptr) Match(d *TMCData, conv bool) bool {
	tyo := d.gotyo
	tystr := d.ctys
	argx := d.goargx

	if isqtptrtymat(tystr, tyo) {
		if conv {
			tvx := reflect.ValueOf(argx)
			if tvx.IsNil() {

			} else {
				// .Elem().FieldByName("Cthis")
			}
			log.Println(tvx, d.Dbgstr())
		}
		return true
	}
	return false
}

type TMCToQStrref struct{}

func (me *TMCToQStrref) Match(d *TMCData, conv bool) bool {
	// QString const& ?<= string
	if "string" == d.gotyo.String() {
		if strings.HasPrefix(d.ctys, "QString ") && strings.HasSuffix(d.ctys, "&") {
			if conv {
				goval := minqt.QStringNew(d.goargx.(string))
				d.ffiargx = goval.Cthis
			}
			return true
		}
	}
	return false
}
