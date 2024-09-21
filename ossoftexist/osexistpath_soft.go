package ossoftexist

import (
	"github.com/yyle88/osexistpath"
	"github.com/yyle88/sure"
)

func IsPathExists(path string) bool {
	res0, err := osexistpath.IsPathExists(path)
	sure.Soft(err)
	return res0
}

func IsFileExists(path string) bool {
	res0, err := osexistpath.IsFileExists(path)
	sure.Soft(err)
	return res0
}

func IsRootExists(path string) bool {
	res0, err := osexistpath.IsRootExists(path)
	sure.Soft(err)
	return res0
}

func Path(path string) bool {
	res0, err := osexistpath.Path(path)
	sure.Soft(err)
	return res0
}

func File(path string) bool {
	res0, err := osexistpath.File(path)
	sure.Soft(err)
	return res0
}

func Root(path string) bool {
	res0, err := osexistpath.Root(path)
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