package files

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

//реализовать чтение любого файла
//реализовать проверку что это json расширение файла

func ReadFile(path string) ([]byte, error){

	if !verifyExtension(path) {
		err := errors.New("file format is not 'json'")		
		return nil, err
	}

	data, err := os.ReadFile(path)
	
	if err != nil {
		return nil, err
	}
	
	return data, nil
}

func verifyExtension (path string) bool{
	fileExtens := filepath.Ext(path)
	if strings.TrimLeft(fileExtens, ".") != "json"{
		return false
	}
	return true
}

func WriteFile(content []byte, name string) error {
	file, err := os.Create(name)
	
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