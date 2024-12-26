# SimulateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | **string** | Entity type should be PAYMENTS or SUBS_PAYMENTS only. | 
**EntityId** | **string** | If the entity type is PAYMENTS, the entity_id will be the transactionId. If the entity type is SUBS_PAYMENTS, the entity_id will be the merchantTxnId | 
**EntitySimulation** | [**EntitySimulationRequest**](EntitySimulationRequest.md) |  | 

## Methods

### NewSimulateRequest

`func NewSimulateRequest(entity string, entityId string, entitySimulation EntitySimulationRequest, ) *SimulateRequest`

NewSimulateRequest instantiates a new SimulateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateRequestWithDefaults

`func NewSimulateRequestWithDefaults() *SimulateRequest`

NewSimulateRequestWithDefaults instantiates a new SimulateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *SimulateRequest) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SimulateRequest) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SimulateRequest) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetEntityId

`func (o *SimulateRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SimulateRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SimulateRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntitySimulation

`func (o *SimulateRequest) GetEntitySimulation() EntitySimulationRequest`

GetEntitySimulation returns the EntitySimulation field if non-nil, zero value otherwise.

### GetEntitySimulationOk

`func (o *SimulateRequest) GetEntitySimulationOk() (*EntitySimulationRequest, bool)`

GetEntitySimulationOk returns a tuple with the EntitySimulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySimulation

`func (o *SimulateRequest) SetEntitySimulation(v EntitySimulationRequest)`

SetEntitySimulation sets EntitySimulation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


