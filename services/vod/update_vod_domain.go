package vod

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

// UpdateVodDomain invokes the vod.UpdateVodDomain API synchronously
// api document: https://help.aliyun.com/api/vod/updatevoddomain.html
func (client *Client) UpdateVodDomain(request *UpdateVodDomainRequest) (response *UpdateVodDomainResponse, err error) {
	response = CreateUpdateVodDomainResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateVodDomainWithChan invokes the vod.UpdateVodDomain API asynchronously
// api document: https://help.aliyun.com/api/vod/updatevoddomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateVodDomainWithChan(request *UpdateVodDomainRequest) (<-chan *UpdateVodDomainResponse, <-chan error) {
	responseChan := make(chan *UpdateVodDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateVodDomain(request)
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

// UpdateVodDomainWithCallback invokes the vod.UpdateVodDomain API asynchronously
// api document: https://help.aliyun.com/api/vod/updatevoddomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateVodDomainWithCallback(request *UpdateVodDomainRequest, callback func(response *UpdateVodDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateVodDomainResponse
		var err error
		defer close(result)
		response, err = client.UpdateVodDomain(request)
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

// UpdateVodDomainRequest is the request struct for api UpdateVodDomain
type UpdateVodDomainRequest struct {
	*requests.RpcRequest
	TopLevelDomain string           `position:"Query" name:"TopLevelDomain"`
	Sources        string           `position:"Query" name:"Sources"`
	SecurityToken  string           `position:"Query" name:"SecurityToken"`
	DomainName     string           `position:"Query" name:"DomainName"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
}

// UpdateVodDomainResponse is the response struct for api UpdateVodDomain
type UpdateVodDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateVodDomainRequest creates a request to invoke UpdateVodDomain API
func CreateUpdateVodDomainRequest() (request *UpdateVodDomainRequest) {
	request = &UpdateVodDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "UpdateVodDomain", "vod", "openAPI")
	return
}

// CreateUpdateVodDomainResponse creates a response to parse from UpdateVodDomain response
func CreateUpdateVodDomainResponse() (response *UpdateVodDomainResponse) {
	response = &UpdateVodDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
