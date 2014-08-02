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

import "testing"

func TestTerminalGuessing(t *testing.T) {

	if termSupportsColor("dumb") != false {
		t.Fatal("forced TERM=dumb, but SupportsColor still returned true")
	}

	if termSupportsColor("randomstuff") != false {
		t.Fatal("forced TERM=randomstuff, but SupportsColor still returned true")
	}

	if termSupportsColor("xterm-color") == false {
		t.Fatal("forced TERM=xterm-color, but SupportsColor still returned true")
	}
}

func TestStringInArgs(t *testing.T) {
	if stringInArgs("--foo", []string{"foo", "bar"}) != false {
		t.Fatal("Unexpected success in finding --foo")
	}

	if stringInArgs("--foo", []string{"foo", "--", "--foo"}) != false {
		t.Fatal("Didn't stop after --")
	}
}
