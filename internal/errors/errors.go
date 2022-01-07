/*
Copyright © 2021, 2022 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package errors

import (
	"fmt"
)

// Error data structure contains the original error object + assigned error code
type Error struct {
	orig error
	msg  string
	code ErrorCode
}

// ErrorCode is enumeration type to specify numeric error code
type ErrorCode uint

const (
	// ErrorCodeUnknown represents numeric error code for unknown error
	ErrorCodeUnknown ErrorCode = iota

	// ErrorCodeNotFound represents numeric error code for error that
	// occurs when the rule data are not found
	ErrorCodeNotFound

	// ErrorCodeInvalidArgument represents numeric error code for error
	// that occurs when invalid argument is provided in request
	ErrorCodeInvalidArgument
)

func (e *Error) Error() string {
	if e.orig != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.orig)
	}

	return e.msg
}

func (e *Error) Unwrap() error {
	return e.orig
}

// Code method returns numeric error code
func (e *Error) Code() ErrorCode {
	return e.code
}

// WrapErrorf function constructs Error data structure with original error
// object, numeric error code, and message.
func WrapErrorf(orig error, code ErrorCode, format string, a ...interface{}) error {
	return &Error{
		orig: orig,
		code: code,
		msg:  fmt.Sprintf(format, a...),
	}
}

// NewErrorf function constructs Error data structure with numeric error code,
// and message.
func NewErrorf(code ErrorCode, format string, a ...interface{}) error {
	return WrapErrorf(nil, code, format, a...)
}
