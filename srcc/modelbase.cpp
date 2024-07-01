
#include <QtQml>

#include "modelbase.h"

#include <dlfcn.h>

auto ListModelBaseQmltype = qmlRegisterType<ListModelBase>("ListModelBase", 1, 0, "ListModelBase");

// extern "C" char* goimplListModelBaseRoleName(qint64, int);

QHash<int, QByteArray> ListModelBase::roleNames() const  {
    QHash<int, QByteArray> names;
    names[Foo0Role] = "foo0";
    names[NameRole] = "name";
    names[ValueRole] = "value";
    names[DeletedRole] = "deleted";

    typedef char* (*fnty)(qint64, int);
    auto symx = dlsym(RTLD_DEFAULT, "goimplListModelBaseRoleName");
    fnty goimplListModelBaseRoleName = (fnty)symx;

    for (int role = Qt::UserRole ; role < Qt::UserRole + 555; role++) {
        char* name = goimplListModelBaseRoleName(goimpl, role);
        if (name == nullptr) { break; }
        names[role] = QString(name).toUtf8();
        // free(name); // cgopp.CStrSmt
    }

    return names;
}

// extern "C" qint64 goimplListModelBaseNew(void*);
ListModelBase::ListModelBase(QObject*parent) : QAbstractListModel(parent) {
    typedef qint64 (*fnty)(void*);
    auto symx = dlsym(RTLD_DEFAULT, "goimplListModelBaseNew");
    fnty goimplListModelBaseNew = (fnty)symx;

    qint64 seqno = goimplListModelBaseNew(this);
    goimpl = seqno;
    QObject::connect(this, &QObject::objectNameChanged, this, &ListModelBase::mySetObjectName);
}
// extern "C" void goimplListModelBaseDtor(qint64);
ListModelBase::~ListModelBase() {
    typedef void (*fnty)(qint64);
    auto symx = dlsym(RTLD_DEFAULT, "goimplListModelBaseDtor");
    fnty goimplListModelBaseDtor = (fnty)symx;

    goimplListModelBaseDtor(goimpl);
    if (goimpl == 12345) {
        emit beginInsertRows(QModelIndex(), 0, 0);
        emit beginRemoveRows(QModelIndex(), 0, 0);
        emit endInsertRows();
        emit endRemoveRows();
        // emit countChanged(0);
    }
}

// extern "C" int goimplListModelBaseGetsetRolecnt(qint64, int, int);
// int ListModelBase::rolecnt() {
//     int rv = goimplListModelBaseGetsetRolecnt(goimpl, 0, 0);
//     return rv;
// }
// void ListModelBase::setRolecnt(int c) {
//     goimplListModelBaseGetsetRolecnt(goimpl, c, 1);
// }

// extern "C" char* goimplListModelBaseGetsetClazz(qint64, char*, int);
// QString ListModelBase::clazz() {
//     char* rv = goimplListModelBaseGetsetClazz(goimpl, 0, 0);
//     return rv;
// }
// void ListModelBase::setClazz(QString c) {
//     goimplListModelBaseGetsetClazz(goimpl, c.toUtf8().data(), 1);
// }

void ListModelBase::mySetObjectName(const QString& c){
    // qDebug()<<__FUNCTION__<<__LINE__<<c;
    // QAbstractListModel::setObjectName(c);

    typedef char* (*fnty)(qint64, char*, int);
    auto symx = dlsym(RTLD_DEFAULT, "goimplListModelBaseGetsetClazz");
    fnty goimplListModelBaseGetsetClazz = (fnty)symx;

    goimplListModelBaseGetsetClazz(goimpl, c.toUtf8().data(), 1);
}

// extern "C" void* goimplListModelBaseData(qint64, int, int);
QVariant ListModelBase::data(const QModelIndex &index, int role) const {
    typedef void* (*fnty)(qint64, int, int);
    auto symx = dlsym(RTLD_DEFAULT, "goimplListModelBaseData");
    fnty goimplListModelBaseData = (fnty)symx;

    void* tv = goimplListModelBaseData(goimpl, index.row(), role);
    QVariant* tv2 = (QVariant*)tv;
    QVariant rv = QVariant(*tv2);
    // delete tv2;
    // qDebug()<<__FUNCTION__<<role <<rv;
    return rv;
}

// extern "C" int goimplListModelBaseRowCount(qint64);
int ListModelBase:: rowCount(const QModelIndex &parent) const {
    typedef int (*fnty)(qint64);
    auto symx = dlsym(RTLD_DEFAULT, "goimplListModelBaseRowCount");
    fnty goimplListModelBaseRowCount = (fnty)symx;

    int rv = goimplListModelBaseRowCount(goimpl);
    // qDebug()<<__FUNCTION__<<rv;
    return rv;
}

void ListModelBase::emitBeginChangeRows(int first, int last, int remove) {    
    if (remove) {
        emit beginRemoveRows(QModelIndex(), first, last);    
    }else{
        emit beginInsertRows(QModelIndex(), first, last);
    }
}
