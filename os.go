package fileable

import (
	"fmt"
	"os"
)

func (f Fileable) IsDir() {
	_, err := os.Stat("hoge.txt")
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
	return ""
}
