package devblog

// I want fast reading from files for the server eventually
// I thought initially it could just be totally preformatted and
// the server could just read/write bytes, but data like
// submission time and tags and stuff need to be template formatted
// so let's just have gob do all the parsing of that data from bytes
// to struct form
import "encoding/gob"

// Okay so what this is doing is making a terrible, terrible
// database.  I realize that.  It _works_ and it's not hard
// to replace with a less terrible (i.e. actual) database later.
//
// Possible problems include:
//  * Each post on a page requires a filesystem
//    call.  5 posts on a page, 5 file reads!
//  * An index that is searched just by iterating through an array
//      this means finding all post with a tag means searching ALL
//      (well, most) TAGS OF ALL POSTS
import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	INDEX_NAME = "index.gob"
)

func GetIndex(dir string) (*Index, error) {
	f, err := os.Open(filepath.Join(dir, INDEX_NAME))
	if err != nil {
		if os.IsNotExist(err) {
			return NewIndex(), nil
		} else {
			return nil, fmt.Errorf("failed to read index file: %v", err)
		}
	}
	defer f.Close()
	index, err := IndexFromReader(f)
	if err != nil {
		return nil, err
	}
	return index, nil
}

func MakeIndex(dir string, index *Index) error {
	f, err := os.Create(filepath.Join(dir, INDEX_NAME))
	if err != nil {
		return fmt.Errorf("failed to create index file: %v", err)
	}
	defer f.Close()
	return index.WriteTo(f)
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
