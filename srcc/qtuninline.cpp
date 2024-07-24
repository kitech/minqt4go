#include <cxxabi.h>

#include <QtCore>
#include <QJSEngine>
#include <QtQml>
#include <QtQuick>
#include <QtQuickControls2>
#include <QtQuickTemplates2>

// private part
#include <QtQuickTemplates2/private/qquickstackview_p.h>
#include <QtQuickTemplates2/private/qquicktoolbutton_p.h>
#include <QtQuickTemplates2/private/qquickbutton_p.h>
// #include <QtQuickTemplates2/private/qquicksplitter_p.h>
#include <QtQml/private/qqmlbuiltinfunctions_p.h>
#include <QtQuickTemplates2/private/qquickapplicationwindow_p.h>
#include <QtQuick/private/qquicklistview_p.h>
#include <QtQuick/private/qquickflickable_p.h>
#include <QtQuick/private/qquickloader_p.h>
#include <QtQuickTemplates2/private/qquickmenubar_p.h>
#include <QtQuickTemplates2/private/qquickmenu_p.h>
#include <QtQuickTemplates2/private/qquickpopup_p.h>
// #include <QtQuickTemplates2/private/qquickpopupitem_p.h>
#include <QtQuickTemplates2/private/qquickpane_p.h>
#include <QtQuickTemplates2/private/qquickswipe_p.h>
#include <QtQuickTemplates2/private/qquicksplitview_p.h>
#include <QtQuickTemplates2/private/qquickscrollview_p.h>
#include <QtQuickTemplates2/private/qquickscrollbar_p.h>
#include <QtQuickTemplates2/private/qquickaction_p.h>
#include <QtQuickTemplates2/private/qquicklabel_p.h>
#include <QtQuickTemplates2/private/qquicktextarea_p.h>
// #include <QtQuickTemplates2/private/qquicktextedit_p.h>
#include <QtQuickTemplates2/private/qquickcombobox_p.h>
#include <QtQuickTemplates2/private/qquickspinbox_p.h>
#include <QtQuickTemplates2/private/qquickcheckbox_p.h>
#include <QtQuickTemplates2/private/qquickslider_p.h>
#include <QtQuickTemplates2/private/qquicktoolbutton_p.h>
#include <QtQuickTemplates2/private/qquicktooltip_p.h>

#include "qtuninline.h"

#include <dlfcn.h>
auto cgoppMallocgcfn = (void* (*)(int))dlsym(RTLD_DEFAULT, "cgoppMallocgc");
// extern "C" void* cgoppMallocgc(int);

char* cxxabi__cxa_demangle(char*a0, char*a1, size_t *length, int *status) {
    return abi::__cxa_demangle(a0, a1, length, status);
}

#define DBGLOG qDebug()<<__FUNCTION__<<__LINE__

static QString dummyqs("dummy&ref");
void* uninlineholder() {
    uintptr_t ptr = 0;
#define nilobj(x) ((x*)0)

    if (nilobj(QObject)->metaObject()!=nullptr) { }
    nilobj(QObject)->parent();
    if (nilobj(QMetaObject)->className()!=nullptr) {}
    if (nilobj(QVariant)->isValid()) {}
    nilobj(QJSEngine)->collectGarbage();
    nilobj(QJSEngine)->objectOwnership(0);
    nilobj(QJSEngine)->setObjectOwnership(0, QJSEngine::CppOwnership);

    (new QQuickMenuBar());
    nilobj(QQuickMenuBar)->addMenu(0);
    nilobj(QQuickMenuBar)->setHeight(0);
    nilobj(QQuickMenuBar)->setWidth(0);

    (new QQuickMenu());
    nilobj(QQuickMenu)->addItem(0);
    nilobj(QQuickMenu)->addAction(0);
    nilobj(QQuickMenu)->setTitle(dummyqs); // 这个很奇怪

    (new QQuickAction());
    nilobj(QQuickAction)->setText(0);

    (new QQuickLabel());
    nilobj(QQuickLabel)->setText(0);
    nilobj(QQuickLabel)->setWidth(0);
    nilobj(QQuickLabel)->setHeight(0);
    nilobj(QQuickButton)->setText(0);

    (new QQuickButton());
    (new QQuickToolButton());
    nilobj(QQuickToolButton)->setText(0);
    nilobj(QQuickPopup)->setZ(123);
    nilobj(QQuickPopup)->z();
    (new QQuickToolTip());
    (delete nilobj(QQuickToolTip));
    nilobj(QQuickToolTip)->setText(0);
    nilobj(QQuickToolTip)->text();
    nilobj(QQuickToolTip)->setDelay(0);
    nilobj(QQuickToolTip)->setTimeout(0);
    nilobj(QQuickToolTip)->setVisible(true);
    nilobj(QQuickItem)->setVisible(true);

    (new QQuickSplitView());
    (new QQuickTextArea());
    // (new QQuickImage());
    // (new QQuickAnimatedImage());

    // 如何取构造函数的地址！！
    // 如何取有重载的方法的地址！！
    // works
    {QString ( QString::*x )(int, int, int, QChar) const = &QString::arg;}
    // {auto x  = &QString::QString;} // not works
    ptr += (uintptr_t)(new QString("abc"));
    QString::fromUtf8(0, 0);
    {delete ((QString*)0);}

    nilobj(QByteArray)->length();

    return (void*)ptr;
}

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
// 有时go获取不到值，可能是dtor太快了
// 分配新内存解决，但是用的go's mallogc
char* QVariantTostr(QVariant*p) {
    // qDebug()<<__FUNCTION__<<__LINE__<<p->toString();
    auto v = p->toString();
    auto rv = (char*)cgoppMallocgcfn(v.length()+1);
    strcpy(rv, qUtf8Printable(v));
    return rv;
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

void* QVariantNewDouble(double v) {
    auto rv = new QVariant(v);
    return rv;
}
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
void* QStringNew(const char*p) {
    auto rv = new QString(p);
    return (void*)rv;
}
const char* QStringToutf8(QString* sp) {
    auto rv = (char*)cgoppMallocgcfn(sp->length()*2+1);
    strcpy(rv, qUtf8Printable((*sp)));
    return rv;
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
const char* QObjectObjectName(QObject*obj) {
    auto rvx = obj->objectName();
    auto rv = (char*)cgoppMallocgcfn(rvx.length()+1);
    strcpy(rv, qUtf8Printable(rv));
    return rv;
}

// 适用于 qml attached property
QQmlProperty* QQmlPropertyNew1(QObject*obj, char*name, void*qe) {
    // qDebug()<<__FUNCTION__<<__LINE__<<obj<<name<<qe;
    if (qe == nullptr) {
        auto ctx = QQmlEngine::contextForObject(obj);
        auto rv = new QQmlProperty(obj, QString(name), ctx);
        return rv;
    }else{
        auto ctx = QQmlEngine::contextForObject(obj);
        // auto rv = new QQmlProperty(obj, QString(name), (QQmlEngine*)qe);
        auto rv = new QQmlProperty(obj, QString(name), ctx);
        return rv;
    }
}
void QQmlPropertyDtor(QQmlProperty*obj) { delete obj; }
QVariant* QQmlPropertyRead(QQmlProperty*obj) {
    auto rv = obj->read();
    // qDebug()<<__FUNCTION__<<__LINE__<<rv<<obj->name()<<obj->isValid();
    return new QVariant(rv);
}
int QQmlPropertyWrite(QQmlProperty*obj, QVariant*val) {
    auto rv = obj->write(*val);
    return rv;
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

void dummyyy() {
    QQuickStackView *stkwin; // not work
}

QQuickItem* QQuickStackView_get(QQuickStackView*me, int idx) {
    QQuickItem* rv = me->get(idx);
    return rv;
}
#if QT_VERSION >= QT_VERSION_CHECK(6, 7, 0)
QQuickItem* QQuickStackView_replaceCurrentItem(QQuickStackView* me, QQuickItem* item) {
    auto olditem = me->currentItem();
    auto rv = me->replaceCurrentItem(item);
    return olditem;
    // DBGLOG<<(rv==olditem)<<(void*)olditem<<(void*)rv;
    // DBGLOG<<(rv==item)<<(void*)item<<(void*)rv;
    // return rv;
}
#endif


/////
#include <dlfcn.h>
void* cgoir_dlsym0(const char* name) {
    return dlsym(RTLD_DEFAULT, name);
}
