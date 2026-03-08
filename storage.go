package main

import (
	"bytes"
	"io"
	"os"
)

// PathTransformFunc is a function which transforms a given key into a path for storage
type PathTransformFunc func(string) string

type StorageOptions struct {
	PathTransformFunc PathTransformFunc
}

type Storage struct {
	StorageOptions
}

// Default Path Transform Function
var DefaultPathTransformFunc = func(key string) string {
	return key
}

func NewStorage(opts StorageOptions) *Storage {
	return &Storage{
		StorageOptions: opts,
	}
}

func (s *Storage) writeSteam(key string, r io.Reader) error {
	pathname := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathname, os.ModePerm); err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	io.Copy(buf, r)

	filename := "something"

	fullPath := pathname + "/" + filename

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	return nil
}
