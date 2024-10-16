package ospath_test

import (
	"testing"

	"github.com/yyle88/done"
	"github.com/yyle88/osexistpath/ospath"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
)

func Test_existNamespace_Path(t *testing.T) {
	for _, path := range []string{
		runpath.CurrentPath(),
		runpath.PARENT.Path(),
	} {
		t.Log(path)
		done.VBE(ospath.EXIST.Path(path)).TRUE()
	}
}

func Test_existNamespace_File(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)
	done.VBE(ospath.EXIST.Path(path)).TRUE()
}

func Test_existNamespace_Root(t *testing.T) {
	path := runpath.PARENT.Path()
	t.Log(path)
	done.VBE(ospath.EXIST.Root(path)).TRUE()
}

func Test_existNamespace_MustPath(t *testing.T) {
	ospath.EXIST.MustPath(runpath.CurrentPath())
	ospath.EXIST.MustPath(runpath.PARENT.Path())
}

func Test_existNamespace_MustFile(t *testing.T) {
	ospath.EXIST.MustFile(runtestpath.SrcPath(t))
}

func Test_existNamespace_MustRoot(t *testing.T) {
	ospath.EXIST.MustRoot(runpath.PARENT.Path())
}
