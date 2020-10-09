package terial

import (
	"bytes"
	"io"
)

// A Decoder reads and decodes values from an input stream.
type Decoder struct {
	r io.Reader
	// s   io.ByteScanner
	err    error
	header []byte
	body   []byte
	buf    *bytes.Buffer
}

// func (d *Decoder) Decode() (RawMessage, error) {
// 	d.rec = make([]byte, 0)
// 	msg := RawMessage(d.rec)
// 	d.rec = nil
// 	return msg, nil
// }

func Unmarshal(data []byte, v interface{}) error {
	return nil
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// returns an array of codes with a 1:1 mapping to the interface
// to avoid reflection.  Assume a simple return structure of either
// struct with embedded fields
// (that are simple types) or a map which points to simple types
func (dec *Decoder) Decode(v interface{}) ([]Code, error) {
	if dec.err != nil {
		return nil, dec.err
	}
	// get header

	// is body compressed? if yes decompress

	// get body

	return nil, dec.err
}
