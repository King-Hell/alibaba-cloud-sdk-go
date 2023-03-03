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

// GetManagedPrometheusStatus invokes the arms.GetManagedPrometheusStatus API synchronously
func (client *Client) GetManagedPrometheusStatus(request *GetManagedPrometheusStatusRequest) (response *GetManagedPrometheusStatusResponse, err error) {
	response = CreateGetManagedPrometheusStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetManagedPrometheusStatusWithChan invokes the arms.GetManagedPrometheusStatus API asynchronously
func (client *Client) GetManagedPrometheusStatusWithChan(request *GetManagedPrometheusStatusRequest) (<-chan *GetManagedPrometheusStatusResponse, <-chan error) {
	responseChan := make(chan *GetManagedPrometheusStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetManagedPrometheusStatus(request)
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

// GetManagedPrometheusStatusWithCallback invokes the arms.GetManagedPrometheusStatus API asynchronously
func (client *Client) GetManagedPrometheusStatusWithCallback(request *GetManagedPrometheusStatusRequest, callback func(response *GetManagedPrometheusStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetManagedPrometheusStatusResponse
		var err error
		defer close(result)
		response, err = client.GetManagedPrometheusStatus(request)
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

// GetManagedPrometheusStatusRequest is the request struct for api GetManagedPrometheusStatus
type GetManagedPrometheusStatusRequest struct {
	*requests.RpcRequest
	ClusterId       string `position:"Query" name:"ClusterId"`
	ClusterType     string `position:"Query" name:"ClusterType"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	VpcId           string `position:"Query" name:"VpcId"`
}

// GetManagedPrometheusStatusResponse is the response struct for api GetManagedPrometheusStatus
type GetManagedPrometheusStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Code      int    `json:"Code" xml:"Code"`
}

// CreateGetManagedPrometheusStatusRequest creates a request to invoke GetManagedPrometheusStatus API
func CreateGetManagedPrometheusStatusRequest() (request *GetManagedPrometheusStatusRequest) {
	request = &GetManagedPrometheusStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetManagedPrometheusStatus", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetManagedPrometheusStatusResponse creates a response to parse from GetManagedPrometheusStatus response
func CreateGetManagedPrometheusStatusResponse() (response *GetManagedPrometheusStatusResponse) {
	response = &GetManagedPrometheusStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
