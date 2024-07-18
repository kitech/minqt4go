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

func (me *LibManager) Open() {
	nowt := time.Now()
	globtmpl := fmt.Sprintf("%s/Qt*.framework/Qt*", libpfx)
	libs, err := filepath.Glob(globtmpl)
	gopp.ErrPrint(err, libs)
	// log.Println(libs, len(libs))

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
			log.Println("found symbol in", mod)
			return voidptr(symptr)
		}
	}
	return nil
}
