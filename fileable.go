package GFileable

import (
	"os"
	"strings"
)

// Fileable is a interface for creating Fileable object.
type Fileable struct {
	Path string
}

// Path is a constructor.
func Path(path string) *Fileable {

	if strings.Contains(path, "~") {

		home := Home()

		hpath := strings.Replace(path, "~", home, 1)

		return &Fileable{Path: hpath}

	}

	if strings.Contains(path, "$HOME") {

		home := Home()

		hpath := strings.Replace(path, "$HOME", home, 1)

		return &Fileable{Path: hpath}

	}

	if path == "" {
		current, _ := os.Getwd()

		return &Fileable{Path: current}
	}

	return &Fileable{Path: path}
}

func Join(path ...string) *Fileable {
	return &Fileable{Path: strings.Join(path, "")}
}
