package terial

import "io"

// A Decoder reads and decodes MessagePack values from an input stream.
type Decoder struct {
	r   io.Reader
	s   io.ByteScanner
	buf []byte
	rec []byte // accumulates read data if not nil
}

func (d *Decoder) Decode() (RawMessage, error) {
	d.rec = make([]byte, 0)
	msg := RawMessage(d.rec)
	d.rec = nil
	return msg, nil
}
