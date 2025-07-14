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

var _ expr = (*exprTransform)(nil)

type exprTransform struct{ expr0 string }

func newExprTransform(expr0 string) *exprTransform { return &exprTransform{expr0} }

// Expr transform: expr
//
// e.g.: transform:1->on|2->off
//
// usage:
//
//	type logger struct {
//		Status byte `log:"Status,transform:1->on|2->off`
//	}
func (t *exprTransform) Expr(_, sv reflect.Value) (values []any) {
	valS := strings.Split(t.expr0, "|")
	var m = map[string]any{}
	for _, val := range valS {
		vv := strings.Split(val, "->")
		if len(vv) > 1 {
			k := strings.TrimSpace(vv[0])
			v := strings.TrimSpace(vv[1])
			m[k] = v
		}
	}
	k := fmt.Sprintf("%v", sv.Interface())
	sa := m[k]
	return []any{sa}
}
