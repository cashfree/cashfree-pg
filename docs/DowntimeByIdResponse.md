# DowntimeByIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the downtime. | [optional] 
**Impact** | Pointer to **string** | Impact level of the downtime (e.g., HIGH, MEDIUM, LOW). | [optional] 
**Message** | Pointer to **string** | Description of the issue. | [optional] 
**StartTime** | Pointer to **time.Time** | Start time of the downtime. | [optional] 
**EndTime** | Pointer to **NullableTime** | End time of the downtime, if applicable; null if still ongoing. | [optional] 
**Status** | Pointer to **string** | Current status of the downtime (e.g., OPEN, UPDATE, RESOLVED). | [optional] 
**Type** | Pointer to **string** | Type/category of the downtime (e.g., SCHEDULED, UNSCHEDULED). | [optional] 
**PaymentMethod** | Pointer to [**DowntimeObjectPaymentMethod**](DowntimeObjectPaymentMethod.md) |  | [optional] 

## Methods

### NewDowntimeByIdResponse

`func NewDowntimeByIdResponse() *DowntimeByIdResponse`

NewDowntimeByIdResponse instantiates a new DowntimeByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDowntimeByIdResponseWithDefaults

`func NewDowntimeByIdResponseWithDefaults() *DowntimeByIdResponse`

NewDowntimeByIdResponseWithDefaults instantiates a new DowntimeByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DowntimeByIdResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DowntimeByIdResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DowntimeByIdResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DowntimeByIdResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpact

`func (o *DowntimeByIdResponse) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *DowntimeByIdResponse) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *DowntimeByIdResponse) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *DowntimeByIdResponse) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetMessage

`func (o *DowntimeByIdResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DowntimeByIdResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DowntimeByIdResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DowntimeByIdResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStartTime

`func (o *DowntimeByIdResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DowntimeByIdResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DowntimeByIdResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DowntimeByIdResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *DowntimeByIdResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DowntimeByIdResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DowntimeByIdResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DowntimeByIdResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *DowntimeByIdResponse) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *DowntimeByIdResponse) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetStatus

`func (o *DowntimeByIdResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DowntimeByIdResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DowntimeByIdResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DowntimeByIdResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *DowntimeByIdResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DowntimeByIdResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DowntimeByIdResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DowntimeByIdResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *DowntimeByIdResponse) GetPaymentMethod() DowntimeObjectPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *DowntimeByIdResponse) GetPaymentMethodOk() (*DowntimeObjectPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *DowntimeByIdResponse) SetPaymentMethod(v DowntimeObjectPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *DowntimeByIdResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


