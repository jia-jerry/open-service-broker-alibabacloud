package scdn

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

// DescribeScdnDomainTopReferVisit invokes the scdn.DescribeScdnDomainTopReferVisit API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomaintoprefervisit.html
func (client *Client) DescribeScdnDomainTopReferVisit(request *DescribeScdnDomainTopReferVisitRequest) (response *DescribeScdnDomainTopReferVisitResponse, err error) {
	response = CreateDescribeScdnDomainTopReferVisitResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainTopReferVisitWithChan invokes the scdn.DescribeScdnDomainTopReferVisit API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomaintoprefervisit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainTopReferVisitWithChan(request *DescribeScdnDomainTopReferVisitRequest) (<-chan *DescribeScdnDomainTopReferVisitResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainTopReferVisitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainTopReferVisit(request)
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

// DescribeScdnDomainTopReferVisitWithCallback invokes the scdn.DescribeScdnDomainTopReferVisit API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomaintoprefervisit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainTopReferVisitWithCallback(request *DescribeScdnDomainTopReferVisitRequest, callback func(response *DescribeScdnDomainTopReferVisitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainTopReferVisitResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainTopReferVisit(request)
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

// DescribeScdnDomainTopReferVisitRequest is the request struct for api DescribeScdnDomainTopReferVisit
type DescribeScdnDomainTopReferVisitRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	SortBy        string           `position:"Query" name:"SortBy"`
}

// DescribeScdnDomainTopReferVisitResponse is the response struct for api DescribeScdnDomainTopReferVisit
type DescribeScdnDomainTopReferVisitResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	DomainName   string       `json:"DomainName" xml:"DomainName"`
	StartTime    string       `json:"StartTime" xml:"StartTime"`
	TopReferList TopReferList `json:"TopReferList" xml:"TopReferList"`
}

// CreateDescribeScdnDomainTopReferVisitRequest creates a request to invoke DescribeScdnDomainTopReferVisit API
func CreateDescribeScdnDomainTopReferVisitRequest() (request *DescribeScdnDomainTopReferVisitRequest) {
	request = &DescribeScdnDomainTopReferVisitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainTopReferVisit", "", "")
	return
}

// CreateDescribeScdnDomainTopReferVisitResponse creates a response to parse from DescribeScdnDomainTopReferVisit response
func CreateDescribeScdnDomainTopReferVisitResponse() (response *DescribeScdnDomainTopReferVisitResponse) {
	response = &DescribeScdnDomainTopReferVisitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
