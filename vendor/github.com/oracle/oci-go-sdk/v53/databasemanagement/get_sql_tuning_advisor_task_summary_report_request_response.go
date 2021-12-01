// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v53/common"
	"net/http"
)

// GetSqlTuningAdvisorTaskSummaryReportRequest wrapper for the GetSqlTuningAdvisorTaskSummaryReport operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetSqlTuningAdvisorTaskSummaryReport.go.html to see an example of how to use GetSqlTuningAdvisorTaskSummaryReportRequest.
type GetSqlTuningAdvisorTaskSummaryReportRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The SQL tuning task identifier. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskId *int64 `mandatory:"true" contributesTo:"path" name:"sqlTuningAdvisorTaskId"`

	// How far back the API will search for begin and end exec id. Unused if neither exec ids nor time filter query params are supplied. This is applicable only for Auto SQL Tuning tasks.
	SearchPeriod GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum `mandatory:"false" contributesTo:"query" name:"searchPeriod" omitEmpty:"true"`

	// The optional greater than or equal to query parameter to filter the timestamp. This is applicable only for Auto SQL Tuning tasks.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp. This is applicable only for Auto SQL Tuning tasks.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The optional greater than or equal to filter on the execution ID related to a specific SQL Tuning Advisor task. This is applicable only for Auto SQL Tuning tasks.
	BeginExecIdGreaterThanOrEqualTo *int64 `mandatory:"false" contributesTo:"query" name:"beginExecIdGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter on the execution ID related to a specific SQL Tuning Advisor task. This is applicable only for Auto SQL Tuning tasks.
	EndExecIdLessThanOrEqualTo *int64 `mandatory:"false" contributesTo:"query" name:"endExecIdLessThanOrEqualTo"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetSqlTuningAdvisorTaskSummaryReportRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetSqlTuningAdvisorTaskSummaryReportRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetSqlTuningAdvisorTaskSummaryReportRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetSqlTuningAdvisorTaskSummaryReportRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetSqlTuningAdvisorTaskSummaryReportResponse wrapper for the GetSqlTuningAdvisorTaskSummaryReport operation
type GetSqlTuningAdvisorTaskSummaryReportResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SqlTuningAdvisorTaskSummaryReport instance
	SqlTuningAdvisorTaskSummaryReport `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetSqlTuningAdvisorTaskSummaryReportResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetSqlTuningAdvisorTaskSummaryReportResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum Enum with underlying type: string
type GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum string

// Set of constants representing the allowable values for GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum
const (
	GetSqlTuningAdvisorTaskSummaryReportSearchPeriodLast24hr  GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum = "LAST_24HR"
	GetSqlTuningAdvisorTaskSummaryReportSearchPeriodLast7day  GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum = "LAST_7DAY"
	GetSqlTuningAdvisorTaskSummaryReportSearchPeriodLast31day GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum = "LAST_31DAY"
	GetSqlTuningAdvisorTaskSummaryReportSearchPeriodSinceLast GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum = "SINCE_LAST"
	GetSqlTuningAdvisorTaskSummaryReportSearchPeriodAll       GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum = "ALL"
)

var mappingGetSqlTuningAdvisorTaskSummaryReportSearchPeriod = map[string]GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum{
	"LAST_24HR":  GetSqlTuningAdvisorTaskSummaryReportSearchPeriodLast24hr,
	"LAST_7DAY":  GetSqlTuningAdvisorTaskSummaryReportSearchPeriodLast7day,
	"LAST_31DAY": GetSqlTuningAdvisorTaskSummaryReportSearchPeriodLast31day,
	"SINCE_LAST": GetSqlTuningAdvisorTaskSummaryReportSearchPeriodSinceLast,
	"ALL":        GetSqlTuningAdvisorTaskSummaryReportSearchPeriodAll,
}

// GetGetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnumValues Enumerates the set of values for GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum
func GetGetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnumValues() []GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum {
	values := make([]GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum, 0)
	for _, v := range mappingGetSqlTuningAdvisorTaskSummaryReportSearchPeriod {
		values = append(values, v)
	}
	return values
}