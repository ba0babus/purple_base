package bins

import (
	"encoding/json"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write(content []byte) (bool, error)
}

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinWithDb struct {
	Bin
	db Db
}

func (b *BinWithDb) Save() error {
	data, err := json.Marshal(b.Bin)
	if err != nil {
		return err
	}
	_, err = b.db.Write(data)
	return err
}

func (b *BinWithDb) Load() error {
	data, err := b.db.Read()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &b.Bin)
}

func NewBin(id string, private bool, name string, db Db) *BinWithDb {
	return &BinWithDb{
		Bin: Bin{
			Id:        id,
			Private:   private,
			CreatedAt: time.Now(),
			Name:      name,
		},
		db: db,
	}
}
