package pairecservice

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

// ListScenes invokes the pairecservice.ListScenes API synchronously
func (client *Client) ListScenes(request *ListScenesRequest) (response *ListScenesResponse, err error) {
	response = CreateListScenesResponse()
	err = client.DoAction(request, response)
	return
}

// ListScenesWithChan invokes the pairecservice.ListScenes API asynchronously
func (client *Client) ListScenesWithChan(request *ListScenesRequest) (<-chan *ListScenesResponse, <-chan error) {
	responseChan := make(chan *ListScenesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListScenes(request)
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

// ListScenesWithCallback invokes the pairecservice.ListScenes API asynchronously
func (client *Client) ListScenesWithCallback(request *ListScenesRequest, callback func(response *ListScenesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListScenesResponse
		var err error
		defer close(result)
		response, err = client.ListScenes(request)
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

// ListScenesRequest is the request struct for api ListScenes
type ListScenesRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
}

// ListScenesResponse is the response struct for api ListScenes
type ListScenesResponse struct {
	*responses.BaseResponse
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	TotalCount int64        `json:"TotalCount" xml:"TotalCount"`
	Scenes     []ScenesItem `json:"Scenes" xml:"Scenes"`
}

// CreateListScenesRequest creates a request to invoke ListScenes API
func CreateListScenesRequest() (request *ListScenesRequest) {
	request = &ListScenesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "ListScenes", "/api/v1/scenes", "", "")
	request.Method = requests.GET
	return
}

// CreateListScenesResponse creates a response to parse from ListScenes response
func CreateListScenesResponse() (response *ListScenesResponse) {
	response = &ListScenesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
