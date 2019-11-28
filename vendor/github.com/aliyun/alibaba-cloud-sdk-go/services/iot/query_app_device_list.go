package iot

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

// QueryAppDeviceList invokes the iot.QueryAppDeviceList API synchronously
// api document: https://help.aliyun.com/api/iot/queryappdevicelist.html
func (client *Client) QueryAppDeviceList(request *QueryAppDeviceListRequest) (response *QueryAppDeviceListResponse, err error) {
	response = CreateQueryAppDeviceListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAppDeviceListWithChan invokes the iot.QueryAppDeviceList API asynchronously
// api document: https://help.aliyun.com/api/iot/queryappdevicelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAppDeviceListWithChan(request *QueryAppDeviceListRequest) (<-chan *QueryAppDeviceListResponse, <-chan error) {
	responseChan := make(chan *QueryAppDeviceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAppDeviceList(request)
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

// QueryAppDeviceListWithCallback invokes the iot.QueryAppDeviceList API asynchronously
// api document: https://help.aliyun.com/api/iot/queryappdevicelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAppDeviceListWithCallback(request *QueryAppDeviceListRequest, callback func(response *QueryAppDeviceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAppDeviceListResponse
		var err error
		defer close(result)
		response, err = client.QueryAppDeviceList(request)
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

// QueryAppDeviceListRequest is the request struct for api QueryAppDeviceList
type QueryAppDeviceListRequest struct {
	*requests.RpcRequest
	CurrentPage     requests.Integer             `position:"Query" name:"CurrentPage"`
	TagList         *[]QueryAppDeviceListTagList `position:"Query" name:"TagList"  type:"Repeated"`
	ProductKeyList  *[]string                    `position:"Query" name:"ProductKeyList"  type:"Repeated"`
	CategoryKeyList *[]string                    `position:"Query" name:"CategoryKeyList"  type:"Repeated"`
	IotInstanceId   string                       `position:"Query" name:"IotInstanceId"`
	PageSize        requests.Integer             `position:"Query" name:"PageSize"`
	AppKey          string                       `position:"Query" name:"AppKey"`
}

// QueryAppDeviceListTagList is a repeated param struct in QueryAppDeviceListRequest
type QueryAppDeviceListTagList struct {
	TagName  string `name:"TagName"`
	TagValue string `name:"TagValue"`
}

// QueryAppDeviceListResponse is the response struct for api QueryAppDeviceList
type QueryAppDeviceListResponse struct {
	*responses.BaseResponse
	RequestId    string                   `json:"RequestId" xml:"RequestId"`
	Success      bool                     `json:"Success" xml:"Success"`
	ErrorMessage string                   `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string                   `json:"Code" xml:"Code"`
	Page         int                      `json:"Page" xml:"Page"`
	PageSize     int                      `json:"PageSize" xml:"PageSize"`
	PageCount    int                      `json:"PageCount" xml:"PageCount"`
	Total        int                      `json:"Total" xml:"Total"`
	Data         DataInQueryAppDeviceList `json:"Data" xml:"Data"`
}

// CreateQueryAppDeviceListRequest creates a request to invoke QueryAppDeviceList API
func CreateQueryAppDeviceListRequest() (request *QueryAppDeviceListRequest) {
	request = &QueryAppDeviceListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryAppDeviceList", "iot", "openAPI")
	return
}

// CreateQueryAppDeviceListResponse creates a response to parse from QueryAppDeviceList response
func CreateQueryAppDeviceListResponse() (response *QueryAppDeviceListResponse) {
	response = &QueryAppDeviceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
