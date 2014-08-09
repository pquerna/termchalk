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

package prettytable

import (
	_ "github.com/kr/text"
	"github.com/pquerna/termchalk/runewidth"
	"math"
	"strings"
)

type PadDirection int

const (
	PadLeft PadDirection = 1 << iota
	PadRight
	PadCenter
)

func padString(s string, l int, padChr string, direction PadDirection) string {
	slen := runewidth.Width(s)
	if l+1 >= slen {
		switch direction {
		case PadLeft:
			s = strings.Repeat(padChr, l-slen) + s
		case PadRight:
			s = s + strings.Repeat(padChr, l-slen)
		case PadCenter:
			padlen := ((float64(l) - float64(slen)) / float64(2.0))
			right := int(math.Ceil(padlen))
			left := int(math.Floor(padlen))
			s = strings.Repeat(padChr, left) + s + strings.Repeat(padChr, right)
		}
	}
	return s
}
