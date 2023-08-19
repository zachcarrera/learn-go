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
	if b.newestPosition == -1 && b.oldestPosition == -1 {
		return 0, errReadFromEmptyBuffer
	}
	readByte := b.buffer[b.oldestPosition]
	b.buffer[b.oldestPosition] = 0
	if b.oldestPosition == b.newestPosition {
		b.Reset()
	}

	b.oldestPosition = (b.oldestPosition + 1) % len(b.buffer)

	if b.oldestPosition > b.newestPosition {
		b.Reset()
	}
	return readByte, nil
}

func (b *Buffer) WriteByte(c byte) error {
	panic("Please implement the WriteByte function")
}

func (b *Buffer) Overwrite(c byte) {
	panic("Please implement the Overwrite function")
}

func (b *Buffer) Reset() {
	b.oldestPosition = -1
	b.newestPosition = -1
}

func (b *Buffer) isFull() bool {
	if b.newestPosition == -1 && b.oldestPosition == -1 {
		return false
	}

	return b.newestPosition-b.oldestPosition == len(b.buffer)-1 || b.oldestPosition-b.newestPosition == 1
}
