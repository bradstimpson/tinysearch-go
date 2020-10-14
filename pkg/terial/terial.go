package terial

// Marshaler TODO
type Marshaler interface {
	Marshal() ([]byte, error)
}

// Unmarshaler TODO
type Unmarshaler interface {
	Unmarshal([]byte, interface{}) error
}

// type Encoder interface {
// 	Encode(*Encoder) error
// }

// type Decoder interface {
// 	Decode(*Decoder) error
// }

//------------------------------------------------------------------------------

// RawMessage slice of bytes alias
type RawMessage []byte

// BaseType are the types we can use
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

// var _ CustomEncoder = (RawMessage)(nil)
// var _ CustomDecoder = (*RawMessage)(nil)

// func (m RawMessage) Encode(enc *Encoder) error {
// 	return enc.write(m)
// }

// func (m *RawMessage) Decode(dec *Decoder) error {
// 	msg, err := dec.Decode()
// 	if err != nil {
// 		return err
// 	}
// 	*m = msg
// 	return nil
// }

// f, err := os.OpenFile("../../build/index.bin", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// if err != nil {
// 	log.Fatal("Couldn't open file")
// }
// defer f.Close()
// err = binary.Write(f, binary.BigEndian, uint64(len(filters)))
// for i, _ := range filters {
// 	var data = struct {
// 		f uint64
// 		u uint64
// 		n uint64
// 		i uint16
// 	}{binary.BigEndian.Uint64(filters[i]), binary.BigEndian.Uint64([]byte(urls[i])), binary.BigEndian.Uint64([]byte(names[i])), uint16(i)}

// 	err = binary.Write(f, binary.BigEndian, data)
// }
// if err != nil {
// 	log.Fatal("Write failed")
// }
