package uniclolde

import (
	// "fmt"
	"testing"
)

func TestFullWidth(t *testing.T) {
	const expected = "ＨｅＬＬｏ　ＷｏＲｌＤ！！！"

	str := "HeLLo WoRlD!!!"
	out := FullWidth(str)

	if out != expected {
		t.Errorf("Expected: %s, returned: %s", expected, out)
	}
}

func TestMathBold(t *testing.T) {
	const expected = "𝐇𝐞𝐋𝐋𝐨 𝐖𝐨𝐑𝐥𝐃!!!"

	str := "HeLLo WoRlD!!!"
	out := MathBold(str)

	if out != expected {
		t.Errorf("Expected: %s, returned: %s", expected, out)
	}
}

func TestMathItalic(t *testing.T) {
	const expected = "𝐻𝑒𝐿𝐿𝑜 𝑊𝑜𝑅𝑙𝐷!!!"

	str := "HeLLo WoRlD!!!"
	out := MathItalic(str)

	if out != expected {
		t.Errorf("Expected: %s, returned: %s", expected, out)
	}
}

func TestMathBoldItalic(t *testing.T) {
	const expected = "𝑯𝒆𝑳𝑳𝒐 𝑾𝒐𝑹𝒍𝑫!!!"

	str := "HeLLo WoRlD!!!"
	out := MathBoldItalic(str)

	if out != expected {
		t.Errorf("Expected: %s, returned: %s", expected, out)
	}
}

func TestSansSerif(t *testing.T) {
	const expected = "𝖧𝖾𝖫𝖫𝗈 𝖶𝗈𝖱𝗅𝖣!!!"

	str := "HeLLo WoRlD!!!"
	out := SansSerif(str)

	if out != expected {
		t.Errorf("Expected: %s, returned: %s", expected, out)
	}
}

func TestSansSerifBold(t *testing.T) {
	const expected = "𝗛𝗲𝗟𝗟𝗼 𝗪𝗼𝗥𝗹𝗗!!!"

	str := "HeLLo WoRlD!!!"
	out := SansSerifBold(str)

	if out != expected {
		t.Errorf("Expected: %s, returned: %s", expected, out)
	}
}

func TestSansSerifItalic(t *testing.T) {
	const expected = "𝘏𝘦𝘓𝘓𝘰 𝘞𝘰𝘙𝘭𝘋!!!"

	str := "HeLLo WoRlD!!!"
	out := SansSerifItalic(str)

	if out != expected {
		t.Errorf("Expected: %s, returned: %s", expected, out)
	}
}

func TestSansSerifBoldItalic(t *testing.T) {
	const expected = "𝙃𝙚𝙇𝙇𝙤 𝙒𝙤𝙍𝙡𝘿!!!"

	str := "HeLLo WoRlD!!!"
	out := SansSerifBoldItalic(str)

	if out != expected {
		t.Errorf("Expected: %s, returned: %s", expected, out)
	}
}

func TestMonospace(t *testing.T) {
	const expected = "𝙷𝚎𝙻𝙻𝚘 𝚆𝚘𝚁𝚕𝙳!!!"

	str := "HeLLo WoRlD!!!"
	out := Monospace(str)

	if out != expected {
		t.Errorf("Expected: %s, returned: %s", expected, out)
	}
}
