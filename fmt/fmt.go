package fmt

import (
	"fmt"
	"io"
	"os"

	"github.com/gsp-lang/gsp/core"
)

func Println(_args ...core.Any) {
	FPrintln(os.Stdout, _args)
}

func FPrintln(w core.Any, _args ...core.Any) {
	args := []interface{}{}
	for _, _arg := range _args {
		args = append(args, _arg)
	}
	fmt.Fprintln(w.(io.Writer), args...)

}

func Printf(f core.Any, _args ...core.Any) {
	Fprintf(os.Stdout, f, _args)
}

func Fprintf(w core.Any, f core.Any, _args ...core.Any) {
	args := []interface{}{}
	for _, _arg := range _args {
		args = append(args, _arg)
	}
	fmt.Fprintf(w.(io.Writer), f.(string), args...)
}
