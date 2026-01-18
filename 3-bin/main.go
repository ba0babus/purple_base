package main

import (
	"encoding/json"
	"fmt"
	"purple_base/bin/bins"
	"purple_base/bin/storage"
)

func main() {
	db := storage.NewFileDb("bin_list.json")
	someBin := bins.NewBin("123", true, "myNewBin", db)

	if err := someBin.Save(); err != nil {
		fmt.Println("Save error:", err)
		return
	}

	if err := someBin.Load(); err != nil {
		fmt.Println("Load error:", err)
		return
	}

	data, err := json.Marshal(someBin.Bin)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println(string(data))
}
