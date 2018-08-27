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
	f := fileable.Path("")

	t.Log(f.Pwd())
	f.Mkdir("")
}
