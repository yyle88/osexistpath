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
	done.VBE(osexistpath.IsPathExists(runpath.CurrentPath(), osexistpath.Sweet)).TRUE()
	done.VBE(osexistpath.IsPathExists(runpath.PARENT.Path(), osexistpath.Sweet)).TRUE()

	done.VBE(osexistpath.IsPathExists(runpath.PARENT.Join("abc.txt"), osexistpath.Sweet)).FALSE()
	done.VBE(osexistpath.IsPathExists(runpath.PARENT.Join("xyz.uvw"), osexistpath.Sweet)).FALSE()
}

func TestIsFileExists(t *testing.T) {
	done.VBE(osexistpath.IsFileExists(runtestpath.SrcPath(t), osexistpath.Sweet)).TRUE()

	done.VBE(osexistpath.IsFileExists(runpath.PARENT.Join("abc.txt"), osexistpath.Sweet)).FALSE()
	done.VBE(osexistpath.IsFileExists(runpath.PARENT.Join("xyz.uvw"), osexistpath.Sweet)).FALSE()

	require.Error(t, done.VBE(osexistpath.IsFileExists(runpath.PARENT.Path(), osexistpath.Sweet)).E)
}

func TestIsRootExists(t *testing.T) {
	done.VBE(osexistpath.IsRootExists(runpath.PARENT.Path(), osexistpath.Sweet)).TRUE()

	done.VBE(osexistpath.IsRootExists(runpath.PARENT.Join("abc.txt"), osexistpath.Sweet)).FALSE()
	done.VBE(osexistpath.IsRootExists(runpath.PARENT.Join("xyz.uvw"), osexistpath.Sweet)).FALSE()

	require.Error(t, done.VBE(osexistpath.IsRootExists(runtestpath.SrcPath(t), osexistpath.Sweet)).E)
}

func TestPath(t *testing.T) {
	done.VBE(osexistpath.IsPath(runpath.CurrentPath())).TRUE()
	done.VBE(osexistpath.IsPath(runpath.PARENT.Path())).TRUE()

	done.VBE(osexistpath.IsPath(runpath.PARENT.Join("abc.txt"))).FALSE()
	done.VBE(osexistpath.IsPath(runpath.PARENT.Join("xyz.uvw"))).FALSE()
}

func TestFile(t *testing.T) {
	done.VBE(osexistpath.IsFile(runtestpath.SrcPath(t))).TRUE()
}

func TestRoot(t *testing.T) {
	done.VBE(osexistpath.IsRoot(runpath.PARENT.Path())).TRUE()
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
