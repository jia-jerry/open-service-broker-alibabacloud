package dyvmsapi

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

// CancelCall invokes the dyvmsapi.CancelCall API synchronously
// api document: https://help.aliyun.com/api/dyvmsapi/cancelcall.html
func (client *Client) CancelCall(request *CancelCallRequest) (response *CancelCallResponse, err error) {
	response = CreateCancelCallResponse()
	err = client.DoAction(request, response)
	return
}

// CancelCallWithChan invokes the dyvmsapi.CancelCall API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/cancelcall.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelCallWithChan(request *CancelCallRequest) (<-chan *CancelCallResponse, <-chan error) {
	responseChan := make(chan *CancelCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelCall(request)
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

// CancelCallWithCallback invokes the dyvmsapi.CancelCall API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/cancelcall.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelCallWithCallback(request *CancelCallRequest, callback func(response *CancelCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelCallResponse
		var err error
		defer close(result)
		response, err = client.CancelCall(request)
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

// CancelCallRequest is the request struct for api CancelCall
type CancelCallRequest struct {
	*requests.RpcRequest
	CallId               string           `position:"Query" name:"CallId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CancelCallResponse is the response struct for api CancelCall
type CancelCallResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    bool   `json:"Status" xml:"Status"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCancelCallRequest creates a request to invoke CancelCall API
func CreateCancelCallRequest() (request *CancelCallRequest) {
	request = &CancelCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "CancelCall", "", "")
	return
}

// CreateCancelCallResponse creates a response to parse from CancelCall response
func CreateCancelCallResponse() (response *CancelCallResponse) {
	response = &CancelCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
