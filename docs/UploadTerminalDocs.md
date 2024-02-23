# UploadTerminalDocs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocType** | **string** | Mention the document type you are uploading. Possible values - ADDRESSPROOF, PHOTOGRAPH. | 
**DocValue** | **string** | Enter the display name of the uploaded file. | 
**File** | **string** | Select the document that should be uploaded or provide the path of that file. You cannot upload a file that is more than 2MB in size. | 

## Methods

### NewUploadTerminalDocs

`func NewUploadTerminalDocs(docType string, docValue string, file string, ) *UploadTerminalDocs`

NewUploadTerminalDocs instantiates a new UploadTerminalDocs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadTerminalDocsWithDefaults

`func NewUploadTerminalDocsWithDefaults() *UploadTerminalDocs`

NewUploadTerminalDocsWithDefaults instantiates a new UploadTerminalDocs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocType

`func (o *UploadTerminalDocs) GetDocType() string`

GetDocType returns the DocType field if non-nil, zero value otherwise.

### GetDocTypeOk

`func (o *UploadTerminalDocs) GetDocTypeOk() (*string, bool)`

GetDocTypeOk returns a tuple with the DocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocType

`func (o *UploadTerminalDocs) SetDocType(v string)`

SetDocType sets DocType field to given value.


### GetDocValue

`func (o *UploadTerminalDocs) GetDocValue() string`

GetDocValue returns the DocValue field if non-nil, zero value otherwise.

### GetDocValueOk

`func (o *UploadTerminalDocs) GetDocValueOk() (*string, bool)`

GetDocValueOk returns a tuple with the DocValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocValue

`func (o *UploadTerminalDocs) SetDocValue(v string)`

SetDocValue sets DocValue field to given value.


### GetFile

`func (o *UploadTerminalDocs) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UploadTerminalDocs) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UploadTerminalDocs) SetFile(v string)`

SetFile sets File field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


