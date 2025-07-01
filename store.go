package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

func CASPathTransformFunc(key string) string {
	// using md5 hash
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])
	blockSize := 5
	sliceLen := len(hashStr) / blockSize

	paths := make([]string, sliceLen)
	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i] = hashStr[from:to]
	}

	return strings.Join(paths, "/")

}

type PathTransformFunc func(string) string

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

var DefaultPathTransformFunc = func(key string) string {
	return key
}

type Store struct {
	storeOpts StoreOpts
}

// maybe this should return an error
func NewStore(storeOpts StoreOpts) *Store {
	return &Store{
		storeOpts: storeOpts,
	}
}

// content addressable storage
// deepnested folder structure

func (s *Store) writeStream(key string, r io.Reader) error {
	pathName := s.storeOpts.PathTransformFunc(key)

	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		return err
	}

	fileName := "somefileName"
	fullPath := pathName + "/" + fileName

	f, err := os.Create(fullPath)

	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	log.Printf("written (%d) bytes to disk: %s", n, fullPath)

	return nil
}
