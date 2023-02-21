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

// CreateAppGroup invokes the schedulerx2.CreateAppGroup API synchronously
func (client *Client) CreateAppGroup(request *CreateAppGroupRequest) (response *CreateAppGroupResponse, err error) {
	response = CreateCreateAppGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAppGroupWithChan invokes the schedulerx2.CreateAppGroup API asynchronously
func (client *Client) CreateAppGroupWithChan(request *CreateAppGroupRequest) (<-chan *CreateAppGroupResponse, <-chan error) {
	responseChan := make(chan *CreateAppGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAppGroup(request)
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

// CreateAppGroupWithCallback invokes the schedulerx2.CreateAppGroup API asynchronously
func (client *Client) CreateAppGroupWithCallback(request *CreateAppGroupRequest, callback func(response *CreateAppGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAppGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateAppGroup(request)
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

// CreateAppGroupRequest is the request struct for api CreateAppGroup
type CreateAppGroupRequest struct {
	*requests.RpcRequest
	NamespaceName        string           `position:"Query" name:"NamespaceName"`
	NamespaceSource      string           `position:"Query" name:"NamespaceSource"`
	ScheduleBusyWorkers  requests.Boolean `position:"Query" name:"ScheduleBusyWorkers"`
	Description          string           `position:"Query" name:"Description"`
	AppName              string           `position:"Query" name:"AppName"`
	AlarmJson            string           `position:"Query" name:"AlarmJson"`
	MonitorContactsJson  string           `position:"Query" name:"MonitorContactsJson"`
	MaxJobs              requests.Integer `position:"Query" name:"MaxJobs"`
	MetricsThresholdJson string           `position:"Query" name:"MetricsThresholdJson"`
	GroupId              string           `position:"Query" name:"GroupId"`
	AppType              requests.Integer `position:"Query" name:"AppType"`
	MonitorConfigJson    string           `position:"Query" name:"MonitorConfigJson"`
	Namespace            string           `position:"Query" name:"Namespace"`
	Xattrs               string           `position:"Query" name:"Xattrs"`
	EnableLog            requests.Boolean `position:"Query" name:"EnableLog"`
	AppKey               string           `position:"Query" name:"AppKey"`
}

// CreateAppGroupResponse is the response struct for api CreateAppGroup
type CreateAppGroupResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateAppGroupRequest creates a request to invoke CreateAppGroup API
func CreateCreateAppGroupRequest() (request *CreateAppGroupRequest) {
	request = &CreateAppGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "CreateAppGroup", "", "")
	request.Method = requests.GET
	return
}

// CreateCreateAppGroupResponse creates a response to parse from CreateAppGroup response
func CreateCreateAppGroupResponse() (response *CreateAppGroupResponse) {
	response = &CreateAppGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
