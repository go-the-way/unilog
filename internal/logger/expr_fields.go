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

// Ensure exprFields implements the expr interface.
var _ expr = (*exprFields)(nil)

// exprFields is a struct that wraps a FieldSlice to implement the expr interface.
type exprFields struct {
	FieldSlice
}

// newExprFields creates a new exprFields instance with the provided FieldSlice.
func newExprFields(fs FieldSlice) *exprFields {
	return &exprFields{fs}
}

// Expr generates a list of values by calling Log() on the embedded FieldSlice.
func (fs exprFields) Expr(_ string, _, _ reflect.Value) (values []any) {
	return []any{fs.Log()}
}
