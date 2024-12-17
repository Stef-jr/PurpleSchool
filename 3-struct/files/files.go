package files

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		fileName: name,
	}
}

func (db JsonDb) Read() ([]byte, error) {

	if !verifyExist(db.fileName) {
		err := errors.New("file not exist")
		return nil, err
	}

	if !verifyExtension(db.fileName) {
		err := errors.New("file format is not 'json'")
		return nil, err
	}

	data, err := os.ReadFile(db.fileName)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db JsonDb) Write(content []byte) error {
	file, err := os.Create(db.fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func verifyExtension(path string) bool {
	fileExtens := filepath.Ext(path)
	return strings.TrimLeft(fileExtens, ".") == "json"
}

func verifyExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}