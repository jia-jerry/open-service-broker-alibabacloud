package mts

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

// QueryVideoPoseJobList invokes the mts.QueryVideoPoseJobList API synchronously
// api document: https://help.aliyun.com/api/mts/queryvideoposejoblist.html
func (client *Client) QueryVideoPoseJobList(request *QueryVideoPoseJobListRequest) (response *QueryVideoPoseJobListResponse, err error) {
	response = CreateQueryVideoPoseJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryVideoPoseJobListWithChan invokes the mts.QueryVideoPoseJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/queryvideoposejoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryVideoPoseJobListWithChan(request *QueryVideoPoseJobListRequest) (<-chan *QueryVideoPoseJobListResponse, <-chan error) {
	responseChan := make(chan *QueryVideoPoseJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryVideoPoseJobList(request)
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

// QueryVideoPoseJobListWithCallback invokes the mts.QueryVideoPoseJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/queryvideoposejoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryVideoPoseJobListWithCallback(request *QueryVideoPoseJobListRequest, callback func(response *QueryVideoPoseJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryVideoPoseJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryVideoPoseJobList(request)
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

// QueryVideoPoseJobListRequest is the request struct for api QueryVideoPoseJobList
type QueryVideoPoseJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string           `position:"Query" name:"JobIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryVideoPoseJobListResponse is the response struct for api QueryVideoPoseJobList
type QueryVideoPoseJobListResponse struct {
	*responses.BaseResponse
	RequestId      string                                `json:"RequestId" xml:"RequestId"`
	NonExistJobIds NonExistJobIdsInQueryVideoPoseJobList `json:"NonExistJobIds" xml:"NonExistJobIds"`
	JobList        JobListInQueryVideoPoseJobList        `json:"JobList" xml:"JobList"`
}

// CreateQueryVideoPoseJobListRequest creates a request to invoke QueryVideoPoseJobList API
func CreateQueryVideoPoseJobListRequest() (request *QueryVideoPoseJobListRequest) {
	request = &QueryVideoPoseJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryVideoPoseJobList", "mts", "openAPI")
	return
}

// CreateQueryVideoPoseJobListResponse creates a response to parse from QueryVideoPoseJobList response
func CreateQueryVideoPoseJobListResponse() (response *QueryVideoPoseJobListResponse) {
	response = &QueryVideoPoseJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}