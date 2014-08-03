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
	"fmt"
	"reflect"
	"sort"
)

type InterfaceArray [][]interface{}

func (s InterfaceArray) Len() int      { return len(s) }
func (s InterfaceArray) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type MungeSorter struct {
	InterfaceArray
	n int
}

func NewMungeSorter(n int, s InterfaceArray) sort.Interface {
	return &MungeSorter{
		s, n,
	}
}

func isNumeric(v interface{}) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64:
		return true
	}
	return false
}

func toFloaty(v interface{}) float64 {
	// TODO: this seems excessively dumb.
	switch vx := v.(type) {
	case int:
		return float64(vx)
	case int8:
		return float64(vx)
	case int16:
		return float64(vx)
	case int32:
		return float64(vx)
	case int64:
		return float64(vx)
	case uint:
		return float64(vx)
	case uint8:
		return float64(vx)
	case uint16:
		return float64(vx)
	case uint32:
		return float64(vx)
	case uint64:
		return float64(vx)
	case uintptr:
		return float64(vx)
	case float32:
		return float64(vx)
	case float64:
		return vx
	}

	return float64(0.0)
}

func toStringy(v interface{}) string {
	if str, ok := v.(string); ok {
		return str
	}

	return fmt.Sprint(v)
}

func (s MungeSorter) Less(i int, j int) bool {
	if isNumeric(s.InterfaceArray[i][s.n]) && isNumeric(s.InterfaceArray[j][s.n]) {
		ifl := toFloaty(s.InterfaceArray[i][s.n])
		jfl := toFloaty(s.InterfaceArray[j][s.n])
		return ifl < jfl
	} else {
		is := toStringy(s.InterfaceArray[i][s.n])
		js := toStringy(s.InterfaceArray[j][s.n])
		return is < js
	}
}
