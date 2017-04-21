Brush, Colorful Writer Interface Wrapper
=========================================

Package brush wraps `io.Writer` interface and supercharge them with high-level
color and formatting API on supported terminal.

This package is similar to Bash's `tput` command, but also works in Windows
system. It will look good on supported terminal but does not do anything on
anything else.

## Usage

Get the package from your favorite terminal by running

```
go get -u github.com/Mandala/brush
```

Then, include it in your package and wrap your writer with `brush.Wrap()`.

```go
import (
    "fmt"
    "os"

    "github.com/mandala/brush"
)

func main() {
    // Wrap stdout with brush wrapper
    color := brush.Wrap(os.Stdout)
    // Switch to red color on supported terminal
    color.Red()
    // Print actual message to user
    fmt.Fprintf(color, "Works with standard fmt.Fprintf function")
    // Don't forget to turn off formatting after writes
    color.Off()
}
```

## API Documentation

Access Brush API documentation on Godoc.org at <https://godoc.org/github.com/Mandala/brush>.

## Bugs and Feature Requests

Please report bugs and feature requests on <https://github.com/Mandala/brush/issues>.

## License

Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.

Source code is governed under MIT license, see LICENSE for more information.
