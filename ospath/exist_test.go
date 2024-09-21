package ospath

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
)

func Test_existNamespace_Path(t *testing.T) {
	for _, path := range []string{
		runpath.CurrentPath(),
		runpath.PARENT.Path(),
	} {
		t.Log(path)

		exist, err := EXIST.Path(path)
		require.NoError(t, err)
		require.True(t, exist)
	}
}

func Test_existNamespace_File(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)

	exist, err := EXIST.Path(path)
	require.NoError(t, err)
	require.True(t, exist)
}

func Test_existNamespace_Root(t *testing.T) {
	path := runpath.PARENT.Path()
	t.Log(path)

	exist, err := EXIST.Root(path)
	require.NoError(t, err)
	require.True(t, exist)
}
