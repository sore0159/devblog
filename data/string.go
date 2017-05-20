package data

import "fmt"

func (d *Data) String() string {
	return fmt.Sprintf("%s Content: %s",
		d.IndexData.String(), string(d.Content),
	)
}

func (id *IndexData) String() string {
	return fmt.Sprintf("UID: %d File:%s Time:%s Tags: %v",
		id.UID, id.FileName, id.Submitted, id.Tags,
	)
}
