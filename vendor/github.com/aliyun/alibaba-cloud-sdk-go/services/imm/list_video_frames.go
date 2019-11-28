package imm

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

// ListVideoFrames invokes the imm.ListVideoFrames API synchronously
// api document: https://help.aliyun.com/api/imm/listvideoframes.html
func (client *Client) ListVideoFrames(request *ListVideoFramesRequest) (response *ListVideoFramesResponse, err error) {
	response = CreateListVideoFramesResponse()
	err = client.DoAction(request, response)
	return
}

// ListVideoFramesWithChan invokes the imm.ListVideoFrames API asynchronously
// api document: https://help.aliyun.com/api/imm/listvideoframes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVideoFramesWithChan(request *ListVideoFramesRequest) (<-chan *ListVideoFramesResponse, <-chan error) {
	responseChan := make(chan *ListVideoFramesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVideoFrames(request)
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

// ListVideoFramesWithCallback invokes the imm.ListVideoFrames API asynchronously
// api document: https://help.aliyun.com/api/imm/listvideoframes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVideoFramesWithCallback(request *ListVideoFramesRequest, callback func(response *ListVideoFramesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVideoFramesResponse
		var err error
		defer close(result)
		response, err = client.ListVideoFrames(request)
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

// ListVideoFramesRequest is the request struct for api ListVideoFrames
type ListVideoFramesRequest struct {
	*requests.RpcRequest
	Project  string `position:"Query" name:"Project"`
	VideoUri string `position:"Query" name:"VideoUri"`
	Marker   string `position:"Query" name:"Marker"`
	SetId    string `position:"Query" name:"SetId"`
}

// ListVideoFramesResponse is the response struct for api ListVideoFrames
type ListVideoFramesResponse struct {
	*responses.BaseResponse
	SetId      string       `json:"SetId" xml:"SetId"`
	VideoUri   string       `json:"VideoUri" xml:"VideoUri"`
	NextMarker string       `json:"NextMarker" xml:"NextMarker"`
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	Frames     []FramesItem `json:"Frames" xml:"Frames"`
}

// CreateListVideoFramesRequest creates a request to invoke ListVideoFrames API
func CreateListVideoFramesRequest() (request *ListVideoFramesRequest) {
	request = &ListVideoFramesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListVideoFrames", "imm", "openAPI")
	return
}

// CreateListVideoFramesResponse creates a response to parse from ListVideoFrames response
func CreateListVideoFramesResponse() (response *ListVideoFramesResponse) {
	response = &ListVideoFramesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
