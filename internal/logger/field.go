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

func (f Field) string() (str string) {
	vv := reflect.ValueOf(f.Value)
	vk := vv.Kind()
	switch {
	case vk >= reflect.Bool && vk <= reflect.Float64 || vk == reflect.Interface || vk == reflect.String:
		return f.basic2string()

	case vk == reflect.Array || vk == reflect.Slice:
		if fs0, ok := f.Value.(FieldSlice); ok {
			return (Field{Name: f.Name, Value: fs0.String()}).basic2string()
		} else {
			return f.array2string(vv)
		}

	case vk == reflect.Map:
		return f.map2string(vv)
	}

	return
}

func (f Field) basic2string() string { return fmt.Sprintf("%s[%v]", f.Name, f.Value) }

func (f Field) array2string(v reflect.Value) string {
	var arrS []string
	for i := 0; i < v.Len(); i++ {
		iv := v.Index(i)
		if iv.CanInterface() {
			arrS = append(arrS, fmt.Sprintf("%v", iv.Interface()))
		}
	}
	return fmt.Sprintf("%s[%v]", f.Name, strings.Join(arrS, ","))
}

func (f Field) map2string(v0 reflect.Value) string {
	var arrS []string
	mr := v0.MapRange()
	for mr.Next() {
		k := mr.Key()
		v := mr.Value()
		if k.CanInterface() && v.CanInterface() {
			arrS = append(arrS, fmt.Sprintf("%s:%v", k, v.Interface()))
		}
	}
	return fmt.Sprintf("%s[%v]", f.Name, strings.Join(arrS, ","))
}
