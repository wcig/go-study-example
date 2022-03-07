package stack

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckStringBalancingSymbols(t *testing.T) {
	str := `func TestBalancingSymbols(t *testing.T) {}`
	err := CheckStringBalancingSymbols(str)
	assert.Nil(t, err)

	str = `()}`
	err = CheckStringBalancingSymbols(str)
	assert.Equal(t, EmptyError, err)

	str = `()[}`
	err = CheckStringBalancingSymbols(str)
	assert.Equal(t, NotMatchError, err)

	str = `(){`
	err = CheckStringBalancingSymbols(str)
	assert.Equal(t, NotMatchError, err)
}

func TestCheckFileBalancingSymbols(t *testing.T) {
	filename := "stack.go"
	err := CheckFileBalancingSymbols(filename)
	assert.Nil(t, err)

	filename = "temp"
	err = ioutil.WriteFile(filename, []byte("()}"), 0644)
	assert.Nil(t, err)
	err = CheckFileBalancingSymbols(filename)
	assert.Equal(t, EmptyError, err)

	err = ioutil.WriteFile(filename, []byte("()[}"), 0644)
	assert.Nil(t, err)
	err = CheckFileBalancingSymbols(filename)
	assert.Equal(t, NotMatchError, err)

	err = ioutil.WriteFile(filename, []byte("(){"), 0644)
	assert.Nil(t, err)
	err = CheckFileBalancingSymbols(filename)
	assert.Equal(t, NotMatchError, err)
	_ = os.Remove(filename)
}
