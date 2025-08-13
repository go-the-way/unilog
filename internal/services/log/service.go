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

package log

import (
	"errors"
	"fmt"

	"github.com/go-the-way/unilog/internal/db"
	"github.com/go-the-way/unilog/internal/models"
	"github.com/go-the-way/unilog/internal/pkg"
	"github.com/go-the-way/unilog/internal/services/base"
)

type service struct{}

func (s *service) GetPage(req GetPageReq) (resp GetPageResp, err error) {
	q := db.GetDB().Model(new(models.Log))
	pkg.IfGt0Func(req.Id, func() { q.Where("id=?", req.Id) })
	pkg.IfGt0Func(req.UserId, func() { q.Where("user_id=?", req.UserId) })
	pkg.IfNotEmptyFunc(req.UserName, func() { q.Where("user_name like concat('%',?,'%')", req.UserName) })
	pkg.IfNotEmptyFunc(req.ClientIP, func() { q.Where("client_ip like concat('%',?,'%')", req.ClientIP) })
	pkg.IfNotEmptyFunc(req.Type1, func() { q.Where("type1 like concat('%',?,'%')", req.Type1) })
	pkg.IfNotEmptyFunc(req.Type2, func() { q.Where("type2 like concat('%',?,'%')", req.Type2) })
	pkg.IfNotEmptyFunc(req.Type3, func() { q.Where("type3 like concat('%',?,'%')", req.Type3) })
	pkg.IfNotEmptyFunc(req.Content, func() { q.Where("content like concat('%',?,'%')", req.Content) })
	pkg.IfNotEmptyFunc(req.CreateTime1, func() { q.Where("create_time>=concat(?,' 00:00:00')", req.CreateTime1) })
	pkg.IfNotEmptyFunc(req.CreateTime2, func() { q.Where("create_time<=concat(?,' 23:59:59')", req.CreateTime2) })
	pkg.IfNotEmptyFunc(req.UpdateTime1, func() { q.Where("update_time>=concat(?,' 00:00:00')", req.UpdateTime1) })
	pkg.IfNotEmptyFunc(req.UpdateTime2, func() { q.Where("update_time<=concat(?,' 23:59:59')", req.UpdateTime2) })
	if req.OrderBy != "" {
		q.Order(req.OrderBy)
	}
	resp.List = make([]models.Log, 0)
	return base.Return(resp, db.GetPagination()(q, req.Page, req.Limit, &resp.Total, &resp.List))
}

func (s *service) Get(req GetReq) (resp GetResp, err error) {
	var list []models.Log
	if err = db.GetDB().Model(new(models.Log)).Where("id=?", req.Id).Find(&list).Error; err != nil {
		return
	}
	if len(list) == 0 {
		err = errors.New(fmt.Sprintf("日志[%d]不存在", req.Id))
		return
	}
	resp.Log = list[0]
	return
}

func (s *service) Add(req AddReq) (err error) {
	return base.Callback1(db.GetDB().Create(req.transform()).Error, req, req.Callback)
}

func (s *service) Update(req UpdateReq) (err error) {
	return base.Callback1(db.GetDB().Model(&models.Log{Id: req.Id}).Updates(req.transform()).Error, req, req.Callback)
}

func (s *service) Delete(req DeleteReq) (err error) {
	return base.Callback1(db.GetDB().Delete(&models.Log{Id: req.Id}).Error, req, req.Callback)
}
