package cache_test

import (
	"testing"

	fileable "github.com/shotastage/FileableGo"
	"github.com/shotastage/FileableGo/cache"
)

func TestCreateFileableCacheStorage(t *testing.T) {

	dirObj1 := cache.Obj(fileable.Home(), ".bash_history")

	dirObj2 := cache.Obj(fileable.Home(), ".bash_profile")

	storage := cache.Storage("testObject", dirObj1, dirObj2)

	defer storage.Close()

	t.Log("Key ID:", storage.Key)
	t.Log("Storage Status:", storage.IsRevoked)

	t.Log("======== Containing Objects ========")

	for i, obj := range storage.Objects {
		t.Log("Found Object", i)
		t.Log("Object ID:", obj.ObjId)
		t.Log("Object Path:", obj.File)
	}
}
