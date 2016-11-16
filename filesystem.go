package main

import "fmt"

type Filesystem interface {
	Mkdir(string)
}

type DefaultFilesystem struct{}

func NewDefaultFilesystem() Filesystem {
	return &DefaultFilesystem{}
}

func (fs *DefaultFilesystem) Mkdir(path string) {
	fmt.Printf("would mkdir %v\n", path)
}
