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

package terminfo

type winsize struct {
	ws_row    uint16 /* rows, in characters */
	ws_col    uint16 /* columns, in characters */
	ws_xpixel uint16 /* horizontal size, pixels */
	ws_ypixel uint16 /* vertical size, pixels */
}

func WindowWidth() uint16 {
	ws := getwinsize()
	return ws.ws_col
}
