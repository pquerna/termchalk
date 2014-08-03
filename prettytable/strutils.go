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
	"github.com/pquerna/termchalk/termwidth"
	"math"
	"strings"
)

type padDirection int

const (
	padLeft padDirection = 1 << iota
	padRight
	padCenter
)

func padString(s string, l int, padChr string, direction padDirection) string {
	slen := termwidth.Width(s)
	if l+1 >= slen {
		switch direction {
		case padLeft:
			s = strings.Repeat(padChr, l+1-slen) + s
		case padRight:
			s = s + strings.Repeat(padChr, l+1-slen)
		case padCenter:
			padlen := ((float64(l) - float64(slen)) / float64(2.0))
			right := int(math.Ceil(padlen))
			left := int(math.Floor(padlen))
			s = strings.Repeat(padChr, left) + s + strings.Repeat(padChr, right)
		}
	}
	return s
}
