all:
	echo "subcm: <run|cxlib>"

run:
	LD_LIBRARY_PATH=../srcc QTGO_DEBUG_FFI_CALL=1 GOGC=10 ./main keyword

cxlib:
	make -C ../srcc
	make -C ../wgtcc
	make -C ../qmlcc
	cp ../wgtcc/libQtWidgetsInline.* ../srcc/
	cp ../qmlcc/libQtQuickInline.* ../srcc/

