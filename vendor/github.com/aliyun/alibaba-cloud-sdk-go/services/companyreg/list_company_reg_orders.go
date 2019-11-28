package companyreg

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

// ListCompanyRegOrders invokes the companyreg.ListCompanyRegOrders API synchronously
// api document: https://help.aliyun.com/api/companyreg/listcompanyregorders.html
func (client *Client) ListCompanyRegOrders(request *ListCompanyRegOrdersRequest) (response *ListCompanyRegOrdersResponse, err error) {
	response = CreateListCompanyRegOrdersResponse()
	err = client.DoAction(request, response)
	return
}

// ListCompanyRegOrdersWithChan invokes the companyreg.ListCompanyRegOrders API asynchronously
// api document: https://help.aliyun.com/api/companyreg/listcompanyregorders.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCompanyRegOrdersWithChan(request *ListCompanyRegOrdersRequest) (<-chan *ListCompanyRegOrdersResponse, <-chan error) {
	responseChan := make(chan *ListCompanyRegOrdersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCompanyRegOrders(request)
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

// ListCompanyRegOrdersWithCallback invokes the companyreg.ListCompanyRegOrders API asynchronously
// api document: https://help.aliyun.com/api/companyreg/listcompanyregorders.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCompanyRegOrdersWithCallback(request *ListCompanyRegOrdersRequest, callback func(response *ListCompanyRegOrdersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCompanyRegOrdersResponse
		var err error
		defer close(result)
		response, err = client.ListCompanyRegOrders(request)
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

// ListCompanyRegOrdersRequest is the request struct for api ListCompanyRegOrders
type ListCompanyRegOrdersRequest struct {
	*requests.RpcRequest
	PageNum       requests.Integer `position:"Query" name:"PageNum"`
	BizCode       string           `position:"Query" name:"BizCode"`
	BizStatus     string           `position:"Query" name:"BizStatus"`
	CompanyName   string           `position:"Query" name:"CompanyName"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	AliyunOrderId string           `position:"Query" name:"AliyunOrderId"`
	BizSubCode    string           `position:"Query" name:"BizSubCode"`
}

// ListCompanyRegOrdersResponse is the response struct for api ListCompanyRegOrders
type ListCompanyRegOrdersResponse struct {
	*responses.BaseResponse
	RequestId      string                     `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int                        `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum int                        `json:"CurrentPageNum" xml:"CurrentPageNum"`
	PageSize       int                        `json:"PageSize" xml:"PageSize"`
	TotalPageNum   int                        `json:"TotalPageNum" xml:"TotalPageNum"`
	PrePage        bool                       `json:"PrePage" xml:"PrePage"`
	NextPage       bool                       `json:"NextPage" xml:"NextPage"`
	Data           DataInListCompanyRegOrders `json:"Data" xml:"Data"`
}

// CreateListCompanyRegOrdersRequest creates a request to invoke ListCompanyRegOrders API
func CreateListCompanyRegOrdersRequest() (request *ListCompanyRegOrdersRequest) {
	request = &ListCompanyRegOrdersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "ListCompanyRegOrders", "companyreg", "openAPI")
	return
}

// CreateListCompanyRegOrdersResponse creates a response to parse from ListCompanyRegOrders response
func CreateListCompanyRegOrdersResponse() (response *ListCompanyRegOrdersResponse) {
	response = &ListCompanyRegOrdersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
