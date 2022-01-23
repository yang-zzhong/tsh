package curr

import (
	"sync"
	"testing"
)

func TestCall(t *testing.T) {
	should := map[int]int{
		24: 12,
		60: 12,
		48: 12,
		0:  12,
		12: 12,
		84: 12,
		36: 12,
		96: 4,
		72: 12}
	var mx sync.Mutex
	Call(100, 9, func(start, size int) {
		mx.Lock()
		defer mx.Unlock()
		if v, ok := should[start]; !ok {
			t.Fatalf("start should has %d but not", start)
		} else if v != size {
			t.Fatalf("size of %d should be %d but not", start, size)
		} else {
			delete(should, start)
		}
	})
	if len(should) > 0 {
		t.Fatalf("something wrong with call")
	}
}

type mmmm struct {
	total, maxC int
}

func TestCurr(t *testing.T) {
	should := map[mmmm]int{
		{10, 11}:  1,
		{10, 10}:  1,
		{100, 11}: 10,
	}
	for k, v := range should {
		if Curr(k.total, k.maxC) != v {
			t.Fatalf("the concurrence unit size should be %d when total %d, maxC %d", v, k.total, k.maxC)
		}
	}
}
