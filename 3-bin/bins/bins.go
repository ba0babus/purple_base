package bins

import "time"

type bin struct { // структура
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBin(id string, private bool, name string) *bin {
	newAcc := &bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}

	return newAcc
}
