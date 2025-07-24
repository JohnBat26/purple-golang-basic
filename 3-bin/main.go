package main

import (
	"demo/3-bin/bins"
	"fmt"
)

func main() {
	bin1 := bins.NewBin("1", "bin1", false)
	bin2 := bins.NewBin("2", "bin2", true)

	bins_list := []bins.Bin{bin1, bin2}

	binList := bins.NewBinList(bins_list)

	fmt.Println(binList)

}
