package smartag

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

// DeleteNetworkOptimization invokes the smartag.DeleteNetworkOptimization API synchronously
// api document: https://help.aliyun.com/api/smartag/deletenetworkoptimization.html
func (client *Client) DeleteNetworkOptimization(request *DeleteNetworkOptimizationRequest) (response *DeleteNetworkOptimizationResponse, err error) {
	response = CreateDeleteNetworkOptimizationResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNetworkOptimizationWithChan invokes the smartag.DeleteNetworkOptimization API asynchronously
// api document: https://help.aliyun.com/api/smartag/deletenetworkoptimization.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteNetworkOptimizationWithChan(request *DeleteNetworkOptimizationRequest) (<-chan *DeleteNetworkOptimizationResponse, <-chan error) {
	responseChan := make(chan *DeleteNetworkOptimizationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNetworkOptimization(request)
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

// DeleteNetworkOptimizationWithCallback invokes the smartag.DeleteNetworkOptimization API asynchronously
// api document: https://help.aliyun.com/api/smartag/deletenetworkoptimization.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteNetworkOptimizationWithCallback(request *DeleteNetworkOptimizationRequest, callback func(response *DeleteNetworkOptimizationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNetworkOptimizationResponse
		var err error
		defer close(result)
		response, err = client.DeleteNetworkOptimization(request)
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

// DeleteNetworkOptimizationRequest is the request struct for api DeleteNetworkOptimization
type DeleteNetworkOptimizationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NetworkOptId         string           `position:"Query" name:"NetworkOptId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteNetworkOptimizationResponse is the response struct for api DeleteNetworkOptimization
type DeleteNetworkOptimizationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteNetworkOptimizationRequest creates a request to invoke DeleteNetworkOptimization API
func CreateDeleteNetworkOptimizationRequest() (request *DeleteNetworkOptimizationRequest) {
	request = &DeleteNetworkOptimizationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DeleteNetworkOptimization", "smartag", "openAPI")
	return
}

// CreateDeleteNetworkOptimizationResponse creates a response to parse from DeleteNetworkOptimization response
func CreateDeleteNetworkOptimizationResponse() (response *DeleteNetworkOptimizationResponse) {
	response = &DeleteNetworkOptimizationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
