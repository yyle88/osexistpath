package osmustexist

import (
	"github.com/yyle88/osexistpath"
	"github.com/yyle88/sure"
)

func IsPathExists(path string, verb osexistpath.CheckMode) bool {
	res0, err := osexistpath.IsPathExists(path, verb)
	sure.Must(err)
	return res0
}

func IsFileExists(path string, verb osexistpath.CheckMode) bool {
	res0, err := osexistpath.IsFileExists(path, verb)
	sure.Must(err)
	return res0
}

func IsRootExists(path string, verb osexistpath.CheckMode) bool {
	res0, err := osexistpath.IsRootExists(path, verb)
	sure.Must(err)
	return res0
}

func IsPathExist(path string) bool {
	res0, err := osexistpath.IsPathExist(path)
	sure.Must(err)
	return res0
}

func IsFileExist(path string) bool {
	res0, err := osexistpath.IsFileExist(path)
	sure.Must(err)
	return res0
}

func IsRootExist(path string) bool {
	res0, err := osexistpath.IsRootExist(path)
	sure.Must(err)
	return res0
}

func IsPath(path string) bool {
	res0, err := osexistpath.IsPath(path)
	sure.Must(err)
	return res0
}

func IsFile(path string) bool {
	res0, err := osexistpath.IsFile(path)
	sure.Must(err)
	return res0
}

func IsRoot(path string) bool {
	res0, err := osexistpath.IsRoot(path)
	sure.Must(err)
	return res0
}

func PATH(path string) string {
	res0, err := osexistpath.PATH(path)
	sure.Must(err)
	return res0
}

func FILE(path string) string {
	res0, err := osexistpath.FILE(path)
	sure.Must(err)
	return res0
}

func ROOT(path string) string {
	res0, err := osexistpath.ROOT(path)
	sure.Must(err)
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
