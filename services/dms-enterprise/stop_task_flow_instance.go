package dms_enterprise

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

// StopTaskFlowInstance invokes the dms_enterprise.StopTaskFlowInstance API synchronously
func (client *Client) StopTaskFlowInstance(request *StopTaskFlowInstanceRequest) (response *StopTaskFlowInstanceResponse, err error) {
	response = CreateStopTaskFlowInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// StopTaskFlowInstanceWithChan invokes the dms_enterprise.StopTaskFlowInstance API asynchronously
func (client *Client) StopTaskFlowInstanceWithChan(request *StopTaskFlowInstanceRequest) (<-chan *StopTaskFlowInstanceResponse, <-chan error) {
	responseChan := make(chan *StopTaskFlowInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopTaskFlowInstance(request)
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

// StopTaskFlowInstanceWithCallback invokes the dms_enterprise.StopTaskFlowInstance API asynchronously
func (client *Client) StopTaskFlowInstanceWithCallback(request *StopTaskFlowInstanceRequest, callback func(response *StopTaskFlowInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopTaskFlowInstanceResponse
		var err error
		defer close(result)
		response, err = client.StopTaskFlowInstance(request)
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

// StopTaskFlowInstanceRequest is the request struct for api StopTaskFlowInstance
type StopTaskFlowInstanceRequest struct {
	*requests.RpcRequest
	DagId         requests.Integer `position:"Query" name:"DagId"`
	Tid           requests.Integer `position:"Query" name:"Tid"`
	DagInstanceId requests.Integer `position:"Query" name:"DagInstanceId"`
}

// StopTaskFlowInstanceResponse is the response struct for api StopTaskFlowInstance
type StopTaskFlowInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateStopTaskFlowInstanceRequest creates a request to invoke StopTaskFlowInstance API
func CreateStopTaskFlowInstanceRequest() (request *StopTaskFlowInstanceRequest) {
	request = &StopTaskFlowInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "StopTaskFlowInstance", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopTaskFlowInstanceResponse creates a response to parse from StopTaskFlowInstance response
func CreateStopTaskFlowInstanceResponse() (response *StopTaskFlowInstanceResponse) {
	response = &StopTaskFlowInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
