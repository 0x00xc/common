package jsonx

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func FromFile(filename string, v interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return FromReader(file, v)
}

func ToFile(v interface{}, filename string) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0644)
}

func FromReader(reader io.Reader, v interface{}) error {
	return json.NewDecoder(reader).Decode(v)
}

func ToWriter(v interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(v)
}
