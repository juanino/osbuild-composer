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
	"os"
	"reflect"
)


// ContentInstallerPackagesAPIService ContentInstallerPackagesAPI service
type ContentInstallerPackagesAPIService service

type ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest struct {
	ctx context.Context
	ApiService *ContentInstallerPackagesAPIService
	repository *string
	artifact *string
	relativePath *string
	file *os.File
	upload *string
}

// A URI of a repository the new content unit should be associated with.
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest) Repository(repository string) ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest {
	r.repository = &repository
	return r
}

// Artifact file representing the physical content
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest) Artifact(artifact string) ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest {
	r.artifact = &artifact
	return r
}

// Path where the artifact is located relative to distributions base_path
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest) RelativePath(relativePath string) ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest {
	r.relativePath = &relativePath
	return r
}

// An uploaded file that may be turned into the artifact of the content unit.
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest) File(file *os.File) ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest {
	r.file = file
	return r
}

// An uncommitted upload that may be turned into the artifact of the content unit.
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest) Upload(upload string) ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest {
	r.upload = &upload
	return r
}

func (r ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentDebInstallerPackagesCreateExecute(r)
}

/*
ContentDebInstallerPackagesCreate Create an installer package

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest
*/
func (a *ContentInstallerPackagesAPIService) ContentDebInstallerPackagesCreate(ctx context.Context) ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest {
	return ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentInstallerPackagesAPIService) ContentDebInstallerPackagesCreateExecute(r ContentInstallerPackagesAPIContentDebInstallerPackagesCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentInstallerPackagesAPIService.ContentDebInstallerPackagesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/installer_packages/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

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
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "repository", r.repository, "")
	}
	if r.artifact != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "artifact", r.artifact, "")
	}
	if r.relativePath != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "relative_path", r.relativePath, "")
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.upload != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "upload", r.upload, "")
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

type ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest struct {
	ctx context.Context
	ApiService *ContentInstallerPackagesAPIService
	architecture *string
	autoBuiltPackage *string
	buildEssential *bool
	builtUsing *string
	essential *bool
	installedSize *int32
	limit *int32
	maintainer *string
	multiArch *string
	offset *int32
	ordering *[]string
	origin *string
	originalMaintainer *string
	package_ *string
	priority *string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	section *string
	sha256 *string
	source *string
	tag *string
	version *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where architecture matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Architecture(architecture string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.architecture = &architecture
	return r
}

// Filter results where auto_built_package matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) AutoBuiltPackage(autoBuiltPackage string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.autoBuiltPackage = &autoBuiltPackage
	return r
}

// Filter results where build_essential matches value  * &#x60;True&#x60; - yes * &#x60;False&#x60; - no
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) BuildEssential(buildEssential bool) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.buildEssential = &buildEssential
	return r
}

// Filter results where built_using matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) BuiltUsing(builtUsing string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.builtUsing = &builtUsing
	return r
}

// Filter results where essential matches value  * &#x60;True&#x60; - yes * &#x60;False&#x60; - no
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Essential(essential bool) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.essential = &essential
	return r
}

// Filter results where installed_size matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) InstalledSize(installedSize int32) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.installedSize = &installedSize
	return r
}

// Number of results to return per page.
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Limit(limit int32) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.limit = &limit
	return r
}

// Filter results where maintainer matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Maintainer(maintainer string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.maintainer = &maintainer
	return r
}

// Filter results where multi_arch matches value  * &#x60;no&#x60; - no * &#x60;same&#x60; - same * &#x60;foreign&#x60; - foreign * &#x60;allowed&#x60; - allowed
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) MultiArch(multiArch string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.multiArch = &multiArch
	return r
}

// The initial index from which to return the results.
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Offset(offset int32) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;package&#x60; - Package * &#x60;-package&#x60; - Package (descending) * &#x60;source&#x60; - Source * &#x60;-source&#x60; - Source (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;architecture&#x60; - Architecture * &#x60;-architecture&#x60; - Architecture (descending) * &#x60;section&#x60; - Section * &#x60;-section&#x60; - Section (descending) * &#x60;priority&#x60; - Priority * &#x60;-priority&#x60; - Priority (descending) * &#x60;origin&#x60; - Origin * &#x60;-origin&#x60; - Origin (descending) * &#x60;tag&#x60; - Tag * &#x60;-tag&#x60; - Tag (descending) * &#x60;bugs&#x60; - Bugs * &#x60;-bugs&#x60; - Bugs (descending) * &#x60;essential&#x60; - Essential * &#x60;-essential&#x60; - Essential (descending) * &#x60;build_essential&#x60; - Build essential * &#x60;-build_essential&#x60; - Build essential (descending) * &#x60;installed_size&#x60; - Installed size * &#x60;-installed_size&#x60; - Installed size (descending) * &#x60;maintainer&#x60; - Maintainer * &#x60;-maintainer&#x60; - Maintainer (descending) * &#x60;original_maintainer&#x60; - Original maintainer * &#x60;-original_maintainer&#x60; - Original maintainer (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;description_md5&#x60; - Description md5 * &#x60;-description_md5&#x60; - Description md5 (descending) * &#x60;homepage&#x60; - Homepage * &#x60;-homepage&#x60; - Homepage (descending) * &#x60;built_using&#x60; - Built using * &#x60;-built_using&#x60; - Built using (descending) * &#x60;auto_built_package&#x60; - Auto built package * &#x60;-auto_built_package&#x60; - Auto built package (descending) * &#x60;multi_arch&#x60; - Multi arch * &#x60;-multi_arch&#x60; - Multi arch (descending) * &#x60;breaks&#x60; - Breaks * &#x60;-breaks&#x60; - Breaks (descending) * &#x60;conflicts&#x60; - Conflicts * &#x60;-conflicts&#x60; - Conflicts (descending) * &#x60;depends&#x60; - Depends * &#x60;-depends&#x60; - Depends (descending) * &#x60;recommends&#x60; - Recommends * &#x60;-recommends&#x60; - Recommends (descending) * &#x60;suggests&#x60; - Suggests * &#x60;-suggests&#x60; - Suggests (descending) * &#x60;enhances&#x60; - Enhances * &#x60;-enhances&#x60; - Enhances (descending) * &#x60;pre_depends&#x60; - Pre depends * &#x60;-pre_depends&#x60; - Pre depends (descending) * &#x60;provides&#x60; - Provides * &#x60;-provides&#x60; - Provides (descending) * &#x60;replaces&#x60; - Replaces * &#x60;-replaces&#x60; - Replaces (descending) * &#x60;relative_path&#x60; - Relative path * &#x60;-relative_path&#x60; - Relative path (descending) * &#x60;sha256&#x60; - Sha256 * &#x60;-sha256&#x60; - Sha256 (descending) * &#x60;custom_fields&#x60; - Custom fields * &#x60;-custom_fields&#x60; - Custom fields (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Ordering(ordering []string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where origin matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Origin(origin string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.origin = &origin
	return r
}

// Filter results where original_maintainer matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) OriginalMaintainer(originalMaintainer string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.originalMaintainer = &originalMaintainer
	return r
}

// Filter results where package matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Package_(package_ string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.package_ = &package_
	return r
}

// Filter results where priority matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Priority(priority string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.priority = &priority
	return r
}

// Multiple values may be separated by commas.
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) PulpHrefIn(pulpHrefIn []string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) PulpIdIn(pulpIdIn []string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) RepositoryVersion(repositoryVersion string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where section matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Section(section string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.section = &section
	return r
}

// Filter results where sha256 matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Sha256(sha256 string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.sha256 = &sha256
	return r
}

// Filter results where source matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Source(source string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.source = &source
	return r
}

// Filter results where tag matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Tag(tag string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.tag = &tag
	return r
}

// Filter results where version matches value
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Version(version string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.version = &version
	return r
}

// A list of fields to include in the response.
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Fields(fields []string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) ExcludeFields(excludeFields []string) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) Execute() (*PaginateddebInstallerPackageResponseList, *http.Response, error) {
	return r.ApiService.ContentDebInstallerPackagesListExecute(r)
}

/*
ContentDebInstallerPackagesList List installer packages

An InstallerPackage represents a '.udeb' installer package.

Associated artifacts: Exactly one '.udeb' installer package file.

Note that installer packages are currently used exclusively for verbatim publications. The APT
publisher (both simple and structured mode) will not include these packages.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest
*/
func (a *ContentInstallerPackagesAPIService) ContentDebInstallerPackagesList(ctx context.Context) ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest {
	return ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebInstallerPackageResponseList
func (a *ContentInstallerPackagesAPIService) ContentDebInstallerPackagesListExecute(r ContentInstallerPackagesAPIContentDebInstallerPackagesListRequest) (*PaginateddebInstallerPackageResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebInstallerPackageResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentInstallerPackagesAPIService.ContentDebInstallerPackagesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/installer_packages/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.architecture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "architecture", r.architecture, "")
	}
	if r.autoBuiltPackage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auto_built_package", r.autoBuiltPackage, "")
	}
	if r.buildEssential != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "build_essential", r.buildEssential, "")
	}
	if r.builtUsing != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "built_using", r.builtUsing, "")
	}
	if r.essential != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "essential", r.essential, "")
	}
	if r.installedSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "installed_size", r.installedSize, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.maintainer != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maintainer", r.maintainer, "")
	}
	if r.multiArch != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "multi_arch", r.multiArch, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.origin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "origin", r.origin, "")
	}
	if r.originalMaintainer != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "original_maintainer", r.originalMaintainer, "")
	}
	if r.package_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "package", r.package_, "")
	}
	if r.priority != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "priority", r.priority, "")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.section != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "section", r.section, "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "")
	}
	if r.source != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "source", r.source, "")
	}
	if r.tag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag", r.tag, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
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

type ContentInstallerPackagesAPIContentDebInstallerPackagesReadRequest struct {
	ctx context.Context
	ApiService *ContentInstallerPackagesAPIService
	debInstallerPackageHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesReadRequest) Fields(fields []string) ContentInstallerPackagesAPIContentDebInstallerPackagesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentInstallerPackagesAPIContentDebInstallerPackagesReadRequest) ExcludeFields(excludeFields []string) ContentInstallerPackagesAPIContentDebInstallerPackagesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentInstallerPackagesAPIContentDebInstallerPackagesReadRequest) Execute() (*DebInstallerPackageResponse, *http.Response, error) {
	return r.ApiService.ContentDebInstallerPackagesReadExecute(r)
}

/*
ContentDebInstallerPackagesRead Inspect an installer package

An InstallerPackage represents a '.udeb' installer package.

Associated artifacts: Exactly one '.udeb' installer package file.

Note that installer packages are currently used exclusively for verbatim publications. The APT
publisher (both simple and structured mode) will not include these packages.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debInstallerPackageHref
 @return ContentInstallerPackagesAPIContentDebInstallerPackagesReadRequest
*/
func (a *ContentInstallerPackagesAPIService) ContentDebInstallerPackagesRead(ctx context.Context, debInstallerPackageHref string) ContentInstallerPackagesAPIContentDebInstallerPackagesReadRequest {
	return ContentInstallerPackagesAPIContentDebInstallerPackagesReadRequest{
		ApiService: a,
		ctx: ctx,
		debInstallerPackageHref: debInstallerPackageHref,
	}
}

// Execute executes the request
//  @return DebInstallerPackageResponse
func (a *ContentInstallerPackagesAPIService) ContentDebInstallerPackagesReadExecute(r ContentInstallerPackagesAPIContentDebInstallerPackagesReadRequest) (*DebInstallerPackageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebInstallerPackageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentInstallerPackagesAPIService.ContentDebInstallerPackagesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_installer_package_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_installer_package_href"+"}", parameterValueToString(r.debInstallerPackageHref, "debInstallerPackageHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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