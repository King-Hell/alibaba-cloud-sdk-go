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

// ResetConsumerGroupPosition invokes the iot.ResetConsumerGroupPosition API synchronously
func (client *Client) ResetConsumerGroupPosition(request *ResetConsumerGroupPositionRequest) (response *ResetConsumerGroupPositionResponse, err error) {
	response = CreateResetConsumerGroupPositionResponse()
	err = client.DoAction(request, response)
	return
}

// ResetConsumerGroupPositionWithChan invokes the iot.ResetConsumerGroupPosition API asynchronously
func (client *Client) ResetConsumerGroupPositionWithChan(request *ResetConsumerGroupPositionRequest) (<-chan *ResetConsumerGroupPositionResponse, <-chan error) {
	responseChan := make(chan *ResetConsumerGroupPositionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetConsumerGroupPosition(request)
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

// ResetConsumerGroupPositionWithCallback invokes the iot.ResetConsumerGroupPosition API asynchronously
func (client *Client) ResetConsumerGroupPositionWithCallback(request *ResetConsumerGroupPositionRequest, callback func(response *ResetConsumerGroupPositionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetConsumerGroupPositionResponse
		var err error
		defer close(result)
		response, err = client.ResetConsumerGroupPosition(request)
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

// ResetConsumerGroupPositionRequest is the request struct for api ResetConsumerGroupPosition
type ResetConsumerGroupPositionRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	GroupId       string `position:"Query" name:"GroupId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// ResetConsumerGroupPositionResponse is the response struct for api ResetConsumerGroupPosition
type ResetConsumerGroupPositionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string `json:"Code" xml:"Code"`
}

// CreateResetConsumerGroupPositionRequest creates a request to invoke ResetConsumerGroupPosition API
func CreateResetConsumerGroupPositionRequest() (request *ResetConsumerGroupPositionRequest) {
	request = &ResetConsumerGroupPositionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ResetConsumerGroupPosition", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResetConsumerGroupPositionResponse creates a response to parse from ResetConsumerGroupPosition response
func CreateResetConsumerGroupPositionResponse() (response *ResetConsumerGroupPositionResponse) {
	response = &ResetConsumerGroupPositionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
