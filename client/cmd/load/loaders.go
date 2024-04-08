package load

import "os"

type Loader interface {
	Load() ([]byte, error)
}

type FileLoader struct {
	Path string `json:"path"`
}

func NewFileLoader(path string) *FileLoader {
	return &FileLoader{Path: path}
}

func (l *FileLoader) Load() ([]byte, error) {
	dat, err := os.ReadFile(l.Path)
	if err != nil {
		return nil, err
	}

	return dat, nil
}
