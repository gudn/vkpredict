package vkpredict

import "sort"

type Entry struct {
	Value string
	Score int
}

type EntryList []Entry

// Implementation of sort.Interface
func (l EntryList) Len() int {
	return len(l)
}

func (l EntryList) Less(i, j int) bool {
	// Default reverse ordering
	return l[i].Score > l[j].Score
}

func (l EntryList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l EntryList) Top(n int) EntryList {
	sort.Sort(l)
	return EntryList(l[:n:n])
}
