package circular

import "errors"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
	buffer         []byte
	newestPosition int
	oldestPosition int
}

var (
	errReadFromEmptyBuffer = errors.New("Cannot read from empty buffer")
	errWriteToFullBuffer   = errors.New("Cannot write to a full buffer")
)

func NewBuffer(size int) *Buffer {
	return &Buffer{
		buffer:         make([]byte, size),
		newestPosition: -1,
		oldestPosition: -1,
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	panic("Please implement the ReadByte function")
}

func (b *Buffer) WriteByte(c byte) error {
	panic("Please implement the WriteByte function")
}

func (b *Buffer) Overwrite(c byte) {
	panic("Please implement the Overwrite function")
}

func (b *Buffer) Reset() {
	panic("Please implement the Reset function")
}
