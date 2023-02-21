package schedulerx2

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

// UpdateAppGroup invokes the schedulerx2.UpdateAppGroup API synchronously
func (client *Client) UpdateAppGroup(request *UpdateAppGroupRequest) (response *UpdateAppGroupResponse, err error) {
	response = CreateUpdateAppGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAppGroupWithChan invokes the schedulerx2.UpdateAppGroup API asynchronously
func (client *Client) UpdateAppGroupWithChan(request *UpdateAppGroupRequest) (<-chan *UpdateAppGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateAppGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAppGroup(request)
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

// UpdateAppGroupWithCallback invokes the schedulerx2.UpdateAppGroup API asynchronously
func (client *Client) UpdateAppGroupWithCallback(request *UpdateAppGroupRequest, callback func(response *UpdateAppGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAppGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateAppGroup(request)
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

// UpdateAppGroupRequest is the request struct for api UpdateAppGroup
type UpdateAppGroupRequest struct {
	*requests.RpcRequest
	Description          string           `position:"Query" name:"Description"`
	AlarmJson            string           `position:"Query" name:"AlarmJson"`
	AppGroupId           requests.Integer `position:"Query" name:"AppGroupId"`
	MaxJobs              requests.Integer `position:"Query" name:"MaxJobs"`
	MetricsThresholdJson string           `position:"Query" name:"MetricsThresholdJson"`
	GroupId              string           `position:"Query" name:"GroupId"`
	Namespace            string           `position:"Query" name:"Namespace"`
	Xattrs               string           `position:"Query" name:"Xattrs"`
	MaxConcurrency       requests.Integer `position:"Query" name:"MaxConcurrency"`
}

// UpdateAppGroupResponse is the response struct for api UpdateAppGroup
type UpdateAppGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateAppGroupRequest creates a request to invoke UpdateAppGroup API
func CreateUpdateAppGroupRequest() (request *UpdateAppGroupRequest) {
	request = &UpdateAppGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "UpdateAppGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateAppGroupResponse creates a response to parse from UpdateAppGroup response
func CreateUpdateAppGroupResponse() (response *UpdateAppGroupResponse) {
	response = &UpdateAppGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
