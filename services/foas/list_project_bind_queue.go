package foas

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

// ListProjectBindQueue invokes the foas.ListProjectBindQueue API synchronously
func (client *Client) ListProjectBindQueue(request *ListProjectBindQueueRequest) (response *ListProjectBindQueueResponse, err error) {
	response = CreateListProjectBindQueueResponse()
	err = client.DoAction(request, response)
	return
}

// ListProjectBindQueueWithChan invokes the foas.ListProjectBindQueue API asynchronously
func (client *Client) ListProjectBindQueueWithChan(request *ListProjectBindQueueRequest) (<-chan *ListProjectBindQueueResponse, <-chan error) {
	responseChan := make(chan *ListProjectBindQueueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProjectBindQueue(request)
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

// ListProjectBindQueueWithCallback invokes the foas.ListProjectBindQueue API asynchronously
func (client *Client) ListProjectBindQueueWithCallback(request *ListProjectBindQueueRequest, callback func(response *ListProjectBindQueueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProjectBindQueueResponse
		var err error
		defer close(result)
		response, err = client.ListProjectBindQueue(request)
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

// ListProjectBindQueueRequest is the request struct for api ListProjectBindQueue
type ListProjectBindQueueRequest struct {
	*requests.RoaRequest
	QueueName   string `position:"Query" name:"queueName"`
	ProjectName string `position:"Path" name:"projectName"`
	ClusterId   string `position:"Query" name:"clusterId"`
}

// ListProjectBindQueueResponse is the response struct for api ListProjectBindQueue
type ListProjectBindQueueResponse struct {
	*responses.BaseResponse
	RequestId string                       `json:"RequestId" xml:"RequestId"`
	Queues    QueuesInListProjectBindQueue `json:"Queues" xml:"Queues"`
}

// CreateListProjectBindQueueRequest creates a request to invoke ListProjectBindQueue API
func CreateListProjectBindQueueRequest() (request *ListProjectBindQueueRequest) {
	request = &ListProjectBindQueueRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "ListProjectBindQueue", "/api/v2/projects/[projectName]/queues", "foas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListProjectBindQueueResponse creates a response to parse from ListProjectBindQueue response
func CreateListProjectBindQueueResponse() (response *ListProjectBindQueueResponse) {
	response = &ListProjectBindQueueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
