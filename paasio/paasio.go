package paasio

import "io"

// Define readCounter and writeCounter types here.
type readCounter struct {
	reader         io.Reader
	byteCount      int64
	operationCount int
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
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	panic("Please implement the NewReadWriteCounter function")
}

func (rc *readCounter) Read(p []byte) (int, error) {
	m, err := rc.reader.Read(p)
	rc.byteCount += int64(m)
	rc.operationCount++
	return m, err
}

func (rc *readCounter) ReadCount() (int64, int) {
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
