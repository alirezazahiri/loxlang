package file

import (
	"os"
)

func (f *File) Read() ([]byte, error) {
	data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
