package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	reader         io.Reader
	byteCount      int64
	operationCount int
	mutex          *sync.Mutex
}

type writeCounter struct {
	writer         io.Writer
	byteCount      int64
	operationCount int
}

type readWriteCounter struct {
	reader ReadCounter
	writer WriteCounter
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		reader: reader,
		mutex:  new(sync.Mutex),
	}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		reader: NewReadCounter(readwriter),
		writer: NewWriteCounter(readwriter),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()
	m, err := rc.reader.Read(p)
	rc.byteCount += int64(m)
	rc.operationCount++
	return m, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()
	return rc.byteCount, rc.operationCount
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	m, err := wc.writer.Write(p)
	wc.byteCount += int64(m)
	wc.operationCount++
	return m, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return wc.byteCount, wc.operationCount
}

func (rwc *readWriteCounter) Read(p []byte) (int, error) {
	return rwc.reader.Read(p)
}

func (rwc *readWriteCounter) ReadCount() (int64, int) {
	return rwc.reader.ReadCount()
}

func (rwc *readWriteCounter) Write(p []byte) (int, error) {
	return rwc.writer.Write(p)
}

func (rwc *readWriteCounter) WriteCount() (int64, int) {
	return rwc.writer.WriteCount()
}
