#ifndef _MODEL_BASE_H_
#define _MODEL_BASE_H_

#include <QAbstractItemModel>
#include <QtQml>

class ListModelBase : public QAbstractListModel {
    Q_OBJECT;

    // virutal column
    // Q_PROPERTY(int rolecnt READ rolecnt WRITE setRolecnt FINAL);
    // Q_PROPERTY(QString clazz READ clazz WRITE setClazz FINAL);
    Q_PROPERTY(quint64 goobj READ goobj CONSTANT FINAL);
    QML_ELEMENT;

public:
    // 一列中有多个字段的情况
    enum RoleNames {
        Foo0Role = Qt::UserRole + 0,
        NameRole = Qt::UserRole + 1,
        ValueRole = Qt::UserRole + 2,
        DeletedRole = Qt::UserRole + 3,
    };
    explicit ListModelBase(QObject* parent=nullptr);
    ~ListModelBase();

    int rolecnt();
    void setRolecnt(int c);
    // QString clazz();
    // void setClazz(QString c);
    quint64 goobj() { return goimpl; }

    virtual void mySetObjectName(const QString& c) ;
    
    virtual QVariant data(const QModelIndex &index, int role = Qt::DisplayRole) const ;
    virtual int rowCount(const QModelIndex &parent = QModelIndex()) const;

    void emitBeginChangeRows(int first, int last, int remove);

protected:
    // return the roles mapping to be used ny QML
    virtual QHash<int, QByteArray> roleNames() const ;

private:
    qint64 goimpl = -1;

};

/*
class ModelBase : public QAbstractItemModel {
    Q_OBJECT;

public:
    explicit ModelBase(QObject* parent = nullptr);
    ~ModelBase();

    int row() const;

    virtual QVariant data(const QModelIndex &index, int role = Qt::DisplayRole);
    virtual int rowCount(const QModelIndex &parent = QModelIndex()) const;
    virtual int columnCount(const QModelIndex &parent = QModelIndex()) const;
const;    
}

*/

#endif  // _MODEL_BASE_H_