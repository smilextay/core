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

var PhoneNumberCountryCodeSourceNames = map[int32]string{
	1:  "from_number_with_plus_sign",
	5:  "from_number_with_idd",
	10: "from_number_without_plus_sign",
	20: "from_default_country",
}

var PhoneNumberCountryCodeSourceValues = map[string]PhoneNumber_CountryCodeSource{
	"from_number_with_plus_sign":    PhoneNumber_COUNTRY_CODE_SOURCE_FROM_NUMBER_WITH_PLUS_SIGN,
	"from_number_with_idd":          PhoneNumber_COUNTRY_CODE_SOURCE_FROM_NUMBER_WITH_IDD,
	"from_number_without_plus_sign": PhoneNumber_COUNTRY_CODE_SOURCE_FROM_NUMBER_WITHOUT_PLUS_SIGN,
	"from_default_country":          PhoneNumber_COUNTRY_CODE_SOURCE_FROM_DEFAULT_COUNTRY,
}

func (x PhoneNumber_CountryCodeSource) Format() string {
	s, ok := PhoneNumberCountryCodeSourceNames[int32(x)]
	if ok {
		return s
	}
	if int(x) == 0 {
		return "unspecified"
	}
	return strconv.Itoa(int(x))
}

func (x *PhoneNumber_CountryCodeSource) Parse(value string) error {
	if x != nil {
		s, ok := PhoneNumberCountryCodeSourceValues[value]
		if ok {
			*x = s
		} else {
			*x = PhoneNumber_COUNTRY_CODE_SOURCE_FROM_NUMBER_WITH_PLUS_SIGN
		}
	} else {
		*x = PhoneNumber_COUNTRY_CODE_SOURCE_FROM_NUMBER_WITH_PLUS_SIGN
	}
	return nil
}
