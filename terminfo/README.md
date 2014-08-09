# terminfo for golang

`terminfo` contains low level utility functions for gathering information about the current terminal, such as the width of the terminal.

```go
package main

import (
	"fmt"
	"github.com/pquerna/termchalk/terminfo"
)

func main() {
	fmt.Println("window width:", terminfo.Width())
}

```

# License

`terminfo` is licensed under the Apache License, Version 2.0

