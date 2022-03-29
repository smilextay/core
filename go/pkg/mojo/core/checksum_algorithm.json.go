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
	RegisterJSONTypeDecoder("core.Checksum_Algorithm", &ChecksumAlgorithmCodec{})
	RegisterJSONTypeEncoder("core.Checksum_Algorithm", &ChecksumAlgorithmCodec{})
}

type ChecksumAlgorithmCodec struct {
}

func (codec *ChecksumAlgorithmCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	e := (*Checksum_Algorithm)(ptr)
	if any.ValueType() == jsoniter.StringValue {
		e.Parse(any.ToString())
	} else if any.ValueType() == jsoniter.NumberValue {
		value := any.ToInt32()
		if _, ok := ChecksumAlgorithmNames[value]; ok {
			*e = Checksum_Algorithm(value)
		}
	}
}

func (codec *ChecksumAlgorithmCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	e := (*Checksum_Algorithm)(ptr)
	stream.WriteString(e.Format())
}

func (codec *ChecksumAlgorithmCodec) IsEmpty(ptr unsafe.Pointer) bool {
	e := (*Checksum_Algorithm)(ptr)
	return e == nil || *e == 0
}
