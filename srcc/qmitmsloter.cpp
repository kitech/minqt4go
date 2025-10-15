
#include <iostream>

#include "qmitmsloter.h"

void QMitmSloter::dummy() {}

// maybe useful, senderSignalIndex(>=qt4)
// QSignalMapper , QSignalSpy

typedef struct ConnCbdata {
	void* sender;
	void* fnptr;
	void* cbptr;
	int argc;
	void* argtys; //int argtys[];
	void* fatptr0;
	void* fatptr1;
} ConnCbdata;

QMitmSloter::QMitmSloter(void* d_) : QObject() {
	this->cbdata = d_;
}

#if QT_VERSION < 0x040000
void QMitmSloter::metacallir(int _id, QUObject* _o) {
	ConnCbdata* d = (ConnCbdata*)this->cbdata;
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
