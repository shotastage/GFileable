package fileable

import "os"

type Fileable struct {
	path string
}

func Path(path string) *Fileable {

	if path == "" {
		current, _ := os.Getwd()

		return &Fileable{path: current}
	}

	return &Fileable{path: path}
}
