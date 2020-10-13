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

// QueryConsumerGroupStatus invokes the iot.QueryConsumerGroupStatus API synchronously
func (client *Client) QueryConsumerGroupStatus(request *QueryConsumerGroupStatusRequest) (response *QueryConsumerGroupStatusResponse, err error) {
	response = CreateQueryConsumerGroupStatusResponse()
	err = client.DoAction(request, response)
	return
}

// QueryConsumerGroupStatusWithChan invokes the iot.QueryConsumerGroupStatus API asynchronously
func (client *Client) QueryConsumerGroupStatusWithChan(request *QueryConsumerGroupStatusRequest) (<-chan *QueryConsumerGroupStatusResponse, <-chan error) {
	responseChan := make(chan *QueryConsumerGroupStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryConsumerGroupStatus(request)
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

// QueryConsumerGroupStatusWithCallback invokes the iot.QueryConsumerGroupStatus API asynchronously
func (client *Client) QueryConsumerGroupStatusWithCallback(request *QueryConsumerGroupStatusRequest, callback func(response *QueryConsumerGroupStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryConsumerGroupStatusResponse
		var err error
		defer close(result)
		response, err = client.QueryConsumerGroupStatus(request)
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

// QueryConsumerGroupStatusRequest is the request struct for api QueryConsumerGroupStatus
type QueryConsumerGroupStatusRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	GroupId       string `position:"Query" name:"GroupId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// QueryConsumerGroupStatusResponse is the response struct for api QueryConsumerGroupStatus
type QueryConsumerGroupStatusResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	Success                    bool                       `json:"Success" xml:"Success"`
	ErrorMessage               string                     `json:"ErrorMessage" xml:"ErrorMessage"`
	AccumulationCount          int                        `json:"AccumulationCount" xml:"AccumulationCount"`
	ConsumerSpeed              int                        `json:"ConsumerSpeed" xml:"ConsumerSpeed"`
	LastConsumerTime           string                     `json:"LastConsumerTime" xml:"LastConsumerTime"`
	Code                       string                     `json:"Code" xml:"Code"`
	ClientConnectionStatusList ClientConnectionStatusList `json:"ClientConnectionStatusList" xml:"ClientConnectionStatusList"`
}

// CreateQueryConsumerGroupStatusRequest creates a request to invoke QueryConsumerGroupStatus API
func CreateQueryConsumerGroupStatusRequest() (request *QueryConsumerGroupStatusRequest) {
	request = &QueryConsumerGroupStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryConsumerGroupStatus", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryConsumerGroupStatusResponse creates a response to parse from QueryConsumerGroupStatus response
func CreateQueryConsumerGroupStatusResponse() (response *QueryConsumerGroupStatusResponse) {
	response = &QueryConsumerGroupStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
