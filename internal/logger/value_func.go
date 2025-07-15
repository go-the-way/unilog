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

var (
	arrayFunc0 ArrayFunc
	mapFunc0   MapFunc
)

func init() { arrayFunc0 = arrayFunc00(); mapFunc0 = mapFunc00() }

func SetArrayFunc(arrayFunc ArrayFunc) { arrayFunc0 = arrayFunc }
func SetMapFunc(mapFunc MapFunc)       { mapFunc0 = mapFunc }

type ArrayFunc func(v reflect.Value) (vv any)
type MapFunc func(v reflect.Value) (vv any)

func arrayFunc00() ArrayFunc { return arrayFunc("%v", ",") }
func mapFunc00() MapFunc     { return mapFunc("%s:%v", ",") }

func arrayFunc(format, sep string) ArrayFunc {
	return func(v reflect.Value) (vv any) {
		var arrS []string
		for i := 0; i < v.Len(); i++ {
			iv := v.Index(i)
			if iv.CanInterface() {
				arrS = append(arrS, fmt.Sprintf(format, iv.Interface()))
			}
		}
		return strings.Join(arrS, sep)
	}
}

func mapFunc(format, sep string) MapFunc {
	return func(v reflect.Value) (vv any) {
		var arrS []string
		mr := v.MapRange()
		for mr.Next() {
			mk := mr.Key()
			mv := mr.Value()
			if mk.CanInterface() && mv.CanInterface() {
				arrS = append(arrS, fmt.Sprintf(format, mk, mv.Interface()))
			}
		}
		return strings.Join(arrS, sep)
	}
}
