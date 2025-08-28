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
	"reflect"
	"strings"
)

// GetFields extracts loggable fields from a given struct, supporting custom log tags and optional default ignore behavior.
// It processes the input struct, validates it, and returns a FieldSlice containing fields to be logged.
// Panics if the input is invalid or not a struct.
func GetFields(struct0 any, defaultIgnore ...bool) (fieldSlice FieldSlice) {
	// Validate the input value and ensure it's a valid reflect.Value.
	v := reflect.ValueOf(struct0)
	if !v.IsValid() {
		panic("unilog: the struct value is invalid.")
	}

	// Dereference pointers and ensure the value is a struct.
	if v = rv(v); v.Kind() != reflect.Struct {
		panic("unilog: the struct value is not supported.")
	}

	// Delegate to getSupportedFields to process the struct fields.
	return getSupportedFields(v, defaultIgnore...)
}

// isStruct0 determines if the provided reflect.Value is a struct type.
func isStruct0(ov reflect.Value) bool {
	return ov.Kind() == reflect.Struct
}

// isArray0 determines if the provided reflect.Value is an array or slice type.
func isArray0(ov reflect.Value) bool {
	return ov.Kind() == reflect.Array || ov.Kind() == reflect.Slice
}

// getSupportedFields processes a struct or array/slice to extract fields suitable for logging.
// It respects struct tags, supports nested structs, arrays, and slices, and applies default ignore rules.
// The defaultIgnore parameter controls whether fields without a log tag are ignored by default.
func getSupportedFields(ov reflect.Value, defaultIgnoreS ...bool) (fieldSlice FieldSlice) {
	// Determine the default ignore behavior for fields without log tags.
	var defaultIgnore bool
	if len(defaultIgnoreS) > 0 {
		defaultIgnore = defaultIgnoreS[0]
	}

	// Dereference the input value to handle pointers.
	ov = rv(ov)
	numField := 0
	isStruct := isStruct0(ov)
	isArray := isArray0(ov)

	// Determine the number of fields to process based on whether it's a struct or array.
	if isStruct {
		numField = ov.NumField()
	} else if isArray {
		numField = ov.Len()
	}

	// Iterate over fields or elements.
	for i := 0; i < numField; i++ {
		if isStruct {
			// Process struct fields.
			fd := ov.Type().Field(i) // Get the struct field definition.
			sv := rv(ov.Field(i))    // Get the dereferenced field value.

			// Skip unexported fields (not accessible for reflection).
			if !fd.IsExported() {
				continue
			}

			// Skip unsupported field types based on the supportedKind map.
			if _, supported := supportedKind[sv.Kind()]; !supported {
				continue
			}

			// Check if the field is a struct.
			fieldIsStruct := isStruct0(sv)
			// Look up the "log" tag in the struct field.
			logTag, ok := fd.Tag.Lookup("log")

			// Skip fields that are structs or marked for default ignore without a log tag,
			// or explicitly ignored with a "-" tag.
			if ((fieldIsStruct || defaultIgnore) && !ok) || logTag == "-" {
				continue
			}

			// Handle inline structs, recursively processing their fields.
			if fieldIsStruct && logTag == ",inline" {
				fieldSlice = append(fieldSlice, getSupportedFields(sv, defaultIgnore)...)
				continue
			}

			// Parse the log tag to extract name, format, and expression.
			logName, format, expr0 := parseTag(fd, logTag)
			// Create a Field instance for logging.
			f := Field{Name: logName, Format: format, expr: expr0, SV: sv, OV: ov}

			// Handle nested structs or arrays/slices of structs without an explicit expression.
			if (fieldIsStruct && expr0 == nil) || ((sv.Kind() == reflect.Array || sv.Kind() == reflect.Slice) && fd.Type.Elem().Kind() == reflect.Struct) {
				f.expr = newExprFields(getSupportedFields(sv, defaultIgnore))
			}
			fieldSlice = append(fieldSlice, f)
		} else if isArray {
			// Process array or slice elements.
			ev := ov.Index(i)      // Get the element at index i.
			et := ev.Type().Kind() // Get the element's type.
			if _, supported := supportedKind[et]; !supported {
				continue
			}
			// Create a Field instance for the array element, using arrayElementFormat and recursive field extraction.
			f := Field{Format: arrayElementFormat, expr: newExprFields(getSupportedFields(ev, defaultIgnore)), SV: ev, OV: ev}
			fieldSlice = append(fieldSlice, f)
		}
	}
	return
}

// parseTag parses a struct field's log tag to extract the log name, format, and expression.
// The log tag is expected to be in the format "name,option1,option2" where options can include
// "ref:<path>", "transform:<mapping>", or a custom format string.
func parseTag(fd reflect.StructField, logTag string) (logName, format string, expr0 expr) {
	logName = fd.Name    // Default to the field name.
	format = fieldFormat // Default to the predefined field format.
	tagS := strings.Split(logTag, ",")
	for i, tag := range tagS {
		if tag = strings.TrimSpace(tag); tag == "" {
			continue
		}
		if i == 0 {
			// First tag is the log name.
			logName = tag
		} else if strings.HasPrefix(tag, "ref:") {
			// Handle reference path expression (e.g., "ref:Ref1.Ref2").
			expr0 = newExprRef(tag[4:])
		} else if strings.HasPrefix(tag, "transform:") {
			// Handle transformation expression (e.g., "transform:1->on|2->off").
			expr0 = newExprTransform(tag[10:])
		} else {
			// Any other tag is treated as the format string.
			format = tag
		}
	}
	return
}

// rv dereferences a reflect.Value until a non-pointer type is reached.
// This ensures the value is usable for reflection operations.
func rv(v reflect.Value) (vv reflect.Value) {
	vv = v
	for vv.Kind() == reflect.Pointer {
		vv = vv.Elem()
	}
	return
}

// supportedKind is a map of reflect.Kind types that are supported for logging.
// It includes basic types (bool, integers, floats, string), structs, arrays, slices, and maps.
var supportedKind = map[reflect.Kind]struct{}{
	reflect.Bool:    {},
	reflect.Int:     {},
	reflect.Int8:    {},
	reflect.Int16:   {},
	reflect.Int32:   {},
	reflect.Int64:   {},
	reflect.Uint:    {},
	reflect.Uint8:   {},
	reflect.Uint16:  {},
	reflect.Uint32:  {},
	reflect.Uint64:  {},
	reflect.Float32: {},
	reflect.Float64: {},
	reflect.Array:   {},
	reflect.Map:     {},
	reflect.Slice:   {},
	reflect.String:  {},
	reflect.Struct:  {},
}
