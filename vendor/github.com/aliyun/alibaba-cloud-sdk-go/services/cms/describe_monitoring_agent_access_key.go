package cms

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

// DescribeMonitoringAgentAccessKey invokes the cms.DescribeMonitoringAgentAccessKey API synchronously
// api document: https://help.aliyun.com/api/cms/describemonitoringagentaccesskey.html
func (client *Client) DescribeMonitoringAgentAccessKey(request *DescribeMonitoringAgentAccessKeyRequest) (response *DescribeMonitoringAgentAccessKeyResponse, err error) {
	response = CreateDescribeMonitoringAgentAccessKeyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMonitoringAgentAccessKeyWithChan invokes the cms.DescribeMonitoringAgentAccessKey API asynchronously
// api document: https://help.aliyun.com/api/cms/describemonitoringagentaccesskey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMonitoringAgentAccessKeyWithChan(request *DescribeMonitoringAgentAccessKeyRequest) (<-chan *DescribeMonitoringAgentAccessKeyResponse, <-chan error) {
	responseChan := make(chan *DescribeMonitoringAgentAccessKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMonitoringAgentAccessKey(request)
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

// DescribeMonitoringAgentAccessKeyWithCallback invokes the cms.DescribeMonitoringAgentAccessKey API asynchronously
// api document: https://help.aliyun.com/api/cms/describemonitoringagentaccesskey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMonitoringAgentAccessKeyWithCallback(request *DescribeMonitoringAgentAccessKeyRequest, callback func(response *DescribeMonitoringAgentAccessKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMonitoringAgentAccessKeyResponse
		var err error
		defer close(result)
		response, err = client.DescribeMonitoringAgentAccessKey(request)
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

// DescribeMonitoringAgentAccessKeyRequest is the request struct for api DescribeMonitoringAgentAccessKey
type DescribeMonitoringAgentAccessKeyRequest struct {
	*requests.RpcRequest
}

// DescribeMonitoringAgentAccessKeyResponse is the response struct for api DescribeMonitoringAgentAccessKey
type DescribeMonitoringAgentAccessKeyResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	AccessKey string `json:"AccessKey" xml:"AccessKey"`
	SecretKey string `json:"SecretKey" xml:"SecretKey"`
}

// CreateDescribeMonitoringAgentAccessKeyRequest creates a request to invoke DescribeMonitoringAgentAccessKey API
func CreateDescribeMonitoringAgentAccessKeyRequest() (request *DescribeMonitoringAgentAccessKeyRequest) {
	request = &DescribeMonitoringAgentAccessKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMonitoringAgentAccessKey", "cms", "openAPI")
	return
}

// CreateDescribeMonitoringAgentAccessKeyResponse creates a response to parse from DescribeMonitoringAgentAccessKey response
func CreateDescribeMonitoringAgentAccessKeyResponse() (response *DescribeMonitoringAgentAccessKeyResponse) {
	response = &DescribeMonitoringAgentAccessKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
