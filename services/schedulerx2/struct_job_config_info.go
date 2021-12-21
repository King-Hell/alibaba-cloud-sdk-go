package schedulerx2

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

// JobConfigInfo is a nested struct in schedulerx2 response
type JobConfigInfo struct {
	Status          int            `json:"Status" xml:"Status"`
	Parameters      string         `json:"Parameters" xml:"Parameters"`
	Description     string         `json:"Description" xml:"Description"`
	ExecuteMode     string         `json:"ExecuteMode" xml:"ExecuteMode"`
	MaxConcurrency  string         `json:"MaxConcurrency" xml:"MaxConcurrency"`
	Name            string         `json:"Name" xml:"Name"`
	MaxAttempt      int            `json:"MaxAttempt" xml:"MaxAttempt"`
	Content         string         `json:"Content" xml:"Content"`
	JarUrl          string         `json:"JarUrl" xml:"JarUrl"`
	ClassName       string         `json:"ClassName" xml:"ClassName"`
	AttemptInterval int            `json:"AttemptInterval" xml:"AttemptInterval"`
	MapTaskXAttrs   MapTaskXAttrs  `json:"MapTaskXAttrs" xml:"MapTaskXAttrs"`
	TimeConfig      TimeConfig     `json:"TimeConfig" xml:"TimeConfig"`
	JobMonitorInfo  JobMonitorInfo `json:"JobMonitorInfo" xml:"JobMonitorInfo"`
}
