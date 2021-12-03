package paasio

import (
	"io"
	"sync"
)

// Define one/more types which implements ReadCounter, WriteCounter and ReadWriteCounter
type ReadWriteCounterImpl struct {
	r      io.Reader
	w      io.Writer
	rBytes int64
	rOps   int
	wBytes int64
	wOps   int
	mut    sync.RWMutex
}

func (rwc *ReadWriteCounterImpl) Read(p []byte) (n int, err error) {
	n, err = rwc.r.Read(p)

	rwc.mut.Lock()
	defer rwc.mut.Unlock()
	rwc.rBytes += int64(n)
	rwc.rOps += 1

	return
}

func (rwc *ReadWriteCounterImpl) ReadCount() (int64, int) {
	rwc.mut.RLock()
	defer rwc.mut.RUnlock()
	return rwc.rBytes, rwc.rOps
}

func (rwc *ReadWriteCounterImpl) Write(p []byte) (n int, err error) {
	n, err = rwc.w.Write(p)

	rwc.mut.Lock()
	defer rwc.mut.Unlock()
	rwc.wBytes += int64(n)
	rwc.wOps += 1

	return
}

func (rwc *ReadWriteCounterImpl) WriteCount() (int64, int) {
	rwc.mut.RLock()
	defer rwc.mut.RUnlock()
	return rwc.wBytes, rwc.wOps
}

// NewReadCounter returns a type which implements ReadCounter
func NewReadCounter(r io.Reader) ReadCounter {
	return &ReadWriteCounterImpl{r: r}
}

// NewWriteCounter returns a type which implements WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	return &ReadWriteCounterImpl{w: w}
}

// NewReadWriteCounter returns a type which implements ReadWriteCounter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &ReadWriteCounterImpl{r: rw, w: rw}
}
