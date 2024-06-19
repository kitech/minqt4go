#ifndef _QTUNINLINE_H_
#define _QTUNINLINE_H_

class QObject;
class QVariant;
class QQmlApplicationEngine;

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

void QMetaObjectInvokeMethod1(void* fnptrx, void* n);

QObject* QObjectFindChild1(QObject*obj, char*str);
QVariant* QObjectProperty1(QObject*obj, char*str);

// qml
QQmlApplicationEngine* QQmlApplicationEngineNew();
void QQmlApplicationEngineLoad1(QQmlApplicationEngine*e, char*str);
QObject* QQmlApplicationEngineRootObject1(QQmlApplicationEngine*e);

#ifdef __cplusplus
};
#endif

#endif // _QTUNINLINE_H_