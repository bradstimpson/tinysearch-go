package terial

import (
	"io"
)

type writer interface {
	io.Writer
	WriteByte(byte) error
}
type Encoder struct {
	w writer

	buf []byte
}

// Marshal takes the interface and converts it into a format
// ready for encoding
func Marshal(v interface{}) ([]byte, error) {
	// enc := NewEncoder()

	// var buf bytes.Buffer
	// enc.Reset(&buf)

	// err := enc.Encode(v)
	// b := buf.Bytes()

	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	e := &Encoder{
		buf: make([]byte, 9),
	}
	// e.Reset(w)
	return e
}

// Writer returns the Encoder's writer.
func (e *Encoder) Writer() io.Writer {
	return e.w
}

func (e *Encoder) write(b []byte) error {
	_, err := e.w.Write(b)
	return err
}
