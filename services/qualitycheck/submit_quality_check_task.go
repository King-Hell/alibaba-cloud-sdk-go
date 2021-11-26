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

// SubmitQualityCheckTask invokes the qualitycheck.SubmitQualityCheckTask API synchronously
func (client *Client) SubmitQualityCheckTask(request *SubmitQualityCheckTaskRequest) (response *SubmitQualityCheckTaskResponse, err error) {
	response = CreateSubmitQualityCheckTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitQualityCheckTaskWithChan invokes the qualitycheck.SubmitQualityCheckTask API asynchronously
func (client *Client) SubmitQualityCheckTaskWithChan(request *SubmitQualityCheckTaskRequest) (<-chan *SubmitQualityCheckTaskResponse, <-chan error) {
	responseChan := make(chan *SubmitQualityCheckTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitQualityCheckTask(request)
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

// SubmitQualityCheckTaskWithCallback invokes the qualitycheck.SubmitQualityCheckTask API asynchronously
func (client *Client) SubmitQualityCheckTaskWithCallback(request *SubmitQualityCheckTaskRequest, callback func(response *SubmitQualityCheckTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitQualityCheckTaskResponse
		var err error
		defer close(result)
		response, err = client.SubmitQualityCheckTask(request)
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

// SubmitQualityCheckTaskRequest is the request struct for api SubmitQualityCheckTask
type SubmitQualityCheckTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// SubmitQualityCheckTaskResponse is the response struct for api SubmitQualityCheckTask
type SubmitQualityCheckTaskResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSubmitQualityCheckTaskRequest creates a request to invoke SubmitQualityCheckTask API
func CreateSubmitQualityCheckTaskRequest() (request *SubmitQualityCheckTaskRequest) {
	request = &SubmitQualityCheckTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "SubmitQualityCheckTask", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitQualityCheckTaskResponse creates a response to parse from SubmitQualityCheckTask response
func CreateSubmitQualityCheckTaskResponse() (response *SubmitQualityCheckTaskResponse) {
	response = &SubmitQualityCheckTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
