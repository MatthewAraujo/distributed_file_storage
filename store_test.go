package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedFilename := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Filename != expectedPathName {
		t.Errorf("have %s want %s", pathKey.Filename, expectedFilename)
	}
}
func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStrore(opts)
	key := "momspicture"
	data := []byte("some data")

	if err := s.writeStream("myspecialpicture", bytes.NewReader(data)); err != nil {
		t.Fatal(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)

	fmt.Println(string(b))
	if string(b) != string(data) {
		t.Errorf("have %s want %s", b, data)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStrore(opts)
	key := "momspicture"
	data := []byte("some data")

	if err := s.writeStream("myspecialpicture", bytes.NewReader(data)); err != nil {
		t.Fatal(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}
