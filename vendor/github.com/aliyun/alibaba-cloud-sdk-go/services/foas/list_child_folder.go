package foas

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

// ListChildFolder invokes the foas.ListChildFolder API synchronously
// api document: https://help.aliyun.com/api/foas/listchildfolder.html
func (client *Client) ListChildFolder(request *ListChildFolderRequest) (response *ListChildFolderResponse, err error) {
	response = CreateListChildFolderResponse()
	err = client.DoAction(request, response)
	return
}

// ListChildFolderWithChan invokes the foas.ListChildFolder API asynchronously
// api document: https://help.aliyun.com/api/foas/listchildfolder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListChildFolderWithChan(request *ListChildFolderRequest) (<-chan *ListChildFolderResponse, <-chan error) {
	responseChan := make(chan *ListChildFolderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListChildFolder(request)
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

// ListChildFolderWithCallback invokes the foas.ListChildFolder API asynchronously
// api document: https://help.aliyun.com/api/foas/listchildfolder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListChildFolderWithCallback(request *ListChildFolderRequest, callback func(response *ListChildFolderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListChildFolderResponse
		var err error
		defer close(result)
		response, err = client.ListChildFolder(request)
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

// ListChildFolderRequest is the request struct for api ListChildFolder
type ListChildFolderRequest struct {
	*requests.RoaRequest
	Path        string `position:"Query" name:"path"`
	ProjectName string `position:"Path" name:"projectName"`
}

// ListChildFolderResponse is the response struct for api ListChildFolder
type ListChildFolderResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Folders   Folders `json:"Folders" xml:"Folders"`
}

// CreateListChildFolderRequest creates a request to invoke ListChildFolder API
func CreateListChildFolderRequest() (request *ListChildFolderRequest) {
	request = &ListChildFolderRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "ListChildFolder", "/api/v2/projects/[projectName]/folders/children", "foas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListChildFolderResponse creates a response to parse from ListChildFolder response
func CreateListChildFolderResponse() (response *ListChildFolderResponse) {
	response = &ListChildFolderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
