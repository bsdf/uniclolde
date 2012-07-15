package uniclolde

import (
	"bytes"
)

// Ｒｅｔｕｒｎｓ　ｓｔｒ　ｉｎ　Ｆｕｌｌ　Ｗｉｄｔｈ
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

// 𝐑𝐞𝐭𝐮𝐫𝐧𝐬 𝐬𝐭𝐫 𝐢𝐧 𝐌𝐚𝐭𝐡 𝐁𝐨𝐥𝐝
func MathBold(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D3BF + c)
		} else if c > 96 && c < 123 {
			buf.WriteRune(0x1D3B9 + c)
		} else if c > 47 && c < 58 {
			buf.WriteRune(0x1D79E + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

// 𝑅𝑒𝑡𝑢𝑟𝑛𝑠 𝑠𝑡𝑟 𝑖𝑛 𝑀𝑎𝑡h 𝐼𝑡𝑎𝑙𝑖𝑐
func MathItalic(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D3F3 + c)
		} else if c > 96 && c < 123 {
			buf.WriteRune(0x1D3ED + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

// 𝑹𝒆𝒕𝒖𝒓𝒏𝒔 𝒔𝒕𝒓 𝒊𝒏 𝑴𝒂𝒕𝒉 𝑩𝒐𝒍𝒅 𝑰𝒕𝒂𝒍𝒊𝒄
func MathBoldItalic(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D427 + c)
		} else if c > 96 && c < 123 {
			buf.WriteRune(0x1D421 + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

// 𝖱𝖾𝗍𝗎𝗋𝗇𝗌 𝗌𝗍𝗋 𝗂𝗇 𝖲𝖺𝗇𝗌 𝖲𝖾𝗋𝗂𝖿
func SansSerif(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D55F + c)
		} else if c > 96 && c < 123 {
			buf.WriteRune(0x1D559 + c)
		} else if c > 47 && c < 58 {
			buf.WriteRune(0x1D7B2 + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

// 𝗥𝗲𝘁𝘂𝗿𝗻𝘀 𝘀𝘁𝗿 𝗶𝗻 𝗦𝗮𝗻𝘀 𝗦𝗲𝗿𝗶𝗳 𝗕𝗼𝗹𝗱
func SansSerifBold(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D593 + c)
		} else if c > 96 && c < 123 {
			buf.WriteRune(0x1D58D + c)
		} else if c > 47 && c < 58 {
			buf.WriteRune(0x1D7BC + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

// 𝘙𝘦𝘵𝘶𝘳𝘯𝘴 𝘴𝘵𝘳 𝘪𝘯 𝘚𝘢𝘯𝘴 𝘚𝘦𝘳𝘪𝘧 𝘐𝘵𝘢𝘭𝘪𝘤
func SansSerifItalic(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D5C7 + c)
		} else if c > 96 && c < 123 {
			buf.WriteRune(0x1D5C1 + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

// 𝙍𝙚𝙩𝙪𝙧𝙣𝙨 𝙨𝙩𝙧 𝙞𝙣 𝙎𝙖𝙣𝙨 𝙎𝙚𝙧𝙞𝙛 𝘽𝙤𝙡𝙙 𝙄𝙩𝙖𝙡𝙞𝙘
func SansSerifBoldItalic(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D5FB + c)
		} else if c > 96 && c < 123 {
			buf.WriteRune(0x1D5F5 + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}

// 𝚁𝚎𝚝𝚞𝚛𝚗𝚜 𝚜𝚝𝚛 𝚒𝚗 𝙼𝚘𝚗𝚘𝚜𝚙𝚊𝚌𝚎
func Monospace(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if c > 64 && c < 91 {
			buf.WriteRune(0x1D62F + c)
		} else if c > 96 && c < 123 {
			buf.WriteRune(0x1D629 + c)
		} else if c > 47 && c < 58 {
			buf.WriteRune(0x1D7C6 + c)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}
