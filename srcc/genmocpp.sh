
# usage: exe <header-file.h>

#set -x

hdrfile=qmitmsloter.h

mocexes="$MOC moc moc-qt3 moc-qt4 moc-qt5 moc-qt6"
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

    hast=$(grep "_t->dummy();" tmpmoc.cxx)
    echo $hast
    if [ x"$hast" == x"" ]; then
        sed -i 's/dummy();/metacallir(_id,_o);/' tmpmoc.cxx
    else
        sed -i 's/_t->dummy();/_t->metacallir(_o,_c,_id,_a);/' tmpmoc.cxx
    fi
    grep "metacallir" tmpmoc.cxx

    dstfile="qmitmsloter_moc_v${verstr}cxx"
    mv -v tmpmoc.cxx $dstfile

    echo
done

