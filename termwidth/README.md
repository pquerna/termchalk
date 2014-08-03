# termwidth for golang

`termwidth` is a low level package for determining the width of a string as rendered by a modern terminal.  It correctly calculates the width of strings that can include embeded ANSI codes or wide unicode characters.

```go
package main

import (
	"fmt"
	"github.com/pquerna/termchalk/ansistyle"
	"github.com/pquerna/termchalk/termwidth"
)

func main() {
	fmt.Println("world length:", termwidth.Width("world"))
	fmt.Println("world length:", termwidth.Width(ansistyle.Bold.Open+"world"+ansistyle.Bold.Close))
}

```

# License

`termwidth` is licensed under the Apache License, Version 2.0

