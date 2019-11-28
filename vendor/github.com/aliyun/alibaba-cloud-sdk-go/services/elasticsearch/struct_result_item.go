package elasticsearch

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

// ResultItem is a nested struct in elasticsearch response
type ResultItem struct {
	Content     string `json:"content" xml:"content"`
	Port        int    `json:"port" xml:"port"`
	Timestamp   int64  `json:"timestamp" xml:"timestamp"`
	Health      string `json:"health" xml:"health"`
	State       string `json:"state" xml:"state"`
	LoadFiveM   string `json:"loadFiveM" xml:"loadFiveM"`
	Host        string `json:"host" xml:"host"`
	NodeType    string `json:"nodeType" xml:"nodeType"`
	InstanceId  string `json:"instanceId" xml:"instanceId"`
	ZoneId      string `json:"zoneId" xml:"zoneId"`
	Description string `json:"description" xml:"description"`
	CpuPercent  string `json:"cpuPercent" xml:"cpuPercent"`
	HeapPercent string `json:"heapPercent" xml:"heapPercent"`
	Source      string `json:"source" xml:"source"`
	Level       string `json:"level" xml:"level"`
	Name        string `json:"name" xml:"name"`
}