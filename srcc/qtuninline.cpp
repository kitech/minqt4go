#include <cxxabi.h>

#include <QtCore>
#include <QtGui>

// #include "qtuninline.h"

#include "qt3macros.h"

#include <dlfcn.h>
auto cgoppMallocgcfn = (void* (*)(int))dlsym(RTLD_DEFAULT, "cgoppMallocgc");
// extern "C" void* cgoppMallocgc(int);

extern "C"
char* cxxabi__cxa_demangle(char*a0, char*a1, size_t *length, int *status) {
    return abi::__cxa_demangle(a0, a1, length, status);
}

#if QT_VERSION > 0x030308
#define DBGLOG qDebug()<<__FUNCTION__<<__LINE__
#else
#include <iostream>
#define DBGLOG std::cout << __FUNCTION__<<__LINE__
#endif
#define nilcxobj(x) ((x*)0)

static QString dummyqs("dummy&ref");
void* uninline_qtcore_holder() {
    uintptr_t retv = 0;
    // retv += (uintptr_t)(qInstallMessageHandler(nullptr));

    //nilcxobj(QMetaObject)->superClass();
    //if (nilcxobj(QObject)->metaObject()!=nullptr) { }
    //nilcxobj(QObject)->setObjectName(QAnyStringView(""));
    //nilcxobj(QObject)->parent();
    //if (nilcxobj(QMetaObject)->className()!=nullptr) {}
    //{auto rv = nilcxobj(QObject)->findChild<QObject*>(""); retv+=(uintptr_t)rv;}

    // 如何取构造函数的地址！！
    // 如何取有重载的方法的地址！！
    // works
    //{QString ( QString::*x )(int, int, int, QChar) const = &QString::arg;}
    // {auto x  = &QString::QString;} // not works
    retv += (uintptr_t)(new QString(""));
    QString::fromUtf8(0, 0);
    {delete ((QString*)0);}

    //retv += nilcxobj(QByteArray)->length();
    if (nilcxobj(QVariant)->isValid()) {}
    //retv+=nilcxobj(QVariant)->typeId();
    //new QAnyStringView("",0);
    //{auto x = sizeof(QAnyStringView);}
    //{retv+=nilcxobj(QAnyStringView)->size();}
    //{retv+=(uintptr_t)nilcxobj(QAnyStringView)->data();}
    //retv += nilcxobj(QByteArrayView)->length();

    new QUrl(QString(dummyqs)); delete nilcxobj(QUrl);
    //nilcxobj(QUrl)->setUrl(dummyqs);
    //nilcxobj(QUrl)->url();

    // 获取类型大小sizeof
    //QMetaType::fromName("QObject").sizeOf();
    //nilcxobj(QMetaType)->metaObject();

    //QCoreApplication::instance();
	//QApplication::instance();
	
	nilcxobj(QButton)->autoRepeat();
	nilcxobj(QMetaObject)->className();
	
    /////
    delete (new QColor("")); new QColor(dummyqs); // 弱符号
    //new QColor(QStringView(dummyqs));
    QColor::colorNames();
    delete nilcxobj(QStringList); delete nilcxobj(QObjectList);

    return (void*)retv;
}

void* uninline_qtgui_holder() {
    //nilcxobj(QGuiApplication)->applicationState();
    return 0;
}

extern "C" const char* QCompileVersion() { return QT_VERSION_STR; }

extern "C" int qtapp_argc() {
    //auto args = QCoreApplication::arguments();
    //return args.size();
    return 0;
}
extern "C" char* qtapp_argat(int idx) {
    //auto args = QCoreApplication::arguments();
    //auto arg  = args.at(idx);
    //return strdup(arg.toUtf8().data());
    return NULL;
}

// __attribute__((noinline))
extern "C" void QVariantDtor(void*p) {
    if (p!= nullptr) delete (QVariant*)p;
}

extern "C" void* QVariantNewInt(int v) {
    auto rv = new QVariant(v);
    return rv;
}
extern "C" int QVariantToint(QVariant*p) {
    bool ok = false;
    return p->toInt(&ok);
}
extern "C" void* QVariantNewInt64(qint64 v) {
    //auto rv = new QVariant(v);
    //return rv;
    return NULL;
}
extern "C" qint64 QVariantToint64(QVariant*p) {
    bool ok = false;
    return p->toLongLong(&ok);
}
extern "C" void* QVariantNewStr(char*str) {
    auto rv = new QVariant(QString(str));
    return rv;
}
// 有时go获取不到值，可能是dtor太快了
// 分配新内存解决，但是用的go's mallogc
extern "C" char* QVariantTostr(QVariant*p) {
    // qDebug()<<__FUNCTION__<<__LINE__<<p->toString();
    auto v = p->toString();
    auto rv = (char*)cgoppMallocgcfn(v.length()+1);
    //strcpy(rv, qUtf8Printable(v));
    return rv;
}

extern "C" QVariant* QVariantNewPtr(void*ptr) {
    // return new QVariant((quint64)(ptr));
    return NULL;
}
// todo only support QObject*
extern "C" void* QVariantToptr(QVariant*p) {
    bool ok = false;
    // auto rv = p->toULongLong(&ok);
    // auto rv = p->value<QQuickItem*>(); // 这种方式对的
    // auto x = p->convert(QMetaType::type("void*")); // not work
    // auto rv = p->value<void*>(); // not work
    // auto rv = p->value<QObject*>(); // works
    // qDebug()<<__FUNCTION__<<(*p)<<(void*)p<<(*p).data()<<rv;
    // return (void*)rv;
    return NULL;
}

extern "C" void* QVariantNewBool(bool v) {
    auto rv = new QVariant(v);
    // qDebug()<<__FUNCTION__<<__LINE__<<v<<*rv<<sizeof(bool);
    return rv;
}
extern "C" bool QVariantTobool(QVariant*p) {
    return p->toBool();
}
extern "C" void* QVariantNewDouble(double v) {
    auto rv = new QVariant(v);
    return rv;
}
extern "C" int QVariantToDouble(QVariant*p, double* v) {
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
void* QStringNew2(const char*p, int len) {
    //QByteArray rv(p, len);
    //auto rv2 = new QString(rv.data());
    //DBGLOG << *rv2 << rv2->length();
    //return (void*)rv2;
    return NULL;
}
extern "C"
const char* QStringToutf8(QString* sp) {
    auto rv = (char*)cgoppMallocgcfn(sp->length()*2+1);
    //strcpy(rv, qUtf8Printable((*sp)));
    return rv;
}

// need client copy
extern "C" const char *QStringToutf8p2(QString *sp) {
    // auto rv = (char *)malloc(sp->length() * 2 + 1);
    //strcpy(rv, qUtf8Printable((*sp)));
    // return qUtf8Printable((*sp));
    return NULL;
}

// new QStringList(); new QObjectList();
extern "C"
void _ZN5QListI7QStringED2Ev_weakwrap(void*px) { delete (QList<QString>*)(px); }
extern "C"
void _ZN5QListIP7QObjectED2Ev_weakwrap(void*px) { delete (QList<QObject*>*)(px); }


///////////
//  QThread::staticMetaObject::metaType().sizeOf()
//  QMetaType::fromName(QTime).sizeOf()
extern "C"
int GetClassSizeByName(const char* clsname) {
    int clzsz = 0;
    char buf[99] = {0};
    sprintf(buf, "_ZN%d%s16staticMetaObjectE", int(strlen(clsname)), clsname);
    void* stmo = dlsym(RTLD_DEFAULT, buf);
    // DBGLOG<<clsname<<buf<<stmo<<(QThread::staticMetaObject.metaType().sizeOf());
    if (stmo != nullptr) {
        // clzsz = ((QMetaObject*)stmo)->metaType().sizeOf();
    } else {
        // QMetaType stmo2 = QMetaType::fromName(QByteArrayView(clsname));
        // clzsz = stmo2.sizeOf();
    }

    return clzsz;
}

// for QtObject::createQmlObject, or other also ok
extern "C"
void QObjectDtor(QObject* o) { delete o; }

extern "C"
void QMetaObjectInvokeMethod1(void* fnptrx, void* n) {
    QObject* o = qApp;
    void (*fnptr)(void*) = (void(*)(void*))fnptrx;
    // QMetaObject::invokeMethod(o, [fnptr,n]{ fnptr(n); }, Qt::QueuedConnection);
}

// 可以直接调用Qt slot，以及qml的函数
extern "C"
int QMetaObjectInvokeMethod2(QObject* obj, char* member, void* retp, void*val0,void*val1,void*val2) {
    //QGenericReturnArgument qret;
    //QGenericArgument a0;
    //QGenericArgument a1;
    //QGenericArgument a2;

    //if (val0 != nullptr) a0 = *((QGenericArgument*)val0);
    //if (val1 != nullptr) a1 = *((QGenericArgument*)val1);
    //if (val2 != nullptr) a2 = *((QGenericArgument*)val2);
    //qret = *((QGenericReturnArgument*)retp);

    // qDebug()<<__FUNCTION__<<__LINE__<<obj<<member<<val0;

    //int rv = QMetaObject::invokeMethod(obj, member, Qt::QueuedConnection, qret, a0, a1, a2);
    // DBGLOG<<rv;
    // *(QGenericReturnArgument*)(retp) = qret;
    //return rv;
    return -1;
}

extern "C"
QObject* QObjectFindChild1(QObject*obj, char*str) {
    // auto rv = obj->findChild<QObject*>(str);
    // return rv;
    return (QObject*)(void*)0x9;
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
    /*auto rvx = obj->objectName();
    auto rv = (char*)cgoppMallocgcfn(rvx.length()+1);
    strcpy(rv, qUtf8Printable(rvx));
    return rv;
    */
    return (const char*)(void*)0x9;
}
extern "C"
char *QObjectObjectName2(QObject *obj, int len, char*rv) {
  //auto rvx = obj->objectName();
  //assert(len > rvx.length());
  //strcpy(rv, qUtf8Printable(rvx));
  //return rv;
  return (char*)(void*)0x9;
}

extern "C"
const char* QMetaObjectNormalizedSignature(const char*signt) {
    //auto rvx = QMetaObject::normalizedSignature(signt);
    //auto rv = (char*)cgoppMallocgcfn(rvx.length()+1);
    //strcpy(rv, qUtf8Printable(rvx));
    //return rv;
    return (const char*)(void*)0x9;
}

extern "C"
const QMetaObject* QObjectVtableMetaObject(QObject*obj) {
    return obj->metaObject();
}

extern "C"
int QObjectSignalArgTypes(QObject*obj, const char*signt, void*fnsym) {
    auto mto = obj->metaObject();
    /*auto mthcnt = mto->methodCount();
    DBGLOG << mto->className() ;
    DBGLOG << signt;
    
    auto signt2 = QMetaObject::normalizedSignature(signt);
    DBGLOG << signt2;
    for (int i = 0; i < mthcnt; i++) {
        auto mth = mto->method(i);
        DBGLOG << signt << i << mth.methodSignature();
    }
    
    auto sigidx = mto->indexOfSignal(signt);
    DBGLOG << sigidx;
    
    int (QMetaObject::*mthadr)(const char *) const = &QMetaObject::indexOfSignal;
    DBGLOG << mthadr;

    auto fnsym2 = dlsym(RTLD_DEFAULT, "_ZNK11QMetaObject13indexOfSignalEPKc");
    DBGLOG << fnsym << fnsym2 << *(void**)(&mthadr); // alleq
    // fnsym = *((void**)fnsym);
    DBGLOG << fnsym;
    DBGLOG << ((mto->*mthadr)(signt)); // works // 'is not a function'
    int (*fno)(void*, const char*) = (int (*)(void*, const char*)) fnsym;
    DBGLOG << fno << (void*)mto << (void*)obj->metaObject() <<(void*)obj;
    sigidx = fno((void*)mto, signt);
    DBGLOG << sigidx;

    return sigidx; */
    return -1;
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
#if QT_VERSION > 0x030308
extern "C"
void _ZN6QColorC1E11QStringView_weakwrap(void*o, QStringView name) {
    new(o) QColor(name);
}
#endif

extern "C"
void _ZN6QColorD2Ev(void*o) {
    delete (QColor*)(o);
}
