/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"reflect"
)


// RepositoriesContainerPushVersionsAPIService RepositoriesContainerPushVersionsAPI service
type RepositoriesContainerPushVersionsAPIService service

type RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *RepositoriesContainerPushVersionsAPIService
	containerContainerPushRepositoryVersionHref string
}

func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesContainerContainerPushVersionsDeleteExecute(r)
}

/*
RepositoriesContainerContainerPushVersionsDelete Delete a repository version

Trigger an asynchronous task to delete a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerContainerPushRepositoryVersionHref
 @return RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsDeleteRequest
*/
func (a *RepositoriesContainerPushVersionsAPIService) RepositoriesContainerContainerPushVersionsDelete(ctx context.Context, containerContainerPushRepositoryVersionHref string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsDeleteRequest {
	return RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		containerContainerPushRepositoryVersionHref: containerContainerPushRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesContainerPushVersionsAPIService) RepositoriesContainerContainerPushVersionsDeleteExecute(r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesContainerPushVersionsAPIService.RepositoriesContainerContainerPushVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{container_container_push_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"container_container_push_repository_version_href"+"}", parameterValueToString(r.containerContainerPushRepositoryVersionHref, "containerContainerPushRepositoryVersionHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest struct {
	ctx context.Context
	ApiService *RepositoriesContainerPushVersionsAPIService
	containerContainerPushRepositoryHref string
	content *string
	contentIn *string
	limit *int32
	number *int32
	numberGt *int32
	numberGte *int32
	numberLt *int32
	numberLte *int32
	numberRange *[]int32
	offset *int32
	ordering *[]string
	pulpCreated *time.Time
	pulpCreatedGt *time.Time
	pulpCreatedGte *time.Time
	pulpCreatedLt *time.Time
	pulpCreatedLte *time.Time
	pulpCreatedRange *[]time.Time
	pulpHrefIn *[]string
	fields *[]string
	excludeFields *[]string
}

// Content Unit referenced by HREF
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) Content(content string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) ContentIn(contentIn string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) Limit(limit int32) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.limit = &limit
	return r
}

// Filter results where number matches value
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) Number(number int32) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.number = &number
	return r
}

// Filter results where number is greater than value
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) NumberGt(numberGt int32) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.numberGt = &numberGt
	return r
}

// Filter results where number is greater than or equal to value
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) NumberGte(numberGte int32) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.numberGte = &numberGte
	return r
}

// Filter results where number is less than value
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) NumberLt(numberLt int32) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.numberLt = &numberLt
	return r
}

// Filter results where number is less than or equal to value
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) NumberLte(numberLte int32) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.numberLte = &numberLte
	return r
}

// Filter results where number is between two comma separated values
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) NumberRange(numberRange []int32) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.numberRange = &numberRange
	return r
}

// The initial index from which to return the results.
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) Offset(offset int32) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;number&#x60; - Number * &#x60;-number&#x60; - Number (descending) * &#x60;complete&#x60; - Complete * &#x60;-complete&#x60; - Complete (descending) * &#x60;info&#x60; - Info * &#x60;-info&#x60; - Info (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) Ordering(ordering []string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pulp_created matches value
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) PulpCreated(pulpCreated time.Time) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// Multiple values may be separated by commas.
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) PulpHrefIn(pulpHrefIn []string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// A list of fields to include in the response.
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) Fields(fields []string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) ExcludeFields(excludeFields []string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) Execute() (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesContainerContainerPushVersionsListExecute(r)
}

/*
RepositoriesContainerContainerPushVersionsList List repository versions

ContainerPushRepositoryVersion represents a single container push repository version.

Repository versions of a push repository are not allowed to be deleted. Versioning of such
repositories, as well as creation/removal, happens automatically without explicit user actions.
Users could make a repository not functional by accident if allowed to delete repository
versions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerContainerPushRepositoryHref
 @return RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest
*/
func (a *RepositoriesContainerPushVersionsAPIService) RepositoriesContainerContainerPushVersionsList(ctx context.Context, containerContainerPushRepositoryHref string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest {
	return RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		containerContainerPushRepositoryHref: containerContainerPushRepositoryHref,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryVersionResponseList
func (a *RepositoriesContainerPushVersionsAPIService) RepositoriesContainerContainerPushVersionsListExecute(r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsListRequest) (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesContainerPushVersionsAPIService.RepositoriesContainerContainerPushVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{container_container_push_repository_href}versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"container_container_push_repository_href"+"}", parameterValueToString(r.containerContainerPushRepositoryHref, "containerContainerPushRepositoryHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.content != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content", r.content, "")
	}
	if r.contentIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content__in", r.contentIn, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "")
	}
	if r.numberGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__gt", r.numberGt, "")
	}
	if r.numberGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__gte", r.numberGte, "")
	}
	if r.numberLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__lt", r.numberLt, "")
	}
	if r.numberLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__lte", r.numberLte, "")
	}
	if r.numberRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__range", r.numberRange, "csv")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created", r.pulpCreated, "")
	}
	if r.pulpCreatedGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gt", r.pulpCreatedGt, "")
	}
	if r.pulpCreatedGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gte", r.pulpCreatedGte, "")
	}
	if r.pulpCreatedLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lt", r.pulpCreatedLt, "")
	}
	if r.pulpCreatedLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lte", r.pulpCreatedLte, "")
	}
	if r.pulpCreatedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__range", r.pulpCreatedRange, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsReadRequest struct {
	ctx context.Context
	ApiService *RepositoriesContainerPushVersionsAPIService
	containerContainerPushRepositoryVersionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsReadRequest) Fields(fields []string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsReadRequest) ExcludeFields(excludeFields []string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsReadRequest) Execute() (*RepositoryVersionResponse, *http.Response, error) {
	return r.ApiService.RepositoriesContainerContainerPushVersionsReadExecute(r)
}

/*
RepositoriesContainerContainerPushVersionsRead Inspect a repository version

ContainerPushRepositoryVersion represents a single container push repository version.

Repository versions of a push repository are not allowed to be deleted. Versioning of such
repositories, as well as creation/removal, happens automatically without explicit user actions.
Users could make a repository not functional by accident if allowed to delete repository
versions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerContainerPushRepositoryVersionHref
 @return RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsReadRequest
*/
func (a *RepositoriesContainerPushVersionsAPIService) RepositoriesContainerContainerPushVersionsRead(ctx context.Context, containerContainerPushRepositoryVersionHref string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsReadRequest {
	return RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		containerContainerPushRepositoryVersionHref: containerContainerPushRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return RepositoryVersionResponse
func (a *RepositoriesContainerPushVersionsAPIService) RepositoriesContainerContainerPushVersionsReadExecute(r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsReadRequest) (*RepositoryVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepositoryVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesContainerPushVersionsAPIService.RepositoriesContainerContainerPushVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{container_container_push_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"container_container_push_repository_version_href"+"}", parameterValueToString(r.containerContainerPushRepositoryVersionHref, "containerContainerPushRepositoryVersionHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsRepairRequest struct {
	ctx context.Context
	ApiService *RepositoriesContainerPushVersionsAPIService
	containerContainerPushRepositoryVersionHref string
	repair *Repair
}

func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsRepairRequest) Repair(repair Repair) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsRepairRequest {
	r.repair = &repair
	return r
}

func (r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsRepairRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesContainerContainerPushVersionsRepairExecute(r)
}

/*
RepositoriesContainerContainerPushVersionsRepair Method for RepositoriesContainerContainerPushVersionsRepair

Trigger an asynchronous task to repair a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerContainerPushRepositoryVersionHref
 @return RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsRepairRequest
*/
func (a *RepositoriesContainerPushVersionsAPIService) RepositoriesContainerContainerPushVersionsRepair(ctx context.Context, containerContainerPushRepositoryVersionHref string) RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsRepairRequest {
	return RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsRepairRequest{
		ApiService: a,
		ctx: ctx,
		containerContainerPushRepositoryVersionHref: containerContainerPushRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesContainerPushVersionsAPIService) RepositoriesContainerContainerPushVersionsRepairExecute(r RepositoriesContainerPushVersionsAPIRepositoriesContainerContainerPushVersionsRepairRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesContainerPushVersionsAPIService.RepositoriesContainerContainerPushVersionsRepair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{container_container_push_repository_version_href}repair/"
	localVarPath = strings.Replace(localVarPath, "{"+"container_container_push_repository_version_href"+"}", parameterValueToString(r.containerContainerPushRepositoryVersionHref, "containerContainerPushRepositoryVersionHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repair == nil {
		return localVarReturnValue, nil, reportError("repair is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.repair
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}