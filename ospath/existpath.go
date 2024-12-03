package ospath

import (
	"github.com/yyle88/osexistpath"
)

type existNamespace struct{}

var EXIST = &existNamespace{}

// IsPath 检查这个路径下是否有东西
func (*existNamespace) IsPath(path string) (bool, error) {
	return osexistpath.IsPath(path)
}

// IsFile 检查文件是否存在
func (*existNamespace) IsFile(path string) (bool, error) {
	return osexistpath.IsFile(path)
}

// IsRoot 这个函数就表示目录是否存在，这里使用root表示目录，虽然不太贴切，但能保持函数名都是4个字母的
// 在我的开源项目里倾向于使用 root 就指目录，让代码更整齐些，虽然这样可能是不恰当的，但就这样吧
func (*existNamespace) IsRoot(path string) (bool, error) {
	return osexistpath.IsRoot(path)
}

// PATH 假如存在就把路径返回，假如不存在就报错，返回路径再转换为单值就有用
func (T *existNamespace) PATH(path string) (string, error) {
	return osexistpath.PATH(path)
}

// FILE 假如存在就把路径返回，假如不存在就报错，返回路径再转换为单值就有用
func (T *existNamespace) FILE(path string) (string, error) {
	return osexistpath.FILE(path)
}

// ROOT 假如存在就把路径返回，假如不存在就报错，返回路径再转换为单值就有用
func (T *existNamespace) ROOT(path string) (string, error) {
	return osexistpath.ROOT(path)
}

func (T *existNamespace) MustPath(path string) {
	osexistpath.MustPath(path)
}

func (T *existNamespace) MustFile(path string) {
	osexistpath.MustFile(path)
}

func (T *existNamespace) MustRoot(path string) {
	osexistpath.MustRoot(path)
}
