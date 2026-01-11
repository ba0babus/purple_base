package main

import (
	"fmt"
	"purple_base/bin/bins"
)

func main() {
	someBin := bins.NewBin("123", true, "myNewBin")
	fmt.Println(someBin)
}
