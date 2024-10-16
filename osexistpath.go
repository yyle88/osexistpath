package osexistpath

import (
	"os"

	"github.com/pkg/errors"
	"github.com/yyle88/erero"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type LogVerb string

const (
	Quiet LogVerb = "QUIET" //比较安静的，即出错时不打印任何日志
	Noisy LogVerb = "NOISY" //比较吵闹的，即无论什么错误都会打印
	Sweet LogVerb = "SWEET" //有点适度的，即只打印预料之外的错误
)

// IsPathExists 检查这个路径下是否有东西
func IsPathExists(path string, verb LogVerb) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if verb == Noisy {
				zaplog.LOGS.P1.Debug("IS_PATH_EXISTS not exists", zap.String("path", path), zap.Bool("exist", false))
			}
			return false, nil // 路径不存在
		}
		if verb != Quiet {
			zaplog.LOGS.P1.Error("IS_PATH_EXISTS", zap.String("path", path), zap.Error(err))
		}
		return false, erero.Wro(err) // 其他的错误
	}
	return true, nil
}

// IsFileExists 检查文件是否存在返回布尔
func IsFileExists(path string, verb LogVerb) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if verb == Noisy {
				zaplog.LOGS.P1.Debug("IS_FILE_EXISTS not exists", zap.String("path", path), zap.Bool("exist", false))
			}
			return false, nil // 路径不存在
		}
		if verb != Quiet {
			zaplog.LOGS.P1.Error("IS_FILE_EXISTS wrong stat", zap.String("path", path), zap.Error(err))
		}
		return false, erero.Wro(err) // 其他的错误
	}
	if info.IsDir() {
		if verb == Noisy || verb == Sweet {
			zaplog.LOGS.P1.Debug("IS_FILE_EXISTS wrong type", zap.String("path", path), zap.String("type", "ROOT"))
		}
		// 这里必须返回错误，否则判定不存在接着创建文件就没法创建，而强行写/删也会出问题
		return false, errors.New("PATH-EXIST-BUT-TYPE-IS-NOT-FILE")
	}
	//这里依然要判断类型避免漏判
	return !info.IsDir(), nil
}

// IsRootExists 这个函数就表示目录是否存在
func IsRootExists(path string, verb LogVerb) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if verb == Noisy {
				zaplog.LOGS.P1.Debug("IS_ROOT_EXISTS not exists", zap.String("path", path), zap.Bool("exist", false))
			}
			return false, nil // 路径不存在
		}
		if verb != Quiet {
			zaplog.LOGS.P1.Error("IS_ROOT_EXISTS wrong stat", zap.String("path", path), zap.Error(err))
		}
		return false, erero.Wro(err) // 其他的错误
	}
	if !info.IsDir() {
		if verb == Noisy || verb == Sweet {
			zaplog.LOGS.P1.Debug("IS_FILE_EXISTS wrong type", zap.String("path", path), zap.String("type", "FILE"))
		}
		// 这里必须返回错误，否则判定不存在接着创建目录就没法创建，而强行建/删也会出问题
		return false, errors.New("PATH-EXIST-BUT-TYPE-IS-NOT-ROOT")
	}
	//这里依然要判断类型避免漏判
	return info.IsDir(), nil
}

// Path 检查这个路径下是否有东西
func Path(path string) (bool, error) {
	return IsPathExists(path, Quiet)
}

// File 检查文件是否存在返回布尔
func File(path string) (bool, error) {
	return IsFileExists(path, Quiet)
}

// Root 这个函数就表示目录是否存在，这里使用root表示目录，虽然不太贴切，但能保持函数名都是4个字母的
// 在我的开源项目里倾向于使用 root 就指目录，让代码更整齐些，虽然这样可能是不恰当的，但就这样吧
func Root(path string) (bool, error) {
	return IsRootExists(path, Quiet)
}

// PATH 假如存在就把路径返回，假如不存在就报错，返回路径再转换为单值就有用
func PATH(path string) (string, error) {
	exist, err := IsPathExists(path, Sweet)
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
	exist, err := IsFileExists(path, Sweet)
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
	exist, err := IsRootExists(path, Sweet)
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
	exist, err := IsPathExists(path, Noisy)
	if err != nil {
		zaplog.LOGS.P1.Panic("must_path", zap.String("path", path), zap.Error(err))
	}
	if !exist {
		zaplog.LOGS.P1.Panic("must_path", zap.String("path", path), zap.Bool("exist", false))
	}
}

// MustFile 检查文件是否存在否则崩掉
func MustFile(path string) {
	exist, err := IsFileExists(path, Noisy)
	if err != nil {
		zaplog.LOGS.P1.Panic("must_file", zap.String("path", path), zap.Error(err))
	}
	if !exist {
		zaplog.LOGS.P1.Panic("must_file", zap.String("path", path), zap.Bool("exist", false))
	}
}

// MustRoot 检查目录是否存在否则崩掉
func MustRoot(path string) {
	exist, err := IsRootExists(path, Noisy)
	if err != nil {
		zaplog.LOGS.P1.Panic("must_root", zap.String("path", path), zap.Error(err))
	}
	if !exist {
		zaplog.LOGS.P1.Panic("must_root", zap.String("path", path), zap.Bool("exist", false))
	}
}
