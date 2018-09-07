package GFileable

import "os"

// Fileable is a interface for creating Fileable object.
type Fileable struct {
	path string
}

// Path is a constructor.
func Path(path string) *Fileable {

	if path == "" {
		current, _ := os.Getwd()

		return &Fileable{path: current}
	}

	return &Fileable{path: path}
}
