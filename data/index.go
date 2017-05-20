package data

import "time"

type IndexData struct {
	FileName  string
	UID       uint64
	Submitted time.Time
	Tags      []string
}

type Data struct {
	IndexData
	Content []byte
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

func (ix *Index) AddData(d *Data) {
	d.UID = ix.NewUID()
	id := new(IndexData)
	*id = d.IndexData
	*ix = append(*ix, id)
}
