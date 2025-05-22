package io

import "io"

// NopCloser transforma um io.Reader em io.ReadCloser sem efeito colateral no Close.
func NopCloser(r io.Reader) io.ReadCloser {
	return nopCloser{r}
}

type nopCloser struct {
	r io.Reader
}

func (n nopCloser) Read(p []byte) (int, error) {
	return n.r.Read(p)
}

func (n nopCloser) Close() error {
	return nil
}
