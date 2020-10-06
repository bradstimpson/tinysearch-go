package terial

type Marshaler interface {
	Marshal() ([]byte, error)
}

type Unmarshaler interface {
	Unmarshal([]byte) error
}

type CustomEncoder interface {
	Encode(*Encoder) error
}

type CustomDecoder interface {
	Decode(*Decoder) error
}

//------------------------------------------------------------------------------

type RawMessage []byte

var _ CustomEncoder = (RawMessage)(nil)
var _ CustomDecoder = (*RawMessage)(nil)

func (m RawMessage) Encode(enc *Encoder) error {
	return enc.write(m)
}

func (m *RawMessage) Decode(dec *Decoder) error {
	msg, err := dec.Decode()
	if err != nil {
		return err
	}
	*m = msg
	return nil
}

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
