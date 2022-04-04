package core

import (
    "google.golang.org/protobuf/proto"
    "reflect"
)

const UnionTypeName = "Union"
const UnionTypeFullName = "mojo.core.Union"

type IsUnion interface {
    IsUnion()
}

// GetUnionFieldIndex make sure the value is the union type, and at least have one field
func GetUnionFieldIndex(value reflect.Value) int {
    v := reflect.Indirect(value)
    if _, ok := value.Interface().(proto.Message); ok {
        if v.NumField() > 3 && v.Field(3).Type().Kind() == reflect.Interface {
            return 3
        }
        panic("GetUnionFieldIndex: the protobuf message is not generated by protoc-gen-go v1.26.0+")
    }
    if v.NumField() > 0 && v.Field(0).Type().Kind() == reflect.Interface {
        return 0
    }
    panic("GetUnionFieldIndex: the first field of the value is not a valid union type")
}

func GetUnionPrimeType(union interface{}) interface{} {
    if union != nil {
        value := reflect.ValueOf(union)
        for {
            if _, ok := value.Interface().(IsUnion); ok {
                value = reflect.Indirect(value).Field(GetUnionFieldIndex(value)).Elem()
                value = reflect.Indirect(value).Field(0)
            } else {
                break
            }
        }
        return value.Interface()
    }
    return nil
}

func GetUnionField(union interface{}, field string) reflect.Value {
    if union != nil {
        value := reflect.ValueOf(union)
        for {
            if _, ok := value.Interface().(IsUnion); ok {
                value = reflect.Indirect(value).Field(GetUnionFieldIndex(value)).Elem()
                value = reflect.Indirect(value).Field(0)
            } else {
                break
            }
        }
        return reflect.Indirect(value).FieldByName(field)
    }
    return reflect.ValueOf(nil)
}

func SetUnionField(union interface{}, field string, val reflect.Value) {
    if union != nil {
        value := reflect.ValueOf(union)
        for {
            if _, ok := value.Interface().(IsUnion); ok {
                value = reflect.Indirect(value).Field(GetUnionFieldIndex(value)).Elem()
                value = reflect.Indirect(value).Field(0)
            } else {
                break
            }
        }

        value = reflect.Indirect(value).FieldByName(field)
        if value.CanSet() {
            value.Set(val)
        }
    }
}
