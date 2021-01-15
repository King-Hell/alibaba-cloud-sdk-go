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

// CreateJob invokes the iot.CreateJob API synchronously
func (client *Client) CreateJob(request *CreateJobRequest) (response *CreateJobResponse, err error) {
	response = CreateCreateJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateJobWithChan invokes the iot.CreateJob API asynchronously
func (client *Client) CreateJobWithChan(request *CreateJobRequest) (<-chan *CreateJobResponse, <-chan error) {
	responseChan := make(chan *CreateJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateJob(request)
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

// CreateJobWithCallback invokes the iot.CreateJob API asynchronously
func (client *Client) CreateJobWithCallback(request *CreateJobRequest, callback func(response *CreateJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateJobResponse
		var err error
		defer close(result)
		response, err = client.CreateJob(request)
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

// CreateJobRequest is the request struct for api CreateJob
type CreateJobRequest struct {
	*requests.RpcRequest
	JobDocument   string `position:"Query" name:"JobDocument"`
	Description   string `position:"Query" name:"Description"`
	Type          string `position:"Query" name:"Type"`
	RolloutConfig string `position:"Query" name:"RolloutConfig"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	JobName       string `position:"Query" name:"JobName"`
	TimeoutConfig string `position:"Query" name:"TimeoutConfig"`
	TargetConfig  string `position:"Query" name:"TargetConfig"`
	JobFile       string `position:"Query" name:"JobFile"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// CreateJobResponse is the response struct for api CreateJob
type CreateJobResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	JobId        string `json:"JobId" xml:"JobId"`
}

// CreateCreateJobRequest creates a request to invoke CreateJob API
func CreateCreateJobRequest() (request *CreateJobRequest) {
	request = &CreateJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateJob", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateJobResponse creates a response to parse from CreateJob response
func CreateCreateJobResponse() (response *CreateJobResponse) {
	response = &CreateJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}