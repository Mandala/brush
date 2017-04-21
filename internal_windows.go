// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// Source code is governed under MIT license, see LICENSE for more information.

// +build windows

// Check terminal support for UNIX-like environment

package brush

import (
	"syscall"
	"unsafe"
)

const (
	enableVirtualTerminalProcessing = 4
)

// Kernel32 DLL loader for GetConsoleMode function
var kernel32 = syscall.NewLazyDLL("kernel32.dll")
var getConsoleMode = kernel32.NewProc("GetConsoleMode")

// OS-dependent color support check function
func checkTerminalColor(fd uintptr) bool {
	var mode uint32
	_, _, e := syscall.Syscall(getConsoleMode.Addr(), 2, uintptr(fd), uintptr(unsafe.Pointer(&mode)), 0)
	if e != 0 {
		return false
	}
	return (mode & enableVirtualTerminalProcessing) != 0
}
