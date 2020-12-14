# ring
Fixed length circular buffer (first-in first-out) in go with 
constant time (O(1)) read & write and minimum memory consumption.
Implements ReaderWriter interface.

## how it works
the data structure has pointers to start and end and fills the array circularly while checking to not overwrite the start with the end

## usage

### api
```go
circle, error := ring.New(int)
```
initiate a new queue with length int (currently not resizing, so choose wisely), returns ``ErrShort`` if ``int<=1``
```go
n, error = circle.Write([]byte{1,2,3})
```
write elements to end of the buffer, returns ``ErrFull`` if max capacity reached
```go
n, error = circle.Read([]byte{})
```
read element from start of the buffer onto another array and remove them from the start of the buffer, returns ``ErrEmpty`` if empty
#### status
```go
int = circle.Capacity()
```
max number of elements the buffer can take
```go
int = circle.Size()
```
count of elements currently in buffer
```go
int = circle.Free()
```
count of free slots in buffer
```go
bool = circle.Full()
```
returns true if no slot is free

## benchmarks
constant time for all input and output operations

```go
for j := 0; j < testSize; j++ {
			err := q.In(i)
			if err != nil {
				panic(err)
			}
		}
		for j := 0; j < testSize; j++ {
			_, err = q.Out()
			if err != nil {
				panic(err)
			}
		}
```

```
testSize =       100          1946 ns/op
testSize =     10000        198770 ns/op
testSize =   1000000      20894134 ns/op
testSize = 100000000    1838633177 ns/op
```

## example
as found in the main folder
needs update
## to do
make the array grow and shrink dynamically without breaking the constant time read / write