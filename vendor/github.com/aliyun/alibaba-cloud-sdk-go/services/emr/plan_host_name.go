package emr

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

// PlanHostName invokes the emr.PlanHostName API synchronously
// api document: https://help.aliyun.com/api/emr/planhostname.html
func (client *Client) PlanHostName(request *PlanHostNameRequest) (response *PlanHostNameResponse, err error) {
	response = CreatePlanHostNameResponse()
	err = client.DoAction(request, response)
	return
}

// PlanHostNameWithChan invokes the emr.PlanHostName API asynchronously
// api document: https://help.aliyun.com/api/emr/planhostname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PlanHostNameWithChan(request *PlanHostNameRequest) (<-chan *PlanHostNameResponse, <-chan error) {
	responseChan := make(chan *PlanHostNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PlanHostName(request)
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

// PlanHostNameWithCallback invokes the emr.PlanHostName API asynchronously
// api document: https://help.aliyun.com/api/emr/planhostname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PlanHostNameWithCallback(request *PlanHostNameRequest, callback func(response *PlanHostNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PlanHostNameResponse
		var err error
		defer close(result)
		response, err = client.PlanHostName(request)
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

// PlanHostNameRequest is the request struct for api PlanHostName
type PlanHostNameRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer         `position:"Query" name:"ResourceOwnerId"`
	HostGroup       *[]PlanHostNameHostGroup `position:"Query" name:"HostGroup"  type:"Repeated"`
	HostInfo        *[]PlanHostNameHostInfo  `position:"Query" name:"HostInfo"  type:"Repeated"`
	ClusterId       string                   `position:"Query" name:"ClusterId"`
}

// PlanHostNameHostGroup is a repeated param struct in PlanHostNameRequest
type PlanHostNameHostGroup struct {
	GroupType string `name:"GroupType"`
	GroupName string `name:"GroupName"`
}

// PlanHostNameHostInfo is a repeated param struct in PlanHostNameRequest
type PlanHostNameHostInfo struct {
	HpHostBizId   string `name:"HpHostBizId"`
	HostGroupName string `name:"HostGroupName"`
}

// PlanHostNameResponse is the response struct for api PlanHostName
type PlanHostNameResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	HostInfoList HostInfoList `json:"HostInfoList" xml:"HostInfoList"`
}

// CreatePlanHostNameRequest creates a request to invoke PlanHostName API
func CreatePlanHostNameRequest() (request *PlanHostNameRequest) {
	request = &PlanHostNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "PlanHostName", "emr", "openAPI")
	return
}

// CreatePlanHostNameResponse creates a response to parse from PlanHostName response
func CreatePlanHostNameResponse() (response *PlanHostNameResponse) {
	response = &PlanHostNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
