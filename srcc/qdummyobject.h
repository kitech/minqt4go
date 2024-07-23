#ifndef _QDUMMYOBJECT_H_
#define _QDUMMYOBJECT_H_

#include <QObject>

class QDynSlotObjecT : public QObject {
    Q_OBJECT;

public:
    QDynSlotObjecT(QObject*parent = nullptr);
    ~QDynSlotObjecT() {}

public slots:
    void dumnyslot();
};

#endif // _QDUMMYOBJECT_H_