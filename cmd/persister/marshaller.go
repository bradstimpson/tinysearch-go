package persister

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"text/template"
)

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
func unmarshalJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// Marshal is a function that marshals the object into an
// io.Reader.
// By default, it uses the JSON marshaller.
func marshalJSON(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// No decoding - if interface []byte convert to io.Reader
func unmarshalBIN(r io.Reader, v interface{}) error {
	if ok, err := name(v); !ok {
		return err
	}

	var b bytes.Buffer
	_, err := b.ReadFrom(r)

	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	for _, val := range b.Bytes() {
		value.Set(reflect.Append(value, reflect.ValueOf(val)))
	}

	return err
}

// No encoding method currently
func marshalBIN(v interface{}) (io.Reader, error) {
	fmt.Println("Marshalling BINARY is currently not supported")
	return nil, nil
}

func unmarshalGOB(r io.Reader, v interface{}) error {
	return gob.NewDecoder(r).Decode(v)
}

// MarshalGOB encodes the interface into a GO GOB
func marshalGOB(v interface{}) (io.Reader, error) {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b.Bytes()), nil
}

func unmarshalGO(r io.Reader, v interface{}) error {
	fmt.Println("Unmarshalling GO is currently not supported")
	return nil
}

func marshalGO(v interface{}) (io.Reader, error) {
	if ok, err := name(v); !ok {
		return nil, err
	}
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// small fix for running tests and normal mode
	if strings.HasSuffix(dir, "tests") {
		dir = strings.TrimSuffix(dir, "tests") + "cmd/persister"
	}

	var b bytes.Buffer
	t := template.Must(template.ParseFiles(dir + "/index.tpl"))
	err = t.Execute(&b, v)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b.Bytes()), err
}

func name(v interface{}) (bool, error) {
	switch v.(type) {
	case []byte:
		return true, nil
	case byte:
		return true, nil
	case *[]byte:
		return true, nil
	default:
		return false, errors.New("Unsupported type")
	}
}
