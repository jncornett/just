package main

import "fmt"

type Filesystem interface {
	Mkdir(string) error
}

type DefaultFilesystem struct{}

func NewDefaultFilesystem() Filesystem {
	return &DefaultFilesystem{}
}

func (fs *DefaultFilesystem) Mkdir(path string) error {
	fmt.Printf("would mkdir %v\n", path)
	return nil
}
