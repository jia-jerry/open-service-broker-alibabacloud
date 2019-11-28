package ehpc

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

// ListCpfsFileSystems invokes the ehpc.ListCpfsFileSystems API synchronously
// api document: https://help.aliyun.com/api/ehpc/listcpfsfilesystems.html
func (client *Client) ListCpfsFileSystems(request *ListCpfsFileSystemsRequest) (response *ListCpfsFileSystemsResponse, err error) {
	response = CreateListCpfsFileSystemsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCpfsFileSystemsWithChan invokes the ehpc.ListCpfsFileSystems API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listcpfsfilesystems.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCpfsFileSystemsWithChan(request *ListCpfsFileSystemsRequest) (<-chan *ListCpfsFileSystemsResponse, <-chan error) {
	responseChan := make(chan *ListCpfsFileSystemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCpfsFileSystems(request)
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

// ListCpfsFileSystemsWithCallback invokes the ehpc.ListCpfsFileSystems API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listcpfsfilesystems.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCpfsFileSystemsWithCallback(request *ListCpfsFileSystemsRequest, callback func(response *ListCpfsFileSystemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCpfsFileSystemsResponse
		var err error
		defer close(result)
		response, err = client.ListCpfsFileSystems(request)
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

// ListCpfsFileSystemsRequest is the request struct for api ListCpfsFileSystems
type ListCpfsFileSystemsRequest struct {
	*requests.RpcRequest
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	FileSystemId string           `position:"Query" name:"FileSystemId"`
}

// ListCpfsFileSystemsResponse is the response struct for api ListCpfsFileSystems
type ListCpfsFileSystemsResponse struct {
	*responses.BaseResponse
	RequestId      string                              `json:"RequestId" xml:"RequestId"`
	TotalCount     int                                 `json:"TotalCount" xml:"TotalCount"`
	PageNumber     int                                 `json:"PageNumber" xml:"PageNumber"`
	PageSize       int                                 `json:"PageSize" xml:"PageSize"`
	FileSystemList FileSystemListInListCpfsFileSystems `json:"FileSystemList" xml:"FileSystemList"`
}

// CreateListCpfsFileSystemsRequest creates a request to invoke ListCpfsFileSystems API
func CreateListCpfsFileSystemsRequest() (request *ListCpfsFileSystemsRequest) {
	request = &ListCpfsFileSystemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListCpfsFileSystems", "ehs", "openAPI")
	return
}

// CreateListCpfsFileSystemsResponse creates a response to parse from ListCpfsFileSystems response
func CreateListCpfsFileSystemsResponse() (response *ListCpfsFileSystemsResponse) {
	response = &ListCpfsFileSystemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
