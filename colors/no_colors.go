package colors

import (
	"bytes"
	"io"
)

type noColors struct {
	out     io.Writer
	lastbuf bytes.Buffer
}

// Uncolored will accept and io.Writer and return a
// new io.Writer that won't include colors.
func Uncolored(w io.Writer) io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (w *noColors) Write(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
