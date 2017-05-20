package data

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
	"os"
	"path"
)

const (
	INDEX_NAME = "index.gob"
)

func GetIndex(dir string) (*Index, error) {
	f, err := os.Open(path.Join(dir, INDEX_NAME))
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

// Don't include directory
func (id IndexData) GobFileName() string {
	return fmt.Sprintf("%d.gob", id.UID)
}

func (id IndexData) GetFromDirectory(dir string) (*Data, error) {
	f, err := os.Open(path.Join(dir, id.GobFileName()))
	if err != nil {
		return nil, err
	}
	d, err := DataFromReader(f)
	f.Close()
	return d, err
}
