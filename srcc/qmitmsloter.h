#ifndef _QMITMSLOTER_H_
#define _QMITMSLOTER_H_

#include <qobject.h>
#include <qmetaobject.h>
#include <qvariant.h>

// 该版本不用再关心moc格式版本问题，非常容易维护
// 用于替换QDynSlotObject
class QMitmSloter : public QObject {
	Q_OBJECT
public:
	QMitmSloter(void* d_);
	void *cbdata;
	
#if QT_VERSION < 0x040000
	void metacallir(int _id, QUObject* _o);
#elif QT_VERSION < 0x070000
	void metacallir(QObject *_o, QMetaObject::Call _c, int _id, void **_a);
#else
#warn "not suported QT_VERSION"
#endif
	
public slots:
	// It's Qt's big feature. Many args connect to none arg.
	// connect(o1, SIGNAL(anynumargsignal(...)), SLOT(dummy()));
	void dummy(/*...*/);
	void slotxx(char, long, float, double, int, short, bool, void*, QObject*, QString, QString&, QVariant, QVariant&){}
};

#endif // _QMITMSLOTER_H_
