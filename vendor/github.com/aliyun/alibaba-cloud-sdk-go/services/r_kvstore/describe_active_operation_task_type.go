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

// DescribeActiveOperationTaskType invokes the r_kvstore.DescribeActiveOperationTaskType API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeactiveoperationtasktype.html
func (client *Client) DescribeActiveOperationTaskType(request *DescribeActiveOperationTaskTypeRequest) (response *DescribeActiveOperationTaskTypeResponse, err error) {
	response = CreateDescribeActiveOperationTaskTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeActiveOperationTaskTypeWithChan invokes the r_kvstore.DescribeActiveOperationTaskType API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeactiveoperationtasktype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeActiveOperationTaskTypeWithChan(request *DescribeActiveOperationTaskTypeRequest) (<-chan *DescribeActiveOperationTaskTypeResponse, <-chan error) {
	responseChan := make(chan *DescribeActiveOperationTaskTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeActiveOperationTaskType(request)
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

// DescribeActiveOperationTaskTypeWithCallback invokes the r_kvstore.DescribeActiveOperationTaskType API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeactiveoperationtasktype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeActiveOperationTaskTypeWithCallback(request *DescribeActiveOperationTaskTypeRequest, callback func(response *DescribeActiveOperationTaskTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeActiveOperationTaskTypeResponse
		var err error
		defer close(result)
		response, err = client.DescribeActiveOperationTaskType(request)
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

// DescribeActiveOperationTaskTypeRequest is the request struct for api DescribeActiveOperationTaskType
type DescribeActiveOperationTaskTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	IsHistory            requests.Integer `position:"Query" name:"IsHistory"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeActiveOperationTaskTypeResponse is the response struct for api DescribeActiveOperationTaskType
type DescribeActiveOperationTaskTypeResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	TypeList  []Items `json:"TypeList" xml:"TypeList"`
}

// CreateDescribeActiveOperationTaskTypeRequest creates a request to invoke DescribeActiveOperationTaskType API
func CreateDescribeActiveOperationTaskTypeRequest() (request *DescribeActiveOperationTaskTypeRequest) {
	request = &DescribeActiveOperationTaskTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeActiveOperationTaskType", "", "")
	return
}

// CreateDescribeActiveOperationTaskTypeResponse creates a response to parse from DescribeActiveOperationTaskType response
func CreateDescribeActiveOperationTaskTypeResponse() (response *DescribeActiveOperationTaskTypeResponse) {
	response = &DescribeActiveOperationTaskTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
