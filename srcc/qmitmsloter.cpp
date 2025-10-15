
#include <iostream>

#include "qmitmsloter.h"

void QMitmSloter::dummy() {}

// maybe useful, senderSignalIndex(>=qt4)
// QSignalMapper , QSignalSpy

// see qtrt.ConnCbdata
typedef struct ConnCbdata {
	void* fnptr;
	void* sender;
	void* fatptr0; // vstring
	void* fatptr1;
} ConnCbdata;

QMitmSloter::QMitmSloter(void* d_) : QObject() {
	this->cbdata = d_;
}

#if QT_VERSION < 0x040000
void QMitmSloter::metacallir(int _id, QUObject* _o) {
	ConnCbdata* d = (ConnCbdata*)this->cbdata;
	// std::cout<<__FUNCTION__<<_id<<_o<<std::endl;
	printf("%s: %d, %p, d=%p\n", __FUNCTION__, _id, _o, this->cbdata);
	
	typedef void (*cbfnty)(void*, int, void*);
	cbfnty fno = (cbfnty)d->fnptr;
	fno(d, _id, _o);
}
#else
void QMitmSloter::metacallir(QObject *_o, QMetaObject::Call _c, int _id, void **_a);
{
	ConnCbdata* d = (ConnCbdata*)this->cbdata;
}
#endif

#if QT_VERSION < 0x040000
#include "qmitmsloter_moc_v3.3.8b.cxx"
#elif QT_VERSION < 0x050000
#include "qmitmsloter_moc_v4.8.7.cxx"
#elif QT_VERSION < 0x060000
#include "qmitmsloter_moc_v5.15.17.cxx"
#elif QT_VERSION < 0x070000
#include "qmitmsloter_moc_v6.0.0.cxx"
#else
#warning "not support QT_VERSION"
#endif
