/****************************************************************************
** QMitmSloter meta object code from reading C++ file 'qmitmsloter.h'
**
** Created: Thu Oct 16 09:45:49 2025
**      by: The Qt MOC ($Id: qt/moc_yacc.cpp   3.3.8   edited Feb 2 14:59 $)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#undef QT_NO_COMPAT
#include "qmitmsloter.h"
#include <qmetaobject.h>
#include <qapplication.h>

#include <private/qucomextra_p.h>
#if !defined(Q_MOC_OUTPUT_REVISION) || (Q_MOC_OUTPUT_REVISION != 26)
#error "This file was generated using the moc from 3.3.8b. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

const char *QMitmSloter::className() const
{
    return "QMitmSloter";
}

QMetaObject *QMitmSloter::metaObj = 0;
static QMetaObjectCleanUp cleanUp_QMitmSloter( "QMitmSloter", &QMitmSloter::staticMetaObject );

#ifndef QT_NO_TRANSLATION
QString QMitmSloter::tr( const char *s, const char *c )
{
    if ( qApp )
	return qApp->translate( "QMitmSloter", s, c, QApplication::DefaultCodec );
    else
	return QString::fromLatin1( s );
}
#ifndef QT_NO_TRANSLATION_UTF8
QString QMitmSloter::trUtf8( const char *s, const char *c )
{
    if ( qApp )
	return qApp->translate( "QMitmSloter", s, c, QApplication::UnicodeUTF8 );
    else
	return QString::fromUtf8( s );
}
#endif // QT_NO_TRANSLATION_UTF8

#endif // QT_NO_TRANSLATION

QMetaObject* QMitmSloter::staticMetaObject()
{
    if ( metaObj )
	return metaObj;
    QMetaObject* parentObject = QObject::staticMetaObject();
    static const QUMethod slot_0 = {"dummy", 0, 0 };
    static const QUParameter param_slot_1[] = {
	{ 0, &static_QUType_int, 0, QUParameter::In }
    };
    static const QUMethod slot_1 = {"dummy", 1, param_slot_1 };
    static const QUParameter param_slot_2[] = {
	{ 0, &static_QUType_ptr, "short", QUParameter::In }
    };
    static const QUMethod slot_2 = {"dummy", 1, param_slot_2 };
    static const QUParameter param_slot_3[] = {
	{ 0, &static_QUType_bool, 0, QUParameter::In }
    };
    static const QUMethod slot_3 = {"dummy", 1, param_slot_3 };
    static const QUParameter param_slot_4[] = {
	{ 0, &static_QUType_ptr, "float", QUParameter::In }
    };
    static const QUMethod slot_4 = {"dummy", 1, param_slot_4 };
    static const QUParameter param_slot_5[] = {
	{ 0, &static_QUType_double, 0, QUParameter::In }
    };
    static const QUMethod slot_5 = {"dummy", 1, param_slot_5 };
    static const QUParameter param_slot_6[] = {
	{ 0, &static_QUType_ptr, "long long", QUParameter::In }
    };
    static const QUMethod slot_6 = {"dummy", 1, param_slot_6 };
    static const QUParameter param_slot_7[] = {
	{ 0, &static_QUType_ptr, "void", QUParameter::In }
    };
    static const QUMethod slot_7 = {"dummy", 1, param_slot_7 };
    static const QUParameter param_slot_8[] = {
	{ 0, &static_QUType_ptr, "QObject", QUParameter::In }
    };
    static const QUMethod slot_8 = {"dummy", 1, param_slot_8 };
    static const QUParameter param_slot_9[] = {
	{ 0, &static_QUType_QString, 0, QUParameter::In }
    };
    static const QUMethod slot_9 = {"dummy", 1, param_slot_9 };
    static const QUParameter param_slot_10[] = {
	{ 0, &static_QUType_QVariant, 0, QUParameter::In }
    };
    static const QUMethod slot_10 = {"dummy", 1, param_slot_10 };
    static const QUParameter param_slot_11[] = {
	{ 0, &static_QUType_ptr, "char", QUParameter::In },
	{ 0, &static_QUType_ptr, "long", QUParameter::In },
	{ 0, &static_QUType_ptr, "float", QUParameter::In },
	{ 0, &static_QUType_double, 0, QUParameter::In },
	{ 0, &static_QUType_int, 0, QUParameter::In },
	{ 0, &static_QUType_ptr, "short", QUParameter::In },
	{ 0, &static_QUType_bool, 0, QUParameter::In },
	{ 0, &static_QUType_ptr, "void", QUParameter::In },
	{ 0, &static_QUType_ptr, "QObject", QUParameter::In },
	{ 0, &static_QUType_QString, 0, QUParameter::In },
	{ 0, &static_QUType_QString, 0, QUParameter::InOut },
	{ 0, &static_QUType_QVariant, 0, QUParameter::In },
	{ 0, &static_QUType_QVariant, 0, QUParameter::InOut }
    };
    static const QUMethod slot_11 = {"slotxx", 13, param_slot_11 };
    static const QMetaData slot_tbl[] = {
	{ "dummy()", &slot_0, QMetaData::Public },
	{ "dummy(int)", &slot_1, QMetaData::Public },
	{ "dummy(short)", &slot_2, QMetaData::Public },
	{ "dummy(bool)", &slot_3, QMetaData::Public },
	{ "dummy(float)", &slot_4, QMetaData::Public },
	{ "dummy(double)", &slot_5, QMetaData::Public },
	{ "dummy(long long)", &slot_6, QMetaData::Public },
	{ "dummy(void*)", &slot_7, QMetaData::Public },
	{ "dummy(QObject*)", &slot_8, QMetaData::Public },
	{ "dummy(QString)", &slot_9, QMetaData::Public },
	{ "dummy(QVariant)", &slot_10, QMetaData::Public },
	{ "slotxx(char,long,float,double,int,short,bool,void*,QObject*,QString,QString&,QVariant,QVariant&)", &slot_11, QMetaData::Public }
    };
    metaObj = QMetaObject::new_metaobject(
	"QMitmSloter", parentObject,
	slot_tbl, 12,
	0, 0,
#ifndef QT_NO_PROPERTIES
	0, 0,
	0, 0,
#endif // QT_NO_PROPERTIES
	0, 0 );
    cleanUp_QMitmSloter.setMetaObject( metaObj );
    return metaObj;
}

void* QMitmSloter::qt_cast( const char* clname )
{
    if ( !qstrcmp( clname, "QMitmSloter" ) )
	return this;
    return QObject::qt_cast( clname );
}

bool QMitmSloter::qt_invoke( int _id, QUObject* _o )
{
    metacallir(_id,_o);
	switch ( _id - staticMetaObject()->slotOffset() ) {
    case 0: dummy(); break;
    case 1: dummy((int)static_QUType_int.get(_o+1)); break;
    case 2: dummy((short)(*((short*)static_QUType_ptr.get(_o+1)))); break;
    case 3: dummy((bool)static_QUType_bool.get(_o+1)); break;
    case 4: dummy((float)(*((float*)static_QUType_ptr.get(_o+1)))); break;
    case 5: dummy((double)static_QUType_double.get(_o+1)); break;
    case 6: dummy((long long)(*((long long*)static_QUType_ptr.get(_o+1)))); break;
    case 7: dummy((void*)static_QUType_ptr.get(_o+1)); break;
    case 8: dummy((QObject*)static_QUType_ptr.get(_o+1)); break;
    case 9: dummy((QString)static_QUType_QString.get(_o+1)); break;
    case 10: dummy((QVariant)static_QUType_QVariant.get(_o+1)); break;
    case 11: slotxx((char)(*((char*)static_QUType_ptr.get(_o+1))),(long)(*((long*)static_QUType_ptr.get(_o+2))),(float)(*((float*)static_QUType_ptr.get(_o+3))),(double)static_QUType_double.get(_o+4),(int)static_QUType_int.get(_o+5),(short)(*((short*)static_QUType_ptr.get(_o+6))),(bool)static_QUType_bool.get(_o+7),(void*)static_QUType_ptr.get(_o+8),(QObject*)static_QUType_ptr.get(_o+9),(QString)static_QUType_QString.get(_o+10),(QString&)static_QUType_QString.get(_o+11),(QVariant)static_QUType_QVariant.get(_o+12),(QVariant&)static_QUType_QVariant.get(_o+13)); break;
    default:
	return QObject::qt_invoke( _id, _o );
    }
    return TRUE;
}

bool QMitmSloter::qt_emit( int _id, QUObject* _o )
{
    return QObject::qt_emit(_id,_o);
}
#ifndef QT_NO_PROPERTIES

bool QMitmSloter::qt_property( int id, int f, QVariant* v)
{
    return QObject::qt_property( id, f, v);
}

bool QMitmSloter::qt_static_property( QObject* , int , int , QVariant* ){ return FALSE; }
#endif // QT_NO_PROPERTIES
