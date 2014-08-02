# prettytable for golang

`prettytable` is a Go package for pretty-printing tables of data on a terminal.

```go
package main

import (
	"fmt"
	"github.com/pquerna/termchalk/prettytable"
)

func main() {
	pt := prettytable.New([]string{"City name", "Area", "Population", "Annual Rainfall"})
	pt.SortBy = "Population"
	pt.AddRow("Adelaide", 1295, 1158259, 600.5)
	pt.AddRow("Brisbane", 5905, 1857594, 1146.4)
	pt.AddRow("Darwin", 112, 120900, 1714.7)
	pt.AddRow("Hobart", 1357, 205556, 619.5)
	pt.AddRow("Sydney", 2058, 4336374, 1214.8)
	pt.AddRow("Melbourne", 1566, 3806092, 646.9)
	pt.AddRow("Perth", 5386, 1554769, 869.4)
	fmt.Println(pt)
}
```

`prettytable` is inspired by Luke Maurits's [prettytable for python](https://code.google.com/p/prettytable/) and @LearnBoost's [cli-table for node.js](https://github.com/LearnBoost/cli-table).

# License

`prettytable` is licensed under the Apache License, Version 2.0
