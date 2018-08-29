package fileable

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

	err := os.RemoveAll(f.path)

	return err
}

// Mv is a function that rename or move directories or files.
func (f Fileable) Mv(to string) error {

	err := os.Rename(f.path, to)

	return err
}

// Chmod is a function that change mode or permission.
func (f Fileable) Chmod(mode os.FileMode) error {

	err := os.Chmod(f.path, mode)

	return err
}

func Touch(name string) {
	println("This function is now under construction.")
}
