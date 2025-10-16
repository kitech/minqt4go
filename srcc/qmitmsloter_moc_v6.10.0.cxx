/****************************************************************************
** Meta object code from reading C++ file 'qmitmsloter.h'
**
** Created by: The Qt Meta Object Compiler version 69 (Qt 6.10.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include "qmitmsloter.h"
#include <QtCore/qmetatype.h>

#include <QtCore/qtmochelpers.h>

#include <memory>


#include <QtCore/qxptype_traits.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'qmitmsloter.h' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 69
#error "This file was generated using the moc from 6.10.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

#ifndef Q_CONSTINIT
#define Q_CONSTINIT
#endif

QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
QT_WARNING_DISABLE_GCC("-Wuseless-cast")
namespace {
struct qt_meta_tag_ZN11QMitmSloterE_t {};
} // unnamed namespace

template <> constexpr inline auto QMitmSloter::qt_create_metaobjectdata<qt_meta_tag_ZN11QMitmSloterE_t>()
{
    namespace QMC = QtMocConstants;
    QtMocHelpers::StringRefStorage qt_stringData {
        "QMitmSloter",
        "dummy",
        "",
        "QVariant",
        "slotxx",
        "QString&",
        "QVariant&"
    };

    QtMocHelpers::UintData qt_methods {
        // Slot 'dummy'
        QtMocHelpers::SlotData<void()>(1, 2, QMC::AccessPublic, QMetaType::Void),
        // Slot 'dummy'
        QtMocHelpers::SlotData<void(int)>(1, 2, QMC::AccessPublic, QMetaType::Void, {{
            { QMetaType::Int, 2 },
        }}),
        // Slot 'dummy'
        QtMocHelpers::SlotData<void(short)>(1, 2, QMC::AccessPublic, QMetaType::Void, {{
            { QMetaType::Short, 2 },
        }}),
        // Slot 'dummy'
        QtMocHelpers::SlotData<void(bool)>(1, 2, QMC::AccessPublic, QMetaType::Void, {{
            { QMetaType::Bool, 2 },
        }}),
        // Slot 'dummy'
        QtMocHelpers::SlotData<void(float)>(1, 2, QMC::AccessPublic, QMetaType::Void, {{
            { QMetaType::Float, 2 },
        }}),
        // Slot 'dummy'
        QtMocHelpers::SlotData<void(double)>(1, 2, QMC::AccessPublic, QMetaType::Void, {{
            { QMetaType::Double, 2 },
        }}),
        // Slot 'dummy'
        QtMocHelpers::SlotData<void(long long)>(1, 2, QMC::AccessPublic, QMetaType::Void, {{
            { QMetaType::LongLong, 2 },
        }}),
        // Slot 'dummy'
        QtMocHelpers::SlotData<void(void *)>(1, 2, QMC::AccessPublic, QMetaType::Void, {{
            { QMetaType::VoidStar, 2 },
        }}),
        // Slot 'dummy'
        QtMocHelpers::SlotData<void(QObject *)>(1, 2, QMC::AccessPublic, QMetaType::Void, {{
            { QMetaType::QObjectStar, 2 },
        }}),
        // Slot 'dummy'
        QtMocHelpers::SlotData<void(QString)>(1, 2, QMC::AccessPublic, QMetaType::Void, {{
            { QMetaType::QString, 2 },
        }}),
        // Slot 'dummy'
        QtMocHelpers::SlotData<void(QVariant)>(1, 2, QMC::AccessPublic, QMetaType::Void, {{
            { 0x80000000 | 3, 2 },
        }}),
        // Slot 'slotxx'
        QtMocHelpers::SlotData<void(char, long, float, double, int, short, bool, void *, QObject *, QString, QString &, QVariant, QVariant &)>(4, 2, QMC::AccessPublic, QMetaType::Void, {{
            { QMetaType::Char, 2 }, { QMetaType::Long, 2 }, { QMetaType::Float, 2 }, { QMetaType::Double, 2 },
            { QMetaType::Int, 2 }, { QMetaType::Short, 2 }, { QMetaType::Bool, 2 }, { QMetaType::VoidStar, 2 },
            { QMetaType::QObjectStar, 2 }, { QMetaType::QString, 2 }, { 0x80000000 | 5, 2 }, { 0x80000000 | 3, 2 },
            { 0x80000000 | 6, 2 },
        }}),
    };
    QtMocHelpers::UintData qt_properties {
    };
    QtMocHelpers::UintData qt_enums {
    };
    return QtMocHelpers::metaObjectData<QMitmSloter, qt_meta_tag_ZN11QMitmSloterE_t>(QMC::MetaObjectFlag{}, qt_stringData,
            qt_methods, qt_properties, qt_enums);
}
Q_CONSTINIT const QMetaObject QMitmSloter::staticMetaObject = { {
    QMetaObject::SuperData::link<QObject::staticMetaObject>(),
    qt_staticMetaObjectStaticContent<qt_meta_tag_ZN11QMitmSloterE_t>.stringdata,
    qt_staticMetaObjectStaticContent<qt_meta_tag_ZN11QMitmSloterE_t>.data,
    qt_static_metacall,
    nullptr,
    qt_staticMetaObjectRelocatingContent<qt_meta_tag_ZN11QMitmSloterE_t>.metaTypes,
    nullptr
} };

void QMitmSloter::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    auto *_t = static_cast<QMitmSloter *>(_o);
    if (_c == QMetaObject::InvokeMetaMethod) {
        metacallir(_o,_c,_id,_a);
	switch (_id) {
        case 0: _t->dummy(); break;
        case 1: _t->dummy((*reinterpret_cast<std::add_pointer_t<int>>(_a[1]))); break;
        case 2: _t->dummy((*reinterpret_cast<std::add_pointer_t<short>>(_a[1]))); break;
        case 3: _t->dummy((*reinterpret_cast<std::add_pointer_t<bool>>(_a[1]))); break;
        case 4: _t->dummy((*reinterpret_cast<std::add_pointer_t<float>>(_a[1]))); break;
        case 5: _t->dummy((*reinterpret_cast<std::add_pointer_t<double>>(_a[1]))); break;
        case 6: _t->dummy((*reinterpret_cast<std::add_pointer_t<qlonglong>>(_a[1]))); break;
        case 7: _t->dummy((*reinterpret_cast<std::add_pointer_t<void*>>(_a[1]))); break;
        case 8: _t->dummy((*reinterpret_cast<std::add_pointer_t<QObject*>>(_a[1]))); break;
        case 9: _t->dummy((*reinterpret_cast<std::add_pointer_t<QString>>(_a[1]))); break;
        case 10: _t->dummy((*reinterpret_cast<std::add_pointer_t<QVariant>>(_a[1]))); break;
        case 11: _t->slotxx((*reinterpret_cast<std::add_pointer_t<char>>(_a[1])),(*reinterpret_cast<std::add_pointer_t<long>>(_a[2])),(*reinterpret_cast<std::add_pointer_t<float>>(_a[3])),(*reinterpret_cast<std::add_pointer_t<double>>(_a[4])),(*reinterpret_cast<std::add_pointer_t<int>>(_a[5])),(*reinterpret_cast<std::add_pointer_t<short>>(_a[6])),(*reinterpret_cast<std::add_pointer_t<bool>>(_a[7])),(*reinterpret_cast<std::add_pointer_t<void*>>(_a[8])),(*reinterpret_cast<std::add_pointer_t<QObject*>>(_a[9])),(*reinterpret_cast<std::add_pointer_t<QString>>(_a[10])),(*reinterpret_cast<std::add_pointer_t<QString&>>(_a[11])),(*reinterpret_cast<std::add_pointer_t<QVariant>>(_a[12])),(*reinterpret_cast<std::add_pointer_t<QVariant&>>(_a[13]))); break;
        default: ;
        }
    }
}

const QMetaObject *QMitmSloter::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QMitmSloter::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_staticMetaObjectStaticContent<qt_meta_tag_ZN11QMitmSloterE_t>.strings))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int QMitmSloter::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 12)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 12;
    }
    if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 12)
            *reinterpret_cast<QMetaType *>(_a[0]) = QMetaType();
        _id -= 12;
    }
    return _id;
}
QT_WARNING_POP
