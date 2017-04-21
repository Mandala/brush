// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// Source code is governed under MIT license, see LICENSE for more information.

// Brush wrapper implementation

package brush

import "io"

// Internal formatting state bit
const (
	// Color state
	colorOff = iota
	colorRed
	colorGreen
	colorOrange
	colorBlue
	colorPurple
	colorCyan
	colorGray
	colorMask = ^((1 << 3) - 1)
	// Format state
	formatBold = 1 << 3
)

// Internal formatting data bytes
var formatBytes = [][]byte{
	[]byte("\033[0m"),
	[]byte("\033[0;31m"),
	[]byte("\033[0;32m"),
	[]byte("\033[0;33m"),
	[]byte("\033[0;34m"),
	[]byte("\033[0;35m"),
	[]byte("\033[0;36m"),
	[]byte("\033[0;37m"),
	[]byte("\033[1m"),
	[]byte("\033[1;31m"),
	[]byte("\033[1;32m"),
	[]byte("\033[1;33m"),
	[]byte("\033[1;34m"),
	[]byte("\033[1;35m"),
	[]byte("\033[1;36m"),
	[]byte("\033[1;37m"),
}

// Writer wraps io.Writer and store related formatting state. User should use
// Wrap function to create new Writer struct instead of manual instantiation.
type Writer struct {
	io.Writer
	format int
	color  bool
}

// Color get underlying writer color support status.
func (w Writer) Color() bool {
	return w.color
}

// Bold set bold formatting
func (w *Writer) Bold() *Writer {
	if w.color {
		w.format |= formatBold
	}
	return w
}

// NoBold unset bold formatting
func (w *Writer) NoBold() *Writer {
	if w.color {
		w.format &= ^formatBold
	}
	return w
}

// Off unset all formatting
func (w *Writer) Off() {
	if !w.color {
		return
	}
	w.format = colorOff
	w.update()
}

// NoColor unset all color
func (w *Writer) NoColor() {
	if !w.color {
		return
	}
	w.format &= colorMask
	w.update()
}

// Red set red color
func (w *Writer) Red() {
	if !w.color {
		return
	}
	w.format &= colorMask
	w.format |= colorRed
	w.update()
}

// Green set green color
func (w *Writer) Green() {
	if !w.color {
		return
	}
	w.format &= colorMask
	w.format |= colorGreen
	w.update()
}

// Orange set orange color
func (w *Writer) Orange() {
	if !w.color {
		return
	}
	w.format &= colorMask
	w.format |= colorOrange
	w.update()
}

// Blue set blue color
func (w *Writer) Blue() {
	if !w.color {
		return
	}
	w.format &= colorMask
	w.format |= colorBlue
	w.update()
}

// Purple set purple color
func (w *Writer) Purple() {
	if !w.color {
		return
	}
	w.format &= colorMask
	w.format |= colorPurple
	w.update()
}

// Cyan set cyan color
func (w *Writer) Cyan() {
	if !w.color {
		return
	}
	w.format &= colorMask
	w.format |= colorCyan
	w.update()
}

// Gray set gray color
func (w *Writer) Gray() {
	if !w.color {
		return
	}
	w.format &= colorMask
	w.format |= colorGray
	w.update()
}

// Wrap create new Writer object by wrapping existing io.Writer interface and
// detect its ability to display color.
func Wrap(w io.Writer) *Writer {
	return &Writer{
		Writer: w,
		color:  isColorEnabled(w),
	}
}

// Update write current formatting state to underlying writer.
func (w Writer) update() error {
	// Write format bytes to underlying writer
	_, err := w.Write(formatBytes[w.format])
	return err
}
