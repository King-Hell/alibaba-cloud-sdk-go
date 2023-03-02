package lto

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

// ListAllPrivacyAlgorithm invokes the lto.ListAllPrivacyAlgorithm API synchronously
func (client *Client) ListAllPrivacyAlgorithm(request *ListAllPrivacyAlgorithmRequest) (response *ListAllPrivacyAlgorithmResponse, err error) {
	response = CreateListAllPrivacyAlgorithmResponse()
	err = client.DoAction(request, response)
	return
}

// ListAllPrivacyAlgorithmWithChan invokes the lto.ListAllPrivacyAlgorithm API asynchronously
func (client *Client) ListAllPrivacyAlgorithmWithChan(request *ListAllPrivacyAlgorithmRequest) (<-chan *ListAllPrivacyAlgorithmResponse, <-chan error) {
	responseChan := make(chan *ListAllPrivacyAlgorithmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAllPrivacyAlgorithm(request)
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

// ListAllPrivacyAlgorithmWithCallback invokes the lto.ListAllPrivacyAlgorithm API asynchronously
func (client *Client) ListAllPrivacyAlgorithmWithCallback(request *ListAllPrivacyAlgorithmRequest, callback func(response *ListAllPrivacyAlgorithmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAllPrivacyAlgorithmResponse
		var err error
		defer close(result)
		response, err = client.ListAllPrivacyAlgorithm(request)
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

// ListAllPrivacyAlgorithmRequest is the request struct for api ListAllPrivacyAlgorithm
type ListAllPrivacyAlgorithmRequest struct {
	*requests.RpcRequest
}

// ListAllPrivacyAlgorithmResponse is the response struct for api ListAllPrivacyAlgorithm
type ListAllPrivacyAlgorithmResponse struct {
	*responses.BaseResponse
	Code           string        `json:"Code" xml:"Code"`
	HttpStatusCode int           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string        `json:"Message" xml:"Message"`
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	Success        bool          `json:"Success" xml:"Success"`
	Data           []AlgTypeInfo `json:"Data" xml:"Data"`
}

// CreateListAllPrivacyAlgorithmRequest creates a request to invoke ListAllPrivacyAlgorithm API
func CreateListAllPrivacyAlgorithmRequest() (request *ListAllPrivacyAlgorithmRequest) {
	request = &ListAllPrivacyAlgorithmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "ListAllPrivacyAlgorithm", "", "")
	request.Method = requests.POST
	return
}

// CreateListAllPrivacyAlgorithmResponse creates a response to parse from ListAllPrivacyAlgorithm response
func CreateListAllPrivacyAlgorithmResponse() (response *ListAllPrivacyAlgorithmResponse) {
	response = &ListAllPrivacyAlgorithmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
