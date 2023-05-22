package GFileable

import (
	"os"
	"os/user"
	"path/filepath"
)

// Fileable represents a file or directory.
type Fileable struct {
	Path string
}

// IsDir returns whether the path is a directory.
func (f Fileable) IsDir() (bool, error) {
	info, err := os.Stat(f.Path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

// IsFile returns whether the path is a file.
func (f Fileable) IsFile() (bool, error) {
	info, err := os.Stat(f.Path)
	if err != nil {
		return false, err
	}
	return !info.IsDir(), nil
}

// Extension returns the file extension.
func (f Fileable) Extension() (string, error) {
	info, err := os.Stat(f.Path)
	if err != nil {
		return "", err
	}
	if info.IsDir() {
		return "", nil
	}
	return filepath.Ext(f.Path), nil
}

// Pwd returns the current directory path.
func Pwd() (string, error) {
	return os.Getwd()
}

// Home returns the home directory path.
func Home() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.HomeDir, nil
}
