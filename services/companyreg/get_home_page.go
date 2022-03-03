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

// GetHomePage invokes the companyreg.GetHomePage API synchronously
func (client *Client) GetHomePage(request *GetHomePageRequest) (response *GetHomePageResponse, err error) {
	response = CreateGetHomePageResponse()
	err = client.DoAction(request, response)
	return
}

// GetHomePageWithChan invokes the companyreg.GetHomePage API asynchronously
func (client *Client) GetHomePageWithChan(request *GetHomePageRequest) (<-chan *GetHomePageResponse, <-chan error) {
	responseChan := make(chan *GetHomePageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetHomePage(request)
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

// GetHomePageWithCallback invokes the companyreg.GetHomePage API asynchronously
func (client *Client) GetHomePageWithCallback(request *GetHomePageRequest, callback func(response *GetHomePageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetHomePageResponse
		var err error
		defer close(result)
		response, err = client.GetHomePage(request)
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

// GetHomePageRequest is the request struct for api GetHomePage
type GetHomePageRequest struct {
	*requests.RpcRequest
	BizId string `position:"Query" name:"BizId"`
}

// GetHomePageResponse is the response struct for api GetHomePage
type GetHomePageResponse struct {
	*responses.BaseResponse
	PageIndex string `json:"PageIndex" xml:"PageIndex"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGetHomePageRequest creates a request to invoke GetHomePage API
func CreateGetHomePageRequest() (request *GetHomePageRequest) {
	request = &GetHomePageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "GetHomePage", "", "")
	request.Method = requests.GET
	return
}

// CreateGetHomePageResponse creates a response to parse from GetHomePage response
func CreateGetHomePageResponse() (response *GetHomePageResponse) {
	response = &GetHomePageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
