// Code generated by mojo. DO NOT EDIT.
// Rerunning mojo will overwrite this file.
//
// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.EmailAddress", &EmailAddressStringCodec{})
	RegisterJSONTypeEncoder("core.EmailAddress", &EmailAddressStringCodec{})
}

// BareEmailAddress will be jsonify to raw, without any codec
type BareEmailAddress EmailAddress

type EmailAddressStringCodec struct {
	IsFieldPointer bool
}

func (codec *EmailAddressStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	s := iter.ReadString()
	emailAddress := codec.emailAddress(ptr)
	if emailAddress == nil {
		emailAddress = &EmailAddress{}
		*(**EmailAddress)(ptr) = emailAddress
	}

	if err := emailAddress.Parse(s); err != nil {
		iter.ReportError("EmailAddressStringCodec", err.Error())
	}
}

func (codec *EmailAddressStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	emailAddress := codec.emailAddress(ptr)
	if emailAddress != nil {
		if checker, ok := interface{}(emailAddress).(EmptyChecker); ok {
			return checker.IsEmpty()
		}
		return false
	}
	return true
}

func (codec *EmailAddressStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	emailAddress := codec.emailAddress(ptr)
	stream.WriteString(emailAddress.Format())
}

func (codec *EmailAddressStringCodec) emailAddress(ptr unsafe.Pointer) *EmailAddress {
	if codec.IsFieldPointer {
		return *(**EmailAddress)(ptr)
	}
	return (*EmailAddress)(ptr)
}

type EmailAddressStructCodec struct {
	IsFieldPointer bool
}

func (codec *EmailAddressStructCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	emailAddress := codec.bareEmailAddress(ptr)
	if a := iter.ReadAny(); a.ValueType() == jsoniter.ObjectValue {
		if emailAddress == nil {
			emailAddress = &BareEmailAddress{}
			*(**BareEmailAddress)(ptr) = emailAddress
		}
		a.ToVal(emailAddress)
	}
}

func (codec *EmailAddressStructCodec) IsEmpty(ptr unsafe.Pointer) bool {
	emailAddress := (*EmailAddress)(codec.bareEmailAddress(ptr))
	if emailAddress != nil {
		if checker, ok := interface{}(emailAddress).(EmptyChecker); ok {
			return checker.IsEmpty()
		}
		return false
	}
	return true
}

func (codec *EmailAddressStructCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteVal(codec.bareEmailAddress(ptr))
}

func (codec *EmailAddressStructCodec) bareEmailAddress(ptr unsafe.Pointer) *BareEmailAddress {
	if codec.IsFieldPointer {
		return *(**BareEmailAddress)(ptr)
	}
	return (*BareEmailAddress)(ptr)
}
