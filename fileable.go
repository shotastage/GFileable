package GFileable

import "os"

// Fileable is a interface for creating Fileable object.
type Fileable struct {
	Path string
}

// Path is a constructor.
func Path(path string) *Fileable {

	if path == "" {
		current, _ := os.Getwd()

		return &Fileable{Path: current}
	}

	return &Fileable{Path: path}
}
