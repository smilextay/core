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
	RegisterJSONTypeDecoder("core.Ordering_Value", &OrderingValueStringCodec{})
	RegisterJSONTypeEncoder("core.Ordering_Value", &OrderingValueStringCodec{})
}

// BareOrderingValue will be jsonify to raw, without any codec
type BareOrderingValue Ordering_Value

type OrderingValueStringCodec struct {
	IsFieldPointer bool
}

func (codec *OrderingValueStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	s := iter.ReadString()
	value := codec.value(ptr)
	if value == nil {
		value = &Ordering_Value{}
		*(**Ordering_Value)(ptr) = value
	}

	if err := value.Parse(s); err != nil {
		iter.ReportError("OrderingValueStringCodec", err.Error())
	}
}

func (codec *OrderingValueStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	value := codec.value(ptr)
	if value != nil {
		if checker, ok := interface{}(value).(EmptyChecker); ok {
			return checker.IsEmpty()
		}
		return false
	}
	return true
}

func (codec *OrderingValueStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	value := codec.value(ptr)
	stream.WriteString(value.Format())
}

func (codec *OrderingValueStringCodec) value(ptr unsafe.Pointer) *Ordering_Value {
	if codec.IsFieldPointer {
		return *(**Ordering_Value)(ptr)
	}
	return (*Ordering_Value)(ptr)
}

type OrderingValueStructCodec struct {
	IsFieldPointer bool
}

func (codec *OrderingValueStructCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := codec.bareOrderingValue(ptr)
	if a := iter.ReadAny(); a.ValueType() == jsoniter.ObjectValue {
		if value == nil {
			value = &BareOrderingValue{}
			*(**BareOrderingValue)(ptr) = value
		}
		a.ToVal(value)
	}
}

func (codec *OrderingValueStructCodec) IsEmpty(ptr unsafe.Pointer) bool {
	value := (*Ordering_Value)(codec.bareOrderingValue(ptr))
	if value != nil {
		if checker, ok := interface{}(value).(EmptyChecker); ok {
			return checker.IsEmpty()
		}
		return false
	}
	return true
}

func (codec *OrderingValueStructCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteVal(codec.bareOrderingValue(ptr))
}

func (codec *OrderingValueStructCodec) bareOrderingValue(ptr unsafe.Pointer) *BareOrderingValue {
	if codec.IsFieldPointer {
		return *(**BareOrderingValue)(ptr)
	}
	return (*BareOrderingValue)(ptr)
}
