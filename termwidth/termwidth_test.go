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
	"github.com/pquerna/termchalk/ansistyle"
	"testing"
)

func TestSimple(t *testing.T) {

	if Width("foo") != 3 {
		t.Fatal("foo should be 3 wide.")
	}

	if Width("\uFF26oo") != 4 {
		t.Fatal("\uFF26oo should be 4 wide.")
	}

	s := ansistyle.Red.Open + "Red!" + ansistyle.Red.Close
	if Width(s) != 4 {
		t.Fatal(s + " should be 4 wide, even with escape codes.")
	}
}
