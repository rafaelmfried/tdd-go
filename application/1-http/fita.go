package server

import io "io"

type fita struct {
	arquivo io.ReadWriteSeeker
}

func (t *fita) Write(p []byte) (n int, err error) {
	t.arquivo.Seek(0, 0)
	return t.arquivo.Write(p)
}