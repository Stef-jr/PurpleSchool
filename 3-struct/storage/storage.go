package storage

import "bins/files"

//реализовать сохранение bin в виде json в локальном файле
//реализовать чтение списка bin в виде json из локального файла

func WriteBins(content []byte, name string ) error {
	
	err := files.WriteFile(content, name)

	if err != nil {
		return err
	}

	return nil
}

func ReadBins(path string) ([]byte, error) {
	
	data, err := files.ReadFile(path)

	if err != nil {
		return nil, err
	}
	return data, nil

}

