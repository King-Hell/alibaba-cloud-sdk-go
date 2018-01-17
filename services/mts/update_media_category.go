package mts

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

func (client *Client) UpdateMediaCategory(request *UpdateMediaCategoryRequest) (response *UpdateMediaCategoryResponse, err error) {
	response = CreateUpdateMediaCategoryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdateMediaCategoryWithChan(request *UpdateMediaCategoryRequest) (<-chan *UpdateMediaCategoryResponse, <-chan error) {
	responseChan := make(chan *UpdateMediaCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMediaCategory(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) UpdateMediaCategoryWithCallback(request *UpdateMediaCategoryRequest, callback func(response *UpdateMediaCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMediaCategoryResponse
		var err error
		defer close(result)
		response, err = client.UpdateMediaCategory(request)
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

type UpdateMediaCategoryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	CateId               requests.Integer `position:"Query" name:"CateId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MediaId              string           `position:"Query" name:"MediaId"`
}

type UpdateMediaCategoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateUpdateMediaCategoryRequest() (request *UpdateMediaCategoryRequest) {
	request = &UpdateMediaCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaCategory", "mts", "openAPI")
	return
}

func CreateUpdateMediaCategoryResponse() (response *UpdateMediaCategoryResponse) {
	response = &UpdateMediaCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
