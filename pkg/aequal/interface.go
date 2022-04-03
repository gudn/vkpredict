package aequal

type Interface interface {
	Equal(i, j int) bool
	Len() (int, int)
}

type Strings struct {
	a, b string
}

func (s *Strings) Len() (int, int) {
	return len(s.a), len(s.b)
}

func (s *Strings) Equal(i, j int) bool {
	return s.a[i] == s.b[j]
}

func NewStrings(a, b string) *Strings {
	return &Strings{a, b}
}
