package storage

import (
	"fmt"
	"os"
)

type FileDb struct {
	filename string
}

func NewFileDb(filename string) *FileDb {
	return &FileDb{filename: filename}
}

func (f *FileDb) Read() ([]byte, error) {
	data, err := os.ReadFile(f.filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read file %s: %w", f.filename, err)
	}
	return data, nil
}

func (f *FileDb) Write(content []byte) (bool, error) {
	file, err := os.OpenFile(f.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return false, fmt.Errorf("unable to open file %s for writing: %w", f.filename, err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return false, fmt.Errorf("unable to write to file %s: %w", f.filename, err)
	}
	return true, nil
}
