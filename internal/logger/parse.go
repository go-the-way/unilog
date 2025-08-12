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

func GetFields(struct0 any, defaultIgnore ...bool) (fieldSlice FieldSlice) {
	v := reflect.ValueOf(struct0)
	if !v.IsValid() {
		panic("unilog: the struct value is invalid.")
	}
	if v = rv(v); v.Kind() != reflect.Struct {
		panic("unilog: the struct value is not supported.")
	}
	return getSupportedFields(v, defaultIgnore...)
}

func getSupportedFields(ov reflect.Value, defaultIgnoreS ...bool) (fieldSlice FieldSlice) {
	var defaultIgnore bool
	if len(defaultIgnoreS) > 0 {
		defaultIgnore = defaultIgnoreS[0]
	}
	ov = rv(ov)
	for i := 0; i < ov.NumField(); i++ {
		fd := ov.Type().Field(i)
		sv := rv(ov.Field(i))
		if !fd.IsExported() {
			continue
		}
		if _, supported := supportedKind[sv.Kind()]; supported {
			isStruct := sv.Kind() == reflect.Struct
			logTag, ok := fd.Tag.Lookup("log")
			if ((isStruct || defaultIgnore) && !ok) || logTag == "-" {
				continue
			}

			if isStruct && logTag == ",inner" {
				fieldSlice = append(fieldSlice, getSupportedFields(sv, defaultIgnore)...)
				continue
			}

			logName, format, expr0 := parseTag(fd, logTag)

			f := Field{Name: logName, Format: format, expr: expr0, SV: sv, OV: ov}
			if isStruct && expr0 == nil {
				f.expr = newExprFields(getSupportedFields(sv, defaultIgnore))
			}
			fieldSlice = append(fieldSlice, f)
		}
	}
	return
}

func parseTag(fd reflect.StructField, logTag string) (logName, format string, expr0 expr) {
	logName = fd.Name
	format = "%s[%v]"
	tagS := strings.Split(logTag, ",")
	for i, tag := range tagS {
		if tag = strings.TrimSpace(tag); tag == "" {
			continue
		}
		if i == 0 {
			logName = tag
		} else if strings.HasPrefix(tag, "ref:") {
			expr0 = newExprRef(tag[4:])
		} else if strings.HasPrefix(tag, "transform:") {
			expr0 = newExprTransform(tag[10:])
		} else {
			format = tag
		}
	}
	return
}

func rv(v reflect.Value) (vv reflect.Value) {
	vv = v
	for vv.Kind() == reflect.Pointer {
		vv = vv.Elem()
	}
	return
}

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
