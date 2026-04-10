# IncidentByIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncidentId** | Pointer to **string** | Unique identifier for the incident. | [optional] 
**IncidentImpact** | Pointer to **string** | Impact level of the incident. | [optional] 
**IncidentMessage** | Pointer to **string** | Description of the issue. | [optional] 
**IncidentStartTime** | Pointer to **time.Time** | Start time of the incident. | [optional] 
**IncidentEndTime** | Pointer to **NullableTime** | End time of the incident, if applicable; null if still ongoing. | [optional] 
**IncidentStatus** | Pointer to **string** | Current status of the incident. | [optional] 
**IncidentType** | Pointer to **string** | Type of the incident. | [optional] 
**PaymentMethod** | Pointer to [**IncidentObjectPaymentMethod**](IncidentObjectPaymentMethod.md) |  | [optional] 

## Methods

### NewIncidentByIdResponse

`func NewIncidentByIdResponse() *IncidentByIdResponse`

NewIncidentByIdResponse instantiates a new IncidentByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentByIdResponseWithDefaults

`func NewIncidentByIdResponseWithDefaults() *IncidentByIdResponse`

NewIncidentByIdResponseWithDefaults instantiates a new IncidentByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidentId

`func (o *IncidentByIdResponse) GetIncidentId() string`

GetIncidentId returns the IncidentId field if non-nil, zero value otherwise.

### GetIncidentIdOk

`func (o *IncidentByIdResponse) GetIncidentIdOk() (*string, bool)`

GetIncidentIdOk returns a tuple with the IncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentId

`func (o *IncidentByIdResponse) SetIncidentId(v string)`

SetIncidentId sets IncidentId field to given value.

### HasIncidentId

`func (o *IncidentByIdResponse) HasIncidentId() bool`

HasIncidentId returns a boolean if a field has been set.

### GetIncidentImpact

`func (o *IncidentByIdResponse) GetIncidentImpact() string`

GetIncidentImpact returns the IncidentImpact field if non-nil, zero value otherwise.

### GetIncidentImpactOk

`func (o *IncidentByIdResponse) GetIncidentImpactOk() (*string, bool)`

GetIncidentImpactOk returns a tuple with the IncidentImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentImpact

`func (o *IncidentByIdResponse) SetIncidentImpact(v string)`

SetIncidentImpact sets IncidentImpact field to given value.

### HasIncidentImpact

`func (o *IncidentByIdResponse) HasIncidentImpact() bool`

HasIncidentImpact returns a boolean if a field has been set.

### GetIncidentMessage

`func (o *IncidentByIdResponse) GetIncidentMessage() string`

GetIncidentMessage returns the IncidentMessage field if non-nil, zero value otherwise.

### GetIncidentMessageOk

`func (o *IncidentByIdResponse) GetIncidentMessageOk() (*string, bool)`

GetIncidentMessageOk returns a tuple with the IncidentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentMessage

`func (o *IncidentByIdResponse) SetIncidentMessage(v string)`

SetIncidentMessage sets IncidentMessage field to given value.

### HasIncidentMessage

`func (o *IncidentByIdResponse) HasIncidentMessage() bool`

HasIncidentMessage returns a boolean if a field has been set.

### GetIncidentStartTime

`func (o *IncidentByIdResponse) GetIncidentStartTime() time.Time`

GetIncidentStartTime returns the IncidentStartTime field if non-nil, zero value otherwise.

### GetIncidentStartTimeOk

`func (o *IncidentByIdResponse) GetIncidentStartTimeOk() (*time.Time, bool)`

GetIncidentStartTimeOk returns a tuple with the IncidentStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentStartTime

`func (o *IncidentByIdResponse) SetIncidentStartTime(v time.Time)`

SetIncidentStartTime sets IncidentStartTime field to given value.

### HasIncidentStartTime

`func (o *IncidentByIdResponse) HasIncidentStartTime() bool`

HasIncidentStartTime returns a boolean if a field has been set.

### GetIncidentEndTime

`func (o *IncidentByIdResponse) GetIncidentEndTime() time.Time`

GetIncidentEndTime returns the IncidentEndTime field if non-nil, zero value otherwise.

### GetIncidentEndTimeOk

`func (o *IncidentByIdResponse) GetIncidentEndTimeOk() (*time.Time, bool)`

GetIncidentEndTimeOk returns a tuple with the IncidentEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentEndTime

`func (o *IncidentByIdResponse) SetIncidentEndTime(v time.Time)`

SetIncidentEndTime sets IncidentEndTime field to given value.

### HasIncidentEndTime

`func (o *IncidentByIdResponse) HasIncidentEndTime() bool`

HasIncidentEndTime returns a boolean if a field has been set.

### SetIncidentEndTimeNil

`func (o *IncidentByIdResponse) SetIncidentEndTimeNil(b bool)`

 SetIncidentEndTimeNil sets the value for IncidentEndTime to be an explicit nil

### UnsetIncidentEndTime
`func (o *IncidentByIdResponse) UnsetIncidentEndTime()`

UnsetIncidentEndTime ensures that no value is present for IncidentEndTime, not even an explicit nil
### GetIncidentStatus

`func (o *IncidentByIdResponse) GetIncidentStatus() string`

GetIncidentStatus returns the IncidentStatus field if non-nil, zero value otherwise.

### GetIncidentStatusOk

`func (o *IncidentByIdResponse) GetIncidentStatusOk() (*string, bool)`

GetIncidentStatusOk returns a tuple with the IncidentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentStatus

`func (o *IncidentByIdResponse) SetIncidentStatus(v string)`

SetIncidentStatus sets IncidentStatus field to given value.

### HasIncidentStatus

`func (o *IncidentByIdResponse) HasIncidentStatus() bool`

HasIncidentStatus returns a boolean if a field has been set.

### GetIncidentType

`func (o *IncidentByIdResponse) GetIncidentType() string`

GetIncidentType returns the IncidentType field if non-nil, zero value otherwise.

### GetIncidentTypeOk

`func (o *IncidentByIdResponse) GetIncidentTypeOk() (*string, bool)`

GetIncidentTypeOk returns a tuple with the IncidentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentType

`func (o *IncidentByIdResponse) SetIncidentType(v string)`

SetIncidentType sets IncidentType field to given value.

### HasIncidentType

`func (o *IncidentByIdResponse) HasIncidentType() bool`

HasIncidentType returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *IncidentByIdResponse) GetPaymentMethod() IncidentObjectPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *IncidentByIdResponse) GetPaymentMethodOk() (*IncidentObjectPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *IncidentByIdResponse) SetPaymentMethod(v IncidentObjectPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *IncidentByIdResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


