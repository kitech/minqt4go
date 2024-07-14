package minqt

import (
	"fmt"
	"log"
	"math/rand"
)

func ShowToastTip(prptr voidptr, tip string) {
	// obj := qmlcpm.rootobj.FindChild("shrinfotip")
	// log.Println(obj.Dbgstr())
	// pr := obj.Parent()
	// log.Println(pr.Dbgstr())

	// po := qmlcpm.rootobj.FindChild("appwin")
	// sym := cgopp.Dlsym0("testtooltipincpp")
	// cgopp.FfiCall[int](sym, qmlcpm.rootobj.Cthis)
	tt := QQuickToolTipNew(prptr)
	tt.SetProperty("height", 33)
	tt.SetProperty("width", 330)
	tt.SetTimeout(2234)
	tt.SetText(fmt.Sprintf("text string %d", rand.Int()))
	tt.SetVisible(true)
	log.Println(tt.Text(), tt.Visible())
}
