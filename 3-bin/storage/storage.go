package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"purple_base/bin/bins"
	"purple_base/bin/file"
)

func SaveBin(content *bins.Bin, filename string) (bool, error) {
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

func ReadBin(filename string) (string, error) {
	content, err := file.ReadFile(filename)
	if err != nil {
		return "", errors.New("UNABLE_TO_READ_BIN")
	}
	return string(content), nil
}
