package terial

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TerialTest struct {
	buf *bytes.Buffer
	enc *Encoder
	dec *Decoder
}

func TestEncode(t *testing.T) {

}

func TestDecode(t *testing.T) {

}

func TestEncodeDecode(t *testing.T) {

}

func TestMarshal(t *testing.T) {

}

func TestUnmarshal(t *testing.T) {

}

func TestBuildHeader(t *testing.T) {

}

func TestBuildBody(t *testing.T) {

}

func TestGzipCompress(t *testing.T) {

}

func TestZstdCompress(t *testing.T) {
	var test map[string]Code = map[string]Code{"testing": Nil, "what": EOC}
	for v, k := range test {
		fmt.Println("Key:", v, "Value:", test[v], "k", k)
	}

	assert.Nil(t, "nil")
}

func TestString(t *testing.T) {
	var testString string = "Testing"
	var outString interface{}
	var buf bytes.Buffer

	err := NewEncoder(&buf, Nil, map[string]Code{"test": Str}).Encode(testString)
	assert.Nil(t, err)
	fmt.Println(buf.Bytes())

	codes, err := NewDecoder(&buf).Decode(outString)
	assert.Nil(t, err)
	assert.NotNil(t, codes)
	if len(codes) > 1 {
		fmt.Println("embdedded structure")
	}
	assert.Equal(t, testString, outString)
}

func TestFloat64(t *testing.T) {

}

func TestBool(t *testing.T) {

}

func TestUint8(t *testing.T) {

}

func TestArrayUint8(t *testing.T) {

}

func TestUint64(t *testing.T) {

}

func TestArrayUint64(t *testing.T) {

}

func TestInt64(t *testing.T) {

}

func TestArrayInt64(t *testing.T) {

}

func TestByte(t *testing.T) {

}

func TestArrayByte(t *testing.T) {

}
