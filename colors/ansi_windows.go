// Copyright 2014 shiena Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package colors

import (
	"bytes"
	"io"
	"syscall"
)

type csiState int

const (
	outsideCsiCode csiState = iota
	firstCsiCode
	secondCsiCode
)

type parseResult int

const (
	noConsole parseResult = iota
	changedColor
	unknown
)

type ansiColorWriter struct {
	w             io.Writer
	mode          outputMode
	state         csiState
	paramStartBuf bytes.Buffer
	paramBuf      bytes.Buffer
}

const (
	firstCsiChar   byte = '\x1b'
	secondeCsiChar byte = '['
	separatorChar  byte = ';'
	sgrCode        byte = 'm'
)

const (
	foregroundBlue      = uint16(0x0001)
	foregroundGreen     = uint16(0x0002)
	foregroundRed       = uint16(0x0004)
	foregroundIntensity = uint16(0x0008)
	backgroundBlue      = uint16(0x0010)
	backgroundGreen     = uint16(0x0020)
	backgroundRed       = uint16(0x0040)
	backgroundIntensity = uint16(0x0080)
	underscore          = uint16(0x8000)

	foregroundMask = foregroundBlue | foregroundGreen | foregroundRed | foregroundIntensity
	backgroundMask = backgroundBlue | backgroundGreen | backgroundRed | backgroundIntensity
)

const (
	ansiReset        = "0"
	ansiIntensityOn  = "1"
	ansiIntensityOff = "21"
	ansiUnderlineOn  = "4"
	ansiUnderlineOff = "24"
	ansiBlinkOn      = "5"
	ansiBlinkOff     = "25"

	ansiForegroundBlack   = "30"
	ansiForegroundRed     = "31"
	ansiForegroundGreen   = "32"
	ansiForegroundYellow  = "33"
	ansiForegroundBlue    = "34"
	ansiForegroundMagenta = "35"
	ansiForegroundCyan    = "36"
	ansiForegroundWhite   = "37"
	ansiForegroundDefault = "39"

	ansiBackgroundBlack   = "40"
	ansiBackgroundRed     = "41"
	ansiBackgroundGreen   = "42"
	ansiBackgroundYellow  = "43"
	ansiBackgroundBlue    = "44"
	ansiBackgroundMagenta = "45"
	ansiBackgroundCyan    = "46"
	ansiBackgroundWhite   = "47"
	ansiBackgroundDefault = "49"

	ansiLightForegroundGray    = "90"
	ansiLightForegroundRed     = "91"
	ansiLightForegroundGreen   = "92"
	ansiLightForegroundYellow  = "93"
	ansiLightForegroundBlue    = "94"
	ansiLightForegroundMagenta = "95"
	ansiLightForegroundCyan    = "96"
	ansiLightForegroundWhite   = "97"

	ansiLightBackgroundGray    = "100"
	ansiLightBackgroundRed     = "101"
	ansiLightBackgroundGreen   = "102"
	ansiLightBackgroundYellow  = "103"
	ansiLightBackgroundBlue    = "104"
	ansiLightBackgroundMagenta = "105"
	ansiLightBackgroundCyan    = "106"
	ansiLightBackgroundWhite   = "107"
)

type drawType int

const (
	foreground drawType = iota
	background
)

type winColor struct {
	code     uint16
	drawType drawType
}

var colorMap = map[string]winColor{
	ansiForegroundBlack:   {0, foreground},
	ansiForegroundRed:     {foregroundRed, foreground},
	ansiForegroundGreen:   {foregroundGreen, foreground},
	ansiForegroundYellow:  {foregroundRed | foregroundGreen, foreground},
	ansiForegroundBlue:    {foregroundBlue, foreground},
	ansiForegroundMagenta: {foregroundRed | foregroundBlue, foreground},
	ansiForegroundCyan:    {foregroundGreen | foregroundBlue, foreground},
	ansiForegroundWhite:   {foregroundRed | foregroundGreen | foregroundBlue, foreground},
	ansiForegroundDefault: {foregroundRed | foregroundGreen | foregroundBlue, foreground},

	ansiBackgroundBlack:   {0, background},
	ansiBackgroundRed:     {backgroundRed, background},
	ansiBackgroundGreen:   {backgroundGreen, background},
	ansiBackgroundYellow:  {backgroundRed | backgroundGreen, background},
	ansiBackgroundBlue:    {backgroundBlue, background},
	ansiBackgroundMagenta: {backgroundRed | backgroundBlue, background},
	ansiBackgroundCyan:    {backgroundGreen | backgroundBlue, background},
	ansiBackgroundWhite:   {backgroundRed | backgroundGreen | backgroundBlue, background},
	ansiBackgroundDefault: {0, background},

	ansiLightForegroundGray:    {foregroundIntensity, foreground},
	ansiLightForegroundRed:     {foregroundIntensity | foregroundRed, foreground},
	ansiLightForegroundGreen:   {foregroundIntensity | foregroundGreen, foreground},
	ansiLightForegroundYellow:  {foregroundIntensity | foregroundRed | foregroundGreen, foreground},
	ansiLightForegroundBlue:    {foregroundIntensity | foregroundBlue, foreground},
	ansiLightForegroundMagenta: {foregroundIntensity | foregroundRed | foregroundBlue, foreground},
	ansiLightForegroundCyan:    {foregroundIntensity | foregroundGreen | foregroundBlue, foreground},
	ansiLightForegroundWhite:   {foregroundIntensity | foregroundRed | foregroundGreen | foregroundBlue, foreground},

	ansiLightBackgroundGray:    {backgroundIntensity, background},
	ansiLightBackgroundRed:     {backgroundIntensity | backgroundRed, background},
	ansiLightBackgroundGreen:   {backgroundIntensity | backgroundGreen, background},
	ansiLightBackgroundYellow:  {backgroundIntensity | backgroundRed | backgroundGreen, background},
	ansiLightBackgroundBlue:    {backgroundIntensity | backgroundBlue, background},
	ansiLightBackgroundMagenta: {backgroundIntensity | backgroundRed | backgroundBlue, background},
	ansiLightBackgroundCyan:    {backgroundIntensity | backgroundGreen | backgroundBlue, background},
	ansiLightBackgroundWhite:   {backgroundIntensity | backgroundRed | backgroundGreen | backgroundBlue, background},
}

var (
	kernel32                       = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTextAttribute    = kernel32.NewProc("SetConsoleTextAttribute")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	defaultAttr                    *textAttributes
)

func init() {
	screenInfo := getConsoleScreenBufferInfo(uintptr(syscall.Stdout))
	if screenInfo != nil {
		colorMap[ansiForegroundDefault] = winColor{
			screenInfo.WAttributes & (foregroundRed | foregroundGreen | foregroundBlue),
			foreground,
		}
		colorMap[ansiBackgroundDefault] = winColor{
			screenInfo.WAttributes & (backgroundRed | backgroundGreen | backgroundBlue),
			background,
		}
		defaultAttr = convertTextAttr(screenInfo.WAttributes)
	}
}

type coord struct {
	X, Y int16
}

type smallRect struct {
	Left, Top, Right, Bottom int16
}

type consoleScreenBufferInfo struct {
	DwSize              coord
	DwCursorPosition    coord
	WAttributes         uint16
	SrWindow            smallRect
	DwMaximumWindowSize coord
}

func getConsoleScreenBufferInfo(hConsoleOutput uintptr) *consoleScreenBufferInfo {
	_ = "STUB: not implemented"
	return nil
}

func setConsoleTextAttribute(hConsoleOutput uintptr, wAttributes uint16) bool {
	_ = "STUB: not implemented"
	return false
}

type textAttributes struct {
	foregroundColor     uint16
	backgroundColor     uint16
	foregroundIntensity uint16
	backgroundIntensity uint16
	underscore          uint16
	otherAttributes     uint16
}

func convertTextAttr(winAttr uint16) *textAttributes { _ = "STUB: not implemented"; return nil }

func convertWinAttr(textAttr *textAttributes) uint16 { _ = "STUB: not implemented"; return 0 }

func changeColor(param []byte) parseResult { _ = "STUB: not implemented"; return *new(parseResult) }

// unknown code

func parseEscapeSequence(command byte, param []byte) parseResult {
	_ = "STUB: not implemented"
	return *new(parseResult)
}

func (cw *ansiColorWriter) flushBuffer() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cw *ansiColorWriter) resetBuffer() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cw *ansiColorWriter) flushTo(w io.Writer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func isParameterChar(b byte) bool { _ = "STUB: not implemented"; return false }

func (cw *ansiColorWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Add one more to the size of the buffer for the last ch
