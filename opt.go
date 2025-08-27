// Copyright 2025 icmpkt Author. All Rights Reserved.
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

func UserId(a uint) CallbackFunc     { return func(req *LogAddReq) { req.UserId = a } }
func UserName(a string) CallbackFunc { return func(req *LogAddReq) { req.UserName = a } }
func ClientIP(a string) CallbackFunc { return func(req *LogAddReq) { req.ClientIP = a } }
func Type1(a string) CallbackFunc    { return func(req *LogAddReq) { req.Type1 = a } }
func Type2(a string) CallbackFunc    { return func(req *LogAddReq) { req.Type2 = a } }
func Type3(a string) CallbackFunc    { return func(req *LogAddReq) { req.Type3 = a } }
func Type4(a string) CallbackFunc    { return func(req *LogAddReq) { req.Type4 = a } }
func Type5(a string) CallbackFunc    { return func(req *LogAddReq) { req.Type5 = a } }
func Content(a string) CallbackFunc  { return func(req *LogAddReq) { req.Content = a } }
