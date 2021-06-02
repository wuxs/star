package star

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"strings"
)

type Star struct {
	output io.ReadWriteCloser
	buffer bytes.Buffer
}

func NewStar(output io.ReadWriteCloser) *Star {
	return &Star{
		output: output,
		buffer: bytes.Buffer{},
	}
}

func (s *Star) Reset() *Star {
	s.buffer.Write([]byte{0x18})
	return s
}

func (s *Star) Init() *Star {
	s.buffer.Write([]byte{0x1B, 0x40})
	return s
}

func (s *Star) Flush() (int, error) {
	data := s.buffer.Bytes()
	s.buffer.Reset()
	fmt.Println(strings.ToUpper(hex.EncodeToString(data)))
	return s.output.Write(data)
}

func (s *Star) SpecifyBold() *Star {
	s.buffer.Write([]byte{0x1B, 0x45})
	return s
}

func (s *Star) CancelBold() *Star {
	s.buffer.Write([]byte{0x1B, 0x46})
	return s
}

func (s *Star) SpecifyEmphasized() *Star {
	s.buffer.Write([]byte{0x1B, 0x45})
	return s
}

func (s *Star) CancelEmphasized() *Star {
	s.buffer.Write([]byte{0x1B, 0x46})
	return s
}

func (s *Star) SpecifyHigLight() *Star {
	s.buffer.Write([]byte{0x1B, 0x34})
	return s
}

func (s *Star) CancelHigLight() *Star {
	s.buffer.Write([]byte{0x1B, 0x35})
	return s
}

func (s *Star) Print(str string) *Star {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GB18030.NewEncoder())
	b, _ := ioutil.ReadAll(reader)
	s.buffer.Write(b)
	return s
}


func (s *Star) PrintWithCodePage(str string, codepage CodePage) *Star {
	s.SpecifyCodePage(codepage)
	s.buffer.Write([]byte(str))
	return s
}

func (s *Star) SpecifyCodePage(codepage CodePage) *Star {
	s.buffer.Write([]byte{0x1B, 0x1D, 0x74, byte(codepage)})
	return s
}

func (s *Star) SpecifyFont(font Font) *Star {
	s.buffer.Write([]byte{0x1B, 0x1E, 0x46, byte(font)})
	return s
}

func (s *Star) SpecifyAlignment(align Align) *Star {
	s.buffer.Write([]byte{0x1B, 0x1D, 0x61, byte(align)})
	return s
}

func (s *Star) CutFull() *Star {
	s.buffer.Write([]byte{0x1B, 0x64, 0x0})
	return s
}

func (s *Star) CutPartial() *Star {
	s.buffer.Write([]byte{0x1B, 0x64, 0x1})
	return s
}

//  1<= n <= 127
func (s *Star) FeedPaperLines(n byte) *Star {
	s.buffer.Write([]byte{0x1B, 0x61, n})
	return s
}

func (s *Star) SpecifyDoubleKanji() *Star {
	s.buffer.Write([]byte{0x1B, 0x77, 0x0})
	return s
}

func (s *Star) SpecifyDoubleWide() *Star {
	s.buffer.Write([]byte{0x1B, 0x57, 0x1})

	return s
}

func (s *Star) CancelDoubleWide() *Star {
	s.buffer.Write([]byte{0x1B, 0x57, 0x0})
	return s
}

// only ANK charset
func (s *Star) SpecifyDoubleTall() *Star {
	s.buffer.Write([]byte{0x1B, 0x68, 0x1})
	return s
}

func (s *Star) CancelDoubleTall() *Star {
	s.buffer.Write([]byte{0x1B, 0x68, 0x0})
	return s
}

func (s *Star) SpecifyBottomMarginLines(n byte) *Star {
	s.buffer.Write([]byte{0x1B, 0x4E, n})
	return s
}

func (s *Star) SpecifyLineSpace(n byte) *Star {
	s.buffer.Write([]byte{0x1B, 0x33, n})
	return s
}

func (s *Star) InquireStatus() *Star {
	s.buffer.Write([]byte{0x1B, 0x06, 0x01})
	return s
}

func (s *Star) SpecifySlashZero() *Star {
	s.buffer.Write([]byte{0x1B, 0x2F, 0x1})
	return s
}

func (s *Star) CancelSlashZero() *Star {
	s.buffer.Write([]byte{0x1B, 0x2F, 0x0})
	return s
}

func (s *Star) SpecifyUnderLine() *Star {
	s.buffer.Write([]byte{0x1B, 0x2D, 0x1})
	return s
}

func (s *Star) CancelUnderLine() *Star {
	s.buffer.Write([]byte{0x1B, 0x2D, 0x0})
	return s
}

func (s *Star) SpecifyUpperLine() *Star {
	s.buffer.Write([]byte{0x1B, 0x5F, 0x1})
	return s
}

func (s *Star) CancelUpperLine() *Star {
	s.buffer.Write([]byte{0x1B, 0x5F, 0x0})
	return s
}

func (s *Star) SpecifyUpsideDown() *Star {
	s.buffer.Write([]byte{0x0F, 0x4})
	return s
}

func (s *Star) CancelUpsideDown() *Star {
	s.buffer.Write([]byte{0x12, 0x5})
	return s
}

func (s *Star) MarginLeft(n byte) *Star {
	s.buffer.Write([]byte{0x1B, 0x6C, n})
	return s
}

func (s *Star) MarginRight(n byte) *Star {
	s.buffer.Write([]byte{0x1B, 0x51, n})
	return s
}
