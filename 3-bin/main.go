package main

import (
	"fmt"
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

func main() {
	bin1 := NewBin("1", "bin1", false)
	bin2 := NewBin("2", "bin2", true)

	bins := []Bin{bin1, bin2}

	binList := NewBinList(bins)

	fmt.Println(binList)

}
