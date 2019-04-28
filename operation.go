package GFileable

import (
	"os"
)

// Mkdir is a function that create a new directory.
func Mkdir(path string) error {

	err := os.MkdirAll(path, 0777)

	return err
}

// Rm is a function that remove files or directories.
func (f Fileable) Rm() error {

	err := os.RemoveAll(f.Path)

	return err
}

// Mv is a function that rename or move directories or files.
func (f Fileable) Mv(to string) error {

	err := os.Rename(f.Path, to)

	return err
}

// Chmod is a function that change mode or permission.
func (f Fileable) Chmod(mode os.FileMode) error {

	err := os.Chmod(f.Path, mode)

	return err
}

// Touch is a function for creating a empty file.
func Touch(name string) error {

	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

// Cd is a function for moving current directory.
func Cd(to string) {
	os.Chdir(to)
}
