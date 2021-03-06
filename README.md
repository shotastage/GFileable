![FileKit](./docs/filekit_readme.png)

[![CircleCI](https://circleci.com/gh/shotastage/GFileable.svg?style=svg)](https://circleci.com/gh/shotastage/GFileable)
![Go](https://img.shields.io/badge/Go-1.12-blue.svg)
![Linux supported](https://img.shields.io/badge/Linux-supported-5D9CEC.svg?style=flat)
[![Go Report Card](https://goreportcard.com/badge/github.com/shotastage/FileableGo?style=flat)](https://goreportcard.com/report/github.com/shotastage/FileableGo)




GFileable is originated to [Fileable for Swift](https://github.com/shotastage/Fileable).
It enables to mange files or directories efficiently and easily.

# Installation

You can also use `go get` command instead of any other package manager.

```:shell
go get -u github.com/shotastage/GFileable
```

# ⌘ APIs

| Function | |
|:--|:--|
| `func Pwd() string` | Get current directory path as a string.|
| `func Home() string` | Get home directory path.|
| `func (f Fileable) IsFile() bool` | Check the file exists or not. |
| `func (f Fileable) IsDir() bool` | Check the directory exists or not.|
| `func (f Fileable) Extension() string` | Get file extension. |
| `func Cd(to string)`| Change directory like a `cd` command.|
| `func Mkdir(path string) error` | Make directory.|
| `func (f Fileable) Rm() error`| Remove directory or file.|
| `func (f Fileable) Mv(to string) error`| Move file or directory.|
| `func Touch(name string) error`| Create empty file.|
| `func (f Fileable) Chmod(mode os.FileMode) error`| Change file permission.|


# License

Fileable is licensed under the `MIT`. 
You can use this library free of charge. See [LICENSE](./LICENSE) for detail.
