# IncidentObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncidentId** | Pointer to **string** | Unique identifier for the incident. | [optional] 
**IncidentImpact** | Pointer to **string** | Impact level of the incident. | [optional] 
**IncidentMessage** | Pointer to **string** | Description of the incident. | [optional] 
**IncidentStartTime** | Pointer to **string** | Start time of the incident. | [optional] 
**IncidentEndTime** | Pointer to **NullableString** | End time of the incident, or null if ongoing. | [optional] 
**IncidentStatus** | Pointer to **string** | Status of the incident. | [optional] 
**IncidentType** | Pointer to **string** | Type of the incident. | [optional] 
**PaymentMethod** | Pointer to [**IncidentObjectPaymentMethod**](IncidentObjectPaymentMethod.md) |  | [optional] 

## Methods

### NewIncidentObject

`func NewIncidentObject() *IncidentObject`

NewIncidentObject instantiates a new IncidentObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentObjectWithDefaults

`func NewIncidentObjectWithDefaults() *IncidentObject`

NewIncidentObjectWithDefaults instantiates a new IncidentObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidentId

`func (o *IncidentObject) GetIncidentId() string`

GetIncidentId returns the IncidentId field if non-nil, zero value otherwise.

### GetIncidentIdOk

`func (o *IncidentObject) GetIncidentIdOk() (*string, bool)`

GetIncidentIdOk returns a tuple with the IncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentId

`func (o *IncidentObject) SetIncidentId(v string)`

SetIncidentId sets IncidentId field to given value.

### HasIncidentId

`func (o *IncidentObject) HasIncidentId() bool`

HasIncidentId returns a boolean if a field has been set.

### GetIncidentImpact

`func (o *IncidentObject) GetIncidentImpact() string`

GetIncidentImpact returns the IncidentImpact field if non-nil, zero value otherwise.

### GetIncidentImpactOk

`func (o *IncidentObject) GetIncidentImpactOk() (*string, bool)`

GetIncidentImpactOk returns a tuple with the IncidentImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentImpact

`func (o *IncidentObject) SetIncidentImpact(v string)`

SetIncidentImpact sets IncidentImpact field to given value.

### HasIncidentImpact

`func (o *IncidentObject) HasIncidentImpact() bool`

HasIncidentImpact returns a boolean if a field has been set.

### GetIncidentMessage

`func (o *IncidentObject) GetIncidentMessage() string`

GetIncidentMessage returns the IncidentMessage field if non-nil, zero value otherwise.

### GetIncidentMessageOk

`func (o *IncidentObject) GetIncidentMessageOk() (*string, bool)`

GetIncidentMessageOk returns a tuple with the IncidentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentMessage

`func (o *IncidentObject) SetIncidentMessage(v string)`

SetIncidentMessage sets IncidentMessage field to given value.

### HasIncidentMessage

`func (o *IncidentObject) HasIncidentMessage() bool`

HasIncidentMessage returns a boolean if a field has been set.

### GetIncidentStartTime

`func (o *IncidentObject) GetIncidentStartTime() string`

GetIncidentStartTime returns the IncidentStartTime field if non-nil, zero value otherwise.

### GetIncidentStartTimeOk

`func (o *IncidentObject) GetIncidentStartTimeOk() (*string, bool)`

GetIncidentStartTimeOk returns a tuple with the IncidentStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentStartTime

`func (o *IncidentObject) SetIncidentStartTime(v string)`

SetIncidentStartTime sets IncidentStartTime field to given value.

### HasIncidentStartTime

`func (o *IncidentObject) HasIncidentStartTime() bool`

HasIncidentStartTime returns a boolean if a field has been set.

### GetIncidentEndTime

`func (o *IncidentObject) GetIncidentEndTime() string`

GetIncidentEndTime returns the IncidentEndTime field if non-nil, zero value otherwise.

### GetIncidentEndTimeOk

`func (o *IncidentObject) GetIncidentEndTimeOk() (*string, bool)`

GetIncidentEndTimeOk returns a tuple with the IncidentEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentEndTime

`func (o *IncidentObject) SetIncidentEndTime(v string)`

SetIncidentEndTime sets IncidentEndTime field to given value.

### HasIncidentEndTime

`func (o *IncidentObject) HasIncidentEndTime() bool`

HasIncidentEndTime returns a boolean if a field has been set.

### SetIncidentEndTimeNil

`func (o *IncidentObject) SetIncidentEndTimeNil(b bool)`

 SetIncidentEndTimeNil sets the value for IncidentEndTime to be an explicit nil

### UnsetIncidentEndTime
`func (o *IncidentObject) UnsetIncidentEndTime()`

UnsetIncidentEndTime ensures that no value is present for IncidentEndTime, not even an explicit nil
### GetIncidentStatus

`func (o *IncidentObject) GetIncidentStatus() string`

GetIncidentStatus returns the IncidentStatus field if non-nil, zero value otherwise.

### GetIncidentStatusOk

`func (o *IncidentObject) GetIncidentStatusOk() (*string, bool)`

GetIncidentStatusOk returns a tuple with the IncidentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentStatus

`func (o *IncidentObject) SetIncidentStatus(v string)`

SetIncidentStatus sets IncidentStatus field to given value.

### HasIncidentStatus

`func (o *IncidentObject) HasIncidentStatus() bool`

HasIncidentStatus returns a boolean if a field has been set.

### GetIncidentType

`func (o *IncidentObject) GetIncidentType() string`

GetIncidentType returns the IncidentType field if non-nil, zero value otherwise.

### GetIncidentTypeOk

`func (o *IncidentObject) GetIncidentTypeOk() (*string, bool)`

GetIncidentTypeOk returns a tuple with the IncidentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentType

`func (o *IncidentObject) SetIncidentType(v string)`

SetIncidentType sets IncidentType field to given value.

### HasIncidentType

`func (o *IncidentObject) HasIncidentType() bool`

HasIncidentType returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *IncidentObject) GetPaymentMethod() IncidentObjectPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *IncidentObject) GetPaymentMethodOk() (*IncidentObjectPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *IncidentObject) SetPaymentMethod(v IncidentObjectPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *IncidentObject) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


