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
	"github.com/go-the-way/unilog/internal/models"
	"github.com/go-the-way/unilog/internal/services/log"
)

// Type aliases for database and model-related types.
// These provide convenient access to types defined in the internal db and models packages.
type (
	// Log represents a log entry model, aliased from the models package.
	Log = models.Log
	// PaginationFunc defines a function for handling pagination in database queries, aliased from the db package.
	PaginationFunc = db.PaginationFunc
)

// Type aliases for logger-related interfaces and types.
// These provide access to logging-related interfaces and structs from the logger package.
type (
	// Logger is an interface for comprehensive logging, aliased from the logger package.
	Logger = logger.Logger
	// LogInfo defines methods for retrieving log name and fields, aliased from the logger package.
	LogInfo = logger.LogInfo
	// LogUserClientIP combines user and client IP logging interfaces, aliased from the logger package.
	LogUserClientIP = logger.LogUserClientIP
	// UC is a shorthand alias for LogUserClientIP, aliased from the logger package.
	UC = logger.LogUserClientIP
	// LogUser defines a method for retrieving user data for logging, aliased from the logger package.
	LogUser = logger.LogUser
	// LogClientIP defines a method for retrieving the client IP address, aliased from the logger package.
	LogClientIP = logger.LogClientIP
	// Userdata represents user information for logging, aliased from the logger package.
	Userdata = logger.Userdata
	// Field represents a single log field with formatting and expression, aliased from the logger package.
	Field = logger.Field
	// FieldSlice is a slice of Field structs for logging multiple fields, aliased from the logger package.
	FieldSlice = logger.FieldSlice
)

// Type aliases for log service request and response types.
// These provide access to request and response structs used in log service operations.
type (
	// LogGetPageReq represents a request for retrieving paginated log entries, aliased from the log package.
	LogGetPageReq = log.GetPageReq
	// LogIdReq represents a request for a log entry by ID, aliased from the log package.
	LogIdReq = log.IdReq
	// LogGetReq represents a request for retrieving a log entry, aliased from the log package.
	LogGetReq = log.GetReq
	// LogAddReq represents a request for adding a new log entry, aliased from the log package.
	LogAddReq = log.AddReq
	// LogUpdateReq represents a request for updating an existing log entry, aliased from the log package.
	LogUpdateReq = log.UpdateReq
	// LogDeleteReq represents a request for deleting a log entry, alliased from the log package.
	LogDeleteReq = log.DeleteReq
	// LogGetPageResp represents the response for paginated log entries, aliased from the log package.
	LogGetPageResp = log.GetPageResp
	// LogGetResp represents the response for a single log entry retrieval, aliased from the log package.
	LogGetResp = log.GetResp
)
