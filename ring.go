package ring

import "fmt"

var (
	ErrShort = fmt.Errorf("ring2 needs to be longer than 1")
	ErrFull  = fmt.Errorf("ring2 full")
	ErrEmpty = fmt.Errorf("ring2 empty")
)

type Ring struct {
	queue     []byte
	maxlength int
	start     int
	end       int
	free      int
}

func New(length int) (Ring, error) {
	if length <= 1 {
		return Ring{}, ErrShort
	}
	return Ring{
		queue:     make([]byte, length),
		maxlength: length,
		start:     0,
		end:       0,
		free:      length,
	}, nil
}
func (q *Ring) Write(input []byte) (int, error) {
	inputSize := len(input)
	if q.free < inputSize {
		return 0, ErrFull // todo make it grow
	}
	for i, x := range input {
		if q.free == 0 {
			return i, ErrFull
		}
		q.queue[q.end] = x
		q.end++
		q.free--
		if q.end >= q.maxlength {
			q.end = 0
		}
	}
	return inputSize, nil
}
func (q *Ring) Read(output []byte) (interface{}, error) {
	var i int
	for i, _ = range output {
		if q.free >= q.maxlength {
			return i, ErrEmpty // todo make it shrink
		}
		output[i] = q.queue[q.start]
		q.start++
		q.free++
		if q.start >= q.maxlength {
			q.start = 0
		}
	}
	return i, nil
}
func (q Ring) Size() int {
	return q.maxlength - q.free
}
func (q Ring) Empty() bool {
	return q.free == q.maxlength
}
func (q Ring) Capacity() int {
	return q.maxlength
}
func (q Ring) Free() int {
	return q.free
}
