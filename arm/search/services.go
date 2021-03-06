package search

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// Services Client
type ServicesClient struct {
	SearchManagementClient
}

func NewServicesClient(subscriptionId string) ServicesClient {
	return NewServicesClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewServicesClientWithBaseUri(baseUri string, subscriptionId string) ServicesClient {
	return ServicesClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// CreateOrUpdate creates or updates a Search service in the given resource
// group. If the Search service already exists, all properties will be
// updated with the given values.
//
// resourceGroupName is the name of the resource group within the current
// subscription. serviceName is the name of the Search service to create or
// update. parameters is the properties to set or update on the Search
// service.
func (client ServicesClient) CreateOrUpdate(resourceGroupName string, serviceName string, parameters SearchServiceCreateOrUpdateParameters) (result SearchServiceResource, ae autorest.Error) {
	req, err := client.NewCreateOrUpdateRequest(resourceGroupName, serviceName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.ServicesClient", "CreateOrUpdate", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.ServicesClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "search.ServicesClient", "CreateOrUpdate", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "search.ServicesClient", "CreateOrUpdate", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the CreateOrUpdate request.
func (client ServicesClient) NewCreateOrUpdateRequest(resourceGroupName string, serviceName string, parameters SearchServiceCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"serviceName":       url.QueryEscape(serviceName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.CreateOrUpdateRequestPreparer(),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the CreateOrUpdate request.
func (client ServicesClient) CreateOrUpdateRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{serviceName}"))
}

// Delete deletes a Search service in the given resource group, along with its
// associated resources.
//
// resourceGroupName is the name of the resource group within the current
// subscription. serviceName is the name of the Search service to delete.
func (client ServicesClient) Delete(resourceGroupName string, serviceName string) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewDeleteRequest(resourceGroupName, serviceName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.ServicesClient", "Delete", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.ServicesClient", "Delete", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound, http.StatusNoContent))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "search.ServicesClient", "Delete", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "search.ServicesClient", "Delete", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the Delete request.
func (client ServicesClient) NewDeleteRequest(resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"serviceName":       url.QueryEscape(serviceName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.DeleteRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Delete request.
func (client ServicesClient) DeleteRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{serviceName}"))
}

// List returns a list of all Search services in the given resource group.
//
// resourceGroupName is the name of the resource group within the current
// subscription.
func (client ServicesClient) List(resourceGroupName string) (result SearchServiceListResult, ae autorest.Error) {
	req, err := client.NewListRequest(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.ServicesClient", "List", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.ServicesClient", "List", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "search.ServicesClient", "List", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "search.ServicesClient", "List", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the List request.
func (client ServicesClient) NewListRequest(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the List request.
func (client ServicesClient) ListRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices"))
}
