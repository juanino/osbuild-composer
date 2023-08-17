/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"encoding/json"
)

// checks if the CollectionNamespaceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionNamespaceResponse{}

// CollectionNamespaceResponse A serializer for a Collection Version namespace field.
type CollectionNamespaceResponse struct {
	Name string `json:"name"`
	MetadataSha256 NullableString `json:"metadata_sha256"`
	AdditionalProperties map[string]interface{}
}

type _CollectionNamespaceResponse CollectionNamespaceResponse

// NewCollectionNamespaceResponse instantiates a new CollectionNamespaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionNamespaceResponse(name string, metadataSha256 NullableString) *CollectionNamespaceResponse {
	this := CollectionNamespaceResponse{}
	this.Name = name
	this.MetadataSha256 = metadataSha256
	return &this
}

// NewCollectionNamespaceResponseWithDefaults instantiates a new CollectionNamespaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionNamespaceResponseWithDefaults() *CollectionNamespaceResponse {
	this := CollectionNamespaceResponse{}
	return &this
}

// GetName returns the Name field value
func (o *CollectionNamespaceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CollectionNamespaceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CollectionNamespaceResponse) SetName(v string) {
	o.Name = v
}

// GetMetadataSha256 returns the MetadataSha256 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CollectionNamespaceResponse) GetMetadataSha256() string {
	if o == nil || o.MetadataSha256.Get() == nil {
		var ret string
		return ret
	}

	return *o.MetadataSha256.Get()
}

// GetMetadataSha256Ok returns a tuple with the MetadataSha256 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionNamespaceResponse) GetMetadataSha256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataSha256.Get(), o.MetadataSha256.IsSet()
}

// SetMetadataSha256 sets field value
func (o *CollectionNamespaceResponse) SetMetadataSha256(v string) {
	o.MetadataSha256.Set(&v)
}

func (o CollectionNamespaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionNamespaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["metadata_sha256"] = o.MetadataSha256.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionNamespaceResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionNamespaceResponse := _CollectionNamespaceResponse{}

	if err = json.Unmarshal(bytes, &varCollectionNamespaceResponse); err == nil {
		*o = CollectionNamespaceResponse(varCollectionNamespaceResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "metadata_sha256")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionNamespaceResponse struct {
	value *CollectionNamespaceResponse
	isSet bool
}

func (v NullableCollectionNamespaceResponse) Get() *CollectionNamespaceResponse {
	return v.value
}

func (v *NullableCollectionNamespaceResponse) Set(val *CollectionNamespaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionNamespaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionNamespaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionNamespaceResponse(val *CollectionNamespaceResponse) *NullableCollectionNamespaceResponse {
	return &NullableCollectionNamespaceResponse{value: val, isSet: true}
}

func (v NullableCollectionNamespaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionNamespaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

