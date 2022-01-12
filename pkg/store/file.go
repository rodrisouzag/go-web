package store

import (
	"encoding/json"
	"errors"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
	AddMock(mock *Mock)
	ClearMock()
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
	if fs.Mock != nil {
		fs.Mock.Data = fileData
		return nil
	}
	err = os.WriteFile(fs.FileName, fileData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileStore) Read(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		fs.Mock.ReadUsed = true
		return json.Unmarshal(fs.Mock.Data, data)
	}
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
		return &FileStore{fileName, nil}
	}
	return nil
}

type FileStore struct {
	FileName string
	Mock     *Mock
}

type Mock struct {
	Data     []byte
	Err      error
	ReadUsed bool
}

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}
func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}
