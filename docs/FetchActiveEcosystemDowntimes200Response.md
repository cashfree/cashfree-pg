# FetchActiveEcosystemDowntimes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Downtimes** | Pointer to [**[]DowntimeObject**](DowntimeObject.md) |  | [optional] 

## Methods

### NewFetchActiveEcosystemDowntimes200Response

`func NewFetchActiveEcosystemDowntimes200Response() *FetchActiveEcosystemDowntimes200Response`

NewFetchActiveEcosystemDowntimes200Response instantiates a new FetchActiveEcosystemDowntimes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchActiveEcosystemDowntimes200ResponseWithDefaults

`func NewFetchActiveEcosystemDowntimes200ResponseWithDefaults() *FetchActiveEcosystemDowntimes200Response`

NewFetchActiveEcosystemDowntimes200ResponseWithDefaults instantiates a new FetchActiveEcosystemDowntimes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FetchActiveEcosystemDowntimes200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FetchActiveEcosystemDowntimes200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FetchActiveEcosystemDowntimes200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FetchActiveEcosystemDowntimes200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDowntimes

`func (o *FetchActiveEcosystemDowntimes200Response) GetDowntimes() []DowntimeObject`

GetDowntimes returns the Downtimes field if non-nil, zero value otherwise.

### GetDowntimesOk

`func (o *FetchActiveEcosystemDowntimes200Response) GetDowntimesOk() (*[]DowntimeObject, bool)`

GetDowntimesOk returns a tuple with the Downtimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimes

`func (o *FetchActiveEcosystemDowntimes200Response) SetDowntimes(v []DowntimeObject)`

SetDowntimes sets Downtimes field to given value.

### HasDowntimes

`func (o *FetchActiveEcosystemDowntimes200Response) HasDowntimes() bool`

HasDowntimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


