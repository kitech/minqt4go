
#include <cassert>
#include <stdint.h>

#if QT_VERSION < 0x040000

#ifndef Q_DECL_EXPORT
#define Q_DECL_EXPORT Q_EXPORT
#endif

#define qint64 int64_t
//#define quint64 uint64_t
//typedef uint64_t quint64; 

#endif
