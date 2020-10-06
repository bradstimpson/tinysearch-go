package persister

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	jsonFile = "../../build/temp.json"
	gobFile  = "../../build/temp.gob"
	goFile   = "../../build/temp.go"
)

type obj struct {
	Name   string
	Number int
	When   time.Time
}

var in *obj = &obj{
	Name:   "Test",
	Number: 42,
	When:   time.Now(),
}
var out obj

func TestWriteReadJSON(t *testing.T) {
	teardownJSON := setupJSON(t)
	defer teardownJSON(t)

	np := NewPersistor(JSON)
	err := np.Save(jsonFile, in)
	assert.Nil(t, err)

	if _, err := os.Stat(jsonFile); !os.IsNotExist(err) {
		assert.Nil(t, err)
	}

	err = np.Load(jsonFile, &out)
	assert.Nil(t, err)
	assert.Equal(t, "Test", out.Name)

}

// func TestWriteReadBin(t *testing.T) {
// 	var o obj
// 	np := NewPersistor()
// 	err := np.SaveGOB("../../build/temp.bin", &o)
// 	assert.Nil(t, err)

// 	assert.False(t, true, "msg")
// }

func TestWriteReadGOB(t *testing.T) {
	teardownGOB := setupGOB(t)
	defer teardownGOB(t)

	np := NewPersistor(GOB)
	err := np.Save(gobFile, in)
	assert.Nil(t, err)

	if _, err := os.Stat(gobFile); !os.IsNotExist(err) {
		assert.Nil(t, err)
	}

	err = np.Load(gobFile, &out)
	assert.Nil(t, err)
	assert.Equal(t, "Test", out.Name)
}

func TestWriteReadGOGoodType(t *testing.T) {
	teardownGO := setupGO(t)
	defer teardownGO(t)

	var data []byte = []byte{1, 2, 3, 4}
	np := NewPersistor(GO)
	err := np.Save(goFile, data)
	assert.Nil(t, err)

	if _, err := os.Stat(goFile); !os.IsNotExist(err) {
		assert.Nil(t, err)
	}
}

func TestWriteReadGOBadType(t *testing.T) {
	var data []int = []int{1, 2, 3, 4}
	np := NewPersistor(GO)
	err := np.Save(goFile, data)
	assert.NotNil(t, err)

	if _, err := os.Stat(goFile); os.IsNotExist(err) {
		assert.NotNil(t, err)
	}
}

func TestMain(m *testing.M) {
	//startup()
	code := m.Run()
	// shutdown(m)
	os.Exit(code)
}

func setupJSON(t *testing.T) func(t *testing.T) {
	// t.Log("setup test case")
	return func(t *testing.T) {
		// t.Log("teardown test case")
		c := exec.Command("bash", "-c", "rm -rf "+jsonFile)
		if err := c.Run(); err != nil {
			fmt.Printf("error deleting temp files %s", err.Error())
		}
	}
}

func setupGOB(t *testing.T) func(t *testing.T) {
	// t.Log("setup test case")
	return func(t *testing.T) {
		// t.Log("teardown test case")
		c := exec.Command("bash", "-c", "rm -rf "+gobFile)
		if err := c.Run(); err != nil {
			fmt.Printf("error deleting temp files %s", err.Error())
		}
	}
}

func setupGO(t *testing.T) func(t *testing.T) {
	// t.Log("setup test case")
	return func(t *testing.T) {
		// t.Log("teardown test case")
		c := exec.Command("bash", "-c", "rm -rf "+goFile)
		if err := c.Run(); err != nil {
			fmt.Printf("error deleting temp files %s", err.Error())
		}
	}
}
