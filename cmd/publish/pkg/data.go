package pkg

import (
	"bytes"
	"encoding/gob"
)

type FilePair struct {
	VideoName string
	VideoData []byte
	ImageName string
	ImageData []byte
}

func EncodeFile(f FilePair) ([]byte, error) {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(f)
	return buffer.Bytes(), err
}

func DecodeFile(data []byte) (FilePair, error) {
	var f FilePair
	var buffer bytes.Buffer

	_, err := buffer.Write(data)
	if err != nil {
		return FilePair{}, err
	}

	decoder := gob.NewDecoder(&buffer)
	err = decoder.Decode(&f)
	return f, err
}
