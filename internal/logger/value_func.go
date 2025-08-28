// Copyright 2025 unilog Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"fmt"
	"reflect"
	"strings"
)

// Package-level variables for configuring logging behavior.
var (
	// arrayFunc0 is the default function for formatting array/slice values.
	arrayFunc0 ArrayFunc
	// mapFunc0 is the default function for formatting map values.
	mapFunc0 MapFunc
	// fieldFormat is the default format string for logging fields (e.g., "%s[%v]").
	fieldFormat = "%s[%v]"
	// arrayElementFormat is the default format string for array/slice elements (e.g., "{%v}").
	arrayElementFormat = "{%v}"
	// fieldJoinSep is the default separator for joining multiple field log strings (e.g., ",").
	fieldJoinSep = ","
)

// init initializes the default array and map formatting functions.
func init() {
	arrayFunc0 = arrayFunc00()
	mapFunc0 = mapFunc00()
}

// SetArrayFunc sets a custom function for formatting array/slice values.
func SetArrayFunc(arrayFunc ArrayFunc) {
	arrayFunc0 = arrayFunc
}

// SetMapFunc sets a custom function for formatting map values.
func SetMapFunc(mapFunc MapFunc) {
	mapFunc0 = mapFunc
}

// SetFieldFormat sets a custom format string for logging fields.
func SetFieldFormat(format string) {
	fieldFormat = format
}

// SetArrayElementFormat sets a custom format string for array/slice elements.
func SetArrayElementFormat(format string) {
	arrayElementFormat = format
}

// SetFieldJoinSep sets a custom separator for joining multiple field log strings.
func SetFieldJoinSep(sep string) {
	fieldJoinSep = sep
}

// ArrayFunc defines a function type for formatting array/slice values into a single value.
type ArrayFunc func(v reflect.Value) (vv any)

// MapFunc defines a function type for formatting map values into a single value.
type MapFunc func(v reflect.Value) (vv any)

// arrayFunc00 returns the default ArrayFunc that formats array/slice elements with a standard format and separator.
func arrayFunc00() ArrayFunc {
	return arrayFunc("%v", ",")
}

// mapFunc00 returns the default MapFunc that formats map key-value pairs with a standard format and separator.
func mapFunc00() MapFunc {
	return mapFunc("%s:%v", ",")
}

// arrayFunc creates an ArrayFunc that formats array/slice elements using the provided format and separator.
// Each element is formatted according to the format string and joined with the separator.
func arrayFunc(format, sep string) ArrayFunc {
	return func(v reflect.Value) (vv any) {
		var arrS []string
		// Iterate over the array/slice elements.
		for i := 0; i < v.Len(); i++ {
			iv := v.Index(i)
			// Only include elements that can be interfaced.
			if iv.CanInterface() {
				arrS = append(arrS, fmt.Sprintf(format, iv.Interface()))
			}
		}
		// Join the formatted elements with the separator.
		return strings.Join(arrS, sep)
	}
}

// mapFunc creates a MapFunc that formats map key-value pairs using the provided format and separator.
// Each key-value pair is formatted according to the format string and joined with the separator.
func mapFunc(format, sep string) MapFunc {
	return func(v reflect.Value) (vv any) {
		var arrS []string
		// Iterate over the map's key-value pairs.
		mr := v.MapRange()
		for mr.Next() {
			mk := mr.Key()
			mv := mr.Value()
			// Only include pairs where both key and value can be interfaced.
			if mk.CanInterface() && mv.CanInterface() {
				arrS = append(arrS, fmt.Sprintf(format, mk, mv.Interface()))
			}
		}
		// Join the formatted pairs with the separator.
		return strings.Join(arrS, sep)
	}
}
