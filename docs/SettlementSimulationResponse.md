# SettlementSimulationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SimulationId** | Pointer to **string** | A unique identifier for the simulation request. | [optional] 
**Entity** | Pointer to **string** | Entity type for which the simulation is performed. Example: \&quot;SETTLEMENTS\&quot;. | [optional] 
**SettlementIds** | Pointer to **[]float32** | List of simulated settlement IDs. | [optional] 
**SimulationStatus** | Pointer to **string** | Status of the simulation request. Example: \&quot;SUCCESS/FAILED/PENDING\&quot;. | [optional] 

## Methods

### NewSettlementSimulationResponse

`func NewSettlementSimulationResponse() *SettlementSimulationResponse`

NewSettlementSimulationResponse instantiates a new SettlementSimulationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementSimulationResponseWithDefaults

`func NewSettlementSimulationResponseWithDefaults() *SettlementSimulationResponse`

NewSettlementSimulationResponseWithDefaults instantiates a new SettlementSimulationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimulationId

`func (o *SettlementSimulationResponse) GetSimulationId() string`

GetSimulationId returns the SimulationId field if non-nil, zero value otherwise.

### GetSimulationIdOk

`func (o *SettlementSimulationResponse) GetSimulationIdOk() (*string, bool)`

GetSimulationIdOk returns a tuple with the SimulationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationId

`func (o *SettlementSimulationResponse) SetSimulationId(v string)`

SetSimulationId sets SimulationId field to given value.

### HasSimulationId

`func (o *SettlementSimulationResponse) HasSimulationId() bool`

HasSimulationId returns a boolean if a field has been set.

### GetEntity

`func (o *SettlementSimulationResponse) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SettlementSimulationResponse) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SettlementSimulationResponse) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *SettlementSimulationResponse) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetSettlementIds

`func (o *SettlementSimulationResponse) GetSettlementIds() []float32`

GetSettlementIds returns the SettlementIds field if non-nil, zero value otherwise.

### GetSettlementIdsOk

`func (o *SettlementSimulationResponse) GetSettlementIdsOk() (*[]float32, bool)`

GetSettlementIdsOk returns a tuple with the SettlementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementIds

`func (o *SettlementSimulationResponse) SetSettlementIds(v []float32)`

SetSettlementIds sets SettlementIds field to given value.

### HasSettlementIds

`func (o *SettlementSimulationResponse) HasSettlementIds() bool`

HasSettlementIds returns a boolean if a field has been set.

### GetSimulationStatus

`func (o *SettlementSimulationResponse) GetSimulationStatus() string`

GetSimulationStatus returns the SimulationStatus field if non-nil, zero value otherwise.

### GetSimulationStatusOk

`func (o *SettlementSimulationResponse) GetSimulationStatusOk() (*string, bool)`

GetSimulationStatusOk returns a tuple with the SimulationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationStatus

`func (o *SettlementSimulationResponse) SetSimulationStatus(v string)`

SetSimulationStatus sets SimulationStatus field to given value.

### HasSimulationStatus

`func (o *SettlementSimulationResponse) HasSimulationStatus() bool`

HasSimulationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


