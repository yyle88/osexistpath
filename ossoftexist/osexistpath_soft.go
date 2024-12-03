package ossoftexist

import (
	"github.com/yyle88/osexistpath"
	"github.com/yyle88/sure"
)

func IsPathExists(path string, verb osexistpath.StatVerb) bool {
	res0, err := osexistpath.IsPathExists(path, verb)
	sure.Soft(err)
	return res0
}

func IsFileExists(path string, verb osexistpath.StatVerb) bool {
	res0, err := osexistpath.IsFileExists(path, verb)
	sure.Soft(err)
	return res0
}

func IsRootExists(path string, verb osexistpath.StatVerb) bool {
	res0, err := osexistpath.IsRootExists(path, verb)
	sure.Soft(err)
	return res0
}

func IsPath(path string) bool {
	res0, err := osexistpath.IsPath(path)
	sure.Soft(err)
	return res0
}

func IsFile(path string) bool {
	res0, err := osexistpath.IsFile(path)
	sure.Soft(err)
	return res0
}

func IsRoot(path string) bool {
	res0, err := osexistpath.IsRoot(path)
	sure.Soft(err)
	return res0
}

func PATH(path string) string {
	res0, err := osexistpath.PATH(path)
	sure.Soft(err)
	return res0
}

func FILE(path string) string {
	res0, err := osexistpath.FILE(path)
	sure.Soft(err)
	return res0
}

func ROOT(path string) string {
	res0, err := osexistpath.ROOT(path)
	sure.Soft(err)
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
