package os

import "os"

func Pwd() string {

	current, err := os.Getwd()

	if err != nil {
		return ""
	}

	return current
}

func Home() string {
	return ""
}
