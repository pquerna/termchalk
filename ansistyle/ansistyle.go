/**
 *  Copyright 2014 Paul Querna
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package ansistyle

const openStr = "\u001b["
const closeStr = "\u001b["

type Code struct {
	Open  string
	Close string
}

var Reset = Code{
	Open:  openStr + "0" + "m",
	Close: openStr + "0" + "m",
}

var Bold = Code{
	Open:  openStr + "1" + "m",
	Close: openStr + "22" + "m",
}

var Dim = Code{
	Open:  openStr + "2" + "m",
	Close: openStr + "22" + "m",
}

// Modes

var Italic = Code{
	Open:  openStr + "3" + "m",
	Close: openStr + "23" + "m",
}

var Underline = Code{
	Open:  openStr + "4" + "m",
	Close: openStr + "24" + "m",
}

var Inverse = Code{
	Open:  openStr + "7" + "m",
	Close: openStr + "27" + "m",
}

var Hidden = Code{
	Open:  openStr + "8" + "m",
	Close: openStr + "28" + "m",
}

var Strikethrough = Code{
	Open:  openStr + "9" + "m",
	Close: openStr + "29" + "m",
}

// Colors

var Black = Code{
	Open:  openStr + "30" + "m",
	Close: openStr + "39" + "m",
}

var Red = Code{
	Open:  openStr + "31" + "m",
	Close: openStr + "39" + "m",
}

var Green = Code{
	Open:  openStr + "32" + "m",
	Close: openStr + "39" + "m",
}

var Yellow = Code{
	Open:  openStr + "33" + "m",
	Close: openStr + "39" + "m",
}

var Blue = Code{
	Open:  openStr + "34" + "m",
	Close: openStr + "39" + "m",
}

var Magenta = Code{
	Open:  openStr + "35" + "m",
	Close: openStr + "39" + "m",
}

var Cyan = Code{
	Open:  openStr + "36" + "m",
	Close: openStr + "39" + "m",
}

var White = Code{
	Open:  openStr + "37" + "m",
	Close: openStr + "39" + "m",
}

var Gray = Code{
	Open:  openStr + "90" + "m",
	Close: openStr + "39" + "m",
}

var Grey = Gray

var BgBlack = Code{
	Open:  openStr + "40" + "m",
	Close: openStr + "49" + "m",
}

var BgRed = Code{
	Open:  openStr + "41" + "m",
	Close: openStr + "49" + "m",
}

var BgGreen = Code{
	Open:  openStr + "42" + "m",
	Close: openStr + "49" + "m",
}

var BgYellow = Code{
	Open:  openStr + "43" + "m",
	Close: openStr + "49" + "m",
}

var BgBlue = Code{
	Open:  openStr + "44" + "m",
	Close: openStr + "49" + "m",
}

var BgMagenta = Code{
	Open:  openStr + "45" + "m",
	Close: openStr + "49" + "m",
}

var BgCyan = Code{
	Open:  openStr + "46" + "m",
	Close: openStr + "49" + "m",
}

var BgWhite = Code{
	Open:  openStr + "47" + "m",
	Close: openStr + "49" + "m",
}
