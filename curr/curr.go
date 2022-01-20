package psh

import "sync"

// CurrCall concurrent handle total, each concurrence handle ps unit size
func CurrCall(total, ps int, call func(start, size int)) {
	var wg sync.WaitGroup
	LineCall(total, ps, func(start, size int) {
		wg.Add(1)
		go func() {
			call(start, size)
			wg.Done()
		}()
	})
	wg.Wait()
}

// Curr calculate each concurrent unit size
// if we wanna handle total 100, max 10 concurrent to handle the total
// then each concurrent will take 10 as the function returned
func Curr(total, maxC int) int {
	if total > maxC {
		if (total % maxC) > 0 {
			return total/maxC + 1
		}
		return total / maxC
	}
	return 1
}

func Call(total, maxC int, call func(start, size int)) {
	CurrCall(total, Curr(total, maxC), call)
}

func LineCall(total, ps int, call func(start, size int)) {
	for i := 0; i < total; i += ps {
		size := ps
		if size > total-i {
			size = total - i
		}
		call(i, size)
	}
}
