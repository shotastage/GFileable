package GFileable

import (
	"os"
	"os/user"
	"strings"
)

// Fileable is a interface for creating Fileable object.
type Fileable struct {
	Path string
}

// Path is a constructor.
func Path(path string) *Fileable {

	if strings.Contains(path, "~") || strings.Contains(path, "$HOME") {

		user, err := user.Current()

		if err != nil {
			panic("Failed to get home dir")
		}

		home := user.HomeDir

		hpath := strings.Replace(path, "~", home, 1)

		hpath = strings.Replace(path, "$HOME", home, 1)

		return &Fileable{Path: hpath}

	}

	if path == "" {
		current, _ := os.Getwd()

		return &Fileable{Path: current}
	}

	return &Fileable{Path: path}
}
