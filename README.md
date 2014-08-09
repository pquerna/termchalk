# termchalk: utilities for making pretty terminal UIs in golang.

![](https://api.travis-ci.org/pquerna/termchalk.svg)

Termchalk packages are divided into two groups: High level packages suitable for easy integration, and low level packages useful for building customized interfaces.

# Overall Status: WIP

# High level packages:

* [prettytable](./prettytable/README.md): Render tables of data to the terminal.

# Low level packages:

* [ansistyle](./ansistyle/README.md): Low level ANSI color codes.
* [runewidth](./runewidth/README.md): Calculate the width of strings as rendered by a terminal.
* [supportscolor](./supportscolor/README.md): Detect if terminal can support color.
* [terminfo](./terminfo/README.md): Gather low level information about the current terminal.

# License

`termchalk` is licensed under the Apache License, Version 2.0.

Portions are derivatives of other open source projects, see [NOTICE](./NOTICE) for details.

