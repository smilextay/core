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
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("core.Int64Values", &Int64ValuesCodec{})
	jsoniter.RegisterTypeEncoder("core.Int64Values", &Int64ValuesCodec{})
}

type Int64ValuesCodec struct {
}

func (codec *Int64ValuesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	int64Values := (*Int64Values)(ptr)
	if any.ValueType() == jsoniter.ArrayValue {
		any.ToVal(&int64Values.Values)
	}
}

func (codec *Int64ValuesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	int64Values := (*Int64Values)(ptr)
	return int64Values == nil || len(int64Values.Values) == 0
}

func (codec *Int64ValuesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	int64Values := (*Int64Values)(ptr)

	stream.WriteArrayStart()
	for i, v := range int64Values.Values {
		if i > 0 {
			stream.WriteMore()
		}
		stream.WriteVal(v)
	}
	stream.WriteArrayEnd()
}