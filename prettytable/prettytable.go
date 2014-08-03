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
	"bytes"
	"fmt"
	"github.com/pquerna/termchalk/termwidth"
	"sort"
	"strings"
)

type PrettyTableSorter func(n int, s InterfaceArray) sort.Interface

type PrettyTable struct {
	header      []string
	rows        [][]interface{}
	rowsStr     [][]string
	SortBy      string
	SortReverse bool
	Sorter      PrettyTableSorter
	Styler      Styler
}

func New(header []string) *PrettyTable {
	return &PrettyTable{
		rows:   make([][]interface{}, 0),
		header: header,
		Sorter: NewMungeSorter,
		Styler: &DefaultStyle{},
	}
}

func pos(needle string, haystack []string) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}

	return -1
}

func (pt *PrettyTable) AddRow(args ...interface{}) {
	pt.rows = append(pt.rows, args)
}

func intMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func cellStr(v interface{}) string {
	return fmt.Sprint(v)
}

func longestLine(s string) int {
	rv := 0
	for _, v := range strings.Split(s, "\n") {
		rv = intMax(termwidth.Width(v), rv)
	}
	return rv
}

func (pt *PrettyTable) styleData() {
	var headers = make([]string, len(pt.header))
	for i, v := range pt.header {
		headers[i] = pt.Styler.Header(i, v)
	}

	var rows = make([][]string, len(pt.rows))
	for i, row := range pt.rows {
		rows[i] = make([]string, len(row))
		for j, cell := range row {
			rows[i][j] = pt.Styler.Cell(pt.header[j], i, j, cellStr(cell))
		}
	}

	pt.header = headers
	pt.rowsStr = rows
}

func (pt *PrettyTable) colWidths() []int {
	widths := make([]int, len(pt.header))
	padding := pt.Styler.Padding() * 2
	for i, v := range pt.header {
		widths[i] = termwidth.Width(v) + padding
	}

	for _, row := range pt.rows {
		for j, cell := range row {
			if j > len(pt.header) {
				continue
			}

			s := cellStr(cell)
			widths[j] = intMax(longestLine(s)+padding, widths[j])
		}
	}

	return widths
}

func (pt *PrettyTable) String() string {
	b := bytes.Buffer{}

	if pt.SortBy != "" {
		n := pos(pt.SortBy, pt.header)
		if n != -1 {
			s := pt.Sorter(n, pt.rows)
			if pt.SortReverse {
				sort.Sort(sort.Reverse(s))
			} else {
				sort.Sort(s)
			}
		}
	}

	pt.styleData()

	colWidths := pt.colWidths()

	// TODO: refactor
	b.WriteString(pt.Styler.TopLeft())
	for i, v := range colWidths {
		b.WriteString(strings.Repeat(pt.Styler.Top(), v))
		if i < len(colWidths)-1 {
			b.WriteString(pt.Styler.TopMid())
		}
	}
	b.WriteString(pt.Styler.TopRight())
	b.WriteString("\n")

	b.WriteString(pt.Styler.Left())
	for i, v := range pt.header {
		// TODO: padding/centers
		cstr := padString(
			v,
			colWidths[i],
			" ",
			pt.Styler.Alignment())

		b.WriteString(cstr)
		if i < len(pt.header)-1 {
			b.WriteString(pt.Styler.Middle())
		}
	}
	b.WriteString(pt.Styler.Right())
	b.WriteString("\n")

	b.WriteString(pt.Styler.Left())
	for i, v := range colWidths {
		b.WriteString(strings.Repeat(pt.Styler.Mid(), v))
		if i < len(colWidths)-1 {
			b.WriteString(pt.Styler.MidMid())
		}
	}
	b.WriteString(pt.Styler.Right())
	b.WriteString("\n")

	for _, v := range pt.rowsStr {
		// TODO: padding/centers
		// TOOD: multi-line strings
		b.WriteString(pt.Styler.Left())
		for i, cell := range v {
			cstr := padString(
				cell,
				colWidths[i],
				" ",
				pt.Styler.Alignment())

			b.WriteString(cstr)
			if i < len(v)-1 {
				b.WriteString(pt.Styler.Middle())
			}
		}
		b.WriteString(pt.Styler.Right())
		b.WriteString("\n")
	}

	// TODO: refactor
	b.WriteString(pt.Styler.BottomLeft())
	for i, v := range colWidths {
		b.WriteString(strings.Repeat(pt.Styler.Bottom(), v))
		if i < len(colWidths)-1 {
			b.WriteString(pt.Styler.BottomMid())
		}
	}
	b.WriteString(pt.Styler.BottomRight())
	b.WriteString("\n")

	return b.String()
}
