package GFileable

import (
	"fmt"
	"os"
)

// Fileable represents a file or directory.
type Fileable struct {
	Path string
}

// Mkdir creates a new directory.
func Mkdir(path string) error {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %v", path, err)
	}
	return nil
}

// Rm removes a file or directory.
func (f Fileable) Rm() error {
	err := os.RemoveAll(f.Path)
	if err != nil {
		return fmt.Errorf("failed to remove %s: %v", f.Path, err)
	}
	return nil
}

// Mv renames or moves a file or directory.
func (f Fileable) Mv(to string) error {
	err := os.Rename(f.Path, to)
	if err != nil {
		return fmt.Errorf("failed to move/rename %s to %s: %v", f.Path, to, err)
	}
	return nil
}

// Chmod changes the mode or permission of a file or directory.
func (f Fileable) Chmod(mode os.FileMode) error {
	err := os.Chmod(f.Path, mode)
	if err != nil {
		return fmt.Errorf("failed to change permissions of %s: %v", f.Path, err)
	}
	return nil
}

// Touch creates a new empty file.
func Touch(name string) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", name, err)
	}
	file.Close() // No need for defer since we don't do anything after this.
	return nil
}

// Cd changes the current directory.
func Cd(to string) error {
	err := os.Chdir(to)
	if err != nil {
		return fmt.Errorf("failed to change directory to %s: %v", to, err)
	}
	return nil
}
