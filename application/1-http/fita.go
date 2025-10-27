package server

import (
	"os"
)

type fita struct {
	file *os.File
}

func NewFita(arquivo *os.File) *fita {
	return &fita{file: arquivo}
}

func (t *fita) Write(p []byte) (n int, err error) {
	t.file.Seek(0, 0)
	return t.file.Write(p)
}