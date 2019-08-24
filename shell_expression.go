package GFileable

import "strings"

func homeExpression(path string) string {

	if strings.Contains(path, "~") {

		home := Home()

		hpath := strings.Replace(path, "~", home, 1)

		return hpath

	}

	if strings.Contains(path, "$HOME") {

		home := Home()

		hpath := strings.Replace(path, "$HOME", home, 1)

		return hpath

	}

	return path
}
