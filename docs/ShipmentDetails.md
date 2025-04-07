# ShipmentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingCompany** | **string** | Tracking company name associated with order. | 
**TrackingUrls** | **[]string** | Tracking Urls associated with order. | 
**TrackingNumbers** | **[]string** | Tracking Numbers associated wih order. | 

## Methods

### NewShipmentDetails

`func NewShipmentDetails(trackingCompany string, trackingUrls []string, trackingNumbers []string, ) *ShipmentDetails`

NewShipmentDetails instantiates a new ShipmentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentDetailsWithDefaults

`func NewShipmentDetailsWithDefaults() *ShipmentDetails`

NewShipmentDetailsWithDefaults instantiates a new ShipmentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingCompany

`func (o *ShipmentDetails) GetTrackingCompany() string`

GetTrackingCompany returns the TrackingCompany field if non-nil, zero value otherwise.

### GetTrackingCompanyOk

`func (o *ShipmentDetails) GetTrackingCompanyOk() (*string, bool)`

GetTrackingCompanyOk returns a tuple with the TrackingCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCompany

`func (o *ShipmentDetails) SetTrackingCompany(v string)`

SetTrackingCompany sets TrackingCompany field to given value.


### GetTrackingUrls

`func (o *ShipmentDetails) GetTrackingUrls() []string`

GetTrackingUrls returns the TrackingUrls field if non-nil, zero value otherwise.

### GetTrackingUrlsOk

`func (o *ShipmentDetails) GetTrackingUrlsOk() (*[]string, bool)`

GetTrackingUrlsOk returns a tuple with the TrackingUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrls

`func (o *ShipmentDetails) SetTrackingUrls(v []string)`

SetTrackingUrls sets TrackingUrls field to given value.


### GetTrackingNumbers

`func (o *ShipmentDetails) GetTrackingNumbers() []string`

GetTrackingNumbers returns the TrackingNumbers field if non-nil, zero value otherwise.

### GetTrackingNumbersOk

`func (o *ShipmentDetails) GetTrackingNumbersOk() (*[]string, bool)`

GetTrackingNumbersOk returns a tuple with the TrackingNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumbers

`func (o *ShipmentDetails) SetTrackingNumbers(v []string)`

SetTrackingNumbers sets TrackingNumbers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


