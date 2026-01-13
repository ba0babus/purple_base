package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(filename string) ([]byte, error) {
	if filepath.Ext(filename) != ".json" {
		return nil, errors.New("FILE_EXTENSION_IS_NOT_JSON")
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
