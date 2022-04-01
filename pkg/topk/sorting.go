package topk

type SortEntry []*Entry

func (s SortEntry) Len() int {
	return len(s)
}

func (s SortEntry) Less(i, j int) bool {
	// reverse order
	return s[i].Score > s[j].Score
}

func (s SortEntry) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
