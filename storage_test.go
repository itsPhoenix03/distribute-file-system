package main

import "testing"

func TestStorage(t *testing.T) {
	opts := StorageOptions{
		PathTransformFunc: DefaultPathTransformFunc,
	}

	s := NewStorage(opts)
}
