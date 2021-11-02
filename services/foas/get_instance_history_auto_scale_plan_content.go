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

// GetInstanceHistoryAutoScalePlanContent invokes the foas.GetInstanceHistoryAutoScalePlanContent API synchronously
func (client *Client) GetInstanceHistoryAutoScalePlanContent(request *GetInstanceHistoryAutoScalePlanContentRequest) (response *GetInstanceHistoryAutoScalePlanContentResponse, err error) {
	response = CreateGetInstanceHistoryAutoScalePlanContentResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceHistoryAutoScalePlanContentWithChan invokes the foas.GetInstanceHistoryAutoScalePlanContent API asynchronously
func (client *Client) GetInstanceHistoryAutoScalePlanContentWithChan(request *GetInstanceHistoryAutoScalePlanContentRequest) (<-chan *GetInstanceHistoryAutoScalePlanContentResponse, <-chan error) {
	responseChan := make(chan *GetInstanceHistoryAutoScalePlanContentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceHistoryAutoScalePlanContent(request)
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

// GetInstanceHistoryAutoScalePlanContentWithCallback invokes the foas.GetInstanceHistoryAutoScalePlanContent API asynchronously
func (client *Client) GetInstanceHistoryAutoScalePlanContentWithCallback(request *GetInstanceHistoryAutoScalePlanContentRequest, callback func(response *GetInstanceHistoryAutoScalePlanContentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceHistoryAutoScalePlanContentResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceHistoryAutoScalePlanContent(request)
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

// GetInstanceHistoryAutoScalePlanContentRequest is the request struct for api GetInstanceHistoryAutoScalePlanContent
type GetInstanceHistoryAutoScalePlanContentRequest struct {
	*requests.RoaRequest
	ProjectName string           `position:"Path" name:"projectName"`
	InstanceId  requests.Integer `position:"Path" name:"instanceId"`
	PlanName    string           `position:"Query" name:"planName"`
	JobName     string           `position:"Path" name:"jobName"`
}

// GetInstanceHistoryAutoScalePlanContentResponse is the response struct for api GetInstanceHistoryAutoScalePlanContent
type GetInstanceHistoryAutoScalePlanContentResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	PlanContent string `json:"PlanContent" xml:"PlanContent"`
}

// CreateGetInstanceHistoryAutoScalePlanContentRequest creates a request to invoke GetInstanceHistoryAutoScalePlanContent API
func CreateGetInstanceHistoryAutoScalePlanContentRequest() (request *GetInstanceHistoryAutoScalePlanContentRequest) {
	request = &GetInstanceHistoryAutoScalePlanContentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "GetInstanceHistoryAutoScalePlanContent", "/api/v2/projects/[projectName]/jobs/[jobName]/instance/[instanceId]/autoscale/plancontent", "foas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetInstanceHistoryAutoScalePlanContentResponse creates a response to parse from GetInstanceHistoryAutoScalePlanContent response
func CreateGetInstanceHistoryAutoScalePlanContentResponse() (response *GetInstanceHistoryAutoScalePlanContentResponse) {
	response = &GetInstanceHistoryAutoScalePlanContentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
