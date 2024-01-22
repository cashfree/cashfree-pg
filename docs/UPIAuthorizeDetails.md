# UPIAuthorizeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproveBy** | Pointer to **string** | Time by which this authorization should be approved by the customer. | [optional] 
**StartTime** | Pointer to **string** | This is the time when the UPI one time mandate will start | [optional] 
**EndTime** | Pointer to **string** | This is the time when the UPI mandate will be over. If the mandate has not been executed by this time, the funds will be returned back to the customer after this time. | [optional] 

## Methods

### NewUPIAuthorizeDetails

`func NewUPIAuthorizeDetails() *UPIAuthorizeDetails`

NewUPIAuthorizeDetails instantiates a new UPIAuthorizeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUPIAuthorizeDetailsWithDefaults

`func NewUPIAuthorizeDetailsWithDefaults() *UPIAuthorizeDetails`

NewUPIAuthorizeDetailsWithDefaults instantiates a new UPIAuthorizeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproveBy

`func (o *UPIAuthorizeDetails) GetApproveBy() string`

GetApproveBy returns the ApproveBy field if non-nil, zero value otherwise.

### GetApproveByOk

`func (o *UPIAuthorizeDetails) GetApproveByOk() (*string, bool)`

GetApproveByOk returns a tuple with the ApproveBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveBy

`func (o *UPIAuthorizeDetails) SetApproveBy(v string)`

SetApproveBy sets ApproveBy field to given value.

### HasApproveBy

`func (o *UPIAuthorizeDetails) HasApproveBy() bool`

HasApproveBy returns a boolean if a field has been set.

### GetStartTime

`func (o *UPIAuthorizeDetails) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UPIAuthorizeDetails) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UPIAuthorizeDetails) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UPIAuthorizeDetails) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *UPIAuthorizeDetails) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UPIAuthorizeDetails) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UPIAuthorizeDetails) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UPIAuthorizeDetails) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


