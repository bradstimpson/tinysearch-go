package terial

import (
	"bytes"
	"fmt"
	"io"
)

// type writer interface {
// 	io.Writer
// 	WriteByte(byte) error
// }
type Encoder struct {
	w      io.Writer
	wrpr   Code
	flds   map[string]Code
	err    error
	header []byte
	body   []byte
	buf    *bytes.Buffer
}

// type Wrapper struct {
// 	st struct
// 	mp map[string]
// }

type BaseType struct {
	b    bool
	f    float64
	s    string
	u8   uint8
	au8  []uint8
	u64  uint64
	au64 []uint64
	i64  int64
	ai64 []int64
	by   byte
	aby  []byte
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
func NewEncoder(w io.Writer, wrapper Code, fields map[string]Code) *Encoder {
	// e := &Encoder{
	// 	buf: make([]byte, 9),
	// }
	// // e.Reset(w)
	// return e
	return &Encoder{w: w, wrpr: wrapper, flds: fields}
}

func (enc *Encoder) Encode(v interface{}) error {
	if enc.err != nil {
		return enc.err
	}
	// var value (type)
	//figure out the type
	// if enc.wrpr == Struct {
	// 	value := v.(struct{})
	// } else if enc.wrpr == Map {
	// 	value := v.(map[string]interface{})
	// } else {
	bt := make(map[string]BaseType)
	for key, val := range enc.flds {
		switch val {
		case Str:
			bt[key] = BaseType{s: v.(string)}
		case Float64:
			bt[key] = BaseType{f: v.(float64)}
		case Bool:
			bt[key] = BaseType{b: v.(bool)}
		case Uint8:
			bt[key] = BaseType{u8: v.(uint8)}
		case Uint64:
			bt[key] = BaseType{u64: v.(uint64)}
		case Int64:
			bt[key] = BaseType{i64: v.(int64)}
		case Byte:
			bt[key] = BaseType{by: v.(byte)}
		case ArrayUint8:
			bt[key] = BaseType{au8: v.([]uint8)}
		case ArrayUint64:
			bt[key] = BaseType{au64: v.([]uint64)}
		case ArrayInt64:
			bt[key] = BaseType{ai64: v.([]int64)}
		case ArrayByte:
			bt[key] = BaseType{aby: v.([]byte)}
		}
		fmt.Println("did this display", val)
	}
	// 	value := v.(string)
	// }

	//build the header

	//build the body
	for _, val := range bt {
		enc.body = []byte(val.s)
	}

	enc.buf = new(bytes.Buffer)
	enc.buf.Reset()
	b := enc.body
	if _, err := enc.w.Write(b); err != nil {
		enc.err = err
	}

	return enc.err
}

func (enc *Encoder) getWrapper() string {
	// if enc.wrpr == Struct {
	// 	return func() {

	// 	}
	// } else if enc.wrpr == Struct {
	// 	return func() struct {
	// 		return struct{}
	// 	}
	// }
	return "nil"
}

// Writer returns the Encoder's writer.
// func (e *Encoder) Writer() io.Writer {
// 	return e.w
// }

// func (e *Encoder) write(b []byte) error {
// 	_, err := e.w.Write(b)
// 	return err
// }
