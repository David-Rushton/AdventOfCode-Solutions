/*
Preforamtted standard out content.
*/
package iostr

import (
	"fmt"
	"os"
	"slices"

	"github.com/David-Rushton/pretty"
)

var (
	verboseMode = slices.Contains(os.Args, "-v") || slices.Contains(os.Args, "--verbose")
)

func Out(a ...any) {
	pretty.Printf(fmt.Sprint(a...), pretty.WithRgb(172, 172, 172))
}

func Outln(a ...any) {
	pretty.Println(fmt.Sprint(a...), pretty.WithRgb(172, 172, 172))
}

func Outf(format string, a ...any) {
	pretty.Println(fmt.Sprintf(format, a...), pretty.WithRgb(172, 172, 172))
}

func Error(a ...any) {
	pretty.Printf(fmt.Sprint(a...), pretty.Red)
}

func Errorln(a ...any) {
	pretty.Println(fmt.Sprint(a...), pretty.Red)
}

func Errorf(format string, a ...any) {
	pretty.Println(fmt.Sprintf(format, a...), pretty.Red)
}

func Verbose(a ...any) {
	if verboseMode {
		pretty.Printf(fmt.Sprint(a...), pretty.Green)
	}
}

func Verboseln(a ...any) {
	if verboseMode {
		pretty.Println(fmt.Sprint(a...), pretty.Green)
	}
}

func Verbosef(format string, a ...any) {
	if verboseMode {
		pretty.Println(fmt.Sprintf(format, a...), pretty.Green)
	}
}
