
#include <QtWidgets>

void* qtwidget_uninline() {
    ((QWidget*)0)->size();
    ((QWidget*)0)->width();
    ((QWidget*)0)->height();
    ((QPushButton*)0)->show();
    return 0;
}