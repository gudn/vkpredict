package pfunc

func Pfunc(s string) []int {
	n := len(s)
	p := make([]int, n+1)
	p[0] = -1
	for i := 0; i < n; i++ {
		k := p[i]
		for k != -1 && s[k] != s[i] {
			k = p[k]
		}
		p[i+1] = k + 1
	}
	return p
}

func MaxPfunc(s string) int {
	p := Pfunc(s)
	mx := -1
	for _, v := range p {
		if v > mx {
			mx = v
		}
	}
	return mx
}
