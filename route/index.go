package route

import (
	"fmt"
	"time"
)

type IndexData struct {
	UID       uint64
	FileName  string
	Title     string
	Submitted time.Time
	Tags      []string
	//? Summary string
}

type Index []*IndexData

func NewIndex() *Index {
	ind := make([]*IndexData, 0)
	return (*Index)(&ind)
}

func (ix Index) NewUID() uint64 {
	var uid uint64 = 1
	for _, id := range ix {
		if uid <= id.UID {
			uid = id.UID + 1
		}
	}
	return uid
}

func (ix *Index) AddData(d *IndexData) error {
	if d.UID == 0 {
		d.UID = ix.NewUID()
	} else {
		for _, existing := range *ix {
			if existing.UID == d.UID {
				return fmt.Errorf("duplicate uid %d for data %s and %s", existing.UID, existing.FileName, d.FileName)
			}
		}
	}
	id := new(IndexData)
	*id = *d
	*ix = append(*ix, id)
	return nil
}
