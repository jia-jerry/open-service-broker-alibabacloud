package trademark

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

// QuerySupplementDetail invokes the trademark.QuerySupplementDetail API synchronously
// api document: https://help.aliyun.com/api/trademark/querysupplementdetail.html
func (client *Client) QuerySupplementDetail(request *QuerySupplementDetailRequest) (response *QuerySupplementDetailResponse, err error) {
	response = CreateQuerySupplementDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySupplementDetailWithChan invokes the trademark.QuerySupplementDetail API asynchronously
// api document: https://help.aliyun.com/api/trademark/querysupplementdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySupplementDetailWithChan(request *QuerySupplementDetailRequest) (<-chan *QuerySupplementDetailResponse, <-chan error) {
	responseChan := make(chan *QuerySupplementDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySupplementDetail(request)
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

// QuerySupplementDetailWithCallback invokes the trademark.QuerySupplementDetail API asynchronously
// api document: https://help.aliyun.com/api/trademark/querysupplementdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySupplementDetailWithCallback(request *QuerySupplementDetailRequest, callback func(response *QuerySupplementDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySupplementDetailResponse
		var err error
		defer close(result)
		response, err = client.QuerySupplementDetail(request)
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

// QuerySupplementDetailRequest is the request struct for api QuerySupplementDetail
type QuerySupplementDetailRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// QuerySupplementDetailResponse is the response struct for api QuerySupplementDetail
type QuerySupplementDetailResponse struct {
	*responses.BaseResponse
	RequestId             string                                  `json:"RequestId" xml:"RequestId"`
	Id                    int64                                   `json:"Id" xml:"Id"`
	SerialNumber          string                                  `json:"SerialNumber" xml:"SerialNumber"`
	Type                  int                                     `json:"Type" xml:"Type"`
	Status                int                                     `json:"Status" xml:"Status"`
	TmNumber              string                                  `json:"TmNumber" xml:"TmNumber"`
	SendTime              int64                                   `json:"SendTime" xml:"SendTime"`
	AcceptTime            int64                                   `json:"AcceptTime" xml:"AcceptTime"`
	SbjDeadTime           int64                                   `json:"SbjDeadTime" xml:"SbjDeadTime"`
	AcceptDeadTime        int64                                   `json:"AcceptDeadTime" xml:"AcceptDeadTime"`
	OperateTime           int64                                   `json:"OperateTime" xml:"OperateTime"`
	UploadFileTemplateUrl string                                  `json:"UploadFileTemplateUrl" xml:"UploadFileTemplateUrl"`
	Content               string                                  `json:"Content" xml:"Content"`
	FileTemplateUrls      FileTemplateUrlsInQuerySupplementDetail `json:"FileTemplateUrls" xml:"FileTemplateUrls"`
}

// CreateQuerySupplementDetailRequest creates a request to invoke QuerySupplementDetail API
func CreateQuerySupplementDetailRequest() (request *QuerySupplementDetailRequest) {
	request = &QuerySupplementDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "QuerySupplementDetail", "trademark", "openAPI")
	return
}

// CreateQuerySupplementDetailResponse creates a response to parse from QuerySupplementDetail response
func CreateQuerySupplementDetailResponse() (response *QuerySupplementDetailResponse) {
	response = &QuerySupplementDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
