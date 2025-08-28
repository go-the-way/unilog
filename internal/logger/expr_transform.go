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

// Ensure exprTransform implements the expr interface.
var _ expr = (*exprTransform)(nil)

// exprTransform is a struct that holds a transformation expression for mapping values.
type exprTransform struct {
	expr0 string // expr0 is the transformation mapping (e.g., "1->on|2->off").
}

// newExprTransform creates a new exprTransform instance with the provided transformation expression.
func newExprTransform(expr0 string) *exprTransform {
	return &exprTransform{expr0}
}

// Expr transforms a value based on a mapping defined in expr0 (e.g., "1->on|2->off").
// The transformation maps an original value to a new value based on the provided expression.
// The format string determines the output structure:
//   - For 2 placeholders (e.g., "%s[%s]"), returns [transformed value].
//   - For 3 placeholders (e.g., "%s[%v=>%s]"), returns [transformed value, original value].
func (t *exprTransform) Expr(format string, _, sv reflect.Value) (values []any) {
	// Split the transformation expression into individual mappings.
	valS := strings.Split(t.expr0, "|")
	var m = map[string]any{}
	for _, val := range valS {
		// Split each mapping into key and value (e.g., "1->on" into "1" and "on").
		vv := strings.Split(val, "->")
		if len(vv) > 1 {
			k := strings.TrimSpace(vv[0])
			v := strings.TrimSpace(vv[1])
			m[k] = v
		}
	}

	// Get the original value and convert it to a string key.
	svv := sv.Interface()
	k := fmt.Sprintf("%v", svv)
	sa := m[k]

	// Count placeholders in the format string to determine the output structure.
	ftc := strings.Count(format, "%")
	switch ftc {
	default:
		// Invalid number of placeholders, return empty slice.
		return
	case 2:
		// Format expects 2 placeholders (e.g., "%s[%s]").
		// Returns the transformed value.
		return []any{sa}
	case 3:
		// Format expects 3 placeholders (e.g., "%s[%v=>%s]").
		// Returns the transformed value and the original value.
		return []any{sa, svv}
	}
}
