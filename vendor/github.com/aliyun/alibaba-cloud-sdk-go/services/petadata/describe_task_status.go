package petadata

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

// DescribeTaskStatus invokes the petadata.DescribeTaskStatus API synchronously
// api document: https://help.aliyun.com/api/petadata/describetaskstatus.html
func (client *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
	response = CreateDescribeTaskStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTaskStatusWithChan invokes the petadata.DescribeTaskStatus API asynchronously
// api document: https://help.aliyun.com/api/petadata/describetaskstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTaskStatusWithChan(request *DescribeTaskStatusRequest) (<-chan *DescribeTaskStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeTaskStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTaskStatus(request)
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

// DescribeTaskStatusWithCallback invokes the petadata.DescribeTaskStatus API asynchronously
// api document: https://help.aliyun.com/api/petadata/describetaskstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTaskStatusWithCallback(request *DescribeTaskStatusRequest, callback func(response *DescribeTaskStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTaskStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeTaskStatus(request)
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

// DescribeTaskStatusRequest is the request struct for api DescribeTaskStatus
type DescribeTaskStatusRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TaskId               string           `position:"Query" name:"TaskId"`
}

// DescribeTaskStatusResponse is the response struct for api DescribeTaskStatus
type DescribeTaskStatusResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TaskId     string `json:"TaskId" xml:"TaskId"`
	Message    string `json:"Message" xml:"Message"`
	BeginTime  string `json:"BeginTime" xml:"BeginTime"`
	FinishTime string `json:"FinishTime" xml:"FinishTime"`
	Status     string `json:"Status" xml:"Status"`
}

// CreateDescribeTaskStatusRequest creates a request to invoke DescribeTaskStatus API
func CreateDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
	request = &DescribeTaskStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PetaData", "2016-01-01", "DescribeTaskStatus", "petadata", "openAPI")
	return
}

// CreateDescribeTaskStatusResponse creates a response to parse from DescribeTaskStatus response
func CreateDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
	response = &DescribeTaskStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
