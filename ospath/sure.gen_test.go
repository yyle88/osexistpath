package ospath

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/sure_cls_gen"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	param := sure_cls_gen.NewGenParam(runpath.PARENT.Path())
	param.SetSubClassNamePartWords("_")
	param.SetSubClassNameStyleEnum(sure_cls_gen.STYLE_SUFFIX_CAMELCASE_TYPE)
	param.SetSureEnum(sure.MUST, sure.SOFT, sure.OMIT)

	cfg := &sure_cls_gen.Config{
		GenParam:      param,
		PkgName:       syntaxgo.CurrentPackageName(),
		ImportOptions: syntaxgo_ast.NewPackageImportOptions(),
		SrcPath:       runtestpath.SrcPath(t),
	}
	sure_cls_gen.Gen(cfg, existNamespace{})
}

func Test_existNamespace_Must_PATH(t *testing.T) {
	{
		//这样只要一行代码就能得到必然存在的路径
		path := EXIST.Must().PATH(runpath.CurrentPath())
		//使用（打印）路径
		t.Log(path)
	}
	{
		//这样只要一行代码就能得到必然存在的路径
		path := EXIST.Must().PATH(runpath.PARENT.Path())
		//使用（打印）路径
		t.Log(path)
	}
}

func Test_existNamespace_Must_FILE(t *testing.T) {
	//这样只要一行代码就能得到必然存在的文件
	path := EXIST.Must().FILE(runtestpath.SrcPath(t))
	//使用（打印）路径
	t.Log(path)
}

func Test_existNamespace_Must_ROOT(t *testing.T) {
	//这样只要一行代码就能得到必然存在的目录
	path := EXIST.Must().ROOT(runpath.PARENT.Path())
	//使用（打印）路径
	t.Log(path)
}
