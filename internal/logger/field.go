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

type Field struct {
	Name, Format string
	expr
	SV, OV reflect.Value
	values []any
}

func (f Field) log() (str string) {
	f.values = append(f.values, f.Name)
	if f.expr != nil {
		f.values = append(f.values, f.Expr(f.Format, f.OV, f.SV)...)
	}
	vk := f.SV.Kind()
	switch {

	case vk >= reflect.Bool && vk <= reflect.Float64 || vk == reflect.String:
		if f.expr == nil {
			f.values = append(f.values, f.SV.Interface())
		}
		return f.basic2string()

	case vk == reflect.Struct:
		return f.basic2string()

	case vk == reflect.Array || vk == reflect.Slice:
		return f.array2string()

	case vk == reflect.Map:
		return f.map2string()
	}

	return
}

func (f Field) basic2string() string { return fmt.Sprintf(f.Format, f.values...) }

func (f Field) array2string() string {
	var arrS []string
	for i := 0; i < f.SV.Len(); i++ {
		iv := f.SV.Index(i)
		if iv.CanInterface() {
			arrS = append(arrS, fmt.Sprintf("%v", iv.Interface()))
		}
	}
	f.values = append(f.values, strings.Join(arrS, ","))
	return fmt.Sprintf(f.Format, f.values...)
}

func (f Field) map2string() string {
	var arrS []string
	mr := f.SV.MapRange()
	for mr.Next() {
		k := mr.Key()
		v := mr.Value()
		if k.CanInterface() && v.CanInterface() {
			arrS = append(arrS, fmt.Sprintf("%s:%v", k, v.Interface()))
		}
	}
	f.values = append(f.values, strings.Join(arrS, ","))
	return fmt.Sprintf(f.Format, f.values...)
}
