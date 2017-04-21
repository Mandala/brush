// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// Source code is governed under MIT license, see LICENSE for more information.

/*
Package brush wraps io.Writer interface and supercharge them with high-level
color and formatting API on supported terminal.

This package is similar to Bash's tput command, but also works in Windows
system. It will look good on supported terminal but does not do anything on
anything else.

Usage

Get the package from your favorite terminal by running

	go get -u github.com/Mandala/brush

Then, include it in your package and wrap your writer with brush.Wrap(). See
example section on how to use the brush wrapper with os.Stdout writer.

Bugs and Feature Requests

Please report them on https://github.com/Mandala/brush/issues.

License

Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.

Source code is governed under MIT license, see LICENSE for more information.
*/
package brush
