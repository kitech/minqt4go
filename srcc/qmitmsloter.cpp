
#include <iostream>

#include "qmitmsloter.h"

// maybe useful, senderSignalIndex(>=qt4)
// QSignalMapper , QSignalSpy
///////
// see qtrt.ConnCbdata
typedef struct ConnCbdata {
	void* fnptr;
	void* sender;
	void* fatptr0; // vstring
	void* fatptr1;
} ConnCbdata;

QMitmSloter::QMitmSloter(void* d_) : QObject() {
	this->cbdata = d_; // d_ would be GCed???
	this->cbdata = malloc(sizeof(ConnCbdata));
	memcpy(this->cbdata, d_, sizeof(ConnCbdata));
}

#if QT_VERSION < 0x040000
void QMitmSloter::metacallir(int _id, QUObject* _a) {
	ConnCbdata* d = (ConnCbdata*)this->cbdata;
	// std::cout<<__FUNCTION__<<_id<<_o<<std::endl;
	printf("%s: %d, %p, d=%p \t:%s:%d\n", __FUNCTION__, _id, _a, this->cbdata, __FILE__, __LINE__);

	typedef void (*cbfnty)(void*, void*, int, int, void*);
	cbfnty fno = (cbfnty)d->fnptr;
	fno(d, this, 0,  _id, _a);
}
#else
// _o: receiver object, this
// _c: direct/queued
// _id: slot inner id/index
// _a: arguments with simple packed format
void QMitmSloter::metacallir(QObject *_o, QMetaObject::Call _c, int _id, void **_a);
{
	ConnCbdata* d = (ConnCbdata*)this->cbdata;
	typedef void (*cbfnty)(void*, void*, int, int, void*);
	cbfnty fno = (cbfnty)d->fnptr;
	fno(d, _o, _c, _id, _a);
}
#endif

#if QT_VERSION < 0x040000
#include "qmitmsloter_moc_v3.3.8b.cxx"
#elif QT_VERSION < 0x050000
#include "qmitmsloter_moc_v4.8.7.cxx"
#elif QT_VERSION < 0x060000
#include "qmitmsloter_moc_v5.15.17.cxx"
#elif QT_VERSION < 0x070000
#include "qmitmsloter_moc_v6.10.0.cxx"
#else
#warning "not support QT_VERSION"
#endif
