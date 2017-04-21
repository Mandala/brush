// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// Source code is governed under MIT license, see LICENSE for more information.

// Internal brush function to detect color support in terminal

package brush

import (
	"io"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// Test color capability on a writer interface
func isColorEnabled(w io.Writer) bool {
	// Try to reflect writer into *os.File
	file, ok := w.(*os.File)
	if !ok {
		return false
	}
	// Check if file is a terminal
	if !terminal.IsTerminal(int(file.Fd())) {
		return false
	}
	// Return from os-dependent terminal color checker
	return checkTerminalColor(file.Fd())
}
