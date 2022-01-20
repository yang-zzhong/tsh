package psh

import (
	"fmt"
	"testing"
)

func TestCall(t *testing.T) {
	Call(100, 9, func(start, size int) {
		fmt.Printf("call: %d - %d\n", start, size)
	})
}

func TestPlan(t *testing.T) {
	fmt.Printf("total: 10, max: 11, curr: %d\n", Curr(10, 11))
	fmt.Printf("total: 10, max: 10, curr: %d\n", Curr(10, 10))
	fmt.Printf("total: 100, max: 11, curr: %d\n", Curr(100, 11))
}
