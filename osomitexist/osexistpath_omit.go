package osomitexist

import (
	"github.com/yyle88/osexistpath"
	"github.com/yyle88/sure"
)

func IsPathExists(path string, verb LogVerb) bool {
	res0, err := osexistpath.IsPathExists(path, verb)
	sure.Omit(err)
	return res0
}

func IsFileExists(path string, verb LogVerb) bool {
	res0, err := osexistpath.IsFileExists(path, verb)
	sure.Omit(err)
	return res0
}

func IsRootExists(path string, verb LogVerb) bool {
	res0, err := osexistpath.IsRootExists(path, verb)
	sure.Omit(err)
	return res0
}

func Path(path string) bool {
	res0, err := osexistpath.Path(path)
	sure.Omit(err)
	return res0
}

func File(path string) bool {
	res0, err := osexistpath.File(path)
	sure.Omit(err)
	return res0
}

func Root(path string) bool {
	res0, err := osexistpath.Root(path)
	sure.Omit(err)
	return res0
}

func PATH(path string) string {
	res0, err := osexistpath.PATH(path)
	sure.Omit(err)
	return res0
}

func FILE(path string) string {
	res0, err := osexistpath.FILE(path)
	sure.Omit(err)
	return res0
}

func ROOT(path string) string {
	res0, err := osexistpath.ROOT(path)
	sure.Omit(err)
	return res0
}

func MustPath(path string) {
	osexistpath.MustPath(path)
}

func MustFile(path string) {
	osexistpath.MustFile(path)
}

func MustRoot(path string) {
	osexistpath.MustRoot(path)
}
