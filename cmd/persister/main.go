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
	SaveJSON(path string, v interface{}) error
	SaveBIN(path string, v interface{}) error
	SaveGOB(path string, v interface{}) error
	SaveGO(path string, v interface{}) error
	LoadJSON(path string, v interface{}) error
	LoadBIN(path string, v interface{}) error
	LoadGOB(path string, v interface{}) error
	LoadGO(path string, v interface{}) error
	setMarshaller(mtype int)
}

type persistor struct {
	lock      sync.Mutex
	Marshal   func(v interface{}) (io.Reader, error)
	Unmarshal func(r io.Reader, v interface{}) error
}

func NewPersistor() Persistor {
	return &persistor{}
}

func (p *persistor) setMarshaller(mtype int) {
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
func (p *persistor) SaveJSON(path string, v interface{}) error {
	p.setMarshaller(JSON)
	err := save(p, path, v)
	return err
}

// SaveBIN saves a binary representation of v to the file at path.
func (p *persistor) SaveBIN(path string, v interface{}) error {
	p.setMarshaller(BIN)
	err := save(p, path, v)
	return err
}

// SaveGOB saves a gob representation of v to the file at path.
func (p *persistor) SaveGOB(path string, v interface{}) error {
	p.setMarshaller(GOB)
	err := save(p, path, v)
	return err
}

// SaveGO saves a go template representation of v to the file at path
func (p *persistor) SaveGO(path string, v interface{}) error {
	p.setMarshaller(GO)
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
func (p *persistor) LoadJSON(path string, v interface{}) error {
	p.setMarshaller(JSON)
	f, err := load(p, path)
	if err != nil {
		return err
	}
	defer f.Close()
	return p.Unmarshal(f, v)
}

func (p *persistor) LoadGOB(path string, v interface{}) error {
	p.setMarshaller(GOB)
	f, err := load(p, path)
	if err != nil {
		return err
	}
	defer f.Close()
	return p.Unmarshal(f, v)
}

func (p *persistor) LoadBIN(path string, v interface{}) error {
	p.setMarshaller(BIN)
	f, err := load(p, path)
	if err != nil {
		return err
	}
	defer f.Close()
	return p.Unmarshal(f, v)
}

func (p *persistor) LoadGO(path string, v interface{}) error {
	p.setMarshaller(GO)
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
