package domain

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

// QueryTransferInByInstanceId invokes the domain.QueryTransferInByInstanceId API synchronously
// api document: https://help.aliyun.com/api/domain/querytransferinbyinstanceid.html
func (client *Client) QueryTransferInByInstanceId(request *QueryTransferInByInstanceIdRequest) (response *QueryTransferInByInstanceIdResponse, err error) {
	response = CreateQueryTransferInByInstanceIdResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTransferInByInstanceIdWithChan invokes the domain.QueryTransferInByInstanceId API asynchronously
// api document: https://help.aliyun.com/api/domain/querytransferinbyinstanceid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTransferInByInstanceIdWithChan(request *QueryTransferInByInstanceIdRequest) (<-chan *QueryTransferInByInstanceIdResponse, <-chan error) {
	responseChan := make(chan *QueryTransferInByInstanceIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTransferInByInstanceId(request)
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

// QueryTransferInByInstanceIdWithCallback invokes the domain.QueryTransferInByInstanceId API asynchronously
// api document: https://help.aliyun.com/api/domain/querytransferinbyinstanceid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTransferInByInstanceIdWithCallback(request *QueryTransferInByInstanceIdRequest, callback func(response *QueryTransferInByInstanceIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTransferInByInstanceIdResponse
		var err error
		defer close(result)
		response, err = client.QueryTransferInByInstanceId(request)
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

// QueryTransferInByInstanceIdRequest is the request struct for api QueryTransferInByInstanceId
type QueryTransferInByInstanceIdRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// QueryTransferInByInstanceIdResponse is the response struct for api QueryTransferInByInstanceId
type QueryTransferInByInstanceIdResponse struct {
	*responses.BaseResponse
	RequestId                                   string `json:"RequestId" xml:"RequestId"`
	SubmissionDate                              string `json:"SubmissionDate" xml:"SubmissionDate"`
	ModificationDate                            string `json:"ModificationDate" xml:"ModificationDate"`
	UserId                                      string `json:"UserId" xml:"UserId"`
	InstanceId                                  string `json:"InstanceId" xml:"InstanceId"`
	DomainName                                  string `json:"DomainName" xml:"DomainName"`
	Status                                      int    `json:"Status" xml:"Status"`
	SimpleTransferInStatus                      string `json:"SimpleTransferInStatus" xml:"SimpleTransferInStatus"`
	ResultCode                                  string `json:"ResultCode" xml:"ResultCode"`
	ResultDate                                  string `json:"ResultDate" xml:"ResultDate"`
	ResultMsg                                   string `json:"ResultMsg" xml:"ResultMsg"`
	TransferAuthorizationCodeSubmissionDate     string `json:"TransferAuthorizationCodeSubmissionDate" xml:"TransferAuthorizationCodeSubmissionDate"`
	NeedMailCheck                               bool   `json:"NeedMailCheck" xml:"NeedMailCheck"`
	Email                                       string `json:"Email" xml:"Email"`
	WhoisMailStatus                             bool   `json:"WhoisMailStatus" xml:"WhoisMailStatus"`
	ExpirationDate                              string `json:"ExpirationDate" xml:"ExpirationDate"`
	ProgressBarType                             int    `json:"ProgressBarType" xml:"ProgressBarType"`
	SubmissionDateLong                          int64  `json:"SubmissionDateLong" xml:"SubmissionDateLong"`
	ModificationDateLong                        int64  `json:"ModificationDateLong" xml:"ModificationDateLong"`
	ResultDateLong                              int64  `json:"ResultDateLong" xml:"ResultDateLong"`
	ExpirationDateLong                          int64  `json:"ExpirationDateLong" xml:"ExpirationDateLong"`
	TransferAuthorizationCodeSubmissionDateLong int64  `json:"TransferAuthorizationCodeSubmissionDateLong" xml:"TransferAuthorizationCodeSubmissionDateLong"`
}

// CreateQueryTransferInByInstanceIdRequest creates a request to invoke QueryTransferInByInstanceId API
func CreateQueryTransferInByInstanceIdRequest() (request *QueryTransferInByInstanceIdRequest) {
	request = &QueryTransferInByInstanceIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryTransferInByInstanceId", "", "")
	return
}

// CreateQueryTransferInByInstanceIdResponse creates a response to parse from QueryTransferInByInstanceId response
func CreateQueryTransferInByInstanceIdResponse() (response *QueryTransferInByInstanceIdResponse) {
	response = &QueryTransferInByInstanceIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
