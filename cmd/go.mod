module main

go 1.22.3

require (
	github.com/kitech/gopp v0.0.0
	github.com/kitech/minqt/qtinline v0.0.0
)

require github.com/orcaman/concurrent-map/v2 v2.0.1 // indirect

// require cmd/internal/objfile v0.0.0
// replace cmd/internal/objfile => /nix/store/1p6vr83cgyfwm8517jhfmf6lypzhy3q2-go-1.22.3/share/go/src/cmd/internal/objfile
require internal/objfile v0.0.0
replace internal/objfile => ./internal/objfile

require (
	github.com/ebitengine/purego v0.7.1 // indirect
	github.com/kitech/dl v0.0.0-20201225001532-be4f4faa4070 // indirect
	github.com/kitech/gopp/cgopp v0.0.0 // indirect
	github.com/qtui/qtclzsz v0.0.0 // indirect
	github.com/qtui/qtsyms v0.0.0
)

require (
	github.com/Workiva/go-datastructures v1.1.3 // indirect
	github.com/bitly/go-simplejson v0.5.1 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/dolthub/maphash v0.1.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/lytics/base62 v0.0.0-20180808010106-0ee4de5a5d6d // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/qtui/miscutil v0.0.0 // indirect
	github.com/qtui/qtmeta v0.0.0
	github.com/qtui/qtqt v0.0.0 // indirect
	github.com/qtui/qtrt v0.0.0
	github.com/smacker/go-tree-sitter v0.0.0-20240625050157-a31a98a7c0f6
	golang.org/x/sys v0.19.0 // indirect
)

replace github.com/qtui/qtrt => ../../qtui/qtrt

replace github.com/qtui/qtclzsz => ../../qtui/qtclzsz

replace github.com/qtui/qtsyms => ../../qtui/qtsyms

replace github.com/qtui/qtqt => ../../qtui/qtqt

replace github.com/qtui/qtmeta => ../../qtui/qtmeta

replace github.com/qtui/miscutil => ../../qtui/miscutil

replace github.com/kitech/minqt => ../

replace github.com/kitech/gopp => ../../goplusplus

replace github.com/kitech/gopp/cgopp => ../../goplusplus/cgopp

replace github.com/kitech/minqt/qtinline => ../qtinline
