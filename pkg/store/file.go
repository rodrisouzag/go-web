package store

import (
	"encoding/json"
	"errors"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "filestorage"
)

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fs.FileName, fileData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) { // si el error NO fue porque el archivo de storage no existe, retorno el error
			return err
		}
		file = []byte("[]") // inicializo un contenido vacio para realizar el unmarshall
	}
	return json.Unmarshal(file, &data)
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName}
	}
	return nil
}

type FileStore struct {
	FileName string
}
