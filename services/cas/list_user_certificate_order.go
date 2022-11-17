package cas

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

// ListUserCertificateOrder invokes the cas.ListUserCertificateOrder API synchronously
func (client *Client) ListUserCertificateOrder(request *ListUserCertificateOrderRequest) (response *ListUserCertificateOrderResponse, err error) {
	response = CreateListUserCertificateOrderResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserCertificateOrderWithChan invokes the cas.ListUserCertificateOrder API asynchronously
func (client *Client) ListUserCertificateOrderWithChan(request *ListUserCertificateOrderRequest) (<-chan *ListUserCertificateOrderResponse, <-chan error) {
	responseChan := make(chan *ListUserCertificateOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserCertificateOrder(request)
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

// ListUserCertificateOrderWithCallback invokes the cas.ListUserCertificateOrder API asynchronously
func (client *Client) ListUserCertificateOrderWithCallback(request *ListUserCertificateOrderRequest, callback func(response *ListUserCertificateOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserCertificateOrderResponse
		var err error
		defer close(result)
		response, err = client.ListUserCertificateOrder(request)
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

// ListUserCertificateOrderRequest is the request struct for api ListUserCertificateOrder
type ListUserCertificateOrderRequest struct {
	*requests.RpcRequest
	ShowSize    requests.Integer `position:"Query" name:"ShowSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Keyword     string           `position:"Query" name:"Keyword"`
	Status      string           `position:"Query" name:"Status"`
	OrderType   string           `position:"Query" name:"OrderType"`
}

// ListUserCertificateOrderResponse is the response struct for api ListUserCertificateOrder
type ListUserCertificateOrderResponse struct {
	*responses.BaseResponse
	ShowSize             int64                      `json:"ShowSize" xml:"ShowSize"`
	CurrentPage          int64                      `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount           int64                      `json:"TotalCount" xml:"TotalCount"`
	RequestId            string                     `json:"RequestId" xml:"RequestId"`
	CertificateOrderList []CertificateOrderListItem `json:"CertificateOrderList" xml:"CertificateOrderList"`
}

// CreateListUserCertificateOrderRequest creates a request to invoke ListUserCertificateOrder API
func CreateListUserCertificateOrderRequest() (request *ListUserCertificateOrderRequest) {
	request = &ListUserCertificateOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cas", "2020-04-07", "ListUserCertificateOrder", "", "")
	request.Method = requests.POST
	return
}

// CreateListUserCertificateOrderResponse creates a response to parse from ListUserCertificateOrder response
func CreateListUserCertificateOrderResponse() (response *ListUserCertificateOrderResponse) {
	response = &ListUserCertificateOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
