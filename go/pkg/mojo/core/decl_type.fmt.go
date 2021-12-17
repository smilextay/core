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
	"strconv"
)

var DeclTypeNames = map[int32]string{
	1: "type",
	2: "value",
	3: "function",
	4: "constructor",
	5: "attribute",
	6: "package",
	7: "generic",
}

var DeclTypeValues = map[string]DeclType{
	"type":        DeclType_DECL_TYPE_TYPE,
	"value":       DeclType_DECL_TYPE_VALUE,
	"function":    DeclType_DECL_TYPE_FUNCTION,
	"constructor": DeclType_DECL_TYPE_CONSTRUCTOR,
	"attribute":   DeclType_DECL_TYPE_ATTRIBUTE,
	"package":     DeclType_DECL_TYPE_PACKAGE,
	"generic":     DeclType_DECL_TYPE_GENERIC,
}

func (x DeclType) Format() string {
	s, ok := DeclTypeNames[int32(x)]
	if ok {
		return s
	}
	if int(x) == 0 {
		return "unspecified"
	}
	return strconv.Itoa(int(x))
}

func (x *DeclType) Parse(value string) error {
	if x != nil {
		s, ok := DeclTypeValues[value]
		if ok {
			*x = s
		} else {
			*x = DeclType_DECL_TYPE_TYPE
		}
	} else {
		*x = DeclType_DECL_TYPE_TYPE
	}
	return nil
}
