package polardb

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

// ModifyDBClusterPrimaryZone invokes the polardb.ModifyDBClusterPrimaryZone API synchronously
func (client *Client) ModifyDBClusterPrimaryZone(request *ModifyDBClusterPrimaryZoneRequest) (response *ModifyDBClusterPrimaryZoneResponse, err error) {
	response = CreateModifyDBClusterPrimaryZoneResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterPrimaryZoneWithChan invokes the polardb.ModifyDBClusterPrimaryZone API asynchronously
func (client *Client) ModifyDBClusterPrimaryZoneWithChan(request *ModifyDBClusterPrimaryZoneRequest) (<-chan *ModifyDBClusterPrimaryZoneResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterPrimaryZoneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterPrimaryZone(request)
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

// ModifyDBClusterPrimaryZoneWithCallback invokes the polardb.ModifyDBClusterPrimaryZone API asynchronously
func (client *Client) ModifyDBClusterPrimaryZoneWithCallback(request *ModifyDBClusterPrimaryZoneRequest, callback func(response *ModifyDBClusterPrimaryZoneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterPrimaryZoneResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterPrimaryZone(request)
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

// ModifyDBClusterPrimaryZoneRequest is the request struct for api ModifyDBClusterPrimaryZone
type ModifyDBClusterPrimaryZoneRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PlannedEndTime       string           `position:"Query" name:"PlannedEndTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	PlannedStartTime     string           `position:"Query" name:"PlannedStartTime"`
	VPCId                string           `position:"Query" name:"VPCId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	FromTimeService      requests.Boolean `position:"Query" name:"FromTimeService"`
}

// ModifyDBClusterPrimaryZoneResponse is the response struct for api ModifyDBClusterPrimaryZone
type ModifyDBClusterPrimaryZoneResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBClusterPrimaryZoneRequest creates a request to invoke ModifyDBClusterPrimaryZone API
func CreateModifyDBClusterPrimaryZoneRequest() (request *ModifyDBClusterPrimaryZoneRequest) {
	request = &ModifyDBClusterPrimaryZoneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterPrimaryZone", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterPrimaryZoneResponse creates a response to parse from ModifyDBClusterPrimaryZone response
func CreateModifyDBClusterPrimaryZoneResponse() (response *ModifyDBClusterPrimaryZoneResponse) {
	response = &ModifyDBClusterPrimaryZoneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
