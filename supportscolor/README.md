# supportscolor for golang

`supportscolor` is a package to test if a terminal should support color codes.  If a terminal is not a TTY, it defaults to returning false.

```go
package main

import (
	"fmt"
	"github.com/pquerna/termchalk/supportscolor"
)

func main() {
	fmt.Println("supports color:", supportscolor.SupportsColor())
}
```

Based on sindresorhus's [supports-color for node.js](https://github.com/sindresorhus/supports-color)

# License

`ansistyle` is licensed under the Apache License, Version 2.0

