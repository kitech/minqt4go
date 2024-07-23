#include <QtCore>

#include "qdummyobject.h"


extern "C" void mystatic_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a) {
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<QDynSlotObjecT *>(_o);
        (void)_t;
        switch (_id) {
        case 0: _t->dumnyslot(); break;
        default: ;
        }
    }
    (void)_a;
}

#include <stdlib.h>
QDynSlotObjecT::QDynSlotObjecT(QObject* parent) : QObject (parent) {
    auto f = &QDynSlotObjecT::qt_static_metacall;
}

void QDynSlotObjecT::dumnyslot() {
    qDebug()<<__FUNCTION__<<__LINE__<<"hehehe";
}

