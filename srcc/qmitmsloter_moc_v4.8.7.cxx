/****************************************************************************
** Meta object code from reading C++ file 'qmitmsloter.h'
**
** Created by: The Qt Meta Object Compiler version 63 (Qt 4.8.7)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include "qmitmsloter.h"
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'qmitmsloter.h' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 63
#error "This file was generated using the moc from 4.8.7. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
static const uint qt_meta_data_QMitmSloter[] = {

 // content:
       6,       // revision
       0,       // classname
       0,    0, // classinfo
       2,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: signature, parameters, type, tag, flags
      13,   12,   12,   12, 0x0a,
      34,   21,   12,   12, 0x0a,

       0        // eod
};

static const char qt_meta_stringdata_QMitmSloter[] = {
    "QMitmSloter\0\0dummy()\0,,,,,,,,,,,,\0"
    "slotxx(char,long,float,double,int,short,bool,void*,QObject*,QString,QS"
    "tring&,QVariant,QVariant&)\0"
};

void QMitmSloter::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        Q_ASSERT(staticMetaObject.cast(_o));
        QMitmSloter *_t = static_cast<QMitmSloter *>(_o);
        switch (_id) {
        case 0: _t->metacallir(_o,_c,_id,_a); break;
        case 1: _t->slotxx((*reinterpret_cast< char(*)>(_a[1])),(*reinterpret_cast< long(*)>(_a[2])),(*reinterpret_cast< float(*)>(_a[3])),(*reinterpret_cast< double(*)>(_a[4])),(*reinterpret_cast< int(*)>(_a[5])),(*reinterpret_cast< short(*)>(_a[6])),(*reinterpret_cast< bool(*)>(_a[7])),(*reinterpret_cast< void*(*)>(_a[8])),(*reinterpret_cast< QObject*(*)>(_a[9])),(*reinterpret_cast< QString(*)>(_a[10])),(*reinterpret_cast< QString(*)>(_a[11])),(*reinterpret_cast< QVariant(*)>(_a[12])),(*reinterpret_cast< QVariant(*)>(_a[13]))); break;
        default: ;
        }
    }
}

const QMetaObjectExtraData QMitmSloter::staticMetaObjectExtraData = {
    0,  qt_static_metacall 
};

const QMetaObject QMitmSloter::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_QMitmSloter,
      qt_meta_data_QMitmSloter, &staticMetaObjectExtraData }
};

#ifdef Q_NO_DATA_RELOCATION
const QMetaObject &QMitmSloter::getStaticMetaObject() { return staticMetaObject; }
#endif //Q_NO_DATA_RELOCATION

const QMetaObject *QMitmSloter::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->metaObject : &staticMetaObject;
}

void *QMitmSloter::qt_metacast(const char *_clname)
{
    if (!_clname) return 0;
    if (!strcmp(_clname, qt_meta_stringdata_QMitmSloter))
        return static_cast<void*>(const_cast< QMitmSloter*>(this));
    return QObject::qt_metacast(_clname);
}

int QMitmSloter::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 2)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    }
    return _id;
}
QT_END_MOC_NAMESPACE
