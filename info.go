package GFileable

import (
	"os"
	"os/user"
	"strings"
)

// IsDir is a function that decide whether path is a directory.
func (f Fileable) IsDir() bool {

	_, err := os.Stat(f.Path)

	if err != nil {
		return false
	}

	return true
}

// IsFile is a function that decide whether path is a file.
func (f Fileable) IsFile() bool {

	_, err := os.Stat(f.Path)

	if err != nil {
		return false
	}

	return true
}

// Extension is a function that return file extension.
func (f Fileable) Extension() string {
	arr := strings.Split(f.Path, "/")

	ext := strings.Split(arr[len(arr)-1], ".")[1]

	return ext
}

// Pwd is a function that return current directory path.
func Pwd() string {

	current, err := os.Getwd()

	if err != nil {
		return ""
	}

	return current
}

// Home is a function that return home directory path.
func Home() string {

	user, err := user.Current()

	if err != nil {
		return ""
	}

	return user.HomeDir
}
