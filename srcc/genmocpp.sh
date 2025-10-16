
# usage: exe <header-file.h>

#set -x

hdrfile=qmitmsloter.h

mocexes="$MOC moc moc-qt3 moc-qt4 moc-qt5 moc-qt6 /usr/lib/qt6/moc"
for mocexe in $mocexes; do
    exepath=$(which $mocexe >/dev/null 2>&1)
    rc=$?
    if [ $rc -eq 1 ]; then
        continue
    fi
    mocver=$($mocexe -v 2>&1)
    echo "$mocexe, ${mocver}..."

    $mocexe $hdrfile > tmpmoc.cxx
    # format m.m.p
    verstr=$(grep "#error \"This file was generated using the moc from" tmpmoc.cxx|awk '{print $10}')
    echo "verstr=${verstr}"

    hasswitch=$(grep "switch ( _id" tmpmoc.cxx)
    echo "*M*: $hasswitch"
    if [ x"$hasswitch" != x"" ]; then # qt3
        sed -i 's/switch ( _id/metacallir(_id,_o);\n\tswitch ( _id/' tmpmoc.cxx
    else # qt4+
        # switch (_id) {
        sed -i 's/switch (_id)/metacallir(_o,_c,_id,_a);\n\tswitch (_id)/' tmpmoc.cxx
    fi

    grep "metacallir" tmpmoc.cxx

    dstfile="qmitmsloter_moc_v${verstr}cxx"
    mv -v tmpmoc.cxx $dstfile

    echo
done

