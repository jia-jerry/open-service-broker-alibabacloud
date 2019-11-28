package cloudapi

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

// DescribeHistoryApis invokes the cloudapi.DescribeHistoryApis API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describehistoryapis.html
func (client *Client) DescribeHistoryApis(request *DescribeHistoryApisRequest) (response *DescribeHistoryApisResponse, err error) {
	response = CreateDescribeHistoryApisResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHistoryApisWithChan invokes the cloudapi.DescribeHistoryApis API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describehistoryapis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHistoryApisWithChan(request *DescribeHistoryApisRequest) (<-chan *DescribeHistoryApisResponse, <-chan error) {
	responseChan := make(chan *DescribeHistoryApisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHistoryApis(request)
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

// DescribeHistoryApisWithCallback invokes the cloudapi.DescribeHistoryApis API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describehistoryapis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHistoryApisWithCallback(request *DescribeHistoryApisRequest, callback func(response *DescribeHistoryApisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHistoryApisResponse
		var err error
		defer close(result)
		response, err = client.DescribeHistoryApis(request)
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

// DescribeHistoryApisRequest is the request struct for api DescribeHistoryApis
type DescribeHistoryApisRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	GroupId       string `position:"Query" name:"GroupId"`
	PageNumber    string `position:"Query" name:"PageNumber"`
	ApiName       string `position:"Query" name:"ApiName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PageSize      string `position:"Query" name:"PageSize"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// DescribeHistoryApisResponse is the response struct for api DescribeHistoryApis
type DescribeHistoryApisResponse struct {
	*responses.BaseResponse
	RequestId   string                           `json:"RequestId" xml:"RequestId"`
	TotalCount  int                              `json:"TotalCount" xml:"TotalCount"`
	PageSize    int                              `json:"PageSize" xml:"PageSize"`
	PageNumber  int                              `json:"PageNumber" xml:"PageNumber"`
	ApiHisItems ApiHisItemsInDescribeHistoryApis `json:"ApiHisItems" xml:"ApiHisItems"`
}

// CreateDescribeHistoryApisRequest creates a request to invoke DescribeHistoryApis API
func CreateDescribeHistoryApisRequest() (request *DescribeHistoryApisRequest) {
	request = &DescribeHistoryApisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeHistoryApis", "apigateway", "openAPI")
	return
}

// CreateDescribeHistoryApisResponse creates a response to parse from DescribeHistoryApis response
func CreateDescribeHistoryApisResponse() (response *DescribeHistoryApisResponse) {
	response = &DescribeHistoryApisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
