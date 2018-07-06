# gorename
[English Document](./README_en.md)

*gorename*是一个修改 golang 包名的小工具
# 安装和使用
要使用`gorename`首先需要安装 golang，并且配置环境变量 `$GOPATH` 和 `$GOBIN`，然后就可以执行以下命令安装：
```
go get github.com/yushuailiu/gorename
```
现在就可以使用 *gorename* 了，比如你想要把目录`/user/code/test`下的所有文件包含的包名 `old/package/name` 修改为 `new/package/name`那么执行以下
命令即可：
```
gorename --source /user/code/test old/package/name new/package/name
```
你也可以执行 `gorename --help` 查看帮助文档：

```
$ gorename --help                                                                                                                                                           [17:14:31]
NAME:
   gorename - Rename golang package

USAGE:
   main [global options] command [command options] [source file or directory path] [old package name] [new package name]

VERSION:
   0.0.1

AUTHOR:
   YuShuai Liu <admin@liuyushuai.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --source value, -s value  source package path or file path (default: "./")
   --help, -h                show help
   --version, -v             print the version
```