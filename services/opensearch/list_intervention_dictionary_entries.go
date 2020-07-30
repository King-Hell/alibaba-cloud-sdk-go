package opensearch

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

// ListInterventionDictionaryEntries invokes the opensearch.ListInterventionDictionaryEntries API synchronously
// api document: https://help.aliyun.com/api/opensearch/listinterventiondictionaryentries.html
func (client *Client) ListInterventionDictionaryEntries(request *ListInterventionDictionaryEntriesRequest) (response *ListInterventionDictionaryEntriesResponse, err error) {
	response = CreateListInterventionDictionaryEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListInterventionDictionaryEntriesWithChan invokes the opensearch.ListInterventionDictionaryEntries API asynchronously
// api document: https://help.aliyun.com/api/opensearch/listinterventiondictionaryentries.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListInterventionDictionaryEntriesWithChan(request *ListInterventionDictionaryEntriesRequest) (<-chan *ListInterventionDictionaryEntriesResponse, <-chan error) {
	responseChan := make(chan *ListInterventionDictionaryEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInterventionDictionaryEntries(request)
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

// ListInterventionDictionaryEntriesWithCallback invokes the opensearch.ListInterventionDictionaryEntries API asynchronously
// api document: https://help.aliyun.com/api/opensearch/listinterventiondictionaryentries.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListInterventionDictionaryEntriesWithCallback(request *ListInterventionDictionaryEntriesRequest, callback func(response *ListInterventionDictionaryEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInterventionDictionaryEntriesResponse
		var err error
		defer close(result)
		response, err = client.ListInterventionDictionaryEntries(request)
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

// ListInterventionDictionaryEntriesRequest is the request struct for api ListInterventionDictionaryEntries
type ListInterventionDictionaryEntriesRequest struct {
	*requests.RoaRequest
	Name string `position:"Path" name:"name"`
	Word string `position:"Query" name:"word"`
}

// ListInterventionDictionaryEntriesResponse is the response struct for api ListInterventionDictionaryEntries
type ListInterventionDictionaryEntriesResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"requestId" xml:"requestId"`
	TotalCount int        `json:"totalCount" xml:"totalCount"`
	Result     []WordItem `json:"result" xml:"result"`
}

// CreateListInterventionDictionaryEntriesRequest creates a request to invoke ListInterventionDictionaryEntries API
func CreateListInterventionDictionaryEntriesRequest() (request *ListInterventionDictionaryEntriesRequest) {
	request = &ListInterventionDictionaryEntriesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListInterventionDictionaryEntries", "/v4/openapi/intervention-dictionaries/[name]/entries", "opensearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListInterventionDictionaryEntriesResponse creates a response to parse from ListInterventionDictionaryEntries response
func CreateListInterventionDictionaryEntriesResponse() (response *ListInterventionDictionaryEntriesResponse) {
	response = &ListInterventionDictionaryEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
