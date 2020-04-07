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

package dts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeSubscriptionInstanceAlert invokes the dts.DescribeSubscriptionInstanceAlert API synchronously
// api document: https://help.aliyun.com/api/dts/describesubscriptioninstancealert.html
func (client *Client) DescribeSubscriptionInstanceAlert(request *DescribeSubscriptionInstanceAlertRequest) (response *DescribeSubscriptionInstanceAlertResponse, err error) {
	response = CreateDescribeSubscriptionInstanceAlertResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSubscriptionInstanceAlertWithChan invokes the dts.DescribeSubscriptionInstanceAlert API asynchronously
// api document: https://help.aliyun.com/api/dts/describesubscriptioninstancealert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSubscriptionInstanceAlertWithChan(request *DescribeSubscriptionInstanceAlertRequest) (<-chan *DescribeSubscriptionInstanceAlertResponse, <-chan error) {
	responseChan := make(chan *DescribeSubscriptionInstanceAlertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSubscriptionInstanceAlert(request)
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

// DescribeSubscriptionInstanceAlertWithCallback invokes the dts.DescribeSubscriptionInstanceAlert API asynchronously
// api document: https://help.aliyun.com/api/dts/describesubscriptioninstancealert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSubscriptionInstanceAlertWithCallback(request *DescribeSubscriptionInstanceAlertRequest, callback func(response *DescribeSubscriptionInstanceAlertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSubscriptionInstanceAlertResponse
		var err error
		defer close(result)
		response, err = client.DescribeSubscriptionInstanceAlert(request)
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

// DescribeSubscriptionInstanceAlertRequest is the request struct for api DescribeSubscriptionInstanceAlert
type DescribeSubscriptionInstanceAlertRequest struct {
	*requests.RpcRequest
	SubscriptionInstanceId string `position:"Query" name:"SubscriptionInstanceId"`
	ClientToken            string `position:"Query" name:"ClientToken"`
	OwnerId                string `position:"Query" name:"OwnerId"`
}

// DescribeSubscriptionInstanceAlertResponse is the response struct for api DescribeSubscriptionInstanceAlert
type DescribeSubscriptionInstanceAlertResponse struct {
	*responses.BaseResponse
	SubscriptionInstanceID   string `json:"SubscriptionInstanceID" xml:"SubscriptionInstanceID"`
	SubscriptionInstanceName string `json:"SubscriptionInstanceName" xml:"SubscriptionInstanceName"`
	DelayAlertStatus         string `json:"DelayAlertStatus" xml:"DelayAlertStatus"`
	DelayAlertPhone          string `json:"DelayAlertPhone" xml:"DelayAlertPhone"`
	DelayOverSeconds         string `json:"DelayOverSeconds" xml:"DelayOverSeconds"`
	ErrorAlertStatus         string `json:"ErrorAlertStatus" xml:"ErrorAlertStatus"`
	ErrorAlertPhone          string `json:"ErrorAlertPhone" xml:"ErrorAlertPhone"`
	Success                  string `json:"Success" xml:"Success"`
	ErrCode                  string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage               string `json:"ErrMessage" xml:"ErrMessage"`
	RequestId                string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeSubscriptionInstanceAlertRequest creates a request to invoke DescribeSubscriptionInstanceAlert API
func CreateDescribeSubscriptionInstanceAlertRequest() (request *DescribeSubscriptionInstanceAlertRequest) {
	request = &DescribeSubscriptionInstanceAlertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "DescribeSubscriptionInstanceAlert", "dts", "openAPI")
	return
}

// CreateDescribeSubscriptionInstanceAlertResponse creates a response to parse from DescribeSubscriptionInstanceAlert response
func CreateDescribeSubscriptionInstanceAlertResponse() (response *DescribeSubscriptionInstanceAlertResponse) {
	response = &DescribeSubscriptionInstanceAlertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}