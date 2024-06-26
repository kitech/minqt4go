
#include <QtCore>
#include <QtQml>
#include <QtQuick>
#include <QtQuickControls2>
#include <QtQuickTemplates2>

#include "qtuninline.h"

const char* QCompileVersion() { return QT_VERSION_STR; }

// __attribute__((noinline))
void QVariantDtor(void*p) {
    if (p!= nullptr) delete (QVariant*)p;
}

void* QVariantNewInt(int v) {
    auto rv = new QVariant(v);
    return rv;
}
int QVariantToint(QVariant*p) {
    bool ok = false;
    return p->toInt(&ok);
}
void* QVariantNewInt64(qint64 v) {
    auto rv = new QVariant(v);
    return rv;
}
qint64 QVariantToint64(QVariant*p) {
    bool ok = false;
    return p->toLongLong(&ok);
}

void* QVariantNewStr(char*str) {
    auto rv = new QVariant(QString(str));
    return rv;
}
char* QVariantTostr(QVariant*p) {
    return p->toString().toUtf8().data();
}

QVariant* QVariantNewPtr(void*ptr) {
    return new QVariant(quint64(ptr));
}
void* QVariantToptr(QVariant*p) {
    bool ok = false;
    // auto rv = p->toULongLong(&ok);
    // auto rv = p->value<QQuickItem*>(); // 这种方式对的
    // auto x = p->convert(QMetaType::type("void*")); // not work
    // auto rv = p->value<void*>(); // not work
    auto rv = p->value<QObject*>(); // works
    // qDebug()<<__FUNCTION__<<(*p)<<(void*)p<<(*p).data()<<rv;
    return (void*)rv;
}

void* QVariantNewBool(bool v) {
    auto rv = new QVariant(v);
    // qDebug()<<__FUNCTION__<<__LINE__<<v<<*rv<<sizeof(bool);
    return rv;
}
bool QVariantTobool(QVariant*p) {
    return p->toBool();
}

///////////
// for QtObject::createQmlObject, or other also ok
void QObjectDtor(QObject* o) { delete o; }

void QMetaObjectInvokeMethod1(void* fnptrx, void* n) {
    QObject* o = qApp;
    void (*fnptr)(void*) = (void(*)(void*))fnptrx;
    QMetaObject::invokeMethod(o, [fnptr,n]{ fnptr(n); }, Qt::QueuedConnection);
}

// 可以直接调用Qt slot，以及qml的函数
int QMetaObjectInvokeMethod2(QObject* obj, char* member, void*val0,void*val1,void*val2) {
    QGenericReturnArgument qret;
    QGenericArgument a0;
    QGenericArgument a1;
    QGenericArgument a2;

    if (val0 != nullptr) a0 = *((QGenericArgument*)val0);
    if (val1 != nullptr) a1 = *((QGenericArgument*)val1);
    if (val2 != nullptr) a2 = *((QGenericArgument*)val2);

    // qDebug()<<__FUNCTION__<<__LINE__<<obj<<member<<val0;

    int rv = QMetaObject::invokeMethod(obj, member, Qt::QueuedConnection, qret, a0, a1, a2);
    return rv;
}

QObject* QObjectFindChild1(QObject*obj, char*str) {
    auto rv = obj->findChild<QObject*>(str);
    return rv;
}
QVariant* QObjectProperty1(QObject*obj, char*str) {
    auto rv = obj->property(str);
    // qDebug()<<__FUNCTION__<<__LINE__<<obj<<rv<<QString(str);
    if (!rv.isValid()) {
        return nullptr;
    }
    auto rv2 = new QVariant(rv); //
    // qDebug()<<__FUNCTION__<<__LINE__<<*rv2<<QString(str)<<(*rv2).data();
    return rv2;
}

// qml
QQmlApplicationEngine* QQmlApplicationEngineNew() {
    auto e = new QQmlApplicationEngine();
    return e;
}
void QQmlApplicationEngineLoad1(QQmlApplicationEngine*e, char*str) {
    auto url = QUrl(str);
    e->load(url);
}
QObject* QQmlApplicationEngineRootObject1(QQmlApplicationEngine*e) {
    auto objs = e->rootObjects();
    return objs.value(0);
}

QQmlComponent* QQmlComponentNew1(void*engine, QObject* parent) {
    auto rv = new QQmlComponent((QQmlEngine*)engine, parent);
    return rv;
}
QObject* QQmlComponentCreate(QQmlComponent*o, void*ctx) {
    auto rv = o->create((QQmlContext*)ctx);
    return rv;
}
void QQmlComponentSetData(QQmlComponent*o, char*data) {
    o->setData(QByteArray(data), QUrl());
}

#include <QtQml/private/qqmlbuiltinfunctions_p.h>
QtObject* QtObjectCreate(QQmlEngine*e) { return QtObject::create(e,e); }

QObject* QtObjectCreateQmlObject(void*o, char* qmltxt, QObject*parent) {
    auto o2 = (QtObject*)o;
    auto rv = o2->createQmlObject(QString(qmltxt), parent);
    return rv;
}

// quick templates
// 6.7 添加了许多方法,但是android的qtsdk现在还是6.6的...
// aqtinstall 大概在6.15日左右发布新版本,支持qt6.7sdk for android了
#include <QtQuickTemplates2/private/qquickstackview_p.h>
void dummyyy() {
    QQuickStackView *stkwin; // not work
}

QQuickItem* QQuickStackView_get(QQuickStackView*me, int idx) {
    QQuickItem* rv = me->get(idx);
    return rv;
}
#if QT_VERSION >= QT_VERSION_CHECK(6, 7, 0)
QQuickItem* QQuickStackView_replaceCurrentItem(QQuickStackView* me, QQuickItem* item) {
    auto rv = me->replaceCurrentItem(item);
    return rv;
}
#endif


/////
#include <dlfcn.h>
void* cgoir_dlsym0(const char* name) {
    return dlsym(RTLD_DEFAULT, name);
}