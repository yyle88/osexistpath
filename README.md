# osexistpath
检查路径是否存在，路径文件是否存在，路径的目录是否存在。因为没有开源包来专门做这件小事，就由我来做吧。

## 具体使用
安装：
```
go get github.com/yyle88/osexistpath
```

使用：
```
osexistpath.MustFile(path) //检查文件是否存在，假如不存在就直接panic
```

```
if osmustexist.File(path) { //检查文件是否存在，假如有错误就直接panic
    //读取内容
}
```

```
if ossoftexist.File(path) { //检查文件是否存在，假如有错误就返回false
    //读取内容
}
```

```
path := osmustexist.FILE(path) //检查文件是否存在，假如已存在就返回路径，否则就直接panic
```

```
path := ossoftexist.FILE(path) //检查文件是否存在，假如已存在就返回路径，否则就返回空字符串
```

## 测试用例
[主要用例](./osexistpath_test.go)

[其它用例](./ospath/exist_test.go)

[其它用例](./ospath/sure.gen_test.go)

## 其它项目
[获取路径](https://github.com/yyle88/runpath)

[软硬兼施](https://github.com/yyle88/sure)

[确保结果](https://github.com/yyle88/done)
