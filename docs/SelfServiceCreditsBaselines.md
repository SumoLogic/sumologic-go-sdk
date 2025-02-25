# SelfServiceCreditsBaselines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuousIngest** | Pointer to **int64** | The amount of continuous logs ingest to allocate to the organization, in GBs. | [optional] [default to 0]
**ContinuousStorage** | Pointer to **int64** | Number of days of continuous logs storage to allocate to the organization, in Days. | [optional] [default to 0]
**Metrics** | Pointer to **int64** | The amount of Metrics usage to allocate to the organization, in DPMs (Data Points per Minute). | [optional] [default to 0]
**TracingIngest** | Pointer to **int64** | The amount of tracing data ingest to allocate to the organization, in GBs. | [optional] [default to 0]
**FlexIngest** | Pointer to **int64** | The amount of flex logs ingest to allocate to the organization, in GBs. | [optional] [default to 0]
**FlexStorage** | Pointer to **int64** | Number of days of flex logs storage to allocate to the organization, in Days. | [optional] [default to 0]
**FlexScanRatio** | Pointer to **int64** | The amount of flex logs ingest scan ratio. | [optional] [default to 0]

## Methods

### NewSelfServiceCreditsBaselines

`func NewSelfServiceCreditsBaselines() *SelfServiceCreditsBaselines`

NewSelfServiceCreditsBaselines instantiates a new SelfServiceCreditsBaselines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServiceCreditsBaselinesWithDefaults

`func NewSelfServiceCreditsBaselinesWithDefaults() *SelfServiceCreditsBaselines`

NewSelfServiceCreditsBaselinesWithDefaults instantiates a new SelfServiceCreditsBaselines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuousIngest

`func (o *SelfServiceCreditsBaselines) GetContinuousIngest() int64`

GetContinuousIngest returns the ContinuousIngest field if non-nil, zero value otherwise.

### GetContinuousIngestOk

`func (o *SelfServiceCreditsBaselines) GetContinuousIngestOk() (*int64, bool)`

GetContinuousIngestOk returns a tuple with the ContinuousIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousIngest

`func (o *SelfServiceCreditsBaselines) SetContinuousIngest(v int64)`

SetContinuousIngest sets ContinuousIngest field to given value.

### HasContinuousIngest

`func (o *SelfServiceCreditsBaselines) HasContinuousIngest() bool`

HasContinuousIngest returns a boolean if a field has been set.

### GetContinuousStorage

`func (o *SelfServiceCreditsBaselines) GetContinuousStorage() int64`

GetContinuousStorage returns the ContinuousStorage field if non-nil, zero value otherwise.

### GetContinuousStorageOk

`func (o *SelfServiceCreditsBaselines) GetContinuousStorageOk() (*int64, bool)`

GetContinuousStorageOk returns a tuple with the ContinuousStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousStorage

`func (o *SelfServiceCreditsBaselines) SetContinuousStorage(v int64)`

SetContinuousStorage sets ContinuousStorage field to given value.

### HasContinuousStorage

`func (o *SelfServiceCreditsBaselines) HasContinuousStorage() bool`

HasContinuousStorage returns a boolean if a field has been set.

### GetMetrics

`func (o *SelfServiceCreditsBaselines) GetMetrics() int64`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *SelfServiceCreditsBaselines) GetMetricsOk() (*int64, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *SelfServiceCreditsBaselines) SetMetrics(v int64)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *SelfServiceCreditsBaselines) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetTracingIngest

`func (o *SelfServiceCreditsBaselines) GetTracingIngest() int64`

GetTracingIngest returns the TracingIngest field if non-nil, zero value otherwise.

### GetTracingIngestOk

`func (o *SelfServiceCreditsBaselines) GetTracingIngestOk() (*int64, bool)`

GetTracingIngestOk returns a tuple with the TracingIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracingIngest

`func (o *SelfServiceCreditsBaselines) SetTracingIngest(v int64)`

SetTracingIngest sets TracingIngest field to given value.

### HasTracingIngest

`func (o *SelfServiceCreditsBaselines) HasTracingIngest() bool`

HasTracingIngest returns a boolean if a field has been set.

### GetFlexIngest

`func (o *SelfServiceCreditsBaselines) GetFlexIngest() int64`

GetFlexIngest returns the FlexIngest field if non-nil, zero value otherwise.

### GetFlexIngestOk

`func (o *SelfServiceCreditsBaselines) GetFlexIngestOk() (*int64, bool)`

GetFlexIngestOk returns a tuple with the FlexIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexIngest

`func (o *SelfServiceCreditsBaselines) SetFlexIngest(v int64)`

SetFlexIngest sets FlexIngest field to given value.

### HasFlexIngest

`func (o *SelfServiceCreditsBaselines) HasFlexIngest() bool`

HasFlexIngest returns a boolean if a field has been set.

### GetFlexStorage

`func (o *SelfServiceCreditsBaselines) GetFlexStorage() int64`

GetFlexStorage returns the FlexStorage field if non-nil, zero value otherwise.

### GetFlexStorageOk

`func (o *SelfServiceCreditsBaselines) GetFlexStorageOk() (*int64, bool)`

GetFlexStorageOk returns a tuple with the FlexStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexStorage

`func (o *SelfServiceCreditsBaselines) SetFlexStorage(v int64)`

SetFlexStorage sets FlexStorage field to given value.

### HasFlexStorage

`func (o *SelfServiceCreditsBaselines) HasFlexStorage() bool`

HasFlexStorage returns a boolean if a field has been set.

### GetFlexScanRatio

`func (o *SelfServiceCreditsBaselines) GetFlexScanRatio() int64`

GetFlexScanRatio returns the FlexScanRatio field if non-nil, zero value otherwise.

### GetFlexScanRatioOk

`func (o *SelfServiceCreditsBaselines) GetFlexScanRatioOk() (*int64, bool)`

GetFlexScanRatioOk returns a tuple with the FlexScanRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexScanRatio

`func (o *SelfServiceCreditsBaselines) SetFlexScanRatio(v int64)`

SetFlexScanRatio sets FlexScanRatio field to given value.

### HasFlexScanRatio

`func (o *SelfServiceCreditsBaselines) HasFlexScanRatio() bool`

HasFlexScanRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


