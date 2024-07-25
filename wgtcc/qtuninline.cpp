
#include <QtWidgets>

#define DBGLOG qDebug()<<__FUNCTION__<<__LINE__
#define nilcxobj(x) ((x*)0)

void* uninline_qtwidgets_holder() {
    QApplication::aboutQt();
    nilcxobj(QWidget)->resize(0,0);

    ((QWidget*)0)->size();
    ((QWidget*)0)->width();
    ((QWidget*)0)->height();
    ((QPushButton*)0)->show();
    ((QLayout*)0)->addItem(0);
    new QSpacerItem(0,0,QSizePolicy::Minimum,QSizePolicy::Minimum);
    return (void*)(uintptr_t(0));
}

// QSpacerItem::QSpacerItem(int,int,QSizePolicy::Policy,QSizePolicy::Policy);
extern "C" 
void _ZN11QSpacerItemC2EiiN11QSizePolicy6PolicyES1__weakwrap(void*o, int w, int h, QSizePolicy::Policy p1, QSizePolicy::Policy p2) {
    // qDebug()<<__FUNCTION__<<o<<w<<h<<p1<<p2;
    auto x = new(o) QSpacerItem(w, h, p1, p2);
    assert(x == o);
    // qDebug()<<__FUNCTION__<<o<<x<<w<<h<<p1<<p2;
}
