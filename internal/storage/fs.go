package storage

import (
	"io/fs"
)

// FS is a wrapper that falls back to `os`.
type FS struct {
	FS fs.FS
}

// Open a file in the provided `fs.FS`. If none provided,
// open via `os.Open`
func (f FS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}
