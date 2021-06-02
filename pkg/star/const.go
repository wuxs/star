package star

type Align byte

const (
	Left   Align = 0x0
	Center Align = 0x1
	Right  Align = 0x2
)

type Font byte

const (
	A Font = 0x0
	B Font = 0x1
	C Font = 0x10
)

type CodePage byte

const (
	Normal    CodePage = 0x0
	Utf8      CodePage = 0x80
	Undefined CodePage = 0xFF
	// TODO more code page
)
