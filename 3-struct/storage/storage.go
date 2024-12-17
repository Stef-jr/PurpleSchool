package storage

type Storage struct {
	path string
}

func NewStorage(path string) *Storage {
	return &Storage{
		path: path,
	}
}

func Write(content []byte) error {
	return nil
}

func Read() ([]byte, error) {

	return []byte{}, nil

}