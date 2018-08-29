package cache

import (
	"strings"

	fileable "github.com/shotastage/FileableGo"
)

type FileableCache struct {
	Objects   []*CacheObj
	Key       string
	IsRevoked bool
}

type CacheObj struct {
	File  string
	ObjId string
}

func Storage(key string, obj ...*CacheObj) *FileableCache {
	return &FileableCache{
		Objects:   obj,
		Key:       key,
		IsRevoked: false,
	}
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
