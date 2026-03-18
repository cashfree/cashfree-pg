# DowntimeObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DowntimeId** | Pointer to **string** | Unique identifier for the downtime. | [optional] 
**DowntimeImpact** | Pointer to **string** | Impact level of the downtime. | [optional] 
**DowntimeMessage** | Pointer to **string** | Description of the downtime. | [optional] 
**DowntimeStartTime** | Pointer to **time.Time** | Start time of the downtime. | [optional] 
**DowntimeEndTime** | Pointer to **NullableTime** | End time of the downtime, or null if ongoing. | [optional] 
**DowntimeStatus** | Pointer to **string** | Status of the downtime (OPEN, UPDATE, RESOLVED). | [optional] 
**DowntimeType** | Pointer to **string** | Type of the downtime (SCHEDULED or UNSCHEDULED). | [optional] 
**PaymentMethod** | Pointer to [**DowntimeObjectPaymentMethod**](DowntimeObjectPaymentMethod.md) |  | [optional] 

## Methods

### NewDowntimeObject

`func NewDowntimeObject() *DowntimeObject`

NewDowntimeObject instantiates a new DowntimeObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDowntimeObjectWithDefaults

`func NewDowntimeObjectWithDefaults() *DowntimeObject`

NewDowntimeObjectWithDefaults instantiates a new DowntimeObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDowntimeId

`func (o *DowntimeObject) GetDowntimeId() string`

GetDowntimeId returns the DowntimeId field if non-nil, zero value otherwise.

### GetDowntimeIdOk

`func (o *DowntimeObject) GetDowntimeIdOk() (*string, bool)`

GetDowntimeIdOk returns a tuple with the DowntimeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeId

`func (o *DowntimeObject) SetDowntimeId(v string)`

SetDowntimeId sets DowntimeId field to given value.

### HasDowntimeId

`func (o *DowntimeObject) HasDowntimeId() bool`

HasDowntimeId returns a boolean if a field has been set.

### GetDowntimeImpact

`func (o *DowntimeObject) GetDowntimeImpact() string`

GetDowntimeImpact returns the DowntimeImpact field if non-nil, zero value otherwise.

### GetDowntimeImpactOk

`func (o *DowntimeObject) GetDowntimeImpactOk() (*string, bool)`

GetDowntimeImpactOk returns a tuple with the DowntimeImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeImpact

`func (o *DowntimeObject) SetDowntimeImpact(v string)`

SetDowntimeImpact sets DowntimeImpact field to given value.

### HasDowntimeImpact

`func (o *DowntimeObject) HasDowntimeImpact() bool`

HasDowntimeImpact returns a boolean if a field has been set.

### GetDowntimeMessage

`func (o *DowntimeObject) GetDowntimeMessage() string`

GetDowntimeMessage returns the DowntimeMessage field if non-nil, zero value otherwise.

### GetDowntimeMessageOk

`func (o *DowntimeObject) GetDowntimeMessageOk() (*string, bool)`

GetDowntimeMessageOk returns a tuple with the DowntimeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeMessage

`func (o *DowntimeObject) SetDowntimeMessage(v string)`

SetDowntimeMessage sets DowntimeMessage field to given value.

### HasDowntimeMessage

`func (o *DowntimeObject) HasDowntimeMessage() bool`

HasDowntimeMessage returns a boolean if a field has been set.

### GetDowntimeStartTime

`func (o *DowntimeObject) GetDowntimeStartTime() time.Time`

GetDowntimeStartTime returns the DowntimeStartTime field if non-nil, zero value otherwise.

### GetDowntimeStartTimeOk

`func (o *DowntimeObject) GetDowntimeStartTimeOk() (*time.Time, bool)`

GetDowntimeStartTimeOk returns a tuple with the DowntimeStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeStartTime

`func (o *DowntimeObject) SetDowntimeStartTime(v time.Time)`

SetDowntimeStartTime sets DowntimeStartTime field to given value.

### HasDowntimeStartTime

`func (o *DowntimeObject) HasDowntimeStartTime() bool`

HasDowntimeStartTime returns a boolean if a field has been set.

### GetDowntimeEndTime

`func (o *DowntimeObject) GetDowntimeEndTime() time.Time`

GetDowntimeEndTime returns the DowntimeEndTime field if non-nil, zero value otherwise.

### GetDowntimeEndTimeOk

`func (o *DowntimeObject) GetDowntimeEndTimeOk() (*time.Time, bool)`

GetDowntimeEndTimeOk returns a tuple with the DowntimeEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeEndTime

`func (o *DowntimeObject) SetDowntimeEndTime(v time.Time)`

SetDowntimeEndTime sets DowntimeEndTime field to given value.

### HasDowntimeEndTime

`func (o *DowntimeObject) HasDowntimeEndTime() bool`

HasDowntimeEndTime returns a boolean if a field has been set.

### SetDowntimeEndTimeNil

`func (o *DowntimeObject) SetDowntimeEndTimeNil(b bool)`

 SetDowntimeEndTimeNil sets the value for DowntimeEndTime to be an explicit nil

### UnsetDowntimeEndTime
`func (o *DowntimeObject) UnsetDowntimeEndTime()`

UnsetDowntimeEndTime ensures that no value is present for DowntimeEndTime, not even an explicit nil
### GetDowntimeStatus

`func (o *DowntimeObject) GetDowntimeStatus() string`

GetDowntimeStatus returns the DowntimeStatus field if non-nil, zero value otherwise.

### GetDowntimeStatusOk

`func (o *DowntimeObject) GetDowntimeStatusOk() (*string, bool)`

GetDowntimeStatusOk returns a tuple with the DowntimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeStatus

`func (o *DowntimeObject) SetDowntimeStatus(v string)`

SetDowntimeStatus sets DowntimeStatus field to given value.

### HasDowntimeStatus

`func (o *DowntimeObject) HasDowntimeStatus() bool`

HasDowntimeStatus returns a boolean if a field has been set.

### GetDowntimeType

`func (o *DowntimeObject) GetDowntimeType() string`

GetDowntimeType returns the DowntimeType field if non-nil, zero value otherwise.

### GetDowntimeTypeOk

`func (o *DowntimeObject) GetDowntimeTypeOk() (*string, bool)`

GetDowntimeTypeOk returns a tuple with the DowntimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeType

`func (o *DowntimeObject) SetDowntimeType(v string)`

SetDowntimeType sets DowntimeType field to given value.

### HasDowntimeType

`func (o *DowntimeObject) HasDowntimeType() bool`

HasDowntimeType returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *DowntimeObject) GetPaymentMethod() DowntimeObjectPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *DowntimeObject) GetPaymentMethodOk() (*DowntimeObjectPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *DowntimeObject) SetPaymentMethod(v DowntimeObjectPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *DowntimeObject) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


