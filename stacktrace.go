package godog

import (
	"fmt"
)

// Frame represents a program counter inside a stack frame.
type stackFrame uintptr

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
func (f stackFrame) pc() uintptr { _ = "STUB: not implemented"; return 0 }

// file returns the full path to the file that contains the
// function for this Frame's pc.
func (f stackFrame) file() string { _ = "STUB: not implemented"; return "" }

func trimGoPath(file string) string { _ = "STUB: not implemented"; return "" }

// line returns the line number of source code of the
// function for this Frame's pc.
func (f stackFrame) line() int { _ = "STUB: not implemented"; return 0 }

// Format formats the frame according to the fmt.Formatter interface.
//
//	%s    source file
//	%d    source line
//	%n    function name
//	%v    equivalent to %s:%d
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//	%+s   path of source file relative to the compile time GOPATH
//	%+v   equivalent to %+s:%d
func (f stackFrame) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// stack represents a stack of program counters.
type stack []uintptr

func (s *stack) Format(st fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func callStack() *stack { _ = "STUB: not implemented"; return nil }

// fundamental is an error that has a message and a stack, but no caller.
type traceError struct {
	msg string
	*stack
}

func (f *traceError) Error() string { _ = "STUB: not implemented"; return "" }

func (f *traceError) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }
