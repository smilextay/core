package core

const StringValuesTypeName = "StringValues"
const StringValuesTypeFullName = "mojo.core.StringValues"

func NewStringValues(values ...string) *StringValues {
    return &StringValues{Vals: values}
}

func (x *StringValues) Unique() *StringValues {
    if x != nil {
        keys := make(map[string]bool)
        var list []string
        for _, entry := range x.Vals {
            if _, ok := keys[entry]; !ok {
                keys[entry] = true
                list = append(list, entry)
            }
        }
        x.Vals = list
    }
    return x
}
