# PayOrderAuthorizeOnlyRequestAuthorizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationToken** | Pointer to **string** |  | [optional] 
**DirectoryServerTransactionId** | Pointer to **string** |  | [optional] 
**ThreeDsServerTransactionId** | Pointer to **string** |  | [optional] 
**Eci** | Pointer to **string** |  | [optional] 
**TokenNumber** | Pointer to **string** |  | [optional] 
**TokenExpiryYear** | Pointer to **string** |  | [optional] 
**TokenExpiryMonth** | Pointer to **string** |  | [optional] 
**TokenCryptogram** | Pointer to **string** |  | [optional] 
**TransactionType** | Pointer to **string** | One of ALT_ID, TOKEN, APPLE_PAY, Indicator for authentication mode | [optional] 

## Methods

### NewPayOrderAuthorizeOnlyRequestAuthorizationData

`func NewPayOrderAuthorizeOnlyRequestAuthorizationData() *PayOrderAuthorizeOnlyRequestAuthorizationData`

NewPayOrderAuthorizeOnlyRequestAuthorizationData instantiates a new PayOrderAuthorizeOnlyRequestAuthorizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayOrderAuthorizeOnlyRequestAuthorizationDataWithDefaults

`func NewPayOrderAuthorizeOnlyRequestAuthorizationDataWithDefaults() *PayOrderAuthorizeOnlyRequestAuthorizationData`

NewPayOrderAuthorizeOnlyRequestAuthorizationDataWithDefaults instantiates a new PayOrderAuthorizeOnlyRequestAuthorizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationToken

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetAuthenticationToken() string`

GetAuthenticationToken returns the AuthenticationToken field if non-nil, zero value otherwise.

### GetAuthenticationTokenOk

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetAuthenticationTokenOk() (*string, bool)`

GetAuthenticationTokenOk returns a tuple with the AuthenticationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationToken

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) SetAuthenticationToken(v string)`

SetAuthenticationToken sets AuthenticationToken field to given value.

### HasAuthenticationToken

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) HasAuthenticationToken() bool`

HasAuthenticationToken returns a boolean if a field has been set.

### GetDirectoryServerTransactionId

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetDirectoryServerTransactionId() string`

GetDirectoryServerTransactionId returns the DirectoryServerTransactionId field if non-nil, zero value otherwise.

### GetDirectoryServerTransactionIdOk

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetDirectoryServerTransactionIdOk() (*string, bool)`

GetDirectoryServerTransactionIdOk returns a tuple with the DirectoryServerTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServerTransactionId

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) SetDirectoryServerTransactionId(v string)`

SetDirectoryServerTransactionId sets DirectoryServerTransactionId field to given value.

### HasDirectoryServerTransactionId

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) HasDirectoryServerTransactionId() bool`

HasDirectoryServerTransactionId returns a boolean if a field has been set.

### GetThreeDsServerTransactionId

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetThreeDsServerTransactionId() string`

GetThreeDsServerTransactionId returns the ThreeDsServerTransactionId field if non-nil, zero value otherwise.

### GetThreeDsServerTransactionIdOk

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetThreeDsServerTransactionIdOk() (*string, bool)`

GetThreeDsServerTransactionIdOk returns a tuple with the ThreeDsServerTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsServerTransactionId

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) SetThreeDsServerTransactionId(v string)`

SetThreeDsServerTransactionId sets ThreeDsServerTransactionId field to given value.

### HasThreeDsServerTransactionId

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) HasThreeDsServerTransactionId() bool`

HasThreeDsServerTransactionId returns a boolean if a field has been set.

### GetEci

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetTokenNumber

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetTokenNumber() string`

GetTokenNumber returns the TokenNumber field if non-nil, zero value otherwise.

### GetTokenNumberOk

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetTokenNumberOk() (*string, bool)`

GetTokenNumberOk returns a tuple with the TokenNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumber

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) SetTokenNumber(v string)`

SetTokenNumber sets TokenNumber field to given value.

### HasTokenNumber

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) HasTokenNumber() bool`

HasTokenNumber returns a boolean if a field has been set.

### GetTokenExpiryYear

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetTokenExpiryYear() string`

GetTokenExpiryYear returns the TokenExpiryYear field if non-nil, zero value otherwise.

### GetTokenExpiryYearOk

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetTokenExpiryYearOk() (*string, bool)`

GetTokenExpiryYearOk returns a tuple with the TokenExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiryYear

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) SetTokenExpiryYear(v string)`

SetTokenExpiryYear sets TokenExpiryYear field to given value.

### HasTokenExpiryYear

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) HasTokenExpiryYear() bool`

HasTokenExpiryYear returns a boolean if a field has been set.

### GetTokenExpiryMonth

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetTokenExpiryMonth() string`

GetTokenExpiryMonth returns the TokenExpiryMonth field if non-nil, zero value otherwise.

### GetTokenExpiryMonthOk

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetTokenExpiryMonthOk() (*string, bool)`

GetTokenExpiryMonthOk returns a tuple with the TokenExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiryMonth

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) SetTokenExpiryMonth(v string)`

SetTokenExpiryMonth sets TokenExpiryMonth field to given value.

### HasTokenExpiryMonth

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) HasTokenExpiryMonth() bool`

HasTokenExpiryMonth returns a boolean if a field has been set.

### GetTokenCryptogram

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetTokenCryptogram() string`

GetTokenCryptogram returns the TokenCryptogram field if non-nil, zero value otherwise.

### GetTokenCryptogramOk

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetTokenCryptogramOk() (*string, bool)`

GetTokenCryptogramOk returns a tuple with the TokenCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCryptogram

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) SetTokenCryptogram(v string)`

SetTokenCryptogram sets TokenCryptogram field to given value.

### HasTokenCryptogram

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) HasTokenCryptogram() bool`

HasTokenCryptogram returns a boolean if a field has been set.

### GetTransactionType

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *PayOrderAuthorizeOnlyRequestAuthorizationData) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


