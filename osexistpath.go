package osexistpath

import (
	"os"

	"github.com/yyle88/erero"
)

// IsPathExists 检查这个路径下是否有东西
func IsPathExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil // 路径不存在
		}
		return false, erero.Wro(err) // 其他的错误
	}
	return true, nil
}

// IsFileExists 检查文件是否存在
func IsFileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // 路径不存在
		}
		return false, erero.Wro(err) // 其他的错误
	}
	return !info.IsDir(), nil
}

// IsRootExists 这个函数就表示目录是否存在
// 因为我觉得 dir 这个单词太丑，让人看着就想吐，而 directory 过长且含义不明
// 因此在我的所有开源项目里，我拒绝使用 dir 和 directory 这俩单词
// 在我的开源项目里 root 就指目录
func IsRootExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // 路径不存在
		}
		return false, erero.Wro(err) // 其他的错误
	}
	return info.IsDir(), nil
}

// Path 检查这个路径下是否有东西
func Path(path string) (bool, error) {
	return IsPathExists(path)
}

// File 检查文件是否存在
func File(path string) (bool, error) {
	return IsFileExists(path)
}

// Root 这个函数就表示目录是否存在
// 因为我觉得 dir 这个单词太丑，让人看着就想吐，而 directory 过长且含义不明
// 因此在我的所有开源项目里，我拒绝使用 dir 和 directory 这俩单词
// 在我的开源项目里 root 就指目录
func Root(path string) (bool, error) {
	return IsRootExists(path)
}

// PATH 假如存在就把路径返回，假如不存在就报错，返回路径再转换为单值就有用
func PATH(path string) (string, error) {
	exist, err := IsPathExists(path)
	if err != nil {
		return "", erero.Wro(err)
	}
	if !exist {
		return "", os.ErrNotExist
	}
	return path, nil
}

// FILE 假如存在就把路径返回，假如不存在就报错，返回路径再转换为单值就有用
func FILE(path string) (string, error) {
	exist, err := IsFileExists(path)
	if err != nil {
		return "", erero.Wro(err)
	}
	if !exist {
		return "", os.ErrNotExist
	}
	return path, nil
}

// ROOT 假如存在就把路径返回，假如不存在就报错，返回路径再转换为单值就有用
func ROOT(path string) (string, error) {
	exist, err := IsRootExists(path)
	if err != nil {
		return "", erero.Wro(err)
	}
	if !exist {
		return "", os.ErrNotExist
	}
	return path, nil
}
