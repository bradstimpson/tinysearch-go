package terial

// Code is an alias for a byte
type Code byte

const (
	// EOC is the end of content code
	EOC   Code = 0xFF
	Nil   Code = 0xc0
	False Code = 0xc2
	True  Code = 0xc3

	Bool    Code = 0xc4
	Float64 Code = 0xc5
	Str     Code = 0xc6
	Struct  Code = 0xc7
	Map     Code = 0xc8

	Uint8       Code = 0xd0
	ArrayUint8  Code = 0xd1
	Uint64      Code = 0xd2
	ArrayUint64 Code = 0xd3
	Int64       Code = 0xd4
	ArrayInt64  Code = 0xd5
	Byte        Code = 0xd6
	ArrayByte   Code = 0xd7
)

// EncFmtID is the format ID used by terial at the start of the
// encoding which in ASCII is equivalent to TER1
const EncFmtID uint32 = 0x54455231

// Versions of the encoding
const (
	// V1 is raw
	V1 = 1 << iota
	// V2 is gzip
	V2
	// V3 is zstd
	V3
)
