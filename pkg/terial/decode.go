package terial

import (
	"bytes"
	"fmt"
	"io"
)

// A Decoder reads and decodes values from an input stream.
type Decoder struct {
	r    io.Reader
	wrpr Code
	flds map[string]Code
	// s   io.ByteScanner
	err    error
	header []byte
	body   []byte
	buf    *bytes.Buffer
	valbuf *bytes.Buffer
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

func NewDecoder(r io.Reader, wrapper Code, fields map[string]Code) *Decoder {
	return &Decoder{r: r, wrpr: wrapper, flds: fields}
}

// returns an array of codes with a 1:1 mapping to the interface
// to avoid reflection.  Assume a simple return structure of either
// struct with embedded fields
// (that are simple types) or a map which points to simple types
func (dec *Decoder) Decode() (interface{}, error) {
	if dec.err != nil {
		return 0, dec.err
	}
	dec.valbuf = new(bytes.Buffer)
	dec.valbuf.Reset()
	// get header

	// is body compressed? if yes decompress

	// get body
	dec.buf = new(bytes.Buffer)
	dec.buf.Reset()

	if _, err := dec.buf.ReadFrom(dec.r); err != nil {
		dec.err = err
	}

	for idx, val := range dec.buf.Bytes() {
		if idx == 0 {
			switch val {
			case byte(Str):
				fmt.Println("string found")
			case byte(Float64):
				fmt.Println("float64 found")
			case byte(Bool):
				fmt.Println("bool found")
			case byte(Uint8):
				fmt.Println("uint8 found")
			case byte(Uint64):
				fmt.Println("uint64 found")
			case byte(Int64):
				fmt.Println("int64 found")
			case byte(Byte):
				fmt.Println("byte found")
			case byte(ArrayUint8):
				fmt.Println("[]uint8 found")
			case byte(ArrayUint64):
				fmt.Println("[]uint64 found")
			case byte(ArrayInt64):
				fmt.Println("[]int64 found")
			case byte(ArrayByte):
				fmt.Println("[]byte found")
			}
			// fmt.Printf("the input value type: 0x%02x\n", val)
		} else {
			dec.valbuf.WriteByte(val)
		}
	}
	return dec.valbuf.String(), dec.err
}
