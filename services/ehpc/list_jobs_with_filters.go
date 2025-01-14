package ehpc

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

// ListJobsWithFilters invokes the ehpc.ListJobsWithFilters API synchronously
func (client *Client) ListJobsWithFilters(request *ListJobsWithFiltersRequest) (response *ListJobsWithFiltersResponse, err error) {
	response = CreateListJobsWithFiltersResponse()
	err = client.DoAction(request, response)
	return
}

// ListJobsWithFiltersWithChan invokes the ehpc.ListJobsWithFilters API asynchronously
func (client *Client) ListJobsWithFiltersWithChan(request *ListJobsWithFiltersRequest) (<-chan *ListJobsWithFiltersResponse, <-chan error) {
	responseChan := make(chan *ListJobsWithFiltersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJobsWithFilters(request)
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

// ListJobsWithFiltersWithCallback invokes the ehpc.ListJobsWithFilters API asynchronously
func (client *Client) ListJobsWithFiltersWithCallback(request *ListJobsWithFiltersRequest, callback func(response *ListJobsWithFiltersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobsWithFiltersResponse
		var err error
		defer close(result)
		response, err = client.ListJobsWithFilters(request)
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

// ListJobsWithFiltersRequest is the request struct for api ListJobsWithFilters
type ListJobsWithFiltersRequest struct {
	*requests.RpcRequest
	JobStatus       string           `position:"Query" name:"JobStatus"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	ExecuteOrder    string           `position:"Query" name:"ExecuteOrder"`
	JobName         string           `position:"Query" name:"JobName"`
	SubmitOrder     string           `position:"Query" name:"SubmitOrder"`
	PendOrder       string           `position:"Query" name:"PendOrder"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	Users           *[]string        `position:"Query" name:"Users"  type:"Repeated"`
	CreateTimeEnd   string           `position:"Query" name:"CreateTimeEnd"`
	Async           requests.Boolean `position:"Query" name:"Async"`
	Nodes           *[]string        `position:"Query" name:"Nodes"  type:"Repeated"`
	Queues          *[]string        `position:"Query" name:"Queues"  type:"Repeated"`
	CreateTimeStart string           `position:"Query" name:"CreateTimeStart"`
}

// ListJobsWithFiltersResponse is the response struct for api ListJobsWithFilters
type ListJobsWithFiltersResponse struct {
	*responses.BaseResponse
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	PageSize   int64     `json:"PageSize" xml:"PageSize"`
	PageNumber int64     `json:"PageNumber" xml:"PageNumber"`
	Success    bool      `json:"Success" xml:"Success"`
	Jobs       []JobInfo `json:"Jobs" xml:"Jobs"`
}

// CreateListJobsWithFiltersRequest creates a request to invoke ListJobsWithFilters API
func CreateListJobsWithFiltersRequest() (request *ListJobsWithFiltersRequest) {
	request = &ListJobsWithFiltersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListJobsWithFilters", "", "")
	request.Method = requests.GET
	return
}

// CreateListJobsWithFiltersResponse creates a response to parse from ListJobsWithFilters response
func CreateListJobsWithFiltersResponse() (response *ListJobsWithFiltersResponse) {
	response = &ListJobsWithFiltersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
