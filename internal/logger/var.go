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

import "reflect"

func GetFieldsFromTag(struct0 any, defaultIgnore ...bool) (fieldSlice FieldSlice) {
	v := reflect.ValueOf(struct0)
	if !v.IsValid() {
		panic("unilog: the struct value is invalid.")
	}
	if v = rv(v); v.Kind() != reflect.Struct {
		panic("unilog: the struct value is not supported.")
	}
	return getSupportedFields(v, defaultIgnore...)
}

func getSupportedFields(v reflect.Value, defaultIgnoreS ...bool) (fieldSlice FieldSlice) {
	var defaultIgnore bool
	if len(defaultIgnoreS) > 0 {
		defaultIgnore = defaultIgnoreS[0]
	}
	for i := 0; i < v.NumField(); i++ {
		fd := v.Type().Field(i)
		fdv := v.Field(i)
		fdv = rv(fdv)
		if _, supported := supportedKind[fdv.Kind()]; supported {
			isStruct := fdv.Kind() == reflect.Struct
			logName, ok := fd.Tag.Lookup("log")
			if ((isStruct || defaultIgnore) && !ok) || logName == "-" {
				continue
			}
			if logName == "" {
				logName = fd.Name
			}
			if fdv.Kind() == reflect.Struct {
				switch logName {
				default:
					fieldSlice = append(fieldSlice, Field{Name: logName, Value: getSupportedFields(fdv)})
				case ",inner":
					fieldSlice = append(fieldSlice, getSupportedFields(fdv)...)
				}
			} else if fdv.CanInterface() {
				fieldSlice = append(fieldSlice, Field{Name: logName, Value: fdv.Interface()})
			}
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
	reflect.Bool:      {},
	reflect.Int:       {},
	reflect.Int8:      {},
	reflect.Int16:     {},
	reflect.Int32:     {},
	reflect.Int64:     {},
	reflect.Uint:      {},
	reflect.Uint8:     {},
	reflect.Uint16:    {},
	reflect.Uint32:    {},
	reflect.Uint64:    {},
	reflect.Float32:   {},
	reflect.Float64:   {},
	reflect.Array:     {},
	reflect.Interface: {},
	reflect.Map:       {},
	reflect.Slice:     {},
	reflect.String:    {},
	reflect.Struct:    {},
}
