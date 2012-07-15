package uniclolde

import (
	"bytes"
)

func FullWidth(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 31 && c < 127 {
			if c == 32 {
				buf.WriteRune(0x3000)
			} else {
				buf.WriteRune(0xFEE0 + c)
			}
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

func MathBold(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D3BF + c)
		} else if c > 96 && c < 122 {
			buf.WriteRune(0x1D3B9 + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

func MathItalic(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D3F3 + c)
		} else if c > 96 && c < 122 {
			buf.WriteRune(0x1D3ED + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

func MathBoldItalic(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D427 + c)
		} else if c > 96 && c < 122 {
			buf.WriteRune(0x1D421 + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

func SansSerif(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D55F + c)
		} else if c > 96 && c < 122 {
			buf.WriteRune(0x1D559 + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

func SansSerifBold(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D593 + c)
		} else if c > 96 && c < 122 {
			buf.WriteRune(0x1D58D + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

func SansSerifItalic(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D5C7 + c)
		} else if c > 96 && c < 122 {
			buf.WriteRune(0x1D5C1 + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

func SansSerifBoldItalic(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D5FB + c)
		} else if c > 96 && c < 122 {
			buf.WriteRune(0x1D5F5 + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

func Monospace(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D62F + c)
		} else if c > 96 && c < 122 {
			buf.WriteRune(0x1D629 + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}
