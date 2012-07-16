`go get -u github.com/bsdf/uniclolde`

```go
package main

import (
	"flag"
	"fmt"
	"github.com/bsdf/uniclolde"
	"strings"
)

var font = flag.String("f", "fullw", "font to use. type -l to list fonts")
var list = flag.Bool("l", false, "list available fonts")

type fontFunc func(string) string

var fonts = map[string]fontFunc{
	"fullw":  uniclolde.FullWidth,
	"mathb":  uniclolde.MathBold,
	"mathi":  uniclolde.MathItalic,
	"mathbi": uniclolde.MathBoldItalic,
	"ss":     uniclolde.SansSerif,
	"ssb":    uniclolde.SansSerifBold,
	"ssi":    uniclolde.SansSerifItalic,
	"ssbi":   uniclolde.SansSerifBoldItalic,
	"mono":   uniclolde.Monospace,
}

func main() {
	flag.Parse()

	if *list || len(flag.Args()) == 0 {
		listFonts()
		return
	}

	if f, ok := fonts[*font]; ok {
		str := strings.Join(flag.Args(), " ")
		fmt.Printf("%s\n", f(str))
	} else {
		fmt.Printf("\"%s\" is not a known font.\n\n", *font)
		listFonts()
	}
}

func listFonts() {
	fmt.Println("available fonts:")
	for k, f := range fonts {
		fmt.Printf("\t%s:\t%s\n", k, f("HeLLo WoRLd!!!!"))
	}
}
```