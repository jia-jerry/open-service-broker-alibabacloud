package r_kvstore

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

// ModifyGuardDomainMode invokes the r_kvstore.ModifyGuardDomainMode API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyguarddomainmode.html
func (client *Client) ModifyGuardDomainMode(request *ModifyGuardDomainModeRequest) (response *ModifyGuardDomainModeResponse, err error) {
	response = CreateModifyGuardDomainModeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyGuardDomainModeWithChan invokes the r_kvstore.ModifyGuardDomainMode API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyguarddomainmode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyGuardDomainModeWithChan(request *ModifyGuardDomainModeRequest) (<-chan *ModifyGuardDomainModeResponse, <-chan error) {
	responseChan := make(chan *ModifyGuardDomainModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyGuardDomainMode(request)
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

// ModifyGuardDomainModeWithCallback invokes the r_kvstore.ModifyGuardDomainMode API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyguarddomainmode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyGuardDomainModeWithCallback(request *ModifyGuardDomainModeRequest, callback func(response *ModifyGuardDomainModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyGuardDomainModeResponse
		var err error
		defer close(result)
		response, err = client.ModifyGuardDomainMode(request)
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

// ModifyGuardDomainModeRequest is the request struct for api ModifyGuardDomainMode
type ModifyGuardDomainModeRequest struct {
	*requests.RpcRequest
	DomainMode           string           `position:"Query" name:"DomainMode"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyGuardDomainModeResponse is the response struct for api ModifyGuardDomainMode
type ModifyGuardDomainModeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyGuardDomainModeRequest creates a request to invoke ModifyGuardDomainMode API
func CreateModifyGuardDomainModeRequest() (request *ModifyGuardDomainModeRequest) {
	request = &ModifyGuardDomainModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyGuardDomainMode", "", "")
	return
}

// CreateModifyGuardDomainModeResponse creates a response to parse from ModifyGuardDomainMode response
func CreateModifyGuardDomainModeResponse() (response *ModifyGuardDomainModeResponse) {
	response = &ModifyGuardDomainModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}