package bins

import (
	"errors"
	"time"
)

type Bin struct{

	Id string `json:"id"`
	Private bool `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name string `json:"name"`

}

func CreateBin(id, name string, private bool) (*Bin, error){

	if id == "" {
		return nil, errors.New("EMPTY_ID")
	}

	if name == "" {
		return nil, errors.New("EMPTY_NAME")
	}
	
	newBin := &Bin{
		Id: id,
		Private: private,
		CreatedAt: time.Now(),
		Name: name,		
	}

	return newBin, nil
}