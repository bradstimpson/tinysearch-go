package terial

import (
	"bytes"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type TerialTest struct {
// 	buf *bytes.Buffer
// 	enc *Encoder
// 	dec *Decoder
// }

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
	// var test map[string]Code = map[string]Code{"testing": Nil, "what": EOC}
	// for v, k := range test {
	// 	fmt.Println("Key:", v, "Value:", test[v], "k", k)
	// }

	// assert.Nil(t, "nil")
}

func TestString(t *testing.T) {
	var teststring string = "Testing"
	var outstring interface{}
	var buf bytes.Buffer

	fields := map[string]Code{"test": Str}

	err := NewEncoder(&buf, Nil, fields).Encode(teststring)
	assert.Nil(t, err)
	fmt.Println(buf.Bytes())

	outstring, err = NewDecoder(&buf, Nil, fields).Decode()
	assert.Nil(t, err)
	assert.NotNil(t, outstring)
	assert.Equal(t, teststring, outstring)
}

func TestFloat64(t *testing.T) {
	var testfloat float64 = float64(-1*math.MaxFloat64 - 1)
	var outfloat interface{}
	var buf bytes.Buffer

	fields := map[string]Code{"test": Float64}

	err := NewEncoder(&buf, Nil, fields).Encode(testfloat)
	assert.Nil(t, err)
	fmt.Println(buf.Bytes())

	outfloat, err = NewDecoder(&buf, Nil, fields).Decode()
	assert.Nil(t, err)
	assert.NotNil(t, outfloat)
	assert.Equal(t, testfloat, outfloat)
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
