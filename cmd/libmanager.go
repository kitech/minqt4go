package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ebitengine/purego"
	"github.com/kitech/gopp"
)

var libpfx = gopp.Mustify(os.UserHomeDir())[0].Str() + "/.nix-profile/lib"

type LibManager struct {
	libs map[string]usize
	// todo
	// current process linked qt shared libs, auto check
	// if not, need search system libpath
	linkcp bool
	opened bool
}

var Libman = LibManagerNew()

func LibManagerNew() *LibManager {
	me := &LibManager{}
	me.libs = map[string]usize{}
	return me
}

// Enter works
func PauseAk() {
	var c [1]byte
	n, err := os.Stdin.Read(c[:])
	gopp.ErrPrint(err, n)
}

func (me *LibManager) Open() {
	nowt := time.Now()
	globtmpl := fmt.Sprintf("%s/Qt*.framework/Qt*", libpfx)
	libs, err := filepath.Glob(globtmpl)
	gopp.ErrPrint(err, libs)
	// log.Println(libs, len(libs))

	{
		// hotfix for QString::QString(char const*)
		file := "~/Downloads/libQt5Inline.dylib" // not work qt6
		file = gopp.Mustify1(os.UserHomeDir()) + "/aprog/fedimqt/libhelloworld.dylib"
		dlh, err := purego.Dlopen(file, purego.RTLD_LAZY)
		gopp.ErrPrint(err, file)
		// log.Println(dlh)
		// PauseAk()
		if dlh != 0 {
			me.libs[filepath.Base(file)] = dlh
		}
	}

	gopp.Mapdo(libs, func(idx int, vx any) any {
		dlh, err := purego.Dlopen(vx.(string), purego.RTLD_LAZY)
		gopp.ErrPrint(err, vx)
		if dlh != 0 {
			me.libs[filepath.Base(vx.(string))] = dlh
		}
		return nil
	})
	log.Println(len(me.libs), me.libs, time.Since(nowt))
	me.opened = true
}

func (me *LibManager) Dlsym(name string) voidptr {
	name = name[1:]
	for mod, dlh := range me.libs {
		symptr, err := purego.Dlsym(dlh, name)
		gopp.ErrPrint(err, name, mod)
		if symptr != 0 {
			log.Println("found symbol in", mod, name)
			return voidptr(symptr)
		}
	}
	return nil
}
