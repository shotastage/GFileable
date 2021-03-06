package GFileable

import "os"

// Write string to file
func (f Fileable) WriteString(str string) error {
	file, err := os.OpenFile(f.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	data := []byte(str)
	_, err = file.Write(data)

	defer file.Close()

	return err
}

// Check files existence
func (f Fileable) Check(files ...string) bool {

	for f := range files {
		_, err := os.Stat(files[f])
		if err != nil {
			return false
		}
	}

	return true
}
