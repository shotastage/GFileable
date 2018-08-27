package fileable

import (
	"fmt"
	"os"
	"os/user"
)

func (f Fileable) IsDir(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		fmt.Println(err)
	}
}

func (f Fileable) IsFile(file string) {

	_, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
	}
}

func (f Fileable) Pwd() string {

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
