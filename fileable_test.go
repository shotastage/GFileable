package fileable_test

import (
	"testing"

	fileable "github.com/shotastage/FileableGo"
)

/*
func TestCaseIsDir(t *testing.T) {

	fileable.Path("/home/")
}
*/

func TestCasePwd(t *testing.T) {
	t.Log("Current directory ==>> ", fileable.Pwd())
}

func TestHome(t *testing.T) {
	t.Log("Home directory ==>> ", fileable.Home())
}

func TestIsDir(t *testing.T) {

	home := fileable.Home()

	f := fileable.Path(fileable.Home())

	if f.IsDir() {
		t.Log(home, "is a directory.")
	} else {
		t.Log(home, "is not a directory.")
	}
}

func TestIsFile(t *testing.T) {
	profile := fileable.Home() + "/.bash_profile"

	f := fileable.Path(profile)

	if f.IsFile() {
		t.Log(profile, "is a file.")
	} else {
		t.Log(profile, "is not a file.")
	}
}
