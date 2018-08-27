![FileKit](./docs/filekit_readme.png)

[![CircleCI](https://circleci.com/gh/shotastage/FileableGo.svg?style=svg)](https://circleci.com/gh/shotastage/FileableGo)
![Go](https://img.shields.io/badge/Go-1.11-blue.svg)
![Linux supported](https://img.shields.io/badge/Linux-supported-5D9CEC.svg?style=flat)
[![Go Report Card](https://goreportcard.com/badge/github.com/shotastage/FileableGo?style=flat)](https://goreportcard.com/report/github.com/shotastage/FileableGo)




FileableGo is originated to [Fileable for Swift](https://github.com/shotastage/Fileable).
It enables to mange files or directories efficiently and easily.

# Installation

## dep

Using dep is strongly recommended to manage required package efficiently.
If you want to add `FileableGo` manually, run following command in Go project root directory.

```
$ dep ensure -add github.com/shotastage/FileableGo
```

## go get

You can also use `go get` command instead of any other package manager.

```
$ go get -u github.com/shotastage/FileableGo
```

# ⌘ APIs

| Function | |
|:--|:--|
| `func (f Fileable) Pwd() string` | Get current directory path as a string.|
| `func Home() string` | Get home directory path.|
| `func (f Fileable) IsFile(file string)` | Check the file exists or not. |
| `func (f Fileable) IsDir(dir string) ` | Check the directory exists or not.|
| `Swift: func ext(path: String) -> Bool ` | Check the directory exists or not. **Not implemented**|
| `Swift: func cd(path: String) throws`| Change directory like a `cd` command. **Not implemented**|
| `Swift: func mkdir(path: String) throws` | Make directory. **Not implemented**|
| `Swift: func rm(target: String) throws`| Remove directory or file. **Not implemented**|
| `Swift: func mv(from fromPath: String, to toPath: String)`| Move file or directory. **Not implemented**|
| `Swift: func touch(_ path: String) throws`| Create empty file.  **Not implemented**|



# License
Fileable is licensed under the `MIT`. 
You can use this library free of charge. See [LICENSE](./LICENSE) for detail.

