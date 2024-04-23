// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// ScheduleTypesEnum Enum with underlying type: string
type ScheduleTypesEnum string

// Set of constants representing the allowable values for ScheduleTypesEnum
const (
	ScheduleTypesOnetime   ScheduleTypesEnum = "ONETIME"
	ScheduleTypesRecurring ScheduleTypesEnum = "RECURRING"
)

var mappingScheduleTypesEnum = map[string]ScheduleTypesEnum{
	"ONETIME":   ScheduleTypesOnetime,
	"RECURRING": ScheduleTypesRecurring,
}

var mappingScheduleTypesEnumLowerCase = map[string]ScheduleTypesEnum{
	"onetime":   ScheduleTypesOnetime,
	"recurring": ScheduleTypesRecurring,
}

// GetScheduleTypesEnumValues Enumerates the set of values for ScheduleTypesEnum
func GetScheduleTypesEnumValues() []ScheduleTypesEnum {
	values := make([]ScheduleTypesEnum, 0)
	for _, v := range mappingScheduleTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleTypesEnumStringValues Enumerates the set of values in String for ScheduleTypesEnum
func GetScheduleTypesEnumStringValues() []string {
	return []string{
		"ONETIME",
		"RECURRING",
	}
}

// GetMappingScheduleTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleTypesEnum(val string) (ScheduleTypesEnum, bool) {
	enum, ok := mappingScheduleTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
