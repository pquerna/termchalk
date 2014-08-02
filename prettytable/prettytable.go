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
)

type PrettyTableSorter func(n int, s InterfaceArray) sort.Interface

type PrettyTable struct {
	header      []string
	rows        [][]interface{}
	SortBy      string
	SortReverse bool
	Sorter      PrettyTableSorter
	Styler      TableStyler
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

func (pt *PrettyTable) colWidths() []int {
	headerWidths := make([]int, len(pt.header))
	for i, v := range pt.header {
		headerWidths[i] = termwidth.Width(v)
	}

	widths := make([]int, len(pt.header))
	return widths
}

func (pt *PrettyTable) String() string {
	b := bytes.Buffer{}

	_ = pt.colWidths()

	b.WriteString(fmt.Sprint(pt.header))
	b.WriteString("\n")

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

	for _, v := range pt.rows {
		b.WriteString(fmt.Sprint(v))
		b.WriteString("\n")
	}
	return b.String()
}
