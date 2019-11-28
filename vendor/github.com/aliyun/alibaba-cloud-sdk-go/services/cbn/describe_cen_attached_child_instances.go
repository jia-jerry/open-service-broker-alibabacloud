package cbn

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

// DescribeCenAttachedChildInstances invokes the cbn.DescribeCenAttachedChildInstances API synchronously
// api document: https://help.aliyun.com/api/cbn/describecenattachedchildinstances.html
func (client *Client) DescribeCenAttachedChildInstances(request *DescribeCenAttachedChildInstancesRequest) (response *DescribeCenAttachedChildInstancesResponse, err error) {
	response = CreateDescribeCenAttachedChildInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCenAttachedChildInstancesWithChan invokes the cbn.DescribeCenAttachedChildInstances API asynchronously
// api document: https://help.aliyun.com/api/cbn/describecenattachedchildinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCenAttachedChildInstancesWithChan(request *DescribeCenAttachedChildInstancesRequest) (<-chan *DescribeCenAttachedChildInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeCenAttachedChildInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCenAttachedChildInstances(request)
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

// DescribeCenAttachedChildInstancesWithCallback invokes the cbn.DescribeCenAttachedChildInstances API asynchronously
// api document: https://help.aliyun.com/api/cbn/describecenattachedchildinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCenAttachedChildInstancesWithCallback(request *DescribeCenAttachedChildInstancesRequest, callback func(response *DescribeCenAttachedChildInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCenAttachedChildInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCenAttachedChildInstances(request)
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

// DescribeCenAttachedChildInstancesRequest is the request struct for api DescribeCenAttachedChildInstances
type DescribeCenAttachedChildInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	CenId                 string           `position:"Query" name:"CenId"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	PageSize              requests.Integer `position:"Query" name:"PageSize"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	ChildInstanceType     string           `position:"Query" name:"ChildInstanceType"`
	PageNumber            requests.Integer `position:"Query" name:"PageNumber"`
	ChildInstanceRegionId string           `position:"Query" name:"ChildInstanceRegionId"`
}

// DescribeCenAttachedChildInstancesResponse is the response struct for api DescribeCenAttachedChildInstances
type DescribeCenAttachedChildInstancesResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	PageNumber     int            `json:"PageNumber" xml:"PageNumber"`
	PageSize       int            `json:"PageSize" xml:"PageSize"`
	ChildInstances ChildInstances `json:"ChildInstances" xml:"ChildInstances"`
}

// CreateDescribeCenAttachedChildInstancesRequest creates a request to invoke DescribeCenAttachedChildInstances API
func CreateDescribeCenAttachedChildInstancesRequest() (request *DescribeCenAttachedChildInstancesRequest) {
	request = &DescribeCenAttachedChildInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCenAttachedChildInstances", "cbn", "openAPI")
	return
}

// CreateDescribeCenAttachedChildInstancesResponse creates a response to parse from DescribeCenAttachedChildInstances response
func CreateDescribeCenAttachedChildInstancesResponse() (response *DescribeCenAttachedChildInstancesResponse) {
	response = &DescribeCenAttachedChildInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
