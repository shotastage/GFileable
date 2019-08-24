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

	path = homeExpression(path)

	if path == "" {
		current, _ := os.Getwd()

		return &Fileable{Path: current}
	}

	return &Fileable{Path: path}
}

func Join(path ...interface{}) *Fileable {

	processing := make([]string, len(path), 10)

	for i := 0; i < len(path); i++ {

		if value, ok := path[i].(string); ok {
			processing[i] = homeExpression(value)
		}

		if value, ok := path[i].(*Fileable); ok {
			processing[i] = value.Path
		}
	}

	return &Fileable{Path: strings.Join(processing, "/")}
}
