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

// Package unilog provides a flexible and extensible logging utility for Go applications.
// It supports structured logging of Go structs with customizable formats, field references,
// value transformations, and user/client metadata integration.
//
// The package uses struct field tags to define logging behavior, allowing customization of
// field names, output formats, references to other fields, and value transformations.
// It supports a wide range of data types, including basic types (integers, floats, strings),
// composite types (maps, arrays, slices), and nested structs, with recursive logging capabilities.
//
// Key features include:
//   - Structured logging with customizable field formats via `log` tags
//   - Support for field references (`ref`) to include related data
//   - Value transformation (`transform`) for human-readable output
//   - Recursive logging of nested structs and arrays/slices
//   - Integration of user data and client IP through the Logger interface
//   - Extensible callback mechanism for custom log processing
//   - Configurable formatting for arrays, maps, and field separators
//
// The primary entry points are the GetFields function, which extracts loggable fields from a struct,
// and the Callback function, which generates a logging function for types implementing the Logger interface.
// For detailed usage and examples, see the README.md file in the repository.
package unilog
