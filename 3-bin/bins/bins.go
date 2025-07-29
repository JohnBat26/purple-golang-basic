package bins

import (
	"time"
)

type Bin struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
}

type BinList struct {
	Bins [](Bin) `json:"bins"`
}

func NewBin(id, name string, private bool) Bin {
	return Bin{
		ID:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}
}

func NewBinList(bins [](Bin)) BinList {
	return BinList{
		bins,
	}
}
