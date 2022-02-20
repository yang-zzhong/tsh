package errors

import (
	"errors"
	"testing"
)

func TestChainedError(t *testing.T) {
	EnableRuntimeInfo()
	e := New("connection reset by peer")
	e1 := e.Cause("send message error")
	e2 := e1.Cause("http post message error")
	e3 := e2.Cause("ios report error")
	e4 := e3.Cause(errors.New("hello world again"))
	em := "ERROR: \"hello world again\" in github.com/yang-zzhong/tsh/errors.TestChainedError\n\tCaused by\n\t  \"ios report error\" in github.com/yang-zzhong/tsh/errors.TestChainedError\n\tCaused by\n\t  \"http post message error\" in github.com/yang-zzhong/tsh/errors.TestChainedError\n\tCaused by\n\t  \"send message error\" in github.com/yang-zzhong/tsh/errors.TestChainedError\n\tCaused by\n\t  \"connection reset by peer\" in github.com/yang-zzhong/tsh/errors.TestChainedError\n"

	if e4.String() != em {
		t.Fatal("something wrong with ChainedError")
	}
}

func testNCE1(err ChainedError) ChainedError {
	return err.Cause("chained error 2")
}

func testNCE2(err ChainedError) ChainedError {
	return err.Cause("chained error 3")
}

func testNCE3(err ChainedError) ChainedError {
	return err.Cause("chained error 4")
}

func TestNestedChainedError(t *testing.T) {
	EnableRuntimeInfo()
	err := testNCE1(testNCE2(testNCE3(New("hello world"))))
	em := "ERROR: \"chained error 2\" in github.com/yang-zzhong/tsh/errors.testNCE1\n\tCaused by\n\t  \"chained error 3\" in github.com/yang-zzhong/tsh/errors.testNCE2\n\tCaused by\n\t  \"chained error 4\" in github.com/yang-zzhong/tsh/errors.testNCE3\n\tCaused by\n\t  \"hello world\" in github.com/yang-zzhong/tsh/errors.TestNestedChainedError\n"
	if em != err.String() {
		t.Fatal("something wrong with ChainedError when nested")
	}
}

func Benchmark_New(b *testing.B) {
	DisableRuntimeInfo()
	for i := 0; i < b.N; i++ {
		New("hello world")
	}
}

func Benchmark_WithRuntimeInfo(b *testing.B) {
	EnableRuntimeInfo()
	for i := 0; i < b.N; i++ {
		New("hello world")
	}
}

func Benchmark_NewChainedError(b *testing.B) {
	DisableRuntimeInfo()
	for i := 0; i < b.N; i++ {
		NewChainedError("hello world", nil)
	}
}

func Benchmark_normalError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		errors.New("hello world")
	}
}
