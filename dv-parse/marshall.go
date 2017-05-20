package main

// I want fast reading from files for the server eventually
// I thought initially it could just be totally preformatted and
// the server could just read/write bytes, but data like
// submission time and tags and stuff need to be template formatted
// so let's just have gob do all the parsing of that data from bytes
// to struct form
import "encoding/gob"

import (
	"fmt"
	"io"
)

func (d *Data) WriteTo(w io.Writer) error {
	return gob.NewEncoder(w).Encode(d)
}

func DataFromReader(r io.Reader) (*Data, error) {
	d := new(Data)
	err := gob.NewDecoder(r).Decode(d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func IndexFromReader(r io.Reader) (*Index, error) {
	data := make([]*IndexData, 0, 4)
	d := gob.NewDecoder(r)
	for {
		id := new(IndexData)
		if err := d.Decode(id); err != nil {
			if err == io.EOF {
				return (*Index)(&data), nil
			} else {
				return nil, fmt.Errorf("gob decode failure on %dth decode: %s", len(data)+1, err)
			}
		}
		data = append(data, id)
	}
}

func (index Index) WriteTo(w io.Writer) error {
	enc := gob.NewEncoder(w)
	for i, id := range index {
		if err := enc.Encode(id); err != nil {
			return fmt.Errorf("gob encode failure on %dth encode", i+1)
		}
	}
	return nil
}
