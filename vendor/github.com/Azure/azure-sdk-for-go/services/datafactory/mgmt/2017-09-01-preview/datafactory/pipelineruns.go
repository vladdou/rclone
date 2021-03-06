package datafactory

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// PipelineRunsClient is the the Azure Data Factory V2 management API provides a RESTful set of web services that
// interact with Azure Data Factory V2 services.
type PipelineRunsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// NewPipelineRunsClient creates an instance of the PipelineRunsClient client.
func NewPipelineRunsClient(subscriptionID string) PipelineRunsClient {
	return NewPipelineRunsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// NewPipelineRunsClientWithBaseURI creates an instance of the PipelineRunsClient client.
func NewPipelineRunsClientWithBaseURI(baseURI string, subscriptionID string) PipelineRunsClient {
	return PipelineRunsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// Get get a pipeline run by its run ID.
//
// resourceGroupName is the resource group name. factoryName is the factory name. runID is the pipeline run
// identifier.
func (client PipelineRunsClient) Get(ctx context.Context, resourceGroupName string, factoryName string, runID string) (result PipelineRun, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.PipelineRunsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, factoryName, runID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelineRunsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.PipelineRunsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelineRunsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// GetPreparer prepares the Get request.
func (client PipelineRunsClient) GetPreparer(ctx context.Context, resourceGroupName string, factoryName string, runID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runId":             autorest.Encode("path", runID),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelineruns/{runId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineRunsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PipelineRunsClient) GetResponder(resp *http.Response) (result PipelineRun, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// QueryByFactory query pipeline runs in the factory based on input filter conditions.
//
// resourceGroupName is the resource group name. factoryName is the factory name. filterParameters is parameters to
// filter the pipeline run.
func (client PipelineRunsClient) QueryByFactory(ctx context.Context, resourceGroupName string, factoryName string, filterParameters PipelineRunFilterParameters) (result PipelineRunQueryResponse, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: filterParameters,
			Constraints: []validation.Constraint{{Target: "filterParameters.LastUpdatedAfter", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "filterParameters.LastUpdatedBefore", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.PipelineRunsClient", "QueryByFactory", err.Error())
	}

	req, err := client.QueryByFactoryPreparer(ctx, resourceGroupName, factoryName, filterParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelineRunsClient", "QueryByFactory", nil, "Failure preparing request")
		return
	}

	resp, err := client.QueryByFactorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.PipelineRunsClient", "QueryByFactory", resp, "Failure sending request")
		return
	}

	result, err = client.QueryByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelineRunsClient", "QueryByFactory", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// QueryByFactoryPreparer prepares the QueryByFactory request.
func (client PipelineRunsClient) QueryByFactoryPreparer(ctx context.Context, resourceGroupName string, factoryName string, filterParameters PipelineRunFilterParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelineruns", pathParameters),
		autorest.WithJSON(filterParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// QueryByFactorySender sends the QueryByFactory request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineRunsClient) QueryByFactorySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// QueryByFactoryResponder handles the response to the QueryByFactory request. The method always
// closes the http.Response Body.
func (client PipelineRunsClient) QueryByFactoryResponder(resp *http.Response) (result PipelineRunQueryResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
