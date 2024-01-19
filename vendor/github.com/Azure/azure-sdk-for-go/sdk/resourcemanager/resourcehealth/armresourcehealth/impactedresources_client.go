//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ImpactedResourcesClient contains the methods for the ImpactedResources group.
// Don't use this type directly, use NewImpactedResourcesClient() instead.
type ImpactedResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewImpactedResourcesClient creates a new instance of ImpactedResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewImpactedResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ImpactedResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ImpactedResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the specific impacted resource in the subscription by an event.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - eventTrackingID - Event Id which uniquely identifies ServiceHealth event.
//   - impactedResourceName - Name of the Impacted Resource.
//   - options - ImpactedResourcesClientGetOptions contains the optional parameters for the ImpactedResourcesClient.Get method.
func (client *ImpactedResourcesClient) Get(ctx context.Context, eventTrackingID string, impactedResourceName string, options *ImpactedResourcesClientGetOptions) (ImpactedResourcesClientGetResponse, error) {
	var err error
	const operationName = "ImpactedResourcesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, eventTrackingID, impactedResourceName, options)
	if err != nil {
		return ImpactedResourcesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImpactedResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImpactedResourcesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ImpactedResourcesClient) getCreateRequest(ctx context.Context, eventTrackingID string, impactedResourceName string, options *ImpactedResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/events/{eventTrackingId}/impactedResources/{impactedResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if eventTrackingID == "" {
		return nil, errors.New("parameter eventTrackingID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventTrackingId}", url.PathEscape(eventTrackingID))
	if impactedResourceName == "" {
		return nil, errors.New("parameter impactedResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{impactedResourceName}", url.PathEscape(impactedResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ImpactedResourcesClient) getHandleResponse(resp *http.Response) (ImpactedResourcesClientGetResponse, error) {
	result := ImpactedResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventImpactedResource); err != nil {
		return ImpactedResourcesClientGetResponse{}, err
	}
	return result, nil
}

// GetByTenantID - Gets the specific impacted resource in the tenant by an event.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - eventTrackingID - Event Id which uniquely identifies ServiceHealth event.
//   - impactedResourceName - Name of the Impacted Resource.
//   - options - ImpactedResourcesClientGetByTenantIDOptions contains the optional parameters for the ImpactedResourcesClient.GetByTenantID
//     method.
func (client *ImpactedResourcesClient) GetByTenantID(ctx context.Context, eventTrackingID string, impactedResourceName string, options *ImpactedResourcesClientGetByTenantIDOptions) (ImpactedResourcesClientGetByTenantIDResponse, error) {
	var err error
	const operationName = "ImpactedResourcesClient.GetByTenantID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByTenantIDCreateRequest(ctx, eventTrackingID, impactedResourceName, options)
	if err != nil {
		return ImpactedResourcesClientGetByTenantIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImpactedResourcesClientGetByTenantIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImpactedResourcesClientGetByTenantIDResponse{}, err
	}
	resp, err := client.getByTenantIDHandleResponse(httpResp)
	return resp, err
}

// getByTenantIDCreateRequest creates the GetByTenantID request.
func (client *ImpactedResourcesClient) getByTenantIDCreateRequest(ctx context.Context, eventTrackingID string, impactedResourceName string, options *ImpactedResourcesClientGetByTenantIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceHealth/events/{eventTrackingId}/impactedResources/{impactedResourceName}"
	if eventTrackingID == "" {
		return nil, errors.New("parameter eventTrackingID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventTrackingId}", url.PathEscape(eventTrackingID))
	if impactedResourceName == "" {
		return nil, errors.New("parameter impactedResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{impactedResourceName}", url.PathEscape(impactedResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByTenantIDHandleResponse handles the GetByTenantID response.
func (client *ImpactedResourcesClient) getByTenantIDHandleResponse(resp *http.Response) (ImpactedResourcesClientGetByTenantIDResponse, error) {
	result := ImpactedResourcesClientGetByTenantIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventImpactedResource); err != nil {
		return ImpactedResourcesClientGetByTenantIDResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionIDAndEventIDPager - Lists impacted resources in the subscription by an event.
//
// Generated from API version 2022-10-01
//   - eventTrackingID - Event Id which uniquely identifies ServiceHealth event.
//   - options - ImpactedResourcesClientListBySubscriptionIDAndEventIDOptions contains the optional parameters for the ImpactedResourcesClient.NewListBySubscriptionIDAndEventIDPager
//     method.
func (client *ImpactedResourcesClient) NewListBySubscriptionIDAndEventIDPager(eventTrackingID string, options *ImpactedResourcesClientListBySubscriptionIDAndEventIDOptions) *runtime.Pager[ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse]{
		More: func(page ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse) (ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ImpactedResourcesClient.NewListBySubscriptionIDAndEventIDPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionIDAndEventIDCreateRequest(ctx, eventTrackingID, options)
			}, nil)
			if err != nil {
				return ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse{}, err
			}
			return client.listBySubscriptionIDAndEventIDHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionIDAndEventIDCreateRequest creates the ListBySubscriptionIDAndEventID request.
func (client *ImpactedResourcesClient) listBySubscriptionIDAndEventIDCreateRequest(ctx context.Context, eventTrackingID string, options *ImpactedResourcesClientListBySubscriptionIDAndEventIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/events/{eventTrackingId}/impactedResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if eventTrackingID == "" {
		return nil, errors.New("parameter eventTrackingID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventTrackingId}", url.PathEscape(eventTrackingID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionIDAndEventIDHandleResponse handles the ListBySubscriptionIDAndEventID response.
func (client *ImpactedResourcesClient) listBySubscriptionIDAndEventIDHandleResponse(resp *http.Response) (ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse, error) {
	result := ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventImpactedResourceListResult); err != nil {
		return ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse{}, err
	}
	return result, nil
}

// NewListByTenantIDAndEventIDPager - Lists impacted resources in the tenant by an event.
//
// Generated from API version 2022-10-01
//   - eventTrackingID - Event Id which uniquely identifies ServiceHealth event.
//   - options - ImpactedResourcesClientListByTenantIDAndEventIDOptions contains the optional parameters for the ImpactedResourcesClient.NewListByTenantIDAndEventIDPager
//     method.
func (client *ImpactedResourcesClient) NewListByTenantIDAndEventIDPager(eventTrackingID string, options *ImpactedResourcesClientListByTenantIDAndEventIDOptions) *runtime.Pager[ImpactedResourcesClientListByTenantIDAndEventIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImpactedResourcesClientListByTenantIDAndEventIDResponse]{
		More: func(page ImpactedResourcesClientListByTenantIDAndEventIDResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImpactedResourcesClientListByTenantIDAndEventIDResponse) (ImpactedResourcesClientListByTenantIDAndEventIDResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ImpactedResourcesClient.NewListByTenantIDAndEventIDPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByTenantIDAndEventIDCreateRequest(ctx, eventTrackingID, options)
			}, nil)
			if err != nil {
				return ImpactedResourcesClientListByTenantIDAndEventIDResponse{}, err
			}
			return client.listByTenantIDAndEventIDHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByTenantIDAndEventIDCreateRequest creates the ListByTenantIDAndEventID request.
func (client *ImpactedResourcesClient) listByTenantIDAndEventIDCreateRequest(ctx context.Context, eventTrackingID string, options *ImpactedResourcesClientListByTenantIDAndEventIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceHealth/events/{eventTrackingId}/impactedResources"
	if eventTrackingID == "" {
		return nil, errors.New("parameter eventTrackingID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventTrackingId}", url.PathEscape(eventTrackingID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTenantIDAndEventIDHandleResponse handles the ListByTenantIDAndEventID response.
func (client *ImpactedResourcesClient) listByTenantIDAndEventIDHandleResponse(resp *http.Response) (ImpactedResourcesClientListByTenantIDAndEventIDResponse, error) {
	result := ImpactedResourcesClientListByTenantIDAndEventIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventImpactedResourceListResult); err != nil {
		return ImpactedResourcesClientListByTenantIDAndEventIDResponse{}, err
	}
	return result, nil
}
