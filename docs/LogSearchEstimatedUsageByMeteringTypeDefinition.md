# LogSearchEstimatedUsageByMeteringTypeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** | Log search Query to compute the estimated volume of data scanned. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**RunByReceiptTime** | Pointer to **bool** | This has the value &#x60;true&#x60; if the search is to be run by receipt time and &#x60;false&#x60; if it is to be run by message time. | [optional] [default to false]
**QueryParameters** | Pointer to [**[]LogSearchQueryParameterSyncDefinitionBase**](LogSearchQueryParameterSyncDefinitionBase.md) | Values for search template used in the search query. Learn more about the search templates here : https://help.sumologic.com/docs/search/get-started-with-search/build-search/search-templates/ | [optional] 
**Timezone** | **string** | Time zone to get the estimated usage details. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).  | 
**EmulateSearchContext** | Pointer to [**LogSearchEstimatedUsageRequestV3AllOfEmulateSearchContext**](LogSearchEstimatedUsageRequestV3AllOfEmulateSearchContext.md) |  | [optional] 
**EstimatedUsageDetails** | [**[]EstimatedUsageDetailsWithMeteringType**](EstimatedUsageDetailsWithMeteringType.md) |  | 

## Methods

### NewLogSearchEstimatedUsageByMeteringTypeDefinition

`func NewLogSearchEstimatedUsageByMeteringTypeDefinition(queryString string, timeRange ResolvableTimeRange, timezone string, estimatedUsageDetails []EstimatedUsageDetailsWithMeteringType, ) *LogSearchEstimatedUsageByMeteringTypeDefinition`

NewLogSearchEstimatedUsageByMeteringTypeDefinition instantiates a new LogSearchEstimatedUsageByMeteringTypeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSearchEstimatedUsageByMeteringTypeDefinitionWithDefaults

`func NewLogSearchEstimatedUsageByMeteringTypeDefinitionWithDefaults() *LogSearchEstimatedUsageByMeteringTypeDefinition`

NewLogSearchEstimatedUsageByMeteringTypeDefinitionWithDefaults instantiates a new LogSearchEstimatedUsageByMeteringTypeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetTimeRange

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetRunByReceiptTime

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetRunByReceiptTime() bool`

GetRunByReceiptTime returns the RunByReceiptTime field if non-nil, zero value otherwise.

### GetRunByReceiptTimeOk

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetRunByReceiptTimeOk() (*bool, bool)`

GetRunByReceiptTimeOk returns a tuple with the RunByReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByReceiptTime

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) SetRunByReceiptTime(v bool)`

SetRunByReceiptTime sets RunByReceiptTime field to given value.

### HasRunByReceiptTime

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) HasRunByReceiptTime() bool`

HasRunByReceiptTime returns a boolean if a field has been set.

### GetQueryParameters

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetQueryParameters() []LogSearchQueryParameterSyncDefinitionBase`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetQueryParametersOk() (*[]LogSearchQueryParameterSyncDefinitionBase, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) SetQueryParameters(v []LogSearchQueryParameterSyncDefinitionBase)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.

### GetTimezone

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetEmulateSearchContext

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetEmulateSearchContext() LogSearchEstimatedUsageRequestV3AllOfEmulateSearchContext`

GetEmulateSearchContext returns the EmulateSearchContext field if non-nil, zero value otherwise.

### GetEmulateSearchContextOk

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetEmulateSearchContextOk() (*LogSearchEstimatedUsageRequestV3AllOfEmulateSearchContext, bool)`

GetEmulateSearchContextOk returns a tuple with the EmulateSearchContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmulateSearchContext

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) SetEmulateSearchContext(v LogSearchEstimatedUsageRequestV3AllOfEmulateSearchContext)`

SetEmulateSearchContext sets EmulateSearchContext field to given value.

### HasEmulateSearchContext

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) HasEmulateSearchContext() bool`

HasEmulateSearchContext returns a boolean if a field has been set.

### GetEstimatedUsageDetails

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetEstimatedUsageDetails() []EstimatedUsageDetailsWithMeteringType`

GetEstimatedUsageDetails returns the EstimatedUsageDetails field if non-nil, zero value otherwise.

### GetEstimatedUsageDetailsOk

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) GetEstimatedUsageDetailsOk() (*[]EstimatedUsageDetailsWithMeteringType, bool)`

GetEstimatedUsageDetailsOk returns a tuple with the EstimatedUsageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedUsageDetails

`func (o *LogSearchEstimatedUsageByMeteringTypeDefinition) SetEstimatedUsageDetails(v []EstimatedUsageDetailsWithMeteringType)`

SetEstimatedUsageDetails sets EstimatedUsageDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


