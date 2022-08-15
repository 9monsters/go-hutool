package hutool_all

import (
	_ "github.com/nine-monsters/go-hutool/hutool-core"
	_ "github.com/nine-monsters/go-hutool/hutool-jwt"
)

type Module struct {
	Path    string  // module path
	Version string  // module version
	Sum     string  // checksum
	Replace *Module // replaced by this module
}

func Hutool() {

}
