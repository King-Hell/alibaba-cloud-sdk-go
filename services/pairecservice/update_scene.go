package pairecservice

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

// UpdateScene invokes the pairecservice.UpdateScene API synchronously
func (client *Client) UpdateScene(request *UpdateSceneRequest) (response *UpdateSceneResponse, err error) {
	response = CreateUpdateSceneResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSceneWithChan invokes the pairecservice.UpdateScene API asynchronously
func (client *Client) UpdateSceneWithChan(request *UpdateSceneRequest) (<-chan *UpdateSceneResponse, <-chan error) {
	responseChan := make(chan *UpdateSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateScene(request)
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

// UpdateSceneWithCallback invokes the pairecservice.UpdateScene API asynchronously
func (client *Client) UpdateSceneWithCallback(request *UpdateSceneRequest, callback func(response *UpdateSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSceneResponse
		var err error
		defer close(result)
		response, err = client.UpdateScene(request)
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

// UpdateSceneRequest is the request struct for api UpdateScene
type UpdateSceneRequest struct {
	*requests.RoaRequest
	SceneId string `position:"Path" name:"SceneId"`
	Body    string `position:"Body" name:"body"`
}

// UpdateSceneResponse is the response struct for api UpdateScene
type UpdateSceneResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateSceneRequest creates a request to invoke UpdateScene API
func CreateUpdateSceneRequest() (request *UpdateSceneRequest) {
	request = &UpdateSceneRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "UpdateScene", "/api/v1/scenes/[SceneId]", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateSceneResponse creates a response to parse from UpdateScene response
func CreateUpdateSceneResponse() (response *UpdateSceneResponse) {
	response = &UpdateSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
