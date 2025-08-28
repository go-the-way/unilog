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

import "strings"

// FieldSlice is a slice of Field structs for logging multiple fields.
type FieldSlice []Field

// Log generates a concatenated string of log entries from all fields in the slice, joined by a separator.
func (fs FieldSlice) Log() string {
	var strS []string
	for _, f := range fs {
		if fStr := f.log(); fStr != "" {
			strS = append(strS, fStr)
		}
	}
	return strings.Join(strS, fieldJoinSep)
}
