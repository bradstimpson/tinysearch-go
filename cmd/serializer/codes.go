package serializer

type Code byte

var (
	EOC   Code = 0xFF
	Nil   Code = 0xc0
	False Code = 0xc2
	True  Code = 0xc3

	Bool    Code = 0xc4
	Float64 Code = 0xc5
	Str     Code = 0xc6

	Uint8       Code = 0xd0
	ArrayUint8  Code = 0xd1
	Uint64      Code = 0xd2
	ArrayUint64 Code = 0xd3
	Int64       Code = 0xd4
	ArrayInt64  Code = 0xd5
	Byte        Code = 0xd6
	ArrayByte   Code = 0xd7
)
