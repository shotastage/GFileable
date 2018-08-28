package fileable

import (
	"os"
)

// Mkdir is a function that create a new directory.
func (f *Fileable) Mkdir(path string) error {

	err := os.MkdirAll(path, 0777)

	return err
}

// Rm is a function that remove files or directories.
func (f *Fileable) Rm(path string) error {

	err := os.RemoveAll(path)

	return err
}

// Mv is a function that rename or move directories or files.
func (f *Fileable) Mv(path, to string) error {

	err := os.Rename(path, to)

	return err
}

// Chmod is a function that change mode or permission.
func (f *Fileable) Chmod(target string, mode os.FileMode) error {

	err := os.Chmod(target, mode)

	return err
}

func (f *Fileable) Touch(name string) {
	println("This function is now under construction.")
}
