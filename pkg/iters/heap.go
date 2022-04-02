package iters

func (it *Iters) less(i, j int) bool {
	return it.heap[i].Top() < it.heap[j].Top()
}

func (it *Iters) swap(i, j int) {
	it.heap[i], it.heap[j] = it.heap[j], it.heap[i]
}

func (it *Iters) siftdown(i int) {
	n := len(it.heap)
	for {
		lc := 2*i + 1
		rc := 2*i + 2
		if rc < n && it.less(rc, lc) {
			lc, rc = rc, lc
		}
		if lc >= n || !it.less(lc, i) {
			return
		}
		it.swap(lc, i)
		i = lc
	}
}

func (it *Iters) siftup(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if !it.less(i, p) {
			return
		}
		it.swap(i, p)
		i = p
	}
}

func (it *Iters) remove(i int) {
	j := len(it.heap) - 1
	if i != j {
		it.swap(i, j)
		it.heap = it.heap[:j]
		it.siftdown(i)
	} else {
		it.heap = it.heap[:j]
	}
}
