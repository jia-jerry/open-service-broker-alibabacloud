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

// SetEdgeInstanceDriverConfigs invokes the iot.SetEdgeInstanceDriverConfigs API synchronously
// api document: https://help.aliyun.com/api/iot/setedgeinstancedriverconfigs.html
func (client *Client) SetEdgeInstanceDriverConfigs(request *SetEdgeInstanceDriverConfigsRequest) (response *SetEdgeInstanceDriverConfigsResponse, err error) {
	response = CreateSetEdgeInstanceDriverConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// SetEdgeInstanceDriverConfigsWithChan invokes the iot.SetEdgeInstanceDriverConfigs API asynchronously
// api document: https://help.aliyun.com/api/iot/setedgeinstancedriverconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetEdgeInstanceDriverConfigsWithChan(request *SetEdgeInstanceDriverConfigsRequest) (<-chan *SetEdgeInstanceDriverConfigsResponse, <-chan error) {
	responseChan := make(chan *SetEdgeInstanceDriverConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetEdgeInstanceDriverConfigs(request)
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

// SetEdgeInstanceDriverConfigsWithCallback invokes the iot.SetEdgeInstanceDriverConfigs API asynchronously
// api document: https://help.aliyun.com/api/iot/setedgeinstancedriverconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetEdgeInstanceDriverConfigsWithCallback(request *SetEdgeInstanceDriverConfigsRequest, callback func(response *SetEdgeInstanceDriverConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetEdgeInstanceDriverConfigsResponse
		var err error
		defer close(result)
		response, err = client.SetEdgeInstanceDriverConfigs(request)
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

// SetEdgeInstanceDriverConfigsRequest is the request struct for api SetEdgeInstanceDriverConfigs
type SetEdgeInstanceDriverConfigsRequest struct {
	*requests.RpcRequest
	Configs       *[]SetEdgeInstanceDriverConfigsConfigs `position:"Query" name:"Configs"  type:"Repeated"`
	InstanceId    string                                 `position:"Query" name:"InstanceId"`
	DriverId      string                                 `position:"Query" name:"DriverId"`
	IotInstanceId string                                 `position:"Query" name:"IotInstanceId"`
}

// SetEdgeInstanceDriverConfigsConfigs is a repeated param struct in SetEdgeInstanceDriverConfigsRequest
type SetEdgeInstanceDriverConfigsConfigs struct {
	Format  string `name:"Format"`
	Content string `name:"Content"`
	Key     string `name:"Key"`
}

// SetEdgeInstanceDriverConfigsResponse is the response struct for api SetEdgeInstanceDriverConfigs
type SetEdgeInstanceDriverConfigsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateSetEdgeInstanceDriverConfigsRequest creates a request to invoke SetEdgeInstanceDriverConfigs API
func CreateSetEdgeInstanceDriverConfigsRequest() (request *SetEdgeInstanceDriverConfigsRequest) {
	request = &SetEdgeInstanceDriverConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "SetEdgeInstanceDriverConfigs", "iot", "openAPI")
	return
}

// CreateSetEdgeInstanceDriverConfigsResponse creates a response to parse from SetEdgeInstanceDriverConfigs response
func CreateSetEdgeInstanceDriverConfigsResponse() (response *SetEdgeInstanceDriverConfigsResponse) {
	response = &SetEdgeInstanceDriverConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
