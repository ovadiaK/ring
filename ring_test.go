package ring

import (
	"fmt"
	"os"
	"runtime/pprof"
	"testing"
)

const (
	testSize   = 200000000
	cpuProfile = "cpu.prof"
	memProfile = "mem.prof"
)

var NoErr = fmt.Errorf("no error")

func expected(exp, got interface{}) string {
	return fmt.Sprintf("expected %s got %s", exp, got)
}
func TestNew(t *testing.T) {
	q, err := New(testSize)
	if err != nil {
		t.Fatal(expected(NoErr, err))
	}
	if q.Capacity() != testSize {
		t.Fatal(expected(testSize, q.Capacity()))
	}
	if q.Size() != 0 {
		t.Fatal(expected(0, q.Size()))
	}
	if !q.Empty() {
		t.Fatal(expected(true, !q.Empty()))
	}
	q, err = New(1)
	if err != ErrShort {
		t.Fatal(expected(ErrShort, err))
	}
}
func TestQueue_In(t *testing.T) {
	q, err := New(testSize)
	if err != nil {
		t.Fatal(expected(NoErr, err))
	}
	for i := 0; i < testSize; i++ {
		_, err := q.Write([]byte{byte(i)})
		if err != nil {
			t.Fatal(expected(NoErr, err))
		}
	}
	if q.Free() != 0 {
		t.Fatal(expected(0, q.Free()))
	}
	_, err = q.Write([]byte{'x'})
	if err != ErrFull {
		t.Fatal(expected(ErrFull, err))
	}
}
func TestQueue_Out(t *testing.T) {
	q, err := New(testSize)
	item := make([]byte, testSize)
	for i := 0; i < testSize; i++ {
		_, err := q.Write([]byte{byte(i)})
		if err != nil {
			t.Fatal(expected(NoErr, err))
		}
		if q.Size() != 1 {
		}
		n, err := q.Read(item)
		if n != 1 {
			t.Fatal(expected(1, n))
		}
		//if item != i {
		//	t.Fatal(expected(i, item))
		//}
	}
	_, err = q.Read(item)
	if err != ErrEmpty {
		t.Fatal(expected(ErrEmpty, err))
	}
}

func BenchmarkQueue_Read(b *testing.B) {
	f, err := os.Create(cpuProfile)
	f2, err := os.Create(memProfile)
	q, err := New(testSize)
	item := make([]byte, 1)
	if err != nil {
		panic(err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	if err := pprof.WriteHeapProfile(f2); err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {

		for j := 0; j < testSize; j++ {
			_, err := q.Write([]byte{byte(i)})
			if err != nil {
				panic(err)
			}
		}
		for j := 0; j < testSize; j++ {
			_, err = q.Read(item)
			if err != nil {
				panic(err)
			}
		}
	}
	pprof.StopCPUProfile()
}
