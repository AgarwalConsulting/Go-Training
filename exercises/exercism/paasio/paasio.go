package paasio

import (
	"io"
)

// Define one/more types which implements ReadCounter, WriteCounter and ReadWriteCounter

// NewReadWriteCounter returns a type which implements ReadWriteCounter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return nil
}

// NewReadCounter returns a type which implements ReadCounter
func NewReadCounter(r io.Reader) ReadCounter {
	return nil
}

// NewWriteCounter returns a type which implements WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	return nil
}
