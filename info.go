package fileable

import (
	"os"
	"os/user"
)

func (f Fileable) IsDir() bool {

	_, err := os.Stat(f.path)

	if err != nil {
		return false
	}

	return true
}

func (f Fileable) IsFile() bool {

	_, err := os.Stat(f.path)

	if err != nil {
		return false
	}

	return true
}

func Pwd() string {

	current, err := os.Getwd()

	if err != nil {
		return ""
	}

	return current
}

func Home() string {

	user, err := user.Current()

	if err != nil {
		return ""
	}

	return user.HomeDir
}
