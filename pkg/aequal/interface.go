package aequal

// Интерфейс для сравнения
type Interface interface {
	// Насколько элементы похожи, должен возвращать число в диапазоне  [0; 1]
	// `i` номер в первой последоватльности, `j` во второй
	Equal(i, j int) float64
	// Длины каждой из двух сравниваемых последовательностей
	Len() (int, int)
}

// Реализация интерфейса для двух строк
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
