# LogSearchQueryEstimationQueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** | Log search Query to compute the estimated volume of data scanned. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**RunByReceiptTime** | Pointer to **bool** | This has the value &#x60;true&#x60; if the search is to be run by receipt time and &#x60;false&#x60; if it is to be run by message time. | [optional] [default to false]
**QueryParameters** | Pointer to [**[]LogSearchQueryParameterSyncDefinitionBase**](LogSearchQueryParameterSyncDefinitionBase.md) | Values for search template used in the search query. Learn more about the search templates here : https://help.sumologic.com/docs/search/get-started-with-search/build-search/search-templates/ | [optional] 

## Methods

### NewLogSearchQueryEstimationQueryDefinition

`func NewLogSearchQueryEstimationQueryDefinition(queryString string, timeRange ResolvableTimeRange, ) *LogSearchQueryEstimationQueryDefinition`

NewLogSearchQueryEstimationQueryDefinition instantiates a new LogSearchQueryEstimationQueryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSearchQueryEstimationQueryDefinitionWithDefaults

`func NewLogSearchQueryEstimationQueryDefinitionWithDefaults() *LogSearchQueryEstimationQueryDefinition`

NewLogSearchQueryEstimationQueryDefinitionWithDefaults instantiates a new LogSearchQueryEstimationQueryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *LogSearchQueryEstimationQueryDefinition) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *LogSearchQueryEstimationQueryDefinition) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *LogSearchQueryEstimationQueryDefinition) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetTimeRange

`func (o *LogSearchQueryEstimationQueryDefinition) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *LogSearchQueryEstimationQueryDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *LogSearchQueryEstimationQueryDefinition) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetRunByReceiptTime

`func (o *LogSearchQueryEstimationQueryDefinition) GetRunByReceiptTime() bool`

GetRunByReceiptTime returns the RunByReceiptTime field if non-nil, zero value otherwise.

### GetRunByReceiptTimeOk

`func (o *LogSearchQueryEstimationQueryDefinition) GetRunByReceiptTimeOk() (*bool, bool)`

GetRunByReceiptTimeOk returns a tuple with the RunByReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByReceiptTime

`func (o *LogSearchQueryEstimationQueryDefinition) SetRunByReceiptTime(v bool)`

SetRunByReceiptTime sets RunByReceiptTime field to given value.

### HasRunByReceiptTime

`func (o *LogSearchQueryEstimationQueryDefinition) HasRunByReceiptTime() bool`

HasRunByReceiptTime returns a boolean if a field has been set.

### GetQueryParameters

`func (o *LogSearchQueryEstimationQueryDefinition) GetQueryParameters() []LogSearchQueryParameterSyncDefinitionBase`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *LogSearchQueryEstimationQueryDefinition) GetQueryParametersOk() (*[]LogSearchQueryParameterSyncDefinitionBase, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *LogSearchQueryEstimationQueryDefinition) SetQueryParameters(v []LogSearchQueryParameterSyncDefinitionBase)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *LogSearchQueryEstimationQueryDefinition) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


