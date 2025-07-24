package bins

import (
	"time"
)

type Bin struct {
	id        string
	name      string
	private   bool
	createdAt time.Time
}

type BinList struct {
	bins [](Bin)
}

func NewBin(id, name string, private bool) Bin {
	return Bin{
		id:        id,
		name:      name,
		private:   private,
		createdAt: time.Now(),
	}
}

func NewBinList(bins [](Bin)) BinList {
	return BinList{
		bins,
	}
}
