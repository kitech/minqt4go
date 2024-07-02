#ifndef _QTUNINLINE_H_
#define _QTUNINLINE_H_

class QObject;
class QVariant;
class QQmlApplicationEngine;
class QQuickItem;
class QQuickStackView;
class QQmlComponent;

#ifdef __cplusplus
extern "C" {
#endif

// core
const char* QCompileVersion();

void QVariantDtor(void*p);
void* QVariantNewInt(int v);
int QVariantToint(QVariant*p);
void* QVariantNewInt64(qint64 v);
qint64 QVariantToint64(QVariant*p);
void* QVariantNewStr(char*str);
char* QVariantTostr(QVariant*p);
QVariant* QVariantNewPtr(void*ptr);
void* QVariantToptr(QVariant*p);
// void* QVariantNewListstr();
void* QVariantNewBool(bool v);
bool QVariantTobool(QVariant*p);


void QStringDtor(void*px);
void* QStringNew(const char*p);

void QObjectDtor(QObject* o);
void QMetaObjectInvokeMethod1(void* fnptrx, void* n);
int QMetaObjectInvokeMethod2(QObject* obj, char* member, void*a0, void*a1, void*a2);

QObject* QObjectFindChild1(QObject*obj, char*str);
QVariant* QObjectProperty1(QObject*obj, char*str);

// qml
QQmlApplicationEngine* QQmlApplicationEngineNew();
void QQmlApplicationEngineLoad1(QQmlApplicationEngine*e, char*str);
QObject* QQmlApplicationEngineRootObject1(QQmlApplicationEngine*e);

QQmlComponent* QQmlComponentNew1(void*engine, QObject* parent);
QObject* QQmlComponentCreate(QQmlComponent*o, void*ctx);
void QQmlComponentSetData(QQmlComponent*o, char*data);
QObject* QtObjectCreateQmlObject(void*o, char* qmltxt, QObject*parent);
class QtObject;
QtObject* QtObjectCreate(QQmlEngine*e);

// quick templates2
QQuickItem* QQuickStackView_get(QQuickStackView*me, int idx);
QQuickItem* QQuickStackView_replaceCurrentItem(QQuickStackView*me, QQuickItem* item);

void* cgoir_dlsym0(const char* name);

#ifdef __cplusplus
};
#endif

#endif // _QTUNINLINE_H_