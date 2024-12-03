# osexistpath - 检查路径、文件和目录是否存在的 Go 库

## 介绍

`osexistpath` 旨在检查路径、文件或目录是否存在。

## 特性

- **IsPathExists**: 检查路径是否存在。
- **IsFileExists**: 检查文件是否存在。
- **IsRootExists**: 检查目录是否存在。
- **PATH, FILE, ROOT**: 假如路径存在，则返回路径，否则会引发 panic。

## 安装

```bash
go get github.com/yyle88/osexistpath
```

## 使用示例

### 示例 1：检查路径是否存在

```go
package main

import (
	"fmt"
	"github.com/yyle88/osexistpath"
)

func main() {
	// 使用 "Might" 详细级别检查路径是否存在
	path := "/some/path"
	exist, err := osexistpath.IsPathExists(path, osexistpath.Might)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("路径存在:", exist)
	}
}
```

### 示例 2：检查文件是否存在

```go
package main

import (
	"fmt"
	"github.com/yyle88/osexistpath"
)

func main() {
	// 使用 "Sweet" 详细级别检查文件是否存在
	file := "/some/file.txt"
	exist, err := osexistpath.IsFileExists(file, osexistpath.Sweet)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("文件存在:", exist)
	}
}
```

### 示例 3：检查目录是否存在

```go
package main

import (
	"fmt"
	"github.com/yyle88/osexistpath"
)

func main() {
	// 使用 "Noisy" 详细级别检查目录是否存在
	dir := "/some/directory"
	exist, err := osexistpath.IsRootExists(dir, osexistpath.Noisy)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("目录存在:", exist)
	}
}
```

---

## 许可

`osexistpath` 是一个开源项目，发布于 MIT 许可证下。有关更多信息，请参阅 [LICENSE](LICENSE) 文件。

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
