package qtinline

/*
#cgo LDFLAGS: -lQtCoreInline -L../srcc/

extern void set_callbackAllInherits(void*);
*/
import "C"

func Keep() {}

func dummy() {
	C.set_callbackAllInherits(nil)
}
