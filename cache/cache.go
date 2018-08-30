package cache

import (
	"strings"

	fileable "github.com/shotastage/FileableGo"
)

// FileableCache is a struct for creating Cache Storage Object.
type FileableCache struct {
	Objects   []*CacheObj
	Key       string
	IsRevoked bool
}

// Storage is a function of constructor for initializing FileableCache object.
func Storage(key string, obj ...*CacheObj) *FileableCache {
	return &FileableCache{
		Objects:   obj,
		Key:       key,
		IsRevoked: false,
	}
}

// CacheObj is a struct for creating file object for caching.
type CacheObj struct {
	File  string
	ObjId string
}

func Obj(path ...string) *CacheObj {

	joinedPath := strings.Join(path, "/")

	return &CacheObj{
		File:  joinedPath,
		ObjId: "obj_id_generator_is_not_implemented",
	}
}

func (obj *CacheObj) Register() error {
	return nil
}

func (f *FileableCache) Revoke() error {

	f.IsRevoked = true

	return nil
}

func (f FileableCache) Read() (*CacheObj, error) {

	obj := &CacheObj{
		File:  "test_path",
		ObjId: "test_id",
	}

	return obj, nil
}

func (obj CacheObj) Out() (string, error) {

	path := obj.File

	return path, nil
}

func (f *FileableCache) Close() error {

	storedObjects := f.Objects

	for _, obj := range storedObjects {
		if err := fileable.Path(obj.File).Rm(); err != nil {
			return err
		}
	}

	return nil
}
