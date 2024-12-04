[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/osexistpath/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/osexistpath/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/osexistpath)](https://pkg.go.dev/github.com/yyle88/osexistpath)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/osexistpath/master.svg)](https://coveralls.io/github/yyle88/osexistpath?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/osexistpath.svg)](https://github.com/yyle88/osexistpath/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/osexistpath)](https://goreportcard.com/report/github.com/yyle88/osexistpath)

# osexistpath - A Go Library for Checking Path, File, and Directory Existence

## Introduction

`osexistpath` is designed to check whether a path, file, or directory exists.

## CHINESE README

[中文说明](README.zh.md)

## Features

- **IsPathExists**: Checks whether a path exists.
- **IsFileExists**: Checks whether a file exists.
- **IsRootExists**: Checks is a directory exists.
- **PATH, FILE, ROOT**: Functions that return the path if it exists, otherwise raise a panic.

## Installation

```bash
go get github.com/yyle88/osexistpath
```

## Usage

### Example 1: Checking if a Path Exists

```go
package main

import (
	"fmt"
	"github.com/yyle88/osexistpath"
)

func main() {
	// Check if a path exists using the "Might" verbosity level
	path := "/some/path"
	exist, err := osexistpath.IsPathExists(path, osexistpath.Might)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Path exists:", exist)
	}
}
```

### Example 2: Checking if a File Exists

```go
package main

import (
	"fmt"
	"github.com/yyle88/osexistpath"
)

func main() {
	// Check if a file exists using the "Sweet" verbosity level
	file := "/some/file.txt"
	exist, err := osexistpath.IsFileExists(file, osexistpath.Sweet)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File exists:", exist)
	}
}
```

### Example 3: Checking if a Directory Exists

```go
package main

import (
	"fmt"
	"github.com/yyle88/osexistpath"
)

func main() {
	// Check if a directory exists using the "Noisy" verbosity level
	dir := "/some/directory"
	exist, err := osexistpath.IsRootExists(dir, osexistpath.Noisy)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Directory exists:", exist)
	}
}
```

---

## License

`osexistpath` is open-source and released under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## Support

Welcome to contribute to this project by submitting pull requests or reporting issues.

If you find this package helpful, give it a star on GitHub!

**Thank you for your support!**

**Happy Coding with `osexistpath`!** 🎉

Give me stars. Thank you!!!

## See stars
[![see stars](https://starchart.cc/yyle88/osexistpath.svg?variant=adaptive)](https://starchart.cc/yyle88/osexistpath)
