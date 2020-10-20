package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMessage string = "Lorem ipsum dolor sit amet."

func TestReadMessage(t *testing.T) {
	r := strings.NewReader(testMessage)
	m, err := ReadMessage(r)

	assert.Nil(t, err)
	assert.Equal(t, testMessage, m.Message)
}
