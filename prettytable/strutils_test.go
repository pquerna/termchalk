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

import "testing"

func TestPad(t *testing.T) {

	x := padString("foo", 7, " ", PadCenter)
	if x != "  foo  " {
		t.Fatal("foo not centered.")
	}

	x = padString("foo", 8, " ", PadCenter)
	if x != "  foo   " {
		t.Fatal("foo not centered.")
	}

	x = padString("foo", 8, " ", PadLeft)
	if x != "     foo" {
		t.Fatal("foo not left'ed.")
	}

	x = padString("foo", 8, " ", PadRight)
	if x != "foo     " {
		t.Fatal("foo not right'ed.")
	}
}
