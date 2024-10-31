package bins

import (
	"bins/files"
	"bins/storage"
	"encoding/json"
	"fmt"
	"time"
)

type BinList struct{
	
	Bins []Bin `json:"bins"`
	CreatedAt time.Time `json:"createdAt"`
	
}

func NewBins() *BinList {
	
	file, err := files.ReadFile("data.json")

	if err != nil {
		return &BinList{
			Bins: []Bin{},
			CreatedAt: time.Now(),
		}
	}

	var binList BinList
	err = json.Unmarshal(file, &binList)
	if err != nil {
		fmt.Println(err)
	}

	return &binList
}

func (binList *BinList) AddBin(bin Bin) {

	binList.Bins = append(binList.Bins, bin)
	//binList.CreatedAt = time.Now()
	binList.saveBinList()
	
}

func (binList *BinList) saveBinList() {
	binList.CreatedAt = time.Now()
	data, err := json.Marshal(binList)
	if err != nil {
		fmt.Println("ошибка преобразования json")
	}	
	err = storage.WriteBins(data, "data.json")
	if err != nil {
		fmt.Println("ошибка записи")
	}
}