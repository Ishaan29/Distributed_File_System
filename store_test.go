package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "pictures"
	pathname := CASPathTransformFunc(key)
	fmt.Println(pathname)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}
	s := NewStore(opts)
	data := bytes.NewReader([]byte("some data"))
	if err := s.writeStream("somekey", data); err != nil {
		t.Error(err)
	}

}
