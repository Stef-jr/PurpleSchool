package bins

import (
	"encoding/json"
	"fmt"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte) error
}
type BinList struct {
	Bins      []Bin     `json:"bins"`
	CreatedAt time.Time `json:"createdAt"`
}

type BinListDb struct {
	BinList
	db Db
}

func NewBinList(db Db) *BinListDb {

	file, err := db.Read()

	if err != nil {
		return &BinListDb{
			BinList: BinList{
				Bins:      []Bin{},
				CreatedAt: time.Now()},
			db: db,
		}
	}

	var binList BinList
	err = json.Unmarshal(file, &binList)
	if err != nil {
		fmt.Println(err)
	}

	return &BinListDb{
		BinList: binList,
		db:      db,
	}
}

func (binList *BinListDb) AddBin(bin Bin) {

	binList.Bins = append(binList.Bins, bin)
	binList.saveBinList()

}

func (binList *BinListDb) saveBinList() {

	binList.CreatedAt = time.Now()
	data, err := json.Marshal(binList)
	if err != nil {
		fmt.Println("Ошибка преобразования json \n" + err.Error())
	}

	err = binList.db.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи данных \n" + err.Error())
	}

}