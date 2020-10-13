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

// DeleteOTAModule invokes the iot.DeleteOTAModule API synchronously
func (client *Client) DeleteOTAModule(request *DeleteOTAModuleRequest) (response *DeleteOTAModuleResponse, err error) {
	response = CreateDeleteOTAModuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteOTAModuleWithChan invokes the iot.DeleteOTAModule API asynchronously
func (client *Client) DeleteOTAModuleWithChan(request *DeleteOTAModuleRequest) (<-chan *DeleteOTAModuleResponse, <-chan error) {
	responseChan := make(chan *DeleteOTAModuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteOTAModule(request)
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

// DeleteOTAModuleWithCallback invokes the iot.DeleteOTAModule API asynchronously
func (client *Client) DeleteOTAModuleWithCallback(request *DeleteOTAModuleRequest, callback func(response *DeleteOTAModuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteOTAModuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteOTAModule(request)
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

// DeleteOTAModuleRequest is the request struct for api DeleteOTAModule
type DeleteOTAModuleRequest struct {
	*requests.RpcRequest
	AuthConfig      string `position:"Query" name:"AuthConfig"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	IotId           string `position:"Query" name:"IotId"`
	IotInstanceId   string `position:"Query" name:"IotInstanceId"`
	ModuleName      string `position:"Query" name:"ModuleName"`
	ProductKey      string `position:"Query" name:"ProductKey"`
	ApiProduct      string `position:"Body" name:"ApiProduct"`
	ApiRevision     string `position:"Body" name:"ApiRevision"`
	DeviceName      string `position:"Query" name:"DeviceName"`
}

// DeleteOTAModuleResponse is the response struct for api DeleteOTAModule
type DeleteOTAModuleResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteOTAModuleRequest creates a request to invoke DeleteOTAModule API
func CreateDeleteOTAModuleRequest() (request *DeleteOTAModuleRequest) {
	request = &DeleteOTAModuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteOTAModule", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteOTAModuleResponse creates a response to parse from DeleteOTAModule response
func CreateDeleteOTAModuleResponse() (response *DeleteOTAModuleResponse) {
	response = &DeleteOTAModuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
