// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// Source code is governed under MIT license, see LICENSE for more information.

// +build !windows

// Check terminal support for UNIX-like environment

package brush

import (
	"os"
	"strings"
)

// OS-dependent color support check function
func checkTerminalColor(fd uintptr) bool {
	// Get the TERM env and check for color capabilities
	return strings.Contains(os.Getenv("TERM"), "color")
}
