package arms

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

// AddPrometheusInstance invokes the arms.AddPrometheusInstance API synchronously
func (client *Client) AddPrometheusInstance(request *AddPrometheusInstanceRequest) (response *AddPrometheusInstanceResponse, err error) {
	response = CreateAddPrometheusInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// AddPrometheusInstanceWithChan invokes the arms.AddPrometheusInstance API asynchronously
func (client *Client) AddPrometheusInstanceWithChan(request *AddPrometheusInstanceRequest) (<-chan *AddPrometheusInstanceResponse, <-chan error) {
	responseChan := make(chan *AddPrometheusInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddPrometheusInstance(request)
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

// AddPrometheusInstanceWithCallback invokes the arms.AddPrometheusInstance API asynchronously
func (client *Client) AddPrometheusInstanceWithCallback(request *AddPrometheusInstanceRequest, callback func(response *AddPrometheusInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddPrometheusInstanceResponse
		var err error
		defer close(result)
		response, err = client.AddPrometheusInstance(request)
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

// AddPrometheusInstanceRequest is the request struct for api AddPrometheusInstance
type AddPrometheusInstanceRequest struct {
	*requests.RpcRequest
	Name string `position:"Query" name:"Name"`
	Type string `position:"Query" name:"Type"`
}

// AddPrometheusInstanceResponse is the response struct for api AddPrometheusInstance
type AddPrometheusInstanceResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateAddPrometheusInstanceRequest creates a request to invoke AddPrometheusInstance API
func CreateAddPrometheusInstanceRequest() (request *AddPrometheusInstanceRequest) {
	request = &AddPrometheusInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "AddPrometheusInstance", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddPrometheusInstanceResponse creates a response to parse from AddPrometheusInstance response
func CreateAddPrometheusInstanceResponse() (response *AddPrometheusInstanceResponse) {
	response = &AddPrometheusInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
