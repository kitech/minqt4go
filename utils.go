package minqt

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kitech/gopp"
)

var qtlibpfx = gopp.Mustify(os.UserHomeDir())[0].Str() + "/.nix-profile/lib"

func GetLibPaths() []string {
	globtmpl := fmt.Sprintf("%s/Qt*.framework/Qt*", qtlibpfx)
	libs, err := filepath.Glob(globtmpl)
	gopp.ErrPrint(err, libs)
	return libs
}
