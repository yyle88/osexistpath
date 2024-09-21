package ospath

import (
	"github.com/yyle88/osexistpath"
)

type existNamespace struct{}

var EXIST = &existNamespace{}

// Path 检查这个路径下是否有东西
func (*existNamespace) Path(path string) (bool, error) {
	return osexistpath.Path(path)
}

// File 检查文件是否存在
func (*existNamespace) File(path string) (bool, error) {
	return osexistpath.File(path)
}

// Root 这个函数就表示目录是否存在
// 因为我觉得 dir 这个单词太丑，让人看着就想吐，而 directory 过长且含义不明
// 因此在我的所有开源项目里，我拒绝使用 dir 和 directory 这俩单词
// 在我的开源项目里 root 就指目录
func (*existNamespace) Root(path string) (bool, error) {
	return osexistpath.Root(path)
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