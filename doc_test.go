// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// Source code is governed under MIT license, see LICENSE for more information.

// Package brush example file for Godoc documentation

package brush_test

import (
	"fmt"
	"os"

	"github.com/mandala/brush"
)

// Wrap os.Stdout writer and print a message with red color.
func Example() {
	// Wrap stdout with brush wrapper
	color := brush.Wrap(os.Stdout)
	// Switch to red color on supported terminal
	color.Red()
	// Print actual message to user
	fmt.Fprintf(color, "Works with standard fmt.Fprintf function")
	// Don't forget to turn off formatting after writes
	color.Off()
}
