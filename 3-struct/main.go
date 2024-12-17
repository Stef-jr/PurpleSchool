package main

import (
	"bins/bins"
	"bins/files"
	"fmt"
)

func main() {

	binList := bins.NewBinList(files.NewJsonDb("file.json"))
	addbin(binList)

}

func addbin(binList *bins.BinListDb) {

	bin, err := bins.CreateBin("A1", "Первый", false)
	if err != nil {
		fmt.Print(err)
	}

	binList.AddBin(*bin)

}
