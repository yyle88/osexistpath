package osexistpath

import (
	"os"

	"github.com/yyle88/erero"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// IsPathExists 检查这个路径下是否有东西
func IsPathExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			zaplog.LOGS.P1.Debug("IS_PATH_EXISTS", zap.String("path", path), zap.Bool("exist", false))
			return false, nil // 路径不存在
		}
		zaplog.LOGS.P1.Error("IS_PATH_EXISTS", zap.String("path", path), zap.Error(err))
		return false, erero.Wro(err) // 其他的错误
	}
	return true, nil
}

// IsFileExists 检查文件是否存在返回布尔
func IsFileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			zaplog.LOGS.P1.Debug("IS_FILE_EXISTS", zap.String("path", path), zap.Bool("exist", false))
			return false, nil // 路径不存在
		}
		zaplog.LOGS.P1.Error("IS_FILE_EXISTS", zap.String("path", path), zap.Error(err))
		return false, erero.Wro(err) // 其他的错误
	}
	if info.IsDir() {
		zaplog.LOGS.P1.Debug("IS_FILE_EXISTS", zap.String("path", path), zap.String("type", "root"))
		// 这里必须返回错误，否则判定不存在接着创建文件就没法创建，而强行写/删也会出问题
		return false, erero.New("PATH-EXIST-BUT-TYPE-IS-NOT-FILE")
	}
	//这里依然要判断类型避免漏判
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
			zaplog.LOGS.P1.Debug("IS_ROOT_EXISTS", zap.String("path", path), zap.Bool("exist", false))
			return false, nil // 路径不存在
		}
		zaplog.LOGS.P1.Error("IS_ROOT_EXISTS", zap.String("path", path), zap.Error(err))
		return false, erero.Wro(err) // 其他的错误
	}
	if !info.IsDir() {
		zaplog.LOGS.P1.Debug("IS_FILE_EXISTS", zap.String("path", path), zap.String("type", "file"))
		// 这里必须返回错误，否则判定不存在接着创建目录就没法创建，而强行建/删也会出问题
		return false, erero.New("PATH-EXIST-BUT-TYPE-IS-NOT-ROOT")
	}
	//这里依然要判断类型避免漏判
	return info.IsDir(), nil
}

// Path 检查这个路径下是否有东西
func Path(path string) (bool, error) {
	return IsPathExists(path)
}

// File 检查文件是否存在返回布尔
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
		zaplog.LOGS.P1.Error("path", zap.String("path", path), zap.Error(err))
		return "", erero.Wro(err)
	}
	if !exist {
		zaplog.LOGS.P1.Debug("path", zap.String("path", path), zap.Bool("exist", false))
		return "", os.ErrNotExist
	}
	return path, nil
}

// FILE 假如存在就把路径返回，假如不存在就报错，返回路径再转换为单值就有用
func FILE(path string) (string, error) {
	exist, err := IsFileExists(path)
	if err != nil {
		zaplog.LOGS.P1.Error("file", zap.String("path", path), zap.Error(err))
		return "", erero.Wro(err)
	}
	if !exist {
		zaplog.LOGS.P1.Debug("file", zap.String("path", path), zap.Bool("exist", false))
		return "", os.ErrNotExist
	}
	return path, nil
}

// ROOT 假如存在就把路径返回，假如不存在就报错，返回路径再转换为单值就有用
func ROOT(path string) (string, error) {
	exist, err := IsRootExists(path)
	if err != nil {
		zaplog.LOGS.P1.Error("root", zap.String("path", path), zap.Error(err))
		return "", erero.Wro(err)
	}
	if !exist {
		zaplog.LOGS.P1.Debug("root", zap.String("path", path), zap.Bool("exist", false))
		return "", os.ErrNotExist
	}
	return path, nil
}

// MustPath 检查路径是否存在否则崩掉
func MustPath(path string) {
	exist, err := Path(path)
	if err != nil {
		zaplog.LOGS.P1.Panic("must_path", zap.String("path", path), zap.Error(err))
	}
	if !exist {
		zaplog.LOGS.P1.Panic("must_path", zap.String("path", path), zap.Bool("exist", false))
	}
}

// MustFile 检查文件是否存在否则崩掉
func MustFile(path string) {
	exist, err := File(path)
	if err != nil {
		zaplog.LOGS.P1.Panic("must_file", zap.String("path", path), zap.Error(err))
	}
	if !exist {
		zaplog.LOGS.P1.Panic("must_file", zap.String("path", path), zap.Bool("exist", false))
	}
}

// MustRoot 检查目录是否存在否则崩掉
func MustRoot(path string) {
	exist, err := Root(path)
	if err != nil {
		zaplog.LOGS.P1.Panic("must_root", zap.String("path", path), zap.Error(err))
	}
	if !exist {
		zaplog.LOGS.P1.Panic("must_root", zap.String("path", path), zap.Bool("exist", false))
	}
}
