package hitsdb

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

// GetLindormInstance invokes the hitsdb.GetLindormInstance API synchronously
func (client *Client) GetLindormInstance(request *GetLindormInstanceRequest) (response *GetLindormInstanceResponse, err error) {
	response = CreateGetLindormInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// GetLindormInstanceWithChan invokes the hitsdb.GetLindormInstance API asynchronously
func (client *Client) GetLindormInstanceWithChan(request *GetLindormInstanceRequest) (<-chan *GetLindormInstanceResponse, <-chan error) {
	responseChan := make(chan *GetLindormInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLindormInstance(request)
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

// GetLindormInstanceWithCallback invokes the hitsdb.GetLindormInstance API asynchronously
func (client *Client) GetLindormInstanceWithCallback(request *GetLindormInstanceRequest, callback func(response *GetLindormInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLindormInstanceResponse
		var err error
		defer close(result)
		response, err = client.GetLindormInstance(request)
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

// GetLindormInstanceRequest is the request struct for api GetLindormInstance
type GetLindormInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// GetLindormInstanceResponse is the response struct for api GetLindormInstance
type GetLindormInstanceResponse struct {
	*responses.BaseResponse
	RequestId           string   `json:"RequestId" xml:"RequestId"`
	InstanceId          string   `json:"InstanceId" xml:"InstanceId"`
	InstanceAlias       string   `json:"InstanceAlias" xml:"InstanceAlias"`
	RegionId            string   `json:"RegionId" xml:"RegionId"`
	ZoneId              string   `json:"ZoneId" xml:"ZoneId"`
	InstanceStatus      string   `json:"InstanceStatus" xml:"InstanceStatus"`
	PayType             string   `json:"PayType" xml:"PayType"`
	NetworkType         string   `json:"NetworkType" xml:"NetworkType"`
	CreateTime          string   `json:"CreateTime" xml:"CreateTime"`
	ExpireTime          string   `json:"ExpireTime" xml:"ExpireTime"`
	InstanceStorage     string   `json:"InstanceStorage" xml:"InstanceStorage"`
	VpcId               string   `json:"VpcId" xml:"VpcId"`
	VswitchId           string   `json:"VswitchId" xml:"VswitchId"`
	AutoRenew           bool     `json:"AutoRenew" xml:"AutoRenew"`
	EngineType          int      `json:"EngineType" xml:"EngineType"`
	ServiceType         string   `json:"ServiceType" xml:"ServiceType"`
	DeletionProtection  string   `json:"DeletionProtection" xml:"DeletionProtection"`
	DiskCategory        string   `json:"DiskCategory" xml:"DiskCategory"`
	ColdStorage         int      `json:"ColdStorage" xml:"ColdStorage"`
	EnableBDS           bool     `json:"EnableBDS" xml:"EnableBDS"`
	AliUid              int64    `json:"AliUid" xml:"AliUid"`
	EnableFS            bool     `json:"EnableFS" xml:"EnableFS"`
	EnablePhoenix       bool     `json:"EnablePhoenix" xml:"EnablePhoenix"`
	DiskUsage           string   `json:"DiskUsage" xml:"DiskUsage"`
	DiskThreshold       string   `json:"DiskThreshold" xml:"DiskThreshold"`
	CreateMilliseconds  int64    `json:"CreateMilliseconds" xml:"CreateMilliseconds"`
	ExpiredMilliseconds int64    `json:"ExpiredMilliseconds" xml:"ExpiredMilliseconds"`
	EnableKms           bool     `json:"EnableKms" xml:"EnableKms"`
	EngineList          []Engine `json:"EngineList" xml:"EngineList"`
}

// CreateGetLindormInstanceRequest creates a request to invoke GetLindormInstance API
func CreateGetLindormInstanceRequest() (request *GetLindormInstanceRequest) {
	request = &GetLindormInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2020-06-15", "GetLindormInstance", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetLindormInstanceResponse creates a response to parse from GetLindormInstance response
func CreateGetLindormInstanceResponse() (response *GetLindormInstanceResponse) {
	response = &GetLindormInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
