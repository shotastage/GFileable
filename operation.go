package fileable

import (
	"os"
)

// Mkdir is a function that create a new directory.
func (f *Fileable) Mkdir(path string) error {

	if err := os.MkdirAll(path, 0777); err != nil {
		return err
	}

	return nil
}

// Rm is a function that remove files or directories.
func (f *Fileable) Rm(path string) error {

	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return nil
}

// Mv is a function that rename or move directories or files.
func (f *Fileable) Mv(path, to string) error {

	if err := os.Rename(path, to); err != nil {
		return err
	}

	return nil
}

func (f *Fileable) Touch(name string) {
	println("This function is now under construction.")
}
