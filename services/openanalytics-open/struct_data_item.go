package openanalytics_open

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

// DataItem is a nested struct in openanalytics_open response
type DataItem struct {
	Name                     string                      `json:"Name" xml:"Name"`
	ChargeType               string                      `json:"ChargeType" xml:"ChargeType"`
	CreateTime               string                      `json:"CreateTime" xml:"CreateTime"`
	DefaultExecutorNumbers   int64                       `json:"DefaultExecutorNumbers" xml:"DefaultExecutorNumbers"`
	MaxMemoryLimit           float64                     `json:"MaxMemoryLimit" xml:"MaxMemoryLimit"`
	AssumeRolePolicyDocument string                      `json:"AssumeRolePolicyDocument" xml:"AssumeRolePolicyDocument"`
	RoleName                 string                      `json:"RoleName" xml:"RoleName"`
	MaxCpuLimit              float64                     `json:"MaxCpuLimit" xml:"MaxCpuLimit"`
	Role                     string                      `json:"Role" xml:"Role"`
	CreatorId                string                      `json:"CreatorId" xml:"CreatorId"`
	MinCpu                   string                      `json:"MinCpu" xml:"MinCpu"`
	Description              string                      `json:"Description" xml:"Description"`
	DefaultExecutorSpecName  string                      `json:"DefaultExecutorSpecName" xml:"DefaultExecutorSpecName"`
	Type                     string                      `json:"Type" xml:"Type"`
	IsServiceLinkRole        bool                        `json:"IsServiceLinkRole" xml:"IsServiceLinkRole"`
	DefaultDriverSpecName    string                      `json:"DefaultDriverSpecName" xml:"DefaultDriverSpecName"`
	RoleId                   string                      `json:"RoleId" xml:"RoleId"`
	ShortName                string                      `json:"ShortName" xml:"ShortName"`
	Remark                   string                      `json:"Remark" xml:"Remark"`
	RamUid                   string                      `json:"RamUid" xml:"RamUid"`
	Version                  string                      `json:"Version" xml:"Version"`
	SparkVersionDescription  string                      `json:"SparkVersionDescription" xml:"SparkVersionDescription"`
	RolePrincipalName        string                      `json:"RolePrincipalName" xml:"RolePrincipalName"`
	MaxMemory                float64                     `json:"MaxMemory" xml:"MaxMemory"`
	CreateDate               string                      `json:"CreateDate" xml:"CreateDate"`
	Arn                      string                      `json:"Arn" xml:"Arn"`
	UserName                 string                      `json:"UserName" xml:"UserName"`
	Status                   string                      `json:"Status" xml:"Status"`
	SparkEngineModuleName    string                      `json:"SparkEngineModuleName" xml:"SparkEngineModuleName"`
	MaxCpu                   float64                     `json:"MaxCpu" xml:"MaxCpu"`
	SparkEngineConfig        SparkEngineConfig           `json:"SparkEngineConfig" xml:"SparkEngineConfig"`
	UserNetWorkConfigList    []UserNetWorkConfigListItem `json:"UserNetWorkConfigList" xml:"UserNetWorkConfigList"`
}