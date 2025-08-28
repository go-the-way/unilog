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

// Ensure exprRef implements the expr interface.
var _ expr = (*exprRef)(nil)

// exprRef is a struct that holds a reference expression string for accessing nested fields.
type exprRef struct {
	expr0 string // expr0 is the reference path (e.g., "Ref1.Ref2").
}

// newExprRef creates a new exprRef instance with the provided reference expression.
func newExprRef(expr0 string) *exprRef {
	return &exprRef{expr0}
}

// Expr generates a list of values based on a format string and reflect values, resolving a reference path.
// The reference path (e.g., "Ref1.Ref2") navigates through nested struct fields.
// The format string determines the output structure:
//   - For 2 placeholders (e.g., "%s[%s]"), returns [reference value].
//   - For 3 placeholders (e.g., "%s[%v=>%s]"), returns [reference value, original value].
func (r *exprRef) Expr(format string, ov, sv reflect.Value) (values []any) {
	expr0 := r.expr0

	// refValFunc resolves the reference path to extract the target field value.
	refValFunc := func() (a any) {
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
				if expr0 != "" && v.Kind() == reflect.Struct {
					fv := v.FieldByName(expr0)
					if fv.IsValid() && fv.CanInterface() {
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

	// Count placeholders in the format string to determine the output structure.
	ftc := strings.Count(format, "%")
	switch ftc {
	default:
		// Invalid number of placeholders, return empty slice.
		return
	case 2:
		// Format expects 2 placeholders (e.g., "%s[%s]").
		// Returns the reference value.
		return []any{refValFunc()}
	case 3:
		// Format expects 3 placeholders (e.g., "%s[%v=>%s]").
		// Returns the reference value and the original value.
		return []any{refValFunc(), sv.Interface()}
	}
}
