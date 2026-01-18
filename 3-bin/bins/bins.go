package bins

import (
	"time"
)

type Db interface {
	Read(string) ([]byte, error)
	Write(string) (bool, error)
}

type Bin struct { // структура
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Db        Db
}

func NewBin(id string, private bool, name string) *Bin {
	newAcc := &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}

	return newAcc
}
