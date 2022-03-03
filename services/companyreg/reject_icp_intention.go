package companyreg

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

// RejectIcpIntention invokes the companyreg.RejectIcpIntention API synchronously
func (client *Client) RejectIcpIntention(request *RejectIcpIntentionRequest) (response *RejectIcpIntentionResponse, err error) {
	response = CreateRejectIcpIntentionResponse()
	err = client.DoAction(request, response)
	return
}

// RejectIcpIntentionWithChan invokes the companyreg.RejectIcpIntention API asynchronously
func (client *Client) RejectIcpIntentionWithChan(request *RejectIcpIntentionRequest) (<-chan *RejectIcpIntentionResponse, <-chan error) {
	responseChan := make(chan *RejectIcpIntentionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RejectIcpIntention(request)
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

// RejectIcpIntentionWithCallback invokes the companyreg.RejectIcpIntention API asynchronously
func (client *Client) RejectIcpIntentionWithCallback(request *RejectIcpIntentionRequest, callback func(response *RejectIcpIntentionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RejectIcpIntentionResponse
		var err error
		defer close(result)
		response, err = client.RejectIcpIntention(request)
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

// RejectIcpIntentionRequest is the request struct for api RejectIcpIntention
type RejectIcpIntentionRequest struct {
	*requests.RpcRequest
	Note  string `position:"Query" name:"Note"`
	BizId string `position:"Query" name:"BizId"`
}

// RejectIcpIntentionResponse is the response struct for api RejectIcpIntention
type RejectIcpIntentionResponse struct {
	*responses.BaseResponse
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateRejectIcpIntentionRequest creates a request to invoke RejectIcpIntention API
func CreateRejectIcpIntentionRequest() (request *RejectIcpIntentionRequest) {
	request = &RejectIcpIntentionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "RejectIcpIntention", "", "")
	request.Method = requests.POST
	return
}

// CreateRejectIcpIntentionResponse creates a response to parse from RejectIcpIntention response
func CreateRejectIcpIntentionResponse() (response *RejectIcpIntentionResponse) {
	response = &RejectIcpIntentionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
