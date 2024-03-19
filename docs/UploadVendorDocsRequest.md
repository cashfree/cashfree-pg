# UploadVendorDocsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocType** | Pointer to **string** | Mention the type of the document you are uploading. Possible values: UIDAI_FRONT, UIDAI_BACK, UIDAI_NUMBER, DL, DL_NUMBER, PASSPORT_FRONT, PASSPORT_BACK, PASSPORT_NUMBER, VOTER_ID, VOTER_ID_NUMBER, PAN, PAN_NUMBER, GST, GSTIN_NUMBER, CIN, CIN_NUMBER, NBFC_CERTIFICATE. If the doc type ends with a number you should add the doc value else upload the doc file. | [optional] 
**DocValue** | Pointer to ***os.File** | Enter the display name of the uploaded file. | [optional] 
**File** | Pointer to **string** | Select the document that should be uploaded or provide the path of that file. You cannot upload a file that is more than 2MB in size. | [optional] 

## Methods

### NewUploadVendorDocsRequest

`func NewUploadVendorDocsRequest() *UploadVendorDocsRequest`

NewUploadVendorDocsRequest instantiates a new UploadVendorDocsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadVendorDocsRequestWithDefaults

`func NewUploadVendorDocsRequestWithDefaults() *UploadVendorDocsRequest`

NewUploadVendorDocsRequestWithDefaults instantiates a new UploadVendorDocsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocType

`func (o *UploadVendorDocsRequest) GetDocType() string`

GetDocType returns the DocType field if non-nil, zero value otherwise.

### GetDocTypeOk

`func (o *UploadVendorDocsRequest) GetDocTypeOk() (*string, bool)`

GetDocTypeOk returns a tuple with the DocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocType

`func (o *UploadVendorDocsRequest) SetDocType(v string)`

SetDocType sets DocType field to given value.

### HasDocType

`func (o *UploadVendorDocsRequest) HasDocType() bool`

HasDocType returns a boolean if a field has been set.

### GetDocValue

`func (o *UploadVendorDocsRequest) GetDocValue() *os.File`

GetDocValue returns the DocValue field if non-nil, zero value otherwise.

### GetDocValueOk

`func (o *UploadVendorDocsRequest) GetDocValueOk() (**os.File, bool)`

GetDocValueOk returns a tuple with the DocValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocValue

`func (o *UploadVendorDocsRequest) SetDocValue(v *os.File)`

SetDocValue sets DocValue field to given value.

### HasDocValue

`func (o *UploadVendorDocsRequest) HasDocValue() bool`

HasDocValue returns a boolean if a field has been set.

### GetFile

`func (o *UploadVendorDocsRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UploadVendorDocsRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UploadVendorDocsRequest) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *UploadVendorDocsRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


