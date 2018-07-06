# gorename
[中文文档](./README.md)

*gorename* is a simple tool to help you change the package name。
# Getting Started
First, you need has install `golang` and has config system environment of `$GOPATH` and `$GOBIN`。
exec the following command install gorename:
```
go get github.com/yushuailiu/gorename
```

Now you can use gorename, if you want change the package `old/package/name`
to `new/package/name` of files which are in directory `/user/code/test`,you can exec the following command:
```
gorename --source /user/code/test old/package/name new/package/name
```

You can also run `gorename --help` from help.

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