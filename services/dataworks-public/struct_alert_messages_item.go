package dataworks_public

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

// AlertMessagesItem is a nested struct in dataworks_public response
type AlertMessagesItem struct {
	RemindId           int64           `json:"RemindId" xml:"RemindId"`
	AlertMessageStatus string          `json:"AlertMessageStatus" xml:"AlertMessageStatus"`
	AlertUser          string          `json:"AlertUser" xml:"AlertUser"`
	AlertTime          int64           `json:"AlertTime" xml:"AlertTime"`
	AlertMethod        string          `json:"AlertMethod" xml:"AlertMethod"`
	Source             string          `json:"Source" xml:"Source"`
	Content            string          `json:"Content" xml:"Content"`
	RemindName         string          `json:"RemindName" xml:"RemindName"`
	AlertId            int64           `json:"AlertId" xml:"AlertId"`
	SlaAlert           SlaAlert        `json:"SlaAlert" xml:"SlaAlert"`
	Instances          []InstancesItem `json:"Instances" xml:"Instances"`
	Topics             []TopicsItem    `json:"Topics" xml:"Topics"`
	Nodes              []NodesItem     `json:"Nodes" xml:"Nodes"`
}
