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

var _ expr = (*exprRef)(nil)

type exprRef struct{ expr0 string }

func newExprRef(expr0 string) *exprRef { return &exprRef{expr0} }

// Expr ref: expr
//
// e.g.: ref:Ref1.Ref2...
//
// usage:
//
//	type logger struct {
//		Name1 string `log:"Name1,ref:Ref1.Name"`
//		Name2 string `log:"Name2,ref:Ref1.Ref11.Name"`
//		Ref1  struct {
//			Name  string
//			Ref11 struct{ Name string }
//		} `log:"-"`
//	}
func (r *exprRef) Expr(ov, sv reflect.Value) (values []any) {
	expr0 := r.expr0
	oldValFunc := func() (a any) {
		v := ov
		if !strings.HasPrefix(expr0, ".") {
			expr0 = "." + expr0
		}
		for {
			dotIdx := strings.IndexByte(expr0, '.')
			if dotIdx == -1 {
				break
			}
			expr0 = expr0[dotIdx+1:]
			dotIdx2 := strings.IndexByte(expr0, '.')
			if dotIdx2 == -1 {
				if expr0 != "" {
					fv := v.FieldByName(expr0)
					if !fv.IsZero() && fv.CanInterface() {
						a = fv.Interface()
					}
				}
				break
			}
			refName := expr0[:dotIdx2]
			v = v.FieldByName(refName)
		}
		return
	}
	newValFunc := func() (a any) { return sv.Interface() }
	oldVal := oldValFunc()
	newVal := newValFunc()
	return []any{oldVal, newVal}
}
