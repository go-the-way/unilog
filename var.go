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

package unilog

import (
	"github.com/go-the-way/unilog/internal/db"
	"github.com/go-the-way/unilog/internal/logger"
	"github.com/go-the-way/unilog/internal/services/log"
)

// Package-level variables for database and pagination configuration.
// These aliases provide convenient access to functions from the internal db package.
var (
	// SetDB assigns the global database instance for use in logging operations.
	SetDB = db.SetDB
	// AutoMigrate performs automatic migration of the database schema for logging-related models.
	AutoMigrate = db.AutoMigrate
	// SetPagination configures the pagination function used for querying log entries.
	SetPagination = db.SetPagination
)

// Package-level variables for logger configuration.
// These aliases provide access to functions for configuring field extraction and formatting.
var (
	// GetFields extracts loggable fields from a struct for logging purposes.
	GetFields = logger.GetFields
	// SetArrayFunc sets a custom function for formatting array/slice values in logs.
	SetArrayFunc = logger.SetArrayFunc
	// SetMapFunc sets a custom function for formatting map values in logs.
	SetMapFunc = logger.SetMapFunc
	// SetFieldFormat sets a custom format string for logging fields.
	SetFieldFormat = logger.SetFieldFormat
	// SetArrayElementFormat sets a custom format string for array/slice elements in logs.
	SetArrayElementFormat = logger.SetArrayElementFormat
	// SetFieldJoinSep sets a custom separator for joining multiple field log strings.
	SetFieldJoinSep = logger.SetFieldJoinSep
)

// Package-level variables for log service operations.
// These aliases provide access to functions for managing log entries.
var (
	// LogGetPage retrieves a paginated list of log entries.
	LogGetPage = log.GetPage
	// LogGet retrieves a single log entry by its identifier or criteria.
	LogGet = log.Get
	// LogAdd creates a new log entry in the database.
	LogAdd = log.Add
	// LogUpdate modifies an existing log entry in the database.
	LogUpdate = log.Update
	// LogDelete removes a log entry from the database.
	LogDelete = log.Delete
)
