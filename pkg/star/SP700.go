package star

import (
	"bytes"
	"io"
)

type SP700 struct {
	Star
}

func NewSP700(output io.ReadWriteCloser) *SP700 {
	return &SP700{
		Star{
			output: output,
			buffer: bytes.Buffer{},
		},
	}
}

func (s *SP700) Print(str string) *SP700 {
	s.buffer.Write([]byte(str))
	return s
}
