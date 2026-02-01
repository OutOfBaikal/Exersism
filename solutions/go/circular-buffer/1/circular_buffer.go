package circular

import (
    "errors"
)

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.

type Buffer struct {
    data   []byte
    size   int
    head   int
    tail   int
    count  int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{make([]byte, size), size, 0, size - 1, 0}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
        return 0, errors.New("Cannot read byte")
    }
    byte := b.data[b.head]
	b.head = (b.head + 1) % b.size
	b.count--
	return byte, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.count == b.size {
        return errors.New("Cannot write byte: buffer is full")
    }
    b.tail = (b.tail + 1) % b.size
	b.data[b.tail] = c
	b.count++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.count == b.size {
        // Перезаписываем старое значение
        b.head = (b.head + 1) % b.size 
    } else {
        b.count++
    }
    b.tail = (b.tail + 1) % b.size
    b.data[b.tail] = c
}

func (b *Buffer) Reset() {
	b.count = 0
	b.head = 0
	b.tail = b.size - 1
}
