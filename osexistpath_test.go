package osexistpath_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/osexistpath"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
)

func TestIsPathExists(t *testing.T) {
	done.VBE(osexistpath.IsPathExists(runpath.CurrentPath())).TRUE()
	done.VBE(osexistpath.IsPathExists(runpath.PARENT.Path())).TRUE()

	done.VBE(osexistpath.IsPathExists(runpath.PARENT.Join("abc.txt"))).FALSE()
	done.VBE(osexistpath.IsPathExists(runpath.PARENT.Join("xyz.uvw"))).FALSE()
}

func TestIsFileExists(t *testing.T) {
	done.VBE(osexistpath.IsFileExists(runtestpath.SrcPath(t))).TRUE()

	done.VBE(osexistpath.IsFileExists(runpath.PARENT.Join("abc.txt"))).FALSE()
	done.VBE(osexistpath.IsFileExists(runpath.PARENT.Join("xyz.uvw"))).FALSE()

	require.Error(t, done.VBE(osexistpath.IsFileExists(runpath.PARENT.Path())).E)
}

func TestIsRootExists(t *testing.T) {
	done.VBE(osexistpath.IsRootExists(runpath.PARENT.Path())).TRUE()

	done.VBE(osexistpath.IsRootExists(runpath.PARENT.Join("abc.txt"))).FALSE()
	done.VBE(osexistpath.IsRootExists(runpath.PARENT.Join("xyz.uvw"))).FALSE()

	require.Error(t, done.VBE(osexistpath.IsRootExists(runtestpath.SrcPath(t))).E)
}

func TestPath(t *testing.T) {
	done.VBE(osexistpath.Path(runpath.CurrentPath())).TRUE()
	done.VBE(osexistpath.Path(runpath.PARENT.Path())).TRUE()

	done.VBE(osexistpath.Path(runpath.PARENT.Join("abc.txt"))).FALSE()
	done.VBE(osexistpath.Path(runpath.PARENT.Join("xyz.uvw"))).FALSE()
}

func TestFile(t *testing.T) {
	done.VBE(osexistpath.File(runtestpath.SrcPath(t))).TRUE()
}

func TestRoot(t *testing.T) {
	done.VBE(osexistpath.Root(runpath.PARENT.Path())).TRUE()
}

func TestMustPath(t *testing.T) {
	osexistpath.MustPath(runpath.CurrentPath())
	osexistpath.MustPath(runpath.PARENT.Path())
}

func TestMustFile(t *testing.T) {
	osexistpath.MustFile(runtestpath.SrcPath(t))
}

func TestMustRoot(t *testing.T) {
	osexistpath.MustRoot(runpath.PARENT.Path())
}
