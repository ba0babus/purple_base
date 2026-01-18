package main

import (
	"fmt"
	"purple_base/bin/bins"
	"purple_base/bin/storage"
)

func main() {
	filename := "bin_list.json"
	someBin := bins.NewBin("123", true, "myNewBin")

	storage.Write(someBin, filename)

	myBin, err := storage.Read(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(myBin))
	}
}
