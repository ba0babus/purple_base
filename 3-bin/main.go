package main

import (
	"fmt"
	"purple_base/bin/bins"
)

func main() {
	filename := "bin_list.json"
	someBin := bins.NewBin("123", true, "myNewBin")

	someBin.SaveBin(filename)

	myBin, err := someBin.ReadBin(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myBin)
	}
}
