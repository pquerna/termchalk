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
}

type DefaultStyle struct {
}

func (ds *DefaultStyle) Top() string {
	return "─"
}
func (ds *DefaultStyle) TopMid() string {
	return "┬"
}
func (ds *DefaultStyle) TopLeft() string {
	return "┌"
}
func (ds *DefaultStyle) TopRight() string {
	return "┐"
}

func (ds *DefaultStyle) Bottom() string {
	return "─"
}
func (ds *DefaultStyle) BottomMid() string {
	return "┴"
}
func (ds *DefaultStyle) BottomLeft() string {
	return "└"
}
func (ds *DefaultStyle) BottomRight() string {
	return "┘"
}

func (ds *DefaultStyle) Left() string {
	return "│"
}
func (ds *DefaultStyle) LeftMid() string {
	return "├"
}

func (ds *DefaultStyle) Right() string {
	return "│"
}
func (ds *DefaultStyle) RightMid() string {
	return "├"
}

func (ds *DefaultStyle) Mid() string {
	return "─"
}
func (ds *DefaultStyle) MidMid() string {
	return "┼"
}

func (ds *DefaultStyle) Middle() string {
	return "│"
}

func (ds *DefaultStyle) Truncate() string {
	return "…"
}
