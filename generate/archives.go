package generate

import (
	"fmt"
	"sort"
)

type ProcessedTaglist struct {
	FileName  string
	Title     string
	Tag       string
	Latest    string
	SplitList []*ListPart
}

type ListPart struct {
	Title string
	Posts []*ProcessedFile
}

func MakeSplitList(list []*ProcessedFile) []*ListPart {
	rlist := make([]*ListPart, 0)
	var title string
	var lp *ListPart
	for _, pf := range list {
		if pf.NoPost {
			continue
		}
		newTitle := pf.PubTime.Format("Jan 2006")
		if newTitle != title {
			lp = new(ListPart)
			rlist = append(rlist, lp)
			lp.Title = newTitle
			title = newTitle
		}
		lp.Posts = append(lp.Posts, pf)
	}
	return rlist
}

type TagNav struct {
	Tag                       [2]string
	Has                       [2]bool
	FirstFN, FirstT, FirstPub string
	PrevFN, PrevT, PrevPub    string
	NextFN, NextT, NextPub    string
	LastFN, LastT, LastPub    string
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
	pTL.SplitList = MakeSplitList(list)
	return pTL
}

// SortByDate sorts from newest to oldest
func SortByDate(list []*ProcessedFile) {
	sort.SliceStable(list, func(i, j int) bool {
		return list[j].PubTime.Before(list[i].PubTime)
	})
}

func AddTagNavs(tg string, list []*ProcessedFile) {
	tgFile := TagLink(tg)
	tn := new(TagNav)
	tn.Tag[0] = tg
	if tg != "" {
		tn.Tag[1] = tgFile
	}
	if len(list) < 2 {
		if len(list) == 1 {
			list[0].TagNavs = append(list[0].TagNavs, *tn)
		}
		return
	}
	tn.LastFN, tn.LastT, tn.LastPub = list[0].FileName, list[0].Title, list[0].Published
	l := len(list) - 1
	tn.FirstFN, tn.FirstT, tn.FirstPub = list[l].FileName, list[l].Title, list[l].Published
	for i, pf := range list {
		tn.Has = [2]bool{true, true}
		p, n := i+1, i-1
		if n < 0 {
			n = 0
			tn.Has[1] = false
		}
		if p > l {
			p = l
			tn.Has[0] = false
		}
		tn.PrevFN, tn.PrevT, tn.PrevPub = list[p].FileName, list[p].Title, list[p].Published
		tn.NextFN, tn.NextT, tn.NextPub = list[n].FileName, list[n].Title, list[n].Published
		pf.TagNavs = append(pf.TagNavs, *tn)
	}
}
