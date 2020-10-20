package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Database struct {
	Directory string
}

func CreateDatabase(directory string) (*Database, error) {
	f, err := os.Stat(directory)
	if err != nil {
		return nil, err
	}

	if !f.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", directory)
	}

	return &Database{
		Directory: directory,
	}, nil
}

func (d *Database) GetMessage(userID int) (Message, error) {
	ID := strconv.Itoa(userID)
	path := filepath.Join(d.Directory, ID)

	f, err := os.Open(path)
	if err != nil {
		return Message{}, err
	}
	defer f.Close()

	return ReadMessage(f)
}

func (d *Database) SetMessage(userID int, message string) error {
	if len(message) > 256 {
		return fmt.Errorf("message must be less than or equal to 256 characters")
	}

	ID := strconv.Itoa(userID)
	path := filepath.Join(d.Directory, ID)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(message)
	if err != nil {
		return err
	}

	return nil
}
