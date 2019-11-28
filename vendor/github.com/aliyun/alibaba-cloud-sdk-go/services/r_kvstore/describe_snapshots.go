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

// DescribeSnapshots invokes the r_kvstore.DescribeSnapshots API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describesnapshots.html
func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
	response = CreateDescribeSnapshotsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSnapshotsWithChan invokes the r_kvstore.DescribeSnapshots API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describesnapshots.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSnapshotsWithChan(request *DescribeSnapshotsRequest) (<-chan *DescribeSnapshotsResponse, <-chan error) {
	responseChan := make(chan *DescribeSnapshotsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSnapshots(request)
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

// DescribeSnapshotsWithCallback invokes the r_kvstore.DescribeSnapshots API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describesnapshots.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSnapshotsWithCallback(request *DescribeSnapshotsRequest, callback func(response *DescribeSnapshotsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSnapshotsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSnapshots(request)
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

// DescribeSnapshotsRequest is the request struct for api DescribeSnapshots
type DescribeSnapshotsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	BeginTime            string           `position:"Query" name:"BeginTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSnapshotsResponse is the response struct for api DescribeSnapshots
type DescribeSnapshotsResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Snapshots Snapshots `json:"Snapshots" xml:"Snapshots"`
}

// CreateDescribeSnapshotsRequest creates a request to invoke DescribeSnapshots API
func CreateDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
	request = &DescribeSnapshotsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeSnapshots", "", "")
	return
}

// CreateDescribeSnapshotsResponse creates a response to parse from DescribeSnapshots response
func CreateDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
	response = &DescribeSnapshotsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}