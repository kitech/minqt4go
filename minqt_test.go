package minqt

import (
	"log"
	"testing"
)

func TestQVarint1(t *testing.T) {

	v := QVarintNew(12345)
	// log.Println(v)
	log.Println(v, v.Toint())
	v.Dtor()
	v = QVarintNew(int64(888))
	log.Println(v, v.Toint64())
	v.Dtor()

	v = QVarintNew("abcde")
	log.Println(v, v.Tostr())
	v.Dtor()

	var x = 123
	v = QVarintNew((voidptr)(&x))
	log.Println(v, v.Toptr(), &x)

}
