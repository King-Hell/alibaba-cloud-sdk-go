package oceanbasepro

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

// DataItem is a nested struct in oceanbasepro response
type DataItem struct {
	StepStatus                string                 `json:"StepStatus" xml:"StepStatus"`
	User                      string                 `json:"User" xml:"User"`
	MaxCpu                    int64                  `json:"MaxCpu" xml:"MaxCpu"`
	SqlText                   string                 `json:"SqlText" xml:"SqlText"`
	ErrorMessage              string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	Metric                    string                 `json:"Metric" xml:"Metric"`
	MinCpu                    int64                  `json:"MinCpu" xml:"MinCpu"`
	SessionId                 int64                  `json:"SessionId" xml:"SessionId"`
	ServerIp                  string                 `json:"ServerIp" xml:"ServerIp"`
	ProjectName               string                 `json:"ProjectName" xml:"ProjectName"`
	ProjectId                 string                 `json:"ProjectId" xml:"ProjectId"`
	StepDescription           string                 `json:"StepDescription" xml:"StepDescription"`
	Tags                      map[string]interface{} `json:"Tags" xml:"Tags"`
	TenantId                  string                 `json:"TenantId" xml:"TenantId"`
	StepName                  string                 `json:"StepName" xml:"StepName"`
	ProjectOwner              string                 `json:"ProjectOwner" xml:"ProjectOwner"`
	StepOrder                 int                    `json:"StepOrder" xml:"StepOrder"`
	UnitNum                   int64                  `json:"UnitNum" xml:"UnitNum"`
	FinishTime                string                 `json:"FinishTime" xml:"FinishTime"`
	StepProgress              int                    `json:"StepProgress" xml:"StepProgress"`
	Command                   string                 `json:"Command" xml:"Command"`
	ExecuteTime               string                 `json:"ExecuteTime" xml:"ExecuteTime"`
	Interactive               bool                   `json:"Interactive" xml:"Interactive"`
	StartTime                 string                 `json:"StartTime" xml:"StartTime"`
	EstimatedRemainingSeconds int64                  `json:"EstimatedRemainingSeconds" xml:"EstimatedRemainingSeconds"`
	Status                    string                 `json:"Status" xml:"Status"`
	Database                  string                 `json:"Database" xml:"Database"`
	BusinessName              string                 `json:"BusinessName" xml:"BusinessName"`
	ClientIp                  string                 `json:"ClientIp" xml:"ClientIp"`
	DestConfig                DestConfig             `json:"DestConfig" xml:"DestConfig"`
	StepInfo                  StepInfo               `json:"StepInfo" xml:"StepInfo"`
	SourceConfig              SourceConfig           `json:"SourceConfig" xml:"SourceConfig"`
	TransferMapping           TransferMapping        `json:"TransferMapping" xml:"TransferMapping"`
	ExtraInfo                 ExtraInfo              `json:"ExtraInfo" xml:"ExtraInfo"`
	TransferStepConfig        TransferStepConfig     `json:"TransferStepConfig" xml:"TransferStepConfig"`
	DataPoints                []DataPoint            `json:"DataPoints" xml:"DataPoints"`
	Steps                     []Step                 `json:"Steps" xml:"Steps"`
	Labels                    []Label                `json:"Labels" xml:"Labels"`
}
