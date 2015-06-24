package fmt

import (
	"fmt"
	"io"

	"github.com/jcla1/gisp/core"
)

func Fprintf(w core.Any, f core.Any, _args ...core.Any) {
	args := []interface{}{}
	for _, _arg := range _args {
		args = append(args, _arg.(interface{}))
	}
	fmt.Fprintf(w.(io.Writer), f.(string), args...)
}
