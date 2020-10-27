package terial

// Code is an alias for a byte
type Code byte

const (
	// EOC is the end of content code
	EOC Code = 0xFF
	// Nil is a noop
	Nil Code = 0xc0
	// False boolean
	False Code = 0xc2
	// True boolean
	True Code = 0xc3

	// Bool type
	Bool Code = 0xc4
	// Float64 type
	Float64 Code = 0xc5
	// Str type
	Str Code = 0xc6
	// Struct type
	Struct Code = 0xc7
	// Map type
	Map Code = 0xc8

	// Uint8 type
	Uint8 Code = 0xd0
	// ArrayUint8 type
	ArrayUint8 Code = 0xd1
	// Uint64 type
	Uint64 Code = 0xd2
	// ArrayUint64 type
	ArrayUint64 Code = 0xd3
	// Int64 type
	Int64 Code = 0xd4
	// ArrayInt64 type
	ArrayInt64 Code = 0xd5
	// Byte type
	Byte Code = 0xd6
	// ArrayByte type
	ArrayByte Code = 0xd7
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
