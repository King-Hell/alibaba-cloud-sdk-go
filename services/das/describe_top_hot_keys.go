package das

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

// DescribeTopHotKeys invokes the das.DescribeTopHotKeys API synchronously
func (client *Client) DescribeTopHotKeys(request *DescribeTopHotKeysRequest) (response *DescribeTopHotKeysResponse, err error) {
	response = CreateDescribeTopHotKeysResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTopHotKeysWithChan invokes the das.DescribeTopHotKeys API asynchronously
func (client *Client) DescribeTopHotKeysWithChan(request *DescribeTopHotKeysRequest) (<-chan *DescribeTopHotKeysResponse, <-chan error) {
	responseChan := make(chan *DescribeTopHotKeysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTopHotKeys(request)
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

// DescribeTopHotKeysWithCallback invokes the das.DescribeTopHotKeys API asynchronously
func (client *Client) DescribeTopHotKeysWithCallback(request *DescribeTopHotKeysRequest, callback func(response *DescribeTopHotKeysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTopHotKeysResponse
		var err error
		defer close(result)
		response, err = client.DescribeTopHotKeys(request)
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

// DescribeTopHotKeysRequest is the request struct for api DescribeTopHotKeys
type DescribeTopHotKeysRequest struct {
	*requests.RpcRequest
	EndTime        string `position:"Query" name:"EndTime"`
	StartTime      string `position:"Query" name:"StartTime"`
	ConsoleContext string `position:"Query" name:"ConsoleContext"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	NodeId         string `position:"Query" name:"NodeId"`
}

// DescribeTopHotKeysResponse is the response struct for api DescribeTopHotKeys
type DescribeTopHotKeysResponse struct {
	*responses.BaseResponse
	Message   string                   `json:"Message" xml:"Message"`
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	Code      string                   `json:"Code" xml:"Code"`
	Success   string                   `json:"Success" xml:"Success"`
	Data      DataInDescribeTopHotKeys `json:"Data" xml:"Data"`
}

// CreateDescribeTopHotKeysRequest creates a request to invoke DescribeTopHotKeys API
func CreateDescribeTopHotKeysRequest() (request *DescribeTopHotKeysRequest) {
	request = &DescribeTopHotKeysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "DescribeTopHotKeys", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeTopHotKeysResponse creates a response to parse from DescribeTopHotKeys response
func CreateDescribeTopHotKeysResponse() (response *DescribeTopHotKeysResponse) {
	response = &DescribeTopHotKeysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
