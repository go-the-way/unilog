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

type PaginationFunc = db.PaginationFunc

type (
	Logger           = logger.Logger
	Info             = logger.Info
	UserClientIP     = logger.UserClientIP
	UC               = UserClientIP
	User             = logger.User
	ClientIP         = logger.ClientIP
	Userdata         = logger.Userdata
	Field            = logger.Field
	FieldSlice       = logger.FieldSlice
	FieldSliceOption = logger.FieldSliceOption
)

type Log = models.Log

type (
	LogGetPageReq = log.GetPageReq
	LogGetReq     = log.GetReq
	LogAddReq     = log.AddReq
	LogUpdateReq  = log.UpdateReq
	LogDeleteReq  = log.DeleteReq

	LogGetPageResp = log.GetPageResp
	LogGetResp     = log.GetResp
)
