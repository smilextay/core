package core

import jsoniter "github.com/json-iterator/go"

const ObjectTypeName = "Object"
const ObjectTypeFullName = "mojo.core.Object"

func NewObject() *Object {
    return &Object{}
}

func (x *Object) To(value interface{}) error {
    if x != nil {
        if json, err := jsoniter.ConfigFastest.Marshal(x); err != nil {
            return err
        } else {
            return jsoniter.ConfigFastest.Unmarshal(json, value)
        }
    }
    return nil
}

func (x *Object) From(value interface{}) error {
    if x != nil {
        if json, err := jsoniter.ConfigFastest.Marshal(value); err != nil {
            return err
        } else {
            return jsoniter.ConfigFastest.Unmarshal(json, x)
        }
    }
    return nil
}

func (x *Object) GetValue(key string) *Value {
    if x != nil && x.Vals != nil {
        return x.Vals[key]
    }
    return nil
}

func (x *Object) GetString(key string) string {
    return x.GetValue(key).GetString()
}

func (x *Object) GetBool(key string) bool {
    return x.GetValue(key).GetBool()
}

func (x *Object) GetInt32(key string) int32 {
    return x.GetValue(key).GetInt32()
}

func (x *Object) GetInt64(key string) int64 {
    return x.GetValue(key).GetInt64()
}

func (x *Object) GetUint64(key string) uint64 {
    return x.GetValue(key).GetUint64()
}

func (x *Object) GetInt32Array(key string) []int32 {
    return x.GetValue(key).GetInt32Array()
}

func (x *Object) GetInt64Array(key string) []int64 {
    return x.GetValue(key).GetInt64Array()
}

func (x *Object) GetUint32Array(key string) []uint32 {
    return x.GetValue(key).GetUint32Array()
}

func (x *Object) GetUint64Array(key string) []uint64 {
    return x.GetValue(key).GetUint64Array()
}

func (x *Object) GetStringArray(key string) []string {
    return x.GetValue(key).GetStringArray()
}

func (x *Object) GetObjectArray(key string) []*Object {
    return x.GetValue(key).GetObjectArray()
}

func (x *Object) init() {
    if x != nil && x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
}

func (x *Object) SetValue(key string, value *Value) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = value
    }

    return x
}

func (x *Object) SetBool(key string, value bool) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewBoolValue(value)
    }
    return x
}

func (x *Object) SetInt(key string, value int) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewIntValue(value)
    }

    return x
}

func (x *Object) SetInt32(key string, value int32) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewInt32Value(value)
    }
    return x
}

func (x *Object) SetInt64(key string, value int64) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewInt64Value(value)
    }
    return x
}

func (x *Object) SetUint32(key string, value uint32) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewUint32Value(value)
    }
    return x
}

func (x *Object) SetUint64(key string, value uint64) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewUint64Value(value)
    }
    return x
}

func (x *Object) SetFloat32(key string, value float32) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewFloat32Value(value)
    }
    return x
}

func (x *Object) SetFloat64(key string, value float64) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewFloat64Value(value)
    }
    return x
}

func (x *Object) SetString(key string, value string) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewStringValue(value)
    }
    return x
}

func (x *Object) SetObject(key string, value *Object) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewObjectValue(value)
    }
    return x
}

func (x *Object) SetInt32Array(key string, vals ...int32) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewInt32ArrayValue(vals...)
    }
    return x
}

func (x *Object) SetInt64Array(key string, vals ...int64) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewInt64ArrayValue(vals...)
    }
    return x
}

func (x *Object) SetUint32Array(key string, vals ...uint32) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewUint32ArrayValue(vals...)
    }
    return x
}

func (x *Object) SetUint64Array(key string, vals ...uint64) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewUint64ArrayValue(vals...)
    }
    return x
}

func (x *Object) SetStringArray(key string, vals ...string) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewStringArrayValue(vals...)
    }
    return x
}

func (x *Object) SetObjectArray(key string, vals ...*Object) *Object {
    if x != nil {
        x.init()
        x.Vals[key] = NewObjectArrayValue(vals...)
    }
    return x
}
