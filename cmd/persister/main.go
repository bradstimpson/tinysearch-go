package persister

import (
	"io"
	"os"
	"sync"
)

const (
	JSON = iota
	BIN
	GOB
	GO
)

type Persistor interface {
	Save(path string, v interface{}) error
	Load(path string, v interface{}) error
	SetMarshaller(mtype int)
}

type persistor struct {
	mtype     int
	lock      sync.Mutex
	Marshal   func(v interface{}) (io.Reader, error)
	Unmarshal func(r io.Reader, v interface{}) error
}

//add mtype int
func NewPersistor(mt int) Persistor {
	return &persistor{
		mtype: mt,
	}
}

// Independent of the constructor, change the marshaller on the fly
func (p *persistor) SetMarshaller(mtype int) {
	switch mtype {
	case JSON:
		p.Unmarshal = unmarshalJSON
		p.Marshal = marshalJSON
	case BIN:
		p.Unmarshal = unmarshalBIN
		p.Marshal = marshalBIN
	case GOB:
		p.Unmarshal = unmarshalGOB
		p.Marshal = marshalGOB
	case GO:
		p.Unmarshal = unmarshalGO
		p.Marshal = marshalGO
	}
}

// SaveJSON saves a JSON representation of v to the file at path.
func (p *persistor) Save(path string, v interface{}) error {
	p.SetMarshaller(p.mtype)
	err := save(p, path, v)
	return err
}

// save function for DRY purposes
func save(p *persistor, path string, v interface{}) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	r, err := p.Marshal(v)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	return err
}

// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func (p *persistor) Load(path string, v interface{}) error {
	p.SetMarshaller(p.mtype)
	f, err := load(p, path)
	if err != nil {
		return err
	}
	defer f.Close()
	return p.Unmarshal(f, v)
}

func load(p *persistor, path string) (*os.File, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, err
}
