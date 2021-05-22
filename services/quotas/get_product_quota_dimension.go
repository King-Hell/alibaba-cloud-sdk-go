package quotas

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

// GetProductQuotaDimension invokes the quotas.GetProductQuotaDimension API synchronously
func (client *Client) GetProductQuotaDimension(request *GetProductQuotaDimensionRequest) (response *GetProductQuotaDimensionResponse, err error) {
	response = CreateGetProductQuotaDimensionResponse()
	err = client.DoAction(request, response)
	return
}

// GetProductQuotaDimensionWithChan invokes the quotas.GetProductQuotaDimension API asynchronously
func (client *Client) GetProductQuotaDimensionWithChan(request *GetProductQuotaDimensionRequest) (<-chan *GetProductQuotaDimensionResponse, <-chan error) {
	responseChan := make(chan *GetProductQuotaDimensionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProductQuotaDimension(request)
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

// GetProductQuotaDimensionWithCallback invokes the quotas.GetProductQuotaDimension API asynchronously
func (client *Client) GetProductQuotaDimensionWithCallback(request *GetProductQuotaDimensionRequest, callback func(response *GetProductQuotaDimensionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProductQuotaDimensionResponse
		var err error
		defer close(result)
		response, err = client.GetProductQuotaDimension(request)
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

// GetProductQuotaDimensionRequest is the request struct for api GetProductQuotaDimension
type GetProductQuotaDimensionRequest struct {
	*requests.RpcRequest
	ProductCode         string                                         `position:"Body" name:"ProductCode"`
	DependentDimensions *[]GetProductQuotaDimensionDependentDimensions `position:"Body" name:"DependentDimensions"  type:"Repeated"`
	DimensionKey        string                                         `position:"Body" name:"DimensionKey"`
}

// GetProductQuotaDimensionDependentDimensions is a repeated param struct in GetProductQuotaDimensionRequest
type GetProductQuotaDimensionDependentDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// GetProductQuotaDimensionResponse is the response struct for api GetProductQuotaDimension
type GetProductQuotaDimensionResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	QuotaDimension QuotaDimension `json:"QuotaDimension" xml:"QuotaDimension"`
}

// CreateGetProductQuotaDimensionRequest creates a request to invoke GetProductQuotaDimension API
func CreateGetProductQuotaDimensionRequest() (request *GetProductQuotaDimensionRequest) {
	request = &GetProductQuotaDimensionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quotas", "2020-05-10", "GetProductQuotaDimension", "quotas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetProductQuotaDimensionResponse creates a response to parse from GetProductQuotaDimension response
func CreateGetProductQuotaDimensionResponse() (response *GetProductQuotaDimensionResponse) {
	response = &GetProductQuotaDimensionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}