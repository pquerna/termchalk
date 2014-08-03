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

package termwidth

import (
	"regexp"
	"strings"
)

var escRegex *regexp.Regexp

func init() {
	escRegex = regexp.MustCompile("\u001b\\[(?:[0-9]{1,3}(?:;[0-9]{1,3})*)?[m|K]")
}

// Determine the width of a string as rendered in the terminal.
func Width(s string) int {
	var rv int
	s = escRegex.ReplaceAllString(s, "")
	rstr := strings.NewReader(s)
	for {
		r, _, err := rstr.ReadRune()
		if err != nil {
			break
		}

		rv += RuneWdith(r)
	}

	return rv
}

// Best effort determine how wide a Rune will be
// rendered on screen.
//
// Ported from prettytable._char_block_width: https://github.com/dprince/python-prettytable/blob/master/prettytable.py
func RuneWdith(r rune) int {
	// Basic Latin, which is probably the most common case
	if 0x0021 <= r && r <= 0x007e {
		return 1
	}

	// Chinese, Japanese, Korean (common)
	if 0x4e00 <= r && r <= 0x9fff {
		return 2
	}

	// Hangul
	if 0xac00 <= r && r <= 0xd7af {
		return 2
	}

	// Hiragana and Katakana
	if 0x3040 <= r && r <= 0x309f {
		return 2
	}
	if 0x30a0 <= r && r <= 0x30ff {
		return 2
	}

	// Full-width Latin characters
	if 0xff01 <= r && r <= 0xff60 {
		return 2
	}

	// CJK punctuation
	if 0x3000 <= r && r <= 0x303e {
		return 2
	}

	// Backspace and delete
	if 0x0008 == r || 0x007f == r {
		return -1
	}

	// Other control characters
	if 0x0000 <= r && r <= 0x001f {
		return 0
	}

	// Take a guess.
	return 1
}
