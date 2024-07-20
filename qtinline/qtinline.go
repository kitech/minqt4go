package qtinline

/*
#cgo LDFLAGS: -lQtInline -L../srcc/

extern void set_callbackAllInherits(void*);
*/
import "C"

func Keep() {}

func dummy() {
	C.set_callbackAllInherits(nil)
}
