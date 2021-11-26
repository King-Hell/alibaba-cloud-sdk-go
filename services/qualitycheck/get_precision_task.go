package qualitycheck

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

// GetPrecisionTask invokes the qualitycheck.GetPrecisionTask API synchronously
func (client *Client) GetPrecisionTask(request *GetPrecisionTaskRequest) (response *GetPrecisionTaskResponse, err error) {
	response = CreateGetPrecisionTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetPrecisionTaskWithChan invokes the qualitycheck.GetPrecisionTask API asynchronously
func (client *Client) GetPrecisionTaskWithChan(request *GetPrecisionTaskRequest) (<-chan *GetPrecisionTaskResponse, <-chan error) {
	responseChan := make(chan *GetPrecisionTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPrecisionTask(request)
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

// GetPrecisionTaskWithCallback invokes the qualitycheck.GetPrecisionTask API asynchronously
func (client *Client) GetPrecisionTaskWithCallback(request *GetPrecisionTaskRequest, callback func(response *GetPrecisionTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPrecisionTaskResponse
		var err error
		defer close(result)
		response, err = client.GetPrecisionTask(request)
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

// GetPrecisionTaskRequest is the request struct for api GetPrecisionTask
type GetPrecisionTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetPrecisionTaskResponse is the response struct for api GetPrecisionTask
type GetPrecisionTaskResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetPrecisionTaskRequest creates a request to invoke GetPrecisionTask API
func CreateGetPrecisionTaskRequest() (request *GetPrecisionTaskRequest) {
	request = &GetPrecisionTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetPrecisionTask", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPrecisionTaskResponse creates a response to parse from GetPrecisionTask response
func CreateGetPrecisionTaskResponse() (response *GetPrecisionTaskResponse) {
	response = &GetPrecisionTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
