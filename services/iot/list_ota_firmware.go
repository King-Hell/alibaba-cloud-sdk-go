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

// ListOTAFirmware invokes the iot.ListOTAFirmware API synchronously
func (client *Client) ListOTAFirmware(request *ListOTAFirmwareRequest) (response *ListOTAFirmwareResponse, err error) {
	response = CreateListOTAFirmwareResponse()
	err = client.DoAction(request, response)
	return
}

// ListOTAFirmwareWithChan invokes the iot.ListOTAFirmware API asynchronously
func (client *Client) ListOTAFirmwareWithChan(request *ListOTAFirmwareRequest) (<-chan *ListOTAFirmwareResponse, <-chan error) {
	responseChan := make(chan *ListOTAFirmwareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOTAFirmware(request)
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

// ListOTAFirmwareWithCallback invokes the iot.ListOTAFirmware API asynchronously
func (client *Client) ListOTAFirmwareWithCallback(request *ListOTAFirmwareRequest, callback func(response *ListOTAFirmwareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOTAFirmwareResponse
		var err error
		defer close(result)
		response, err = client.ListOTAFirmware(request)
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

// ListOTAFirmwareRequest is the request struct for api ListOTAFirmware
type ListOTAFirmwareRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DestVersion   string           `position:"Query" name:"DestVersion"`
}

// ListOTAFirmwareResponse is the response struct for api ListOTAFirmware
type ListOTAFirmwareResponse struct {
	*responses.BaseResponse
	RequestId    string                        `json:"RequestId" xml:"RequestId"`
	Success      bool                          `json:"Success" xml:"Success"`
	Code         string                        `json:"Code" xml:"Code"`
	ErrorMessage string                        `json:"ErrorMessage" xml:"ErrorMessage"`
	Total        int                           `json:"Total" xml:"Total"`
	PageSize     int                           `json:"PageSize" xml:"PageSize"`
	PageCount    int                           `json:"PageCount" xml:"PageCount"`
	CurrentPage  int                           `json:"CurrentPage" xml:"CurrentPage"`
	FirmwareInfo FirmwareInfoInListOTAFirmware `json:"FirmwareInfo" xml:"FirmwareInfo"`
}

// CreateListOTAFirmwareRequest creates a request to invoke ListOTAFirmware API
func CreateListOTAFirmwareRequest() (request *ListOTAFirmwareRequest) {
	request = &ListOTAFirmwareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ListOTAFirmware", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListOTAFirmwareResponse creates a response to parse from ListOTAFirmware response
func CreateListOTAFirmwareResponse() (response *ListOTAFirmwareResponse) {
	response = &ListOTAFirmwareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
