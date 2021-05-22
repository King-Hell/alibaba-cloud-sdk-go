package iot

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

// DeleteEdgeInstanceMessageRouting invokes the iot.DeleteEdgeInstanceMessageRouting API synchronously
func (client *Client) DeleteEdgeInstanceMessageRouting(request *DeleteEdgeInstanceMessageRoutingRequest) (response *DeleteEdgeInstanceMessageRoutingResponse, err error) {
	response = CreateDeleteEdgeInstanceMessageRoutingResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEdgeInstanceMessageRoutingWithChan invokes the iot.DeleteEdgeInstanceMessageRouting API asynchronously
func (client *Client) DeleteEdgeInstanceMessageRoutingWithChan(request *DeleteEdgeInstanceMessageRoutingRequest) (<-chan *DeleteEdgeInstanceMessageRoutingResponse, <-chan error) {
	responseChan := make(chan *DeleteEdgeInstanceMessageRoutingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEdgeInstanceMessageRouting(request)
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

// DeleteEdgeInstanceMessageRoutingWithCallback invokes the iot.DeleteEdgeInstanceMessageRouting API asynchronously
func (client *Client) DeleteEdgeInstanceMessageRoutingWithCallback(request *DeleteEdgeInstanceMessageRoutingRequest, callback func(response *DeleteEdgeInstanceMessageRoutingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEdgeInstanceMessageRoutingResponse
		var err error
		defer close(result)
		response, err = client.DeleteEdgeInstanceMessageRouting(request)
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

// DeleteEdgeInstanceMessageRoutingRequest is the request struct for api DeleteEdgeInstanceMessageRouting
type DeleteEdgeInstanceMessageRoutingRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	RouteId       requests.Integer `position:"Query" name:"RouteId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// DeleteEdgeInstanceMessageRoutingResponse is the response struct for api DeleteEdgeInstanceMessageRouting
type DeleteEdgeInstanceMessageRoutingResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteEdgeInstanceMessageRoutingRequest creates a request to invoke DeleteEdgeInstanceMessageRouting API
func CreateDeleteEdgeInstanceMessageRoutingRequest() (request *DeleteEdgeInstanceMessageRoutingRequest) {
	request = &DeleteEdgeInstanceMessageRoutingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteEdgeInstanceMessageRouting", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEdgeInstanceMessageRoutingResponse creates a response to parse from DeleteEdgeInstanceMessageRouting response
func CreateDeleteEdgeInstanceMessageRoutingResponse() (response *DeleteEdgeInstanceMessageRoutingResponse) {
	response = &DeleteEdgeInstanceMessageRoutingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}