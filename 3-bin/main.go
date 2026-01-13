package main

import (
	"fmt"
	"purple_base/bin/bins"
	"purple_base/bin/storage"
)

func main() {
	filename := "bin_list.json"
	someBin := bins.NewBin("123", true, "myNewBin")

	storage.SaveBin(someBin, filename)

	myBin, err := storage.ReadBin(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myBin)
	}
}
