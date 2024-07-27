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
#include <QtQuickControls2Material/private/qquickmaterialstyle_p.h>
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


#define DBGLOG qDebug()<<__FUNCTION__<<__LINE__
#define nilcxobj(x) ((x*)0)

static QString dummyqs("dummy&ref");
void* uninline_qtquick_holder() {
    uintptr_t ptr = 0;

    nilcxobj(QQmlApplicationEngine)->load(dummyqs);
    nilcxobj(QQuickApplicationWindow)->contentItem();
    nilcxobj(QQmlEngine)->contextForObject(0);

    nilcxobj(QJSEngine)->collectGarbage();
    nilcxobj(QJSEngine)->objectOwnership(0);
    nilcxobj(QJSEngine)->setObjectOwnership(0, QJSEngine::CppOwnership);

    (new QQuickMenuBar());
    nilcxobj(QQuickMenuBar)->addMenu(0);
    nilcxobj(QQuickMenuBar)->setHeight(0);
    nilcxobj(QQuickMenuBar)->setWidth(0);

    (new QQuickMenu());
    nilcxobj(QQuickMenu)->addItem(0);
    nilcxobj(QQuickMenu)->addAction(0);
    nilcxobj(QQuickMenu)->setTitle(dummyqs); // 这个很奇怪

    (new QQuickAction());
    nilcxobj(QQuickAction)->setText(0);

    (new QQuickLabel());
    nilcxobj(QQuickLabel)->setText(0);
    nilcxobj(QQuickLabel)->setWidth(0);
    nilcxobj(QQuickLabel)->setHeight(0);
    nilcxobj(QQuickButton)->setText(0);

    (new QQuickButton());
    (new QQuickToolButton());
    nilcxobj(QQuickToolButton)->setText(0);
    nilcxobj(QQuickPopup)->setZ(123);
    nilcxobj(QQuickPopup)->z();
    (new QQuickToolTip());
    (delete nilcxobj(QQuickToolTip));
    nilcxobj(QQuickToolTip)->setText(0);
    nilcxobj(QQuickToolTip)->text();
    nilcxobj(QQuickToolTip)->setDelay(0);
    nilcxobj(QQuickToolTip)->setTimeout(0);
    nilcxobj(QQuickToolTip)->setVisible(true);
    nilcxobj(QQuickItem)->setVisible(true);

    (new QQuickSplitView());
    (new QQuickTextArea());
    // (new QQuickImage());
    // (new QQuickAnimatedImage());



    return (void*)ptr;
}



// 适用于 qml attached property
extern "C"
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
extern "C"
void QQmlPropertyDtor(QQmlProperty*obj) { delete obj; }
extern "C"
QVariant* QQmlPropertyRead(QQmlProperty*obj) {
    auto rv = obj->read();
    // qDebug()<<__FUNCTION__<<__LINE__<<rv<<obj->name()<<obj->isValid();
    return new QVariant(rv);
}
extern "C"
int QQmlPropertyWrite(QQmlProperty*obj, QVariant*val) {
    auto rv = obj->write(*val);
    return rv;
}

// qml
extern "C"
QQmlApplicationEngine* QQmlApplicationEngineNew() {
    auto e = new QQmlApplicationEngine();
    return e;
}
extern "C"
void QQmlApplicationEngineLoad1(QQmlApplicationEngine*e, char*str) {
    auto url = QUrl(str);
    e->load(url);
}
extern "C"
QObject* QQmlApplicationEngineRootObject1(QQmlApplicationEngine*e) {
    auto objs = e->rootObjects();
    return objs.value(0);
}
extern "C"
QQmlComponent* QQmlComponentNew1(void*engine, QObject* parent) {
    auto rv = new QQmlComponent((QQmlEngine*)engine, parent);
    return rv;
}
extern "C"
QObject* QQmlComponentCreate(QQmlComponent*o, void*ctx) {
    auto rv = o->create((QQmlContext*)ctx);
    return rv;
}
extern "C"
void QQmlComponentSetData(QQmlComponent*o, char*data) {
    o->setData(QByteArray(data), QUrl());
}

#include <QtQml/private/qqmlbuiltinfunctions_p.h>
extern "C"
QtObject* QtObjectCreate(QQmlEngine*e) { return QtObject::create(e,e); }

extern "C"
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

extern "C"
QQuickItem* QQuickStackView_get(QQuickStackView*me, int idx) {
    QQuickItem* rv = me->get(idx);
    return rv;
}
#if QT_VERSION >= QT_VERSION_CHECK(6, 7, 0)
extern "C"
QQuickItem* QQuickStackView_replaceCurrentItem(QQuickStackView* me, QQuickItem* item) {
    auto olditem = me->currentItem();
    auto rv = me->replaceCurrentItem(item);
    return olditem;
    // DBGLOG<<(rv==olditem)<<(void*)olditem<<(void*)rv;
    // DBGLOG<<(rv==item)<<(void*)item<<(void*)rv;
    // return rv;
}
#endif

