package osexistpath

import (
	"os"

	"github.com/pkg/errors"
	"github.com/yyle88/erero"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type CheckMode string

const (
	Noisy CheckMode = "NOISY" // Noisy mode, logs all errors // 比较吵闹的，即无论什么错误都会打印
	Sweet CheckMode = "SWEET" // Sweet mode, logs unexpected errors // 有点适度的，即只打印预料之外的错误
	Quiet CheckMode = "QUIET" // Quiet mode, no logs on errors // 比较安静的，即出错时不打印任何日志
	Might CheckMode = "MIGHT" // Might mode, checks type without raising errors // 表示询问的，即询问文件是何种类型的，在调用函数前没有预期，因此不抛出错误
)

// IsPathExists checks if there is something in this path
// IsPathExists 检查这个路径下是否有东西
func IsPathExists(path string, verb CheckMode) (bool, error) {
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

// IsFileExists checks if a file exists and returns a boolean
// IsFileExists 检查文件是否存在返回布尔
func IsFileExists(path string, verb CheckMode) (bool, error) {
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
		//这里不要直接打错误日志，需要根据情况而定
		erx := errors.New("PATH-EXIST-BUT-TYPE-IS-NOT-FILE")

		switch verb {
		case Noisy, Sweet:
			zaplog.LOGS.P1.Debug("IS_FILE_EXISTS wrong type", zap.String("path", path), zap.String("type", "ROOT"))
			return false, erx
		case Quiet:
			return false, erx
		case Might:
			// 这里不要返回错误，因为这里的意图是判定是不是这个类型，在逻辑中是if询问性质的，没有明确的预期
			return false, nil
		default:
			// 这里必须返回错误，否则判定不存在接着创建文件就没法创建，而强行写/删也会出问题
			return false, erx
		}
	}
	//这里依然要判断类型避免漏判
	return !info.IsDir(), nil
}

// IsRootExists checks if the directory exists
// IsRootExists 这个函数就表示目录是否存在
func IsRootExists(path string, verb CheckMode) (bool, error) {
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
		//这里不要直接打错误日志，需要根据情况而定
		erx := errors.New("PATH-EXIST-BUT-TYPE-IS-NOT-ROOT")

		switch verb {
		case Noisy, Sweet:
			zaplog.LOGS.P1.Debug("IS_FILE_EXISTS wrong type", zap.String("path", path), zap.String("type", "FILE"))
			return false, erx
		case Quiet:
			return false, erx
		case Might:
			// 这里不要返回错误，因为这里的意图是判定是不是这个类型，在逻辑中是if询问性质的，没有明确的预期
			return false, nil
		default:
			// 这里必须返回错误，否则判定不存在接着创建目录就没法创建，而强行建/删也会出问题
			return false, erx
		}
	}
	//这里依然要判断类型避免漏判
	return info.IsDir(), nil
}

// IsPathExist checks if there is something in the path
// IsPathExist 检查这个路径下是否有东西
func IsPathExist(path string) (bool, error) {
	return IsPathExists(path, Quiet)
}

// IsFileExist checks if the file exists and returns a boolean
// IsFileExist 检查文件是否存在返回布尔
func IsFileExist(path string) (bool, error) {
	return IsFileExists(path, Quiet)
}

// IsRootExist checks if the directory exists
// IsRootExist 这个函数就表示目录是否存在，这里使用root表示目录，保持函数名都是4个字母的
func IsRootExist(path string) (bool, error) {
	return IsRootExists(path, Quiet)
}

// IsPath checks if there is something in the path
// IsPath 检查这个路径下是否有东西
func IsPath(path string) (bool, error) {
	return IsPathExists(path, Might)
}

// IsFile checks if the file exists and returns a boolean
// IsFile 检查文件是否存在返回布尔
func IsFile(path string) (bool, error) {
	return IsFileExists(path, Might)
}

// IsRoot checks if the directory exists
// IsRoot 这个函数就表示目录是否存在，这里使用root表示目录，保持函数名都是4个字母的
func IsRoot(path string) (bool, error) {
	return IsRootExists(path, Might)
}

// PATH checks if the path exists, otherwise returns an error
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

// FILE checks if the file exists, otherwise returns an error
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

// ROOT checks if the root directory exists, otherwise returns an error
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

// MustPath checks if the path exists, otherwise panics
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

// MustFile checks if the file exists, otherwise panics
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

// MustRoot checks if the directory exists, otherwise panics
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
