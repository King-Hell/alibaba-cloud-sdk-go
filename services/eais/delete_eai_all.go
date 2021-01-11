package eais

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

// DeleteEaiAll invokes the eais.DeleteEaiAll API synchronously
func (client *Client) DeleteEaiAll(request *DeleteEaiAllRequest) (response *DeleteEaiAllResponse, err error) {
	response = CreateDeleteEaiAllResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEaiAllWithChan invokes the eais.DeleteEaiAll API asynchronously
func (client *Client) DeleteEaiAllWithChan(request *DeleteEaiAllRequest) (<-chan *DeleteEaiAllResponse, <-chan error) {
	responseChan := make(chan *DeleteEaiAllResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEaiAll(request)
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

// DeleteEaiAllWithCallback invokes the eais.DeleteEaiAll API asynchronously
func (client *Client) DeleteEaiAllWithCallback(request *DeleteEaiAllRequest, callback func(response *DeleteEaiAllResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEaiAllResponse
		var err error
		defer close(result)
		response, err = client.DeleteEaiAll(request)
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

// DeleteEaiAllRequest is the request struct for api DeleteEaiAll
type DeleteEaiAllRequest struct {
	*requests.RpcRequest
	ClientInstanceId             string `position:"Query" name:"ClientInstanceId"`
	ElasticAcceleratedInstanceId string `position:"Query" name:"ElasticAcceleratedInstanceId"`
}

// DeleteEaiAllResponse is the response struct for api DeleteEaiAll
type DeleteEaiAllResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEaiAllRequest creates a request to invoke DeleteEaiAll API
func CreateDeleteEaiAllRequest() (request *DeleteEaiAllRequest) {
	request = &DeleteEaiAllRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eais", "2019-06-24", "DeleteEaiAll", "eais", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEaiAllResponse creates a response to parse from DeleteEaiAll response
func CreateDeleteEaiAllResponse() (response *DeleteEaiAllResponse) {
	response = &DeleteEaiAllResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
