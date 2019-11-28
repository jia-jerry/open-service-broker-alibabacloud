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

// CreateAlertUserGroup invokes the emr.CreateAlertUserGroup API synchronously
// api document: https://help.aliyun.com/api/emr/createalertusergroup.html
func (client *Client) CreateAlertUserGroup(request *CreateAlertUserGroupRequest) (response *CreateAlertUserGroupResponse, err error) {
	response = CreateCreateAlertUserGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAlertUserGroupWithChan invokes the emr.CreateAlertUserGroup API asynchronously
// api document: https://help.aliyun.com/api/emr/createalertusergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAlertUserGroupWithChan(request *CreateAlertUserGroupRequest) (<-chan *CreateAlertUserGroupResponse, <-chan error) {
	responseChan := make(chan *CreateAlertUserGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAlertUserGroup(request)
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

// CreateAlertUserGroupWithCallback invokes the emr.CreateAlertUserGroup API asynchronously
// api document: https://help.aliyun.com/api/emr/createalertusergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAlertUserGroupWithCallback(request *CreateAlertUserGroupRequest, callback func(response *CreateAlertUserGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAlertUserGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateAlertUserGroup(request)
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

// CreateAlertUserGroupRequest is the request struct for api CreateAlertUserGroup
type CreateAlertUserGroupRequest struct {
	*requests.RpcRequest
	UserList        string           `position:"Query" name:"UserList"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Name            string           `position:"Query" name:"Name"`
	Description     string           `position:"Query" name:"Description"`
}

// CreateAlertUserGroupResponse is the response struct for api CreateAlertUserGroup
type CreateAlertUserGroupResponse struct {
	*responses.BaseResponse
	Id int64 `json:"Id" xml:"Id"`
}

// CreateCreateAlertUserGroupRequest creates a request to invoke CreateAlertUserGroup API
func CreateCreateAlertUserGroupRequest() (request *CreateAlertUserGroupRequest) {
	request = &CreateAlertUserGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateAlertUserGroup", "emr", "openAPI")
	return
}

// CreateCreateAlertUserGroupResponse creates a response to parse from CreateAlertUserGroup response
func CreateCreateAlertUserGroupResponse() (response *CreateAlertUserGroupResponse) {
	response = &CreateAlertUserGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
