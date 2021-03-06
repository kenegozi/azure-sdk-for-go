package authorization

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

// ManagementLocks Client
type ManagementLocksClient struct {
	AuthorizationClient
}

func NewManagementLocksClient(subscriptionId string) ManagementLocksClient {
	return NewManagementLocksClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewManagementLocksClientWithBaseUri(baseUri string, subscriptionId string) ManagementLocksClient {
	return ManagementLocksClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// CreateOrUpdateAtResourceGroupLevel create or update a management lock at
// the resource group level.
//
// resourceGroupName is the resource group name. lockName is the lock name.
// parameters is the management lock parameters.
func (client ManagementLocksClient) CreateOrUpdateAtResourceGroupLevel(resourceGroupName string, lockName string, parameters ManagementLock) (result ManagementLock, ae autorest.Error) {
	req, err := client.NewCreateOrUpdateAtResourceGroupLevelRequest(resourceGroupName, lockName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtResourceGroupLevel", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtResourceGroupLevel", "Failure preparing request")
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
			ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtResourceGroupLevel", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtResourceGroupLevel", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the CreateOrUpdateAtResourceGroupLevel request.
func (client ManagementLocksClient) NewCreateOrUpdateAtResourceGroupLevelRequest(resourceGroupName string, lockName string, parameters ManagementLock) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":          url.QueryEscape(lockName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.CreateOrUpdateAtResourceGroupLevelRequestPreparer(),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the CreateOrUpdateAtResourceGroupLevel request.
func (client ManagementLocksClient) CreateOrUpdateAtResourceGroupLevelRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/locks/{lockName}"))
}

// CreateOrUpdateAtResourceLevel create or update a management lock at the
// resource level or any level below resource.
//
// resourceGroupName is the name of the resource group.
// resourceProviderNamespace is resource identity. parentResourcePath is
// resource identity. resourceType is resource identity. resourceName is
// resource identity. lockName is the name of lock. parameters is create or
// update management lock parameters.
func (client ManagementLocksClient) CreateOrUpdateAtResourceLevel(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, lockName string, parameters ManagementLock) (result ManagementLock, ae autorest.Error) {
	req, err := client.NewCreateOrUpdateAtResourceLevelRequest(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, lockName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtResourceLevel", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtResourceLevel", "Failure preparing request")
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
			ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtResourceLevel", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtResourceLevel", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the CreateOrUpdateAtResourceLevel request.
func (client ManagementLocksClient) NewCreateOrUpdateAtResourceLevelRequest(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, lockName string, parameters ManagementLock) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":                  url.QueryEscape(lockName),
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.CreateOrUpdateAtResourceLevelRequestPreparer(),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the CreateOrUpdateAtResourceLevel request.
func (client ManagementLocksClient) CreateOrUpdateAtResourceLevelRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/locks/{lockName}"))
}

// DeleteAtResourceLevel deletes the management lock of a resource or any
// level below resource.
//
// resourceGroupName is the name of the resource group.
// resourceProviderNamespace is resource identity. parentResourcePath is
// resource identity. resourceType is resource identity. resourceName is
// resource identity. lockName is the name of lock.
func (client ManagementLocksClient) DeleteAtResourceLevel(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, lockName string) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewDeleteAtResourceLevelRequest(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, lockName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtResourceLevel", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtResourceLevel", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK, http.StatusAccepted))
	if err == nil {
		err = client.IsPollingAllowed(resp)
		if err == nil {
			resp, err = client.PollAsNeeded(resp)
		}
	}

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtResourceLevel", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtResourceLevel", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the DeleteAtResourceLevel request.
func (client ManagementLocksClient) NewDeleteAtResourceLevelRequest(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, lockName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":                  url.QueryEscape(lockName),
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.DeleteAtResourceLevelRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the DeleteAtResourceLevel request.
func (client ManagementLocksClient) DeleteAtResourceLevelRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/locks/{lockName}"))
}

// CreateOrUpdateAtSubscriptionLevel create or update a management lock at the
// subscription level.
//
// lockName is the name of lock. parameters is the management lock parameters.
func (client ManagementLocksClient) CreateOrUpdateAtSubscriptionLevel(lockName string, parameters ManagementLock) (result ManagementLock, ae autorest.Error) {
	req, err := client.NewCreateOrUpdateAtSubscriptionLevelRequest(lockName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtSubscriptionLevel", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtSubscriptionLevel", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusCreated, http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtSubscriptionLevel", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "CreateOrUpdateAtSubscriptionLevel", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the CreateOrUpdateAtSubscriptionLevel request.
func (client ManagementLocksClient) NewCreateOrUpdateAtSubscriptionLevelRequest(lockName string, parameters ManagementLock) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":       url.QueryEscape(lockName),
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.CreateOrUpdateAtSubscriptionLevelRequestPreparer(),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the CreateOrUpdateAtSubscriptionLevel request.
func (client ManagementLocksClient) CreateOrUpdateAtSubscriptionLevelRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks/{lockName}"))
}

// DeleteAtSubscriptionLevel deletes the management lock of a subscription.
//
// lockName is the name of lock.
func (client ManagementLocksClient) DeleteAtSubscriptionLevel(lockName string) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewDeleteAtSubscriptionLevelRequest(lockName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtSubscriptionLevel", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtSubscriptionLevel", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK, http.StatusAccepted))
	if err == nil {
		err = client.IsPollingAllowed(resp)
		if err == nil {
			resp, err = client.PollAsNeeded(resp)
		}
	}

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtSubscriptionLevel", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtSubscriptionLevel", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the DeleteAtSubscriptionLevel request.
func (client ManagementLocksClient) NewDeleteAtSubscriptionLevelRequest(lockName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":       url.QueryEscape(lockName),
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.DeleteAtSubscriptionLevelRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the DeleteAtSubscriptionLevel request.
func (client ManagementLocksClient) DeleteAtSubscriptionLevelRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks/{lockName}"))
}

// Get gets the management lock of a scope.
//
// lockName is name of the management lock.
func (client ManagementLocksClient) Get(lockName string) (result ManagementLock, ae autorest.Error) {
	req, err := client.NewGetRequest(lockName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "Get", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "Get", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "Get", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "Get", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Get request.
func (client ManagementLocksClient) NewGetRequest(lockName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":       url.QueryEscape(lockName),
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.GetRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Get request.
func (client ManagementLocksClient) GetRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks/{lockName}"))
}

// DeleteAtResourceGroupLevel deletes the management lock of a resource group.
//
// resourceGroup is the resource group names. lockName is the name of lock.
func (client ManagementLocksClient) DeleteAtResourceGroupLevel(resourceGroup string, lockName string) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewDeleteAtResourceGroupLevelRequest(resourceGroup, lockName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtResourceGroupLevel", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtResourceGroupLevel", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK, http.StatusAccepted))
	if err == nil {
		err = client.IsPollingAllowed(resp)
		if err == nil {
			resp, err = client.PollAsNeeded(resp)
		}
	}

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtResourceGroupLevel", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "DeleteAtResourceGroupLevel", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the DeleteAtResourceGroupLevel request.
func (client ManagementLocksClient) NewDeleteAtResourceGroupLevelRequest(resourceGroup string, lockName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":       url.QueryEscape(lockName),
		"resourceGroup":  url.QueryEscape(resourceGroup),
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.DeleteAtResourceGroupLevelRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the DeleteAtResourceGroupLevel request.
func (client ManagementLocksClient) DeleteAtResourceGroupLevelRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Authorization/locks/{lockName}"))
}

// ListAtResourceGroupLevel gets all the management locks of a resource group.
//
// resourceGroupName is resource group name. filter is the filter to apply on
// the operation.
func (client ManagementLocksClient) ListAtResourceGroupLevel(resourceGroupName string, filter string) (result ManagementLockListResult, ae autorest.Error) {
	req, err := client.NewListAtResourceGroupLevelRequest(resourceGroupName, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtResourceGroupLevel", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtResourceGroupLevel", "Failure preparing request")
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
			ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtResourceGroupLevel", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtResourceGroupLevel", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the ListAtResourceGroupLevel request.
func (client ManagementLocksClient) NewListAtResourceGroupLevelRequest(resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"$filter":     filter,
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListAtResourceGroupLevelRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the ListAtResourceGroupLevel request.
func (client ManagementLocksClient) ListAtResourceGroupLevelRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/locks"))
}

// ListAtResourceLevel gets all the management locks of a resource or any
// level below resource.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. resourceProviderNamespace is resource identity.
// parentResourcePath is resource identity. resourceType is resource
// identity. resourceName is resource identity. filter is the filter to apply
// on the operation.
func (client ManagementLocksClient) ListAtResourceLevel(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result ManagementLockListResult, ae autorest.Error) {
	req, err := client.NewListAtResourceLevelRequest(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtResourceLevel", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtResourceLevel", "Failure preparing request")
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
			ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtResourceLevel", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtResourceLevel", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the ListAtResourceLevel request.
func (client ManagementLocksClient) NewListAtResourceLevelRequest(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"$filter":     filter,
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListAtResourceLevelRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the ListAtResourceLevel request.
func (client ManagementLocksClient) ListAtResourceLevelRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/locks"))
}

// ListNext get a list of management locks at resource level or below.
//
// nextLink is nextLink from the previous successful call to List operation.
func (client ManagementLocksClient) ListNext(nextLink string) (result ManagementLockListResult, ae autorest.Error) {
	req, err := client.NewListNextRequest(nextLink)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListNext", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListNext", "Failure preparing request")
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
			ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListNext", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListNext", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the ListNext request.
func (client ManagementLocksClient) NewListNextRequest(nextLink string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nextLink":       nextLink,
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	return autorest.DecoratePreparer(
		client.ListNextRequestPreparer(),
		autorest.WithPathParameters(pathParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the ListNext request.
func (client ManagementLocksClient) ListNextRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/{nextLink}"))
}

// ListAtSubscriptionLevel gets all the management locks of a subscription.
//
// filter is the filter to apply on the operation.
func (client ManagementLocksClient) ListAtSubscriptionLevel(filter string) (result ManagementLockListResult, ae autorest.Error) {
	req, err := client.NewListAtSubscriptionLevelRequest(filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtSubscriptionLevel", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtSubscriptionLevel", "Failure preparing request")
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
			ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtSubscriptionLevel", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "authorization.ManagementLocksClient", "ListAtSubscriptionLevel", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the ListAtSubscriptionLevel request.
func (client ManagementLocksClient) NewListAtSubscriptionLevelRequest(filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"$filter":     filter,
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListAtSubscriptionLevelRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the ListAtSubscriptionLevel request.
func (client ManagementLocksClient) ListAtSubscriptionLevelRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks"))
}
