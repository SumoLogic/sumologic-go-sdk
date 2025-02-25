/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on Collector and Search Job APIs, see our [API home page](https://help.sumologic.com/docs/api). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment, see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> KR </td>       <td> https://api.kr.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/docs/manage/security/access-keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---301-error---moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---401-error---credential-could-not-be-verified) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"fmt"
)

// checks if the LogSearchDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogSearchDefinition{}

// LogSearchDefinition struct for LogSearchDefinition
type LogSearchDefinition struct {
	// Query to perform.
	QueryString string `json:"queryString"`
	TimeRange ResolvableTimeRange `json:"timeRange"`
	// This has the value `true` if the search is to be run by receipt time and `false` if it is to be run by message time.
	RunByReceiptTime *bool `json:"runByReceiptTime,omitempty"`
	// Values for search template used in the search query. Learn more about the search templates here : https://help.sumologic.com/docs/search/get-started-with-search/build-search/search-templates/
	QueryParameters []LogSearchQueryParameterSyncDefinitionBase `json:"queryParameters,omitempty"`
	// Define the parsing mode to scan the JSON format log messages. Possible values are:   1. `AutoParse`   2. `Manual` In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid=0011).
	ParsingMode *string `json:"parsingMode,omitempty" validate:"regexp=^(AutoParse|Manual)$"`
	// Name of the item in the content library.
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9 +%-@.,_()\\\\\\\\]+$"`
	// Item description in the content library.
	Description *string `json:"description,omitempty"`
	Schedule *LogSearchScheduleSyncDefinition `json:"schedule,omitempty"`
	// Aggregate Results Settings and View configurations, Legends settings, and different visualisation settings overrides. Leave this field empty to use the defaults. This property contains JSON object encoded as a string. 
	Properties *string `json:"properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogSearchDefinition LogSearchDefinition

// NewLogSearchDefinition instantiates a new LogSearchDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogSearchDefinition(queryString string, timeRange ResolvableTimeRange, name string) *LogSearchDefinition {
	this := LogSearchDefinition{}
	this.QueryString = queryString
	this.TimeRange = timeRange
	var runByReceiptTime bool = false
	this.RunByReceiptTime = &runByReceiptTime
	var parsingMode string = "Manual"
	this.ParsingMode = &parsingMode
	this.Name = name
	return &this
}

// NewLogSearchDefinitionWithDefaults instantiates a new LogSearchDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogSearchDefinitionWithDefaults() *LogSearchDefinition {
	this := LogSearchDefinition{}
	var runByReceiptTime bool = false
	this.RunByReceiptTime = &runByReceiptTime
	var parsingMode string = "Manual"
	this.ParsingMode = &parsingMode
	return &this
}

// GetQueryString returns the QueryString field value
func (o *LogSearchDefinition) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *LogSearchDefinition) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value
func (o *LogSearchDefinition) SetQueryString(v string) {
	o.QueryString = v
}

// GetTimeRange returns the TimeRange field value
func (o *LogSearchDefinition) GetTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *LogSearchDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *LogSearchDefinition) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = v
}

// GetRunByReceiptTime returns the RunByReceiptTime field value if set, zero value otherwise.
func (o *LogSearchDefinition) GetRunByReceiptTime() bool {
	if o == nil || IsNil(o.RunByReceiptTime) {
		var ret bool
		return ret
	}
	return *o.RunByReceiptTime
}

// GetRunByReceiptTimeOk returns a tuple with the RunByReceiptTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchDefinition) GetRunByReceiptTimeOk() (*bool, bool) {
	if o == nil || IsNil(o.RunByReceiptTime) {
		return nil, false
	}
	return o.RunByReceiptTime, true
}

// HasRunByReceiptTime returns a boolean if a field has been set.
func (o *LogSearchDefinition) HasRunByReceiptTime() bool {
	if o != nil && !IsNil(o.RunByReceiptTime) {
		return true
	}

	return false
}

// SetRunByReceiptTime gets a reference to the given bool and assigns it to the RunByReceiptTime field.
func (o *LogSearchDefinition) SetRunByReceiptTime(v bool) {
	o.RunByReceiptTime = &v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *LogSearchDefinition) GetQueryParameters() []LogSearchQueryParameterSyncDefinitionBase {
	if o == nil || IsNil(o.QueryParameters) {
		var ret []LogSearchQueryParameterSyncDefinitionBase
		return ret
	}
	return o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchDefinition) GetQueryParametersOk() ([]LogSearchQueryParameterSyncDefinitionBase, bool) {
	if o == nil || IsNil(o.QueryParameters) {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *LogSearchDefinition) HasQueryParameters() bool {
	if o != nil && !IsNil(o.QueryParameters) {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given []LogSearchQueryParameterSyncDefinitionBase and assigns it to the QueryParameters field.
func (o *LogSearchDefinition) SetQueryParameters(v []LogSearchQueryParameterSyncDefinitionBase) {
	o.QueryParameters = v
}

// GetParsingMode returns the ParsingMode field value if set, zero value otherwise.
func (o *LogSearchDefinition) GetParsingMode() string {
	if o == nil || IsNil(o.ParsingMode) {
		var ret string
		return ret
	}
	return *o.ParsingMode
}

// GetParsingModeOk returns a tuple with the ParsingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchDefinition) GetParsingModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParsingMode) {
		return nil, false
	}
	return o.ParsingMode, true
}

// HasParsingMode returns a boolean if a field has been set.
func (o *LogSearchDefinition) HasParsingMode() bool {
	if o != nil && !IsNil(o.ParsingMode) {
		return true
	}

	return false
}

// SetParsingMode gets a reference to the given string and assigns it to the ParsingMode field.
func (o *LogSearchDefinition) SetParsingMode(v string) {
	o.ParsingMode = &v
}

// GetName returns the Name field value
func (o *LogSearchDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogSearchDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LogSearchDefinition) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LogSearchDefinition) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LogSearchDefinition) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LogSearchDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *LogSearchDefinition) GetSchedule() LogSearchScheduleSyncDefinition {
	if o == nil || IsNil(o.Schedule) {
		var ret LogSearchScheduleSyncDefinition
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchDefinition) GetScheduleOk() (*LogSearchScheduleSyncDefinition, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *LogSearchDefinition) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given LogSearchScheduleSyncDefinition and assigns it to the Schedule field.
func (o *LogSearchDefinition) SetSchedule(v LogSearchScheduleSyncDefinition) {
	o.Schedule = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *LogSearchDefinition) GetProperties() string {
	if o == nil || IsNil(o.Properties) {
		var ret string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchDefinition) GetPropertiesOk() (*string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *LogSearchDefinition) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given string and assigns it to the Properties field.
func (o *LogSearchDefinition) SetProperties(v string) {
	o.Properties = &v
}

func (o LogSearchDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogSearchDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryString"] = o.QueryString
	toSerialize["timeRange"] = o.TimeRange
	if !IsNil(o.RunByReceiptTime) {
		toSerialize["runByReceiptTime"] = o.RunByReceiptTime
	}
	if !IsNil(o.QueryParameters) {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	if !IsNil(o.ParsingMode) {
		toSerialize["parsingMode"] = o.ParsingMode
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogSearchDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryString",
		"timeRange",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLogSearchDefinition := _LogSearchDefinition{}

	err = json.Unmarshal(data, &varLogSearchDefinition)

	if err != nil {
		return err
	}

	*o = LogSearchDefinition(varLogSearchDefinition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "queryString")
		delete(additionalProperties, "timeRange")
		delete(additionalProperties, "runByReceiptTime")
		delete(additionalProperties, "queryParameters")
		delete(additionalProperties, "parsingMode")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "properties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogSearchDefinition struct {
	value *LogSearchDefinition
	isSet bool
}

func (v NullableLogSearchDefinition) Get() *LogSearchDefinition {
	return v.value
}

func (v *NullableLogSearchDefinition) Set(val *LogSearchDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSearchDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSearchDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSearchDefinition(val *LogSearchDefinition) *NullableLogSearchDefinition {
	return &NullableLogSearchDefinition{value: val, isSet: true}
}

func (v NullableLogSearchDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSearchDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


