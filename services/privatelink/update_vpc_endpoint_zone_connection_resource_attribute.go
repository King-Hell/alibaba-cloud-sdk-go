package privatelink

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateVpcEndpointZoneConnectionResourceAttribute invokes the privatelink.UpdateVpcEndpointZoneConnectionResourceAttribute API synchronously
func (client *Client) UpdateVpcEndpointZoneConnectionResourceAttribute(request *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) (response *UpdateVpcEndpointZoneConnectionResourceAttributeResponse, err error) {
	response = CreateUpdateVpcEndpointZoneConnectionResourceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateVpcEndpointZoneConnectionResourceAttributeWithChan invokes the privatelink.UpdateVpcEndpointZoneConnectionResourceAttribute API asynchronously
func (client *Client) UpdateVpcEndpointZoneConnectionResourceAttributeWithChan(request *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) (<-chan *UpdateVpcEndpointZoneConnectionResourceAttributeResponse, <-chan error) {
	responseChan := make(chan *UpdateVpcEndpointZoneConnectionResourceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateVpcEndpointZoneConnectionResourceAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateVpcEndpointZoneConnectionResourceAttributeWithCallback invokes the privatelink.UpdateVpcEndpointZoneConnectionResourceAttribute API asynchronously
func (client *Client) UpdateVpcEndpointZoneConnectionResourceAttributeWithCallback(request *UpdateVpcEndpointZoneConnectionResourceAttributeRequest, callback func(response *UpdateVpcEndpointZoneConnectionResourceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateVpcEndpointZoneConnectionResourceAttributeResponse
		var err error
		defer close(result)
		response, err = client.UpdateVpcEndpointZoneConnectionResourceAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateVpcEndpointZoneConnectionResourceAttributeRequest is the request struct for api UpdateVpcEndpointZoneConnectionResourceAttribute
type UpdateVpcEndpointZoneConnectionResourceAttributeRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	EndpointId           string           `position:"Query" name:"EndpointId"`
	ResourceId           string           `position:"Query" name:"ResourceId"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceAllocateMode string           `position:"Query" name:"ResourceAllocateMode"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
	ResourceReplaceMode  string           `position:"Query" name:"ResourceReplaceMode"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	ServiceId            string           `position:"Query" name:"ServiceId"`
}

// UpdateVpcEndpointZoneConnectionResourceAttributeResponse is the response struct for api UpdateVpcEndpointZoneConnectionResourceAttribute
type UpdateVpcEndpointZoneConnectionResourceAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateVpcEndpointZoneConnectionResourceAttributeRequest creates a request to invoke UpdateVpcEndpointZoneConnectionResourceAttribute API
func CreateUpdateVpcEndpointZoneConnectionResourceAttributeRequest() (request *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) {
	request = &UpdateVpcEndpointZoneConnectionResourceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Privatelink", "2020-04-15", "UpdateVpcEndpointZoneConnectionResourceAttribute", "privatelink", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateVpcEndpointZoneConnectionResourceAttributeResponse creates a response to parse from UpdateVpcEndpointZoneConnectionResourceAttribute response
func CreateUpdateVpcEndpointZoneConnectionResourceAttributeResponse() (response *UpdateVpcEndpointZoneConnectionResourceAttributeResponse) {
	response = &UpdateVpcEndpointZoneConnectionResourceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
