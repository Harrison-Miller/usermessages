package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testString = "golang is super cool"
var longTestString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func SetupDatabase(t *testing.T) (string, *Database) {
	// setup filesystem
	dir, err := ioutil.TempDir("", "data")
	assert.Nil(t, err)

	// setup database
	d, err := CreateDatabase(dir)
	assert.Nil(t, err)
	assert.NotNil(t, d)

	return dir, d
}

func TestCreateDatabaseNotExist(t *testing.T) {
	d, err := CreateDatabase("data")
	assert.NotNil(t, err)
	assert.Nil(t, d)
}

func TestGetMessage(t *testing.T) {
	dir, d := SetupDatabase(t)
	defer os.RemoveAll(dir)

	err := ioutil.WriteFile(dir+"/1", []byte(testString), 0666)
	assert.Nil(t, err)

	// test getting message
	m, err := d.GetMessage(1)
	assert.Nil(t, err)
	assert.Equal(t, testString, m.Message)
}

func TestGetMessageNotExist(t *testing.T) {
	dir, d := SetupDatabase(t)
	defer os.RemoveAll(dir)

	// test getting message that doesn't exist
	_, err := d.GetMessage(1)
	assert.NotNil(t, err)
}

func TestSetMessage(t *testing.T) {
	dir, d := SetupDatabase(t)
	defer os.RemoveAll(dir)

	// test writing a message
	err := d.SetMessage(1, testString)
	assert.Nil(t, err)

	contents, err := ioutil.ReadFile(dir + "/1")
	assert.Nil(t, err)
	assert.Equal(t, testString, string(contents))
}

func TestSetMessageLong(t *testing.T) {
	dir, d := SetupDatabase(t)
	defer os.RemoveAll(dir)

	// test writing a message longer than 256 characters
	err := d.SetMessage(1, longTestString)
	assert.NotNil(t, err)
}
