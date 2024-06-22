#include <QtCore>

auto oldqtmsgoutfn = qInstallMessageHandler(nullptr);
extern "C" void qtMessageOutputGoimpl(int, const char*, const char*, const char*);
void qtMessageOutput(QtMsgType mtype, const QMessageLogContext& ctx, const QString& msg) {
    // auto oldfn = (DeclType(qInstallMessageHandler(nullptr)))oldqtmsgoutfn;
    // oldqtmsgoutfn(mtype, ctx, msg);

    int itype = int(mtype);
    const char* file = ctx.file;
    const char* funcname = ctx.function;
    QByteArray cmsg = msg.toUtf8();
    qtMessageOutputGoimpl(itype, file, funcname, cmsg.data());
}

void initQtmsgout() {
    oldqtmsgoutfn = qInstallMessageHandler(qtMessageOutput);
}
// not work
// oldqtmsgoutfn = qInstallMessageHandler(qtMessageOutput);
