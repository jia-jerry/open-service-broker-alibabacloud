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

// DeleteFaceSearchImageById invokes the imm.DeleteFaceSearchImageById API synchronously
// api document: https://help.aliyun.com/api/imm/deletefacesearchimagebyid.html
func (client *Client) DeleteFaceSearchImageById(request *DeleteFaceSearchImageByIdRequest) (response *DeleteFaceSearchImageByIdResponse, err error) {
	response = CreateDeleteFaceSearchImageByIdResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFaceSearchImageByIdWithChan invokes the imm.DeleteFaceSearchImageById API asynchronously
// api document: https://help.aliyun.com/api/imm/deletefacesearchimagebyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFaceSearchImageByIdWithChan(request *DeleteFaceSearchImageByIdRequest) (<-chan *DeleteFaceSearchImageByIdResponse, <-chan error) {
	responseChan := make(chan *DeleteFaceSearchImageByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFaceSearchImageById(request)
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

// DeleteFaceSearchImageByIdWithCallback invokes the imm.DeleteFaceSearchImageById API asynchronously
// api document: https://help.aliyun.com/api/imm/deletefacesearchimagebyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFaceSearchImageByIdWithCallback(request *DeleteFaceSearchImageByIdRequest, callback func(response *DeleteFaceSearchImageByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFaceSearchImageByIdResponse
		var err error
		defer close(result)
		response, err = client.DeleteFaceSearchImageById(request)
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

// DeleteFaceSearchImageByIdRequest is the request struct for api DeleteFaceSearchImageById
type DeleteFaceSearchImageByIdRequest struct {
	*requests.RpcRequest
	ImageId   string `position:"Query" name:"ImageId"`
	Project   string `position:"Query" name:"Project"`
	GroupName string `position:"Query" name:"GroupName"`
	User      string `position:"Query" name:"User"`
}

// DeleteFaceSearchImageByIdResponse is the response struct for api DeleteFaceSearchImageById
type DeleteFaceSearchImageByIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteFaceSearchImageByIdRequest creates a request to invoke DeleteFaceSearchImageById API
func CreateDeleteFaceSearchImageByIdRequest() (request *DeleteFaceSearchImageByIdRequest) {
	request = &DeleteFaceSearchImageByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DeleteFaceSearchImageById", "imm", "openAPI")
	return
}

// CreateDeleteFaceSearchImageByIdResponse creates a response to parse from DeleteFaceSearchImageById response
func CreateDeleteFaceSearchImageByIdResponse() (response *DeleteFaceSearchImageByIdResponse) {
	response = &DeleteFaceSearchImageByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
