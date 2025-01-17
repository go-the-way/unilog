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

package base

func Callback1[T any](err0 error, t T, fns ...func(t T)) (err error) {
	if err = err0; err != nil {
		return
	}
	for _, fn := range fns {
		if fn != nil {
			fn(t)
		}
	}
	return
}

func Callback2[T1, T2 any](err0 error, t1 T1, t2 T2, fns ...func(t1 T1, t2 T2)) (err error) {
	if err = err0; err != nil {
		return
	}
	for _, fn := range fns {
		if fn != nil {
			fn(t1, t2)
		}
	}
	return
}

func Callback3[T1, T2, T3 any](err0 error, t1 T1, t2 T2, t3 T3, fns ...func(t1 T1, t2 T2, t3 T3)) (err error) {
	if err = err0; err != nil {
		return
	}
	for _, fn := range fns {
		if fn != nil {
			fn(t1, t2, t3)
		}
	}
	return
}

func Callback1Err[T any](err0 error, t T, fns ...func(t T) (err error)) (err error) {
	if err = err0; err != nil {
		return
	}
	for _, fn := range fns {
		if fn != nil {
			if err = fn(t); err != nil {
				return
			}
		}
	}
	return
}

func Return[T any](t T, err0 error) (T, error) { return t, err0 }
