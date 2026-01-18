package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"purple_base/bin/bins"
	"purple_base/bin/file"
)

func Write(content *bins.Bin, filename string) (bool, error) {
	var isSaved bool
	data, _ := json.Marshal(content)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("UNABLE_TO_OPEN_FILE")
	}
	defer file.Close()
	_, err = file.WriteString(string(data))
	if err != nil {
		fmt.Println(err)
		return false, errors.New("UNABLE_TO_WRITE_FILE")
	}
	file.Close()
	isSaved = true
	return isSaved, nil
}

func Read(filename string) ([]byte, error) {
	content, err := file.Read(filename)
	if err != nil {
		return nil, errors.New("UNABLE_TO_READ_BIN")
	}
	return content, nil
}
