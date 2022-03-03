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

// SaveOpeningBalanceStatus invokes the companyreg.SaveOpeningBalanceStatus API synchronously
func (client *Client) SaveOpeningBalanceStatus(request *SaveOpeningBalanceStatusRequest) (response *SaveOpeningBalanceStatusResponse, err error) {
	response = CreateSaveOpeningBalanceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// SaveOpeningBalanceStatusWithChan invokes the companyreg.SaveOpeningBalanceStatus API asynchronously
func (client *Client) SaveOpeningBalanceStatusWithChan(request *SaveOpeningBalanceStatusRequest) (<-chan *SaveOpeningBalanceStatusResponse, <-chan error) {
	responseChan := make(chan *SaveOpeningBalanceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveOpeningBalanceStatus(request)
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

// SaveOpeningBalanceStatusWithCallback invokes the companyreg.SaveOpeningBalanceStatus API asynchronously
func (client *Client) SaveOpeningBalanceStatusWithCallback(request *SaveOpeningBalanceStatusRequest, callback func(response *SaveOpeningBalanceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveOpeningBalanceStatusResponse
		var err error
		defer close(result)
		response, err = client.SaveOpeningBalanceStatus(request)
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

// SaveOpeningBalanceStatusRequest is the request struct for api SaveOpeningBalanceStatus
type SaveOpeningBalanceStatusRequest struct {
	*requests.RpcRequest
	OpeningBalanceStatus string `position:"Query" name:"OpeningBalanceStatus"`
	BizId                string `position:"Query" name:"BizId"`
}

// SaveOpeningBalanceStatusResponse is the response struct for api SaveOpeningBalanceStatus
type SaveOpeningBalanceStatusResponse struct {
	*responses.BaseResponse
	Result    bool   `json:"Result" xml:"Result"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSaveOpeningBalanceStatusRequest creates a request to invoke SaveOpeningBalanceStatus API
func CreateSaveOpeningBalanceStatusRequest() (request *SaveOpeningBalanceStatusRequest) {
	request = &SaveOpeningBalanceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "SaveOpeningBalanceStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveOpeningBalanceStatusResponse creates a response to parse from SaveOpeningBalanceStatus response
func CreateSaveOpeningBalanceStatusResponse() (response *SaveOpeningBalanceStatusResponse) {
	response = &SaveOpeningBalanceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
