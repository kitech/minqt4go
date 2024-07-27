#include <cxxabi.h>

#include <QtCore>
#include <QtGui>

#include "qtuninline.h"

#include <dlfcn.h>
auto cgoppMallocgcfn = (void* (*)(int))dlsym(RTLD_DEFAULT, "cgoppMallocgc");
// extern "C" void* cgoppMallocgc(int);

extern "C"
char* cxxabi__cxa_demangle(char*a0, char*a1, size_t *length, int *status) {
    return abi::__cxa_demangle(a0, a1, length, status);
}

#define DBGLOG qDebug()<<__FUNCTION__<<__LINE__
#define nilcxobj(x) ((x*)0)

static QString dummyqs("dummy&ref");
void* uninline_qtcore_holder() {
    uintptr_t retv = 0;
    retv += (uintptr_t)(qInstallMessageHandler(nullptr));

    nilcxobj(QMetaObject)->superClass();
    if (nilcxobj(QObject)->metaObject()!=nullptr) { }
    nilcxobj(QObject)->setObjectName(QAnyStringView(""));
    nilcxobj(QObject)->parent();
    if (nilcxobj(QMetaObject)->className()!=nullptr) {}
    {auto rv = nilcxobj(QObject)->findChild<QObject*>(""); retv+=(uintptr_t)rv;}

    // 如何取构造函数的地址！！
    // 如何取有重载的方法的地址！！
    // works
    {QString ( QString::*x )(int, int, int, QChar) const = &QString::arg;}
    // {auto x  = &QString::QString;} // not works
    retv += (uintptr_t)(new QString(""));
    QString::fromUtf8(0, 0);
    {delete ((QString*)0);}

    nilcxobj(QByteArray)->length();
    if (nilcxobj(QVariant)->isValid()) {}
    new QAnyStringView("",0);
    {auto x = sizeof(QAnyStringView);}
    {retv+=nilcxobj(QAnyStringView)->size();}
    {retv+=(uintptr_t)nilcxobj(QAnyStringView)->data();}

    new QUrl(QString(dummyqs)); delete nilcxobj(QUrl);
    nilcxobj(QUrl)->setUrl(dummyqs);

    QCoreApplication::instance();

    /////
    delete (new QColor("")); new QColor(dummyqs); // 弱符号
    new QColor(QStringView(dummyqs));
    QColor::colorNames();

    return (void*)retv;
}

void* uninline_qtgui_holder() {
    nilcxobj(QGuiApplication)->applicationState();
    return 0;
}

extern "C"
const char* QCompileVersion() { return QT_VERSION_STR; }

// __attribute__((noinline))
extern "C"
void QVariantDtor(void*p) {
    if (p!= nullptr) delete (QVariant*)p;
}

extern "C"
void* QVariantNewInt(int v) {
    auto rv = new QVariant(v);
    return rv;
}
extern "C"
int QVariantToint(QVariant*p) {
    bool ok = false;
    return p->toInt(&ok);
}
extern "C"
void* QVariantNewInt64(qint64 v) {
    auto rv = new QVariant(v);
    return rv;
}
extern "C"
qint64 QVariantToint64(QVariant*p) {
    bool ok = false;
    return p->toLongLong(&ok);
}
extern "C"
void* QVariantNewStr(char*str) {
    auto rv = new QVariant(QString(str));
    return rv;
}
// 有时go获取不到值，可能是dtor太快了
// 分配新内存解决，但是用的go's mallogc
extern "C"
char* QVariantTostr(QVariant*p) {
    // qDebug()<<__FUNCTION__<<__LINE__<<p->toString();
    auto v = p->toString();
    auto rv = (char*)cgoppMallocgcfn(v.length()+1);
    strcpy(rv, qUtf8Printable(v));
    return rv;
}

extern "C"
QVariant* QVariantNewPtr(void*ptr) {
    return new QVariant(quint64(ptr));
}
extern "C"
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

extern "C"
void* QVariantNewBool(bool v) {
    auto rv = new QVariant(v);
    // qDebug()<<__FUNCTION__<<__LINE__<<v<<*rv<<sizeof(bool);
    return rv;
}
extern "C"
bool QVariantTobool(QVariant*p) {
    return p->toBool();
}
extern "C"
void* QVariantNewDouble(double v) {
    auto rv = new QVariant(v);
    return rv;
}
extern "C"
int QVariantToDouble(QVariant*p, double* v) {
    auto rv = p->toDouble();
    *v = rv;
    return 1;
}

// 直接使用 extern "C"，不在头文件中写了
extern "C" void QByteArrayDtor(void*px) { delete (QByteArray*)px; }
extern "C" void _ZN10QByteArrayD2Ev_weakwrap(void*px) { delete (QByteArray*)px; }

extern "C" void QStringDtor(void*px) { delete (QString*)px; }
extern "C" void _ZN7QStringD2Ev_weakwrap(void*px)  { delete (QString*)px; }
extern "C" void _ZN7QStringC1EPKc_weakwrap(void*o, const char*p) {
    new(o) QString(p);
}
extern "C"
void* QStringNew(const char*p) {
    auto rv = new QString(p);
    return (void*)rv;
}
extern "C"
const char* QStringToutf8(QString* sp) {
    auto rv = (char*)cgoppMallocgcfn(sp->length()*2+1);
    strcpy(rv, qUtf8Printable((*sp)));
    return rv;
}

// new QStringList(); new QObjectList();

///////////
// for QtObject::createQmlObject, or other also ok
extern "C"
void QObjectDtor(QObject* o) { delete o; }

extern "C"
void QMetaObjectInvokeMethod1(void* fnptrx, void* n) {
    QObject* o = qApp;
    void (*fnptr)(void*) = (void(*)(void*))fnptrx;
    QMetaObject::invokeMethod(o, [fnptr,n]{ fnptr(n); }, Qt::QueuedConnection);
}

// 可以直接调用Qt slot，以及qml的函数
extern "C"
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

extern "C"
QObject* QObjectFindChild1(QObject*obj, char*str) {
    auto rv = obj->findChild<QObject*>(str);
    return rv;
}
extern "C"
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
extern "C"
const char* QObjectObjectName(QObject*obj) {
    auto rvx = obj->objectName();
    auto rv = (char*)cgoppMallocgcfn(rvx.length()+1);
    strcpy(rv, qUtf8Printable(rv));
    return rv;
}

/////
#include <dlfcn.h>
extern "C"
void* cgoir_dlsym0(const char* name) {
    return dlsym(RTLD_DEFAULT, name);
}


///////
extern "C"
void _ZN6QColorC1ERK7QString_weakwrap(void*o, void*namex) {
    auto name = *(QString*)namex;
    new(o) QColor(name);
}
extern "C"
void _ZN6QColorC1EPKc_weakwrap(void*o, char*name) {
    new(o) QColor(name);
}
extern "C"
void _ZN6QColorC1E11QStringView_weakwrap(void*o, QStringView name) {
    new(o) QColor(name);
}

extern "C"
void _ZN6QColorD2Ev(void*o) {
    delete (QColor*)(o);
}