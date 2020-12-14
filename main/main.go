package main

import (
	"fmt"
	"github.com/ovadiaK/ring"
)

const size = 10000

func main() {
	var q, err = ring.New(size)
	if err != nil {
		panic(err)
	}

	fmt.Println("capacity:", q.Capacity())
	fmt.Println("free:", q.Free())

	for i := 0; i < size; i++ {
		if _, err := q.Write([]byte{byte(i)}); err != nil {
			panic(err)
		}
	}

	fmt.Println("size:", q.Size())
	fmt.Println("free:", q.Free())
	buf := make([]byte, size)

	if _, err := q.Read(buf); err != nil {
		panic(err)
	}

	fmt.Println("size:", q.Size())
	fmt.Println("free:", q.Free())
	fmt.Println("empty?:", q.Empty())
}
