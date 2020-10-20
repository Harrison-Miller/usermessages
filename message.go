package main

import (
	"io"
	"io/ioutil"
)

type Message struct {
	Message string `json:"message"`
}

func ReadMessage(r io.Reader) (Message, error) {
	contents, err := ioutil.ReadAll(r)
	return Message{
		Message: string(contents),
	}, err
}
