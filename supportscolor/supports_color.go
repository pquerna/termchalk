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

package supportscolor

import (
	"github.com/mattn/go-isatty"
	"os"
	"regexp"
)

func stringInArgs(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}

		// if we see a --, assume we are chain calling into
		// another program, and we don't want to assume we can
		// interpert it.
		if v == "--" {
			break
		}
	}

	return false
}

func termSupportsColor(term string) bool {
	if term == "dumb" {
		return false
	}

	rv, err := regexp.MatchString("(^screen|^xterm|^vt100|color|ansi|cygwin|linux)", term)
	if err != nil {
		return false
	}

	return rv
}

func SupportsColor() bool {
	if stringInArgs("--no-color", os.Args) {
		return false
	}

	if stringInArgs("--color", os.Args) {
		return true
	}

	if !isatty.IsTerminal(os.Stdout.Fd()) {
		return false
	}

	if os.Getenv("COLORTERM") != "" {
		return true
	}

	var term = os.Getenv("TERM")

	return termSupportsColor(term)
}
