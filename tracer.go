package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	io.Writer
}

func New(w io.Writer) Tracer {
	return &tracer{Writer: w}
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.Writer, a...)
	fmt.Fprintln(t.Writer)
}
