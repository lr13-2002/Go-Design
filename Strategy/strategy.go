package strategy

import (
	"fmt"
	"io/ioutil"
	"os"
)

var strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

type StorageStrategy interface {
	Save(name string, data []byte) error
}

func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := strategys[t]

	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}

	return s, nil
}

type fileStorage struct{}

func (fileStorage) Save(name string, data []byte) error {
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

type encryptFileStorage struct{}

func (encryptFileStorage) Save(name string, data []byte) error {
	data, err := encrypt(data)
	if err != nil {
		return nil
	}

	return ioutil.WriteFile(name, data, os.ModeAppend)
}

func encrypt(data []byte) ([]byte, error) {
	return data, nil
}
