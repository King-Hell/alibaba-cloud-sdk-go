package iot

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

// SpeechBySynthesis invokes the iot.SpeechBySynthesis API synchronously
func (client *Client) SpeechBySynthesis(request *SpeechBySynthesisRequest) (response *SpeechBySynthesisResponse, err error) {
	response = CreateSpeechBySynthesisResponse()
	err = client.DoAction(request, response)
	return
}

// SpeechBySynthesisWithChan invokes the iot.SpeechBySynthesis API asynchronously
func (client *Client) SpeechBySynthesisWithChan(request *SpeechBySynthesisRequest) (<-chan *SpeechBySynthesisResponse, <-chan error) {
	responseChan := make(chan *SpeechBySynthesisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SpeechBySynthesis(request)
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

// SpeechBySynthesisWithCallback invokes the iot.SpeechBySynthesis API asynchronously
func (client *Client) SpeechBySynthesisWithCallback(request *SpeechBySynthesisRequest, callback func(response *SpeechBySynthesisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SpeechBySynthesisResponse
		var err error
		defer close(result)
		response, err = client.SpeechBySynthesis(request)
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

// SpeechBySynthesisRequest is the request struct for api SpeechBySynthesis
type SpeechBySynthesisRequest struct {
	*requests.RpcRequest
	Voice         string           `position:"Body" name:"Voice"`
	SpeechId      string           `position:"Body" name:"SpeechId"`
	AudioFormat   string           `position:"Body" name:"AudioFormat"`
	IotId         string           `position:"Body" name:"IotId"`
	IotInstanceId string           `position:"Body" name:"IotInstanceId"`
	Text          string           `position:"Body" name:"Text"`
	ProductKey    string           `position:"Body" name:"ProductKey"`
	Volume        requests.Integer `position:"Body" name:"Volume"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Body" name:"DeviceName"`
	SpeechRate    requests.Integer `position:"Body" name:"SpeechRate"`
}

// SpeechBySynthesisResponse is the response struct for api SpeechBySynthesis
type SpeechBySynthesisResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateSpeechBySynthesisRequest creates a request to invoke SpeechBySynthesis API
func CreateSpeechBySynthesisRequest() (request *SpeechBySynthesisRequest) {
	request = &SpeechBySynthesisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "SpeechBySynthesis", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSpeechBySynthesisResponse creates a response to parse from SpeechBySynthesis response
func CreateSpeechBySynthesisResponse() (response *SpeechBySynthesisResponse) {
	response = &SpeechBySynthesisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
