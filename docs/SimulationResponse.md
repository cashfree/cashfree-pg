# SimulationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SimulationId** | Pointer to **string** |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **int64** |  | [optional] 
**EntitySimulation** | Pointer to [**EntitySimulationResponse**](EntitySimulationResponse.md) |  | [optional] 

## Methods

### NewSimulationResponse

`func NewSimulationResponse() *SimulationResponse`

NewSimulationResponse instantiates a new SimulationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulationResponseWithDefaults

`func NewSimulationResponseWithDefaults() *SimulationResponse`

NewSimulationResponseWithDefaults instantiates a new SimulationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimulationId

`func (o *SimulationResponse) GetSimulationId() string`

GetSimulationId returns the SimulationId field if non-nil, zero value otherwise.

### GetSimulationIdOk

`func (o *SimulationResponse) GetSimulationIdOk() (*string, bool)`

GetSimulationIdOk returns a tuple with the SimulationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationId

`func (o *SimulationResponse) SetSimulationId(v string)`

SetSimulationId sets SimulationId field to given value.

### HasSimulationId

`func (o *SimulationResponse) HasSimulationId() bool`

HasSimulationId returns a boolean if a field has been set.

### GetEntity

`func (o *SimulationResponse) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SimulationResponse) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SimulationResponse) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *SimulationResponse) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEntityId

`func (o *SimulationResponse) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SimulationResponse) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SimulationResponse) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SimulationResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntitySimulation

`func (o *SimulationResponse) GetEntitySimulation() EntitySimulationResponse`

GetEntitySimulation returns the EntitySimulation field if non-nil, zero value otherwise.

### GetEntitySimulationOk

`func (o *SimulationResponse) GetEntitySimulationOk() (*EntitySimulationResponse, bool)`

GetEntitySimulationOk returns a tuple with the EntitySimulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySimulation

`func (o *SimulationResponse) SetEntitySimulation(v EntitySimulationResponse)`

SetEntitySimulation sets EntitySimulation field to given value.

### HasEntitySimulation

`func (o *SimulationResponse) HasEntitySimulation() bool`

HasEntitySimulation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


