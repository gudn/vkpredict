package aequal

type Interface interface {
	Equal(i, j int) float64
	Len() (int, int)
}

type Strings struct {
	a, b string
}

func (s *Strings) Len() (int, int) {
	return len(s.a), len(s.b)
}

func (s *Strings) Equal(i, j int) float64 {
	if s.a[i] == s.b[j] {
		return 1
	}
	return 0
}

func NewStrings(a, b string) *Strings {
	return &Strings{a, b}
}
