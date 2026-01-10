package main

import (
	"fmt"
	"time"
)

type bin struct { // структура
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id string, private bool, name string) *bin {
	newAcc := &bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}

	return newAcc
}

func main() {
	var BinList []bin
	someBin := newBin("123", true, "myNewBin")
	fmt.Println(BinList, someBin)
}
