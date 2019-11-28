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

// DescribeParameterTemplates invokes the r_kvstore.DescribeParameterTemplates API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeparametertemplates.html
func (client *Client) DescribeParameterTemplates(request *DescribeParameterTemplatesRequest) (response *DescribeParameterTemplatesResponse, err error) {
	response = CreateDescribeParameterTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeParameterTemplatesWithChan invokes the r_kvstore.DescribeParameterTemplates API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeparametertemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeParameterTemplatesWithChan(request *DescribeParameterTemplatesRequest) (<-chan *DescribeParameterTemplatesResponse, <-chan error) {
	responseChan := make(chan *DescribeParameterTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeParameterTemplates(request)
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

// DescribeParameterTemplatesWithCallback invokes the r_kvstore.DescribeParameterTemplates API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeparametertemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeParameterTemplatesWithCallback(request *DescribeParameterTemplatesRequest, callback func(response *DescribeParameterTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeParameterTemplatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeParameterTemplates(request)
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

// DescribeParameterTemplatesRequest is the request struct for api DescribeParameterTemplates
type DescribeParameterTemplatesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	Engine               string           `position:"Query" name:"Engine"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	CharacterType        string           `position:"Query" name:"CharacterType"`
}

// DescribeParameterTemplatesResponse is the response struct for api DescribeParameterTemplates
type DescribeParameterTemplatesResponse struct {
	*responses.BaseResponse
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Engine         string     `json:"Engine" xml:"Engine"`
	EngineVersion  string     `json:"EngineVersion" xml:"EngineVersion"`
	ParameterCount string     `json:"ParameterCount" xml:"ParameterCount"`
	Parameters     Parameters `json:"Parameters" xml:"Parameters"`
}

// CreateDescribeParameterTemplatesRequest creates a request to invoke DescribeParameterTemplates API
func CreateDescribeParameterTemplatesRequest() (request *DescribeParameterTemplatesRequest) {
	request = &DescribeParameterTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeParameterTemplates", "", "")
	return
}

// CreateDescribeParameterTemplatesResponse creates a response to parse from DescribeParameterTemplates response
func CreateDescribeParameterTemplatesResponse() (response *DescribeParameterTemplatesResponse) {
	response = &DescribeParameterTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
