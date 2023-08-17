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


// PublicationsVerbatimAPIService PublicationsVerbatimAPI service
type PublicationsVerbatimAPIService service

type PublicationsVerbatimAPIPublicationsDebVerbatimCreateRequest struct {
	ctx context.Context
	ApiService *PublicationsVerbatimAPIService
	debVerbatimPublication *DebVerbatimPublication
}

func (r PublicationsVerbatimAPIPublicationsDebVerbatimCreateRequest) DebVerbatimPublication(debVerbatimPublication DebVerbatimPublication) PublicationsVerbatimAPIPublicationsDebVerbatimCreateRequest {
	r.debVerbatimPublication = &debVerbatimPublication
	return r
}

func (r PublicationsVerbatimAPIPublicationsDebVerbatimCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PublicationsDebVerbatimCreateExecute(r)
}

/*
PublicationsDebVerbatimCreate Create a verbatim publication

Trigger an asynchronous task to publish content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PublicationsVerbatimAPIPublicationsDebVerbatimCreateRequest
*/
func (a *PublicationsVerbatimAPIService) PublicationsDebVerbatimCreate(ctx context.Context) PublicationsVerbatimAPIPublicationsDebVerbatimCreateRequest {
	return PublicationsVerbatimAPIPublicationsDebVerbatimCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *PublicationsVerbatimAPIService) PublicationsDebVerbatimCreateExecute(r PublicationsVerbatimAPIPublicationsDebVerbatimCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsVerbatimAPIService.PublicationsDebVerbatimCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/publications/deb/verbatim/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debVerbatimPublication == nil {
		return localVarReturnValue, nil, reportError("debVerbatimPublication is required and must be specified")
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
	localVarPostBody = r.debVerbatimPublication
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

type PublicationsVerbatimAPIPublicationsDebVerbatimDeleteRequest struct {
	ctx context.Context
	ApiService *PublicationsVerbatimAPIService
	debVerbatimPublicationHref string
}

func (r PublicationsVerbatimAPIPublicationsDebVerbatimDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.PublicationsDebVerbatimDeleteExecute(r)
}

/*
PublicationsDebVerbatimDelete Delete a verbatim publication

An VerbatimPublication is the Pulp-internal representation of a "mirrored" AptRepositoryVersion.

In other words, the verbatim publisher will recreate the synced subset of some a APT
repository using the exact same metadata files and signatures as used by the upstream original.
Once a Pulp publication has been created, it can be served by creating a Pulp distribution (in
a near atomic action).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debVerbatimPublicationHref
 @return PublicationsVerbatimAPIPublicationsDebVerbatimDeleteRequest
*/
func (a *PublicationsVerbatimAPIService) PublicationsDebVerbatimDelete(ctx context.Context, debVerbatimPublicationHref string) PublicationsVerbatimAPIPublicationsDebVerbatimDeleteRequest {
	return PublicationsVerbatimAPIPublicationsDebVerbatimDeleteRequest{
		ApiService: a,
		ctx: ctx,
		debVerbatimPublicationHref: debVerbatimPublicationHref,
	}
}

// Execute executes the request
func (a *PublicationsVerbatimAPIService) PublicationsDebVerbatimDeleteExecute(r PublicationsVerbatimAPIPublicationsDebVerbatimDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsVerbatimAPIService.PublicationsDebVerbatimDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_verbatim_publication_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_verbatim_publication_href"+"}", parameterValueToString(r.debVerbatimPublicationHref, "debVerbatimPublicationHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PublicationsVerbatimAPIPublicationsDebVerbatimListRequest struct {
	ctx context.Context
	ApiService *PublicationsVerbatimAPIService
	content *string
	contentIn *string
	limit *int32
	offset *int32
	ordering *[]string
	pulpCreated *time.Time
	pulpCreatedGt *time.Time
	pulpCreatedGte *time.Time
	pulpCreatedLt *time.Time
	pulpCreatedLte *time.Time
	pulpCreatedRange *[]time.Time
	pulpHrefIn *[]string
	pulpIdIn *[]string
	repository *string
	repositoryVersion *string
	fields *[]string
	excludeFields *[]string
}

// Content Unit referenced by HREF
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) Content(content string) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) ContentIn(contentIn string) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) Limit(limit int32) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) Offset(offset int32) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;complete&#x60; - Complete * &#x60;-complete&#x60; - Complete (descending) * &#x60;pass_through&#x60; - Pass through * &#x60;-pass_through&#x60; - Pass through (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) Ordering(ordering []string) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pulp_created matches value
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) PulpCreated(pulpCreated time.Time) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) PulpCreatedGt(pulpCreatedGt time.Time) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) PulpCreatedGte(pulpCreatedGte time.Time) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) PulpCreatedLt(pulpCreatedLt time.Time) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) PulpCreatedLte(pulpCreatedLte time.Time) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// Multiple values may be separated by commas.
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) PulpHrefIn(pulpHrefIn []string) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) PulpIdIn(pulpIdIn []string) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository referenced by HREF
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) Repository(repository string) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.repository = &repository
	return r
}

// Repository Version referenced by HREF
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) RepositoryVersion(repositoryVersion string) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// A list of fields to include in the response.
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) Fields(fields []string) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) ExcludeFields(excludeFields []string) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) Execute() (*PaginateddebVerbatimPublicationResponseList, *http.Response, error) {
	return r.ApiService.PublicationsDebVerbatimListExecute(r)
}

/*
PublicationsDebVerbatimList List verbatim publications

An VerbatimPublication is the Pulp-internal representation of a "mirrored" AptRepositoryVersion.

In other words, the verbatim publisher will recreate the synced subset of some a APT
repository using the exact same metadata files and signatures as used by the upstream original.
Once a Pulp publication has been created, it can be served by creating a Pulp distribution (in
a near atomic action).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PublicationsVerbatimAPIPublicationsDebVerbatimListRequest
*/
func (a *PublicationsVerbatimAPIService) PublicationsDebVerbatimList(ctx context.Context) PublicationsVerbatimAPIPublicationsDebVerbatimListRequest {
	return PublicationsVerbatimAPIPublicationsDebVerbatimListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebVerbatimPublicationResponseList
func (a *PublicationsVerbatimAPIService) PublicationsDebVerbatimListExecute(r PublicationsVerbatimAPIPublicationsDebVerbatimListRequest) (*PaginateddebVerbatimPublicationResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebVerbatimPublicationResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsVerbatimAPIService.PublicationsDebVerbatimList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/publications/deb/verbatim/"

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
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
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

type PublicationsVerbatimAPIPublicationsDebVerbatimReadRequest struct {
	ctx context.Context
	ApiService *PublicationsVerbatimAPIService
	debVerbatimPublicationHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PublicationsVerbatimAPIPublicationsDebVerbatimReadRequest) Fields(fields []string) PublicationsVerbatimAPIPublicationsDebVerbatimReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PublicationsVerbatimAPIPublicationsDebVerbatimReadRequest) ExcludeFields(excludeFields []string) PublicationsVerbatimAPIPublicationsDebVerbatimReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PublicationsVerbatimAPIPublicationsDebVerbatimReadRequest) Execute() (*DebVerbatimPublicationResponse, *http.Response, error) {
	return r.ApiService.PublicationsDebVerbatimReadExecute(r)
}

/*
PublicationsDebVerbatimRead Inspect a verbatim publication

An VerbatimPublication is the Pulp-internal representation of a "mirrored" AptRepositoryVersion.

In other words, the verbatim publisher will recreate the synced subset of some a APT
repository using the exact same metadata files and signatures as used by the upstream original.
Once a Pulp publication has been created, it can be served by creating a Pulp distribution (in
a near atomic action).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debVerbatimPublicationHref
 @return PublicationsVerbatimAPIPublicationsDebVerbatimReadRequest
*/
func (a *PublicationsVerbatimAPIService) PublicationsDebVerbatimRead(ctx context.Context, debVerbatimPublicationHref string) PublicationsVerbatimAPIPublicationsDebVerbatimReadRequest {
	return PublicationsVerbatimAPIPublicationsDebVerbatimReadRequest{
		ApiService: a,
		ctx: ctx,
		debVerbatimPublicationHref: debVerbatimPublicationHref,
	}
}

// Execute executes the request
//  @return DebVerbatimPublicationResponse
func (a *PublicationsVerbatimAPIService) PublicationsDebVerbatimReadExecute(r PublicationsVerbatimAPIPublicationsDebVerbatimReadRequest) (*DebVerbatimPublicationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebVerbatimPublicationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsVerbatimAPIService.PublicationsDebVerbatimRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_verbatim_publication_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_verbatim_publication_href"+"}", parameterValueToString(r.debVerbatimPublicationHref, "debVerbatimPublicationHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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