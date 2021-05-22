package scsp

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

// UpdateEntityTagRelation invokes the scsp.UpdateEntityTagRelation API synchronously
func (client *Client) UpdateEntityTagRelation(request *UpdateEntityTagRelationRequest) (response *UpdateEntityTagRelationResponse, err error) {
	response = CreateUpdateEntityTagRelationResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEntityTagRelationWithChan invokes the scsp.UpdateEntityTagRelation API asynchronously
func (client *Client) UpdateEntityTagRelationWithChan(request *UpdateEntityTagRelationRequest) (<-chan *UpdateEntityTagRelationResponse, <-chan error) {
	responseChan := make(chan *UpdateEntityTagRelationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEntityTagRelation(request)
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

// UpdateEntityTagRelationWithCallback invokes the scsp.UpdateEntityTagRelation API asynchronously
func (client *Client) UpdateEntityTagRelationWithCallback(request *UpdateEntityTagRelationRequest, callback func(response *UpdateEntityTagRelationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEntityTagRelationResponse
		var err error
		defer close(result)
		response, err = client.UpdateEntityTagRelation(request)
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

// UpdateEntityTagRelationRequest is the request struct for api UpdateEntityTagRelation
type UpdateEntityTagRelationRequest struct {
	*requests.RpcRequest
	InstanceId     string `position:"Body"`
	EntityTagParam string `position:"Body"`
}

// UpdateEntityTagRelationResponse is the response struct for api UpdateEntityTagRelation
type UpdateEntityTagRelationResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateEntityTagRelationRequest creates a request to invoke UpdateEntityTagRelation API
func CreateUpdateEntityTagRelationRequest() (request *UpdateEntityTagRelationRequest) {
	request = &UpdateEntityTagRelationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "UpdateEntityTagRelation", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateEntityTagRelationResponse creates a response to parse from UpdateEntityTagRelation response
func CreateUpdateEntityTagRelationResponse() (response *UpdateEntityTagRelationResponse) {
	response = &UpdateEntityTagRelationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}