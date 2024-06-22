package minqt

import (
	"fmt"
	"sync/atomic"

	"github.com/ebitengine/purego"
	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
	cmap "github.com/orcaman/concurrent-map/v2"
)

/*
 */
import "C"

var lmrefs = cmap.New[*ListModelBase]()
var lmseq int64 = 10000

// 以下是与go交互的代码
type Datar interface {
	Data(name string) any
	DedupKey() string
	OrderKey() int64
}

func (me *ListModelBase) Add(d Datar) bool {
	me.BeginChangeRows(me.datas.Count(), me.datas.Count(), false)
	ok := me.datas.Put(d.DedupKey(), d)
	me.EndChangeRows(false)
	if ok {

	}

	return true
}

// like coloumn but in list
// used when ListModel.RoleName() call
var clazzrolenames = cmap.New[[]string]()

func RegisterModelRoleNames(clazz string, names ...string) {
	clazzrolenames.Set(clazz, names)
}

////// 以下是与cpp交互的代码

type ListModelBase struct {
	cppimpl voidptr
	seq     *int64

	clazz string

	///// data container
	datas *gopp.ListMap0[string, Datar]
}

func ListModelBaseof(seqptrx int64) *ListModelBase {
	seqptr := (seqptrx)
	key := fmt.Sprintf("%d", seqptr)
	p, ok := lmrefs.Get(key)
	gopp.FalsePrint(ok, "wtf", key)
	return p
}

//export goimplListModelBaseNew
func goimplListModelBaseNew(px voidptr) int64 {
	gopp.Info(px)
	me := ListModelBaseNew()
	me.cppimpl = px

	rv := new(int64)
	*rv = atomic.AddInt64(&lmseq, 3)
	me.seq = rv

	key := fmt.Sprintf("%d", *rv)
	lmrefs.Set(key, me)
	return *rv
}
func ListModelBaseNew() *ListModelBase {
	me := &ListModelBase{}
	me.datas = gopp.ListMap0New[string, Datar]()
	// 怎么确定 roleNames???

	return me
}

//export goimplListModelBaseDtor
func goimplListModelBaseDtor(px int64) {
	me := ListModelBaseof(px)
	me.Dtor()
}
func (me *ListModelBase) Dtor() {
	key := fmt.Sprintf("%d", *me.seq)
	lmrefs.Remove(key)
}

// //export goimplListModelBaseGetsetRolecnt
// func goimplListModelBaseGetsetRolecnt(px int64, c int, set int) int {
// 	gopp.Info(px, c, set)
// 	me := ListModelBaseof(px)

// 	a, ok := clazzrolenames.Get(me.clazz)
// 	if !ok {
// 		return -1
// 	}

// 	if set == 1 {
// 		// 最先调用的

// 	} else {
// 		return len(a)
// 	}
// 	return 0
// }

//export goimplListModelBaseGetsetClazz
func goimplListModelBaseGetsetClazz(px int64, c voidptr, set int) voidptr {
	gopp.Info(px, c, set, cgopp.GoString(c))
	me := ListModelBaseof(px)
	if set == 1 {
		me.clazz = cgopp.GoString(c)
	} else {
		return cgopp.CString(me.clazz)
	}
	return nil
}

//export goimplListModelBaseRoleName
func goimplListModelBaseRoleName(px int64, c int) voidptr {
	gopp.Info(px, c)

	me := ListModelBaseof(px)
	rv := me.RoleName(c)
	if len(rv) == 0 {
		return nil
	}
	return cgopp.CString(rv)
}

func init() {
	// var roleNames = map[int]string{256: "foo0", 257: "name", 258: "value", 259: "deleted"}
	// RegisterModelRoleNames("msglstmdl", gopp.MapValues(roleNames)...)
}

func (me *ListModelBase) RoleName(c int) string {
	a, ok := clazzrolenames.Get(me.clazz)
	if !ok {
		return gopp.ZeroStr
	}

	return gopp.ValueAt(a, c-256)
}

//export goimplListModelBaseRowCount
func goimplListModelBaseRowCount(px int64) int {
	// gopp.Info(px)
	me := ListModelBaseof(px)
	return me.RowCount()
}

func (me *ListModelBase) RowCount() int {
	return me.datas.Len()
	// return 3
}

//export goimplListModelBaseData
func goimplListModelBaseData(px int64, row int, role int) voidptr {
	gopp.Info(px, row, role)
	me := ListModelBaseof(px)
	return me.Data(row, role)
}

func (me *ListModelBase) Data(row, role int) voidptr {
	rv := QVarintNew(fmt.Sprintf("r%d of %d", row, role))
	_, dv, ok := me.datas.GetIndex(row)
	gopp.Info(rv, me.RoleName(role), dv, ok, row, role, me.datas.Len())
	if dv != nil {
		v2 := dv.Data(me.RoleName(role))
		rv = QVarintNew(fmt.Sprintf("%v", v2))
	}
	return rv.Cthis
}

func (me *ListModelBase) BeginChangeRows(first, last int, remove bool) {
	sym := dlsym("_ZN13ListModelBase19emitBeginChangeRowsEiii")
	var fno func(voidptr, int, int, int)
	purego.RegisterFunc(&fno, usize(sym))
	fno(me.cppimpl, first, last, gopp.IfElse2(remove, 1, 0))
}

func (me *ListModelBase) EndChangeRows(remove bool) {
	const name0 = "_ZN18QAbstractItemModel13endInsertRowsEv"
	const name1 = "_ZN18QAbstractItemModel13endRemoveRowsEv"
	sym := gopp.IfElse2(remove, dlsym(name1), dlsym(name0))

	var fno func(voidptr)
	purego.RegisterFunc(&fno, usize(sym))
	fno(me.cppimpl)
}
