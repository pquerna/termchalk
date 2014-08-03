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

type TableStyler interface {
	Top() string
	TopMid() string
	TopLeft() string
	TopRight() string
	Bottom() string
	BottomMid() string
	BottomLeft() string
	BottomRight() string
	Left() string
	LeftMid() string
	Right() string
	RightMid() string
	Mid() string
	MidMid() string
	Middle() string
	Truncate() string
	Padding() int
	Alignment() PadDirection
}

type DefaultStyle struct {
	UnicodeBox
}

type UnicodeBox struct {
}

type ClassicBox struct {
	UnicodeBox
}

func (ds *UnicodeBox) Alignment() PadDirection {
	return PadCenter
}

func (ds *UnicodeBox) Padding() int {
	return 2
}

func (ds *UnicodeBox) Top() string {
	return "─"
}
func (ds *UnicodeBox) TopMid() string {
	return "┬"
}
func (ds *UnicodeBox) TopLeft() string {
	return "┌"
}
func (ds *UnicodeBox) TopRight() string {
	return "┐"
}

func (ds *UnicodeBox) Bottom() string {
	return "─"
}
func (ds *UnicodeBox) BottomMid() string {
	return "┴"
}
func (ds *UnicodeBox) BottomLeft() string {
	return "└"
}
func (ds *UnicodeBox) BottomRight() string {
	return "┘"
}

func (ds *UnicodeBox) Left() string {
	return "│"
}
func (ds *UnicodeBox) LeftMid() string {
	return "├"
}

func (ds *UnicodeBox) Right() string {
	return "│"
}
func (ds *UnicodeBox) RightMid() string {
	return "├"
}

func (ds *UnicodeBox) Mid() string {
	return "─"
}
func (ds *UnicodeBox) MidMid() string {
	return "┼"
}

func (ds *UnicodeBox) Middle() string {
	return "│"
}

func (ds *UnicodeBox) Truncate() string {
	return "…"
}

func (ds *ClassicBox) Top() string {
	return "-"
}
func (ds *ClassicBox) TopMid() string {
	return "+"
}
func (ds *ClassicBox) TopLeft() string {
	return "+"
}
func (ds *ClassicBox) TopRight() string {
	return "+"
}

func (ds *ClassicBox) Bottom() string {
	return "-"
}
func (ds *ClassicBox) BottomMid() string {
	return "+"
}
func (ds *ClassicBox) BottomLeft() string {
	return "+"
}
func (ds *ClassicBox) BottomRight() string {
	return "+"
}

func (ds *ClassicBox) Left() string {
	return "|"
}
func (ds *ClassicBox) LeftMid() string {
	return "+"
}

func (ds *ClassicBox) Right() string {
	return "|"
}
func (ds *ClassicBox) RightMid() string {
	return "+"
}

func (ds *ClassicBox) Mid() string {
	return "-"
}
func (ds *ClassicBox) MidMid() string {
	return "+"
}

func (ds *ClassicBox) Middle() string {
	return "|"
}

func (ds *ClassicBox) Truncate() string {
	return "..."
}
