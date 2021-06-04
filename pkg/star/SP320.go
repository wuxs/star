package star

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
)

type SP320 struct {
	Star
}

func NewSP320(output io.ReadWriteCloser) *SP320 {
	return &SP320{
		Star{
			output: output,
			buffer: bytes.Buffer{},
		},
	}
}

func (s *SP320) Print(str string) *SP320 {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GB18030.NewEncoder())
	b, _ := ioutil.ReadAll(reader)
	s.buffer.Write(b)
	return s
}
