package file

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

// Client is a method accessofr for file methods
type Client struct {
}

// NewClient returns a new instance of client
func NewClient() Client {
	return Client{}
}

// GetFile gets the io.Reader of a file
func (c *Client) GetFile(path string) (io.Reader, error) {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(fileBytes), nil
}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func (c *Client) WriteToFile(filename string, data string) error {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
