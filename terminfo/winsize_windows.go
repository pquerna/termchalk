// +build windows
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

import (
	"syscall"
	"unsafe"
)

// Windows get window size code from Konstantin Khomouto go-nuts thread:
// https://groups.google.com/d/msg/golang-nuts/lQRDFwhS650/ZH7GMEj-h2gJ

func getwinsize() winsize {
	ws := winsize{}

	handle, err := syscall.Open("CONOUT$", syscall.O_RDONLY, 0)
	if err != nil {
		return ws
	}
	defer syscall.Close(handle)

	sb, err := getConsoleScreenBufferInfo(hCon)
	if err != nil {
		return ws
	}

	ws.ws_row = uint16(sb.size.x)
	ws.ws_col = uint16(sb.size.y)
	return
}

var modkernel32 = syscall.NewLazyDLL("kernel32.dll")
var procGetConScrBufInfo = modkernel32.NewProc("GetConsoleScreenBufferInfo")

type coord struct {
	x int16
	y int16
}

type smallRect struct {
	left   int16
	top    int16
	right  int16
	bottom int16
}

type consoleScreenBuffer struct {
	size       coord
	cursorPos  coord
	attrs      int32
	window     smallRect
	maxWinSize coord
}

func getConsoleScreenBufferInfo(hCon syscall.Handle) (sb consoleScreenBuffer, err error) {
	rc, _, ec := syscall.Syscall(procGetConScrBufInfo.Addr(), 2,
		uintptr(hCon), uintptr(unsafe.Pointer(&sb)), 0)
	if rc == 0 {
		err = syscall.Errno(ec)
	}
	return
}
