package array

import (
	"math/rand"
	"strconv"
	"testing"
)

func Test_ArrayContain(t *testing.T) {
	arr1 := make([]string, 10)
	arr2 := make([]string, 10)
	for i := 0; i < 10; i++ {
		val := strconv.Itoa(rand.Intn(100))
		arr1[i] = val
		arr2[i] = val
	}
	if !Contain(arr1, arr2) {
		t.Fatalf("array full contain with order error")
	}
	if !Contain(arr1, arr2[:5]) {
		t.Fatalf("array contain with order error")
	}
	if Contain(arr1, append(arr2[:5], "101")) {
		t.Fatalf("array contain with order error: shouldn't contain")
	}
	if !Contain(arr1, arr1[0]) {
		t.Fatalf("array contain elem error")
	}
	if Contain(arr1, "101") {
		t.Fatalf("array contain elem error: shouldn't contain")
	}
}

func TestArrayEqual(t *testing.T) {
	arr1 := make([]string, 10)
	arr2 := make([]string, 10)
	for i := 0; i < 10; i++ {
		val := strconv.Itoa(rand.Intn(100))
		arr1[i] = val
		arr2[i] = val
	}
	if !Equal(arr1, arr2) {
		t.Fatalf("should equal")
	}
}

func Test_ArraySubFrom(t *testing.T) {
	arr1 := make([]string, 10)
	arr2 := make([]string, 10)
	for i := 0; i < 10; i++ {
		val := strconv.Itoa(rand.Intn(100))
		arr1[i] = val
		arr2[i] = val
	}
	if SubFrom(arr1, arr2) == -1 {
		t.Fatalf("sub from 0 error")
	}
	if SubFrom(arr1, arr2[2:]) != 2 {
		t.Fatalf("sub from 2 error")
	}
	if SubFrom(arr1, append(arr2[2:], "101")) > 0 {
		t.Fatalf("not sub from 2 error")
	}
}

func Benchmark_ArrayContain(b *testing.B) {
	arr1 := make([]string, 100)
	arr2 := make([]string, 100)
	for i := 0; i < 100; i++ {
		val := strconv.Itoa(rand.Intn(100))
		arr1[i] = val
		arr2[i] = val
	}
	for i := 0; i < b.N; i++ {
		Contain(arr1, arr2)
	}
}

func Benchmark_ArraySubFrom(b *testing.B) {
	arr1 := make([]string, 100)
	arr2 := make([]string, 100)
	for i := 0; i < 100; i++ {
		val := strconv.Itoa(rand.Intn(100))
		arr1[i] = val
		arr2[i] = val
	}
	for i := 0; i < b.N; i++ {
		SubFrom(arr1, arr2)
	}
}

func Benchmark_ArrayEqual(b *testing.B) {
	arr1 := make([]string, 100)
	arr2 := make([]string, 100)
	for i := 0; i < 100; i++ {
		val := strconv.Itoa(rand.Intn(100))
		arr1[i] = val
		arr2[i] = val
	}
	for i := 0; i < b.N; i++ {
		Equal(arr1, arr2)
	}
}
