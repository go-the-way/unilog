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
)

type Field struct {
	Name, Format string
	expr
	SV, OV reflect.Value
}

func (f Field) log() (str string) {
	var values []any
	values = append(values, f.Name)
	values = append(values, f.eval()...)
	return fmt.Sprintf(f.Format, values...)
}

func (f Field) eval() (values []any) {
	if f.expr != nil {
		return f.Expr(f.Format, f.OV, f.SV)
	}
	return []any{f.value()}
}

func (f Field) value() (v any) {
	vk := f.SV.Kind()
	switch {
	case vk >= reflect.Bool && vk <= reflect.Float64 || vk == reflect.String || vk == reflect.Struct:
		return f.SV.Interface()

	case vk == reflect.Array || vk == reflect.Slice:
		return arrayFunc0(f.SV)

	case vk == reflect.Map:
		return mapFunc0(f.SV)
	}

	return
}
