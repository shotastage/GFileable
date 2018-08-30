package fileable

import (
	"os"
	"os/user"
)

// IsDir is a function that decide whether path is a directory.
func (f Fileable) IsDir() bool {

	_, err := os.Stat(f.path)

	if err != nil {
		return false
	}

	return true
}

// IsFile is a function that decide whether path is a file.
func (f Fileable) IsFile() bool {

	_, err := os.Stat(f.path)

	if err != nil {
		return false
	}

	return true
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
