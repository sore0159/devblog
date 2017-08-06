package generate

import (
	"fmt"
	"sort"
)

type ProcessedTaglist struct {
	FileName string
	Title    string
	Tag      string
	Latest   string
	Posts    []*ProcessedFile
}

func ProcessTaglist(tg string, list []*ProcessedFile) *ProcessedTaglist {
	if len(list) == 0 {
		return nil
	}
	SortByDate(list)
	pTL := new(ProcessedTaglist)
	pTL.FileName = TagLink(tg)
	pTL.Title = fmt.Sprintf("Archive: %s", tg)
	pTL.Tag = tg
	pTL.Latest = list[0].PubTime.Format("Jan 2, 2006 (15:04)")
	pTL.Posts = list
	return pTL
}

func SortByDate(list []*ProcessedFile) {
	sort.SliceStable(list, func(i, j int) bool {
		return list[j].PubTime.Before(list[i].PubTime)
	})
}
