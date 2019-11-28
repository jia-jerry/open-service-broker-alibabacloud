package green

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

// UpdateOssCallbackSetting invokes the green.UpdateOssCallbackSetting API synchronously
// api document: https://help.aliyun.com/api/green/updateosscallbacksetting.html
func (client *Client) UpdateOssCallbackSetting(request *UpdateOssCallbackSettingRequest) (response *UpdateOssCallbackSettingResponse, err error) {
	response = CreateUpdateOssCallbackSettingResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateOssCallbackSettingWithChan invokes the green.UpdateOssCallbackSetting API asynchronously
// api document: https://help.aliyun.com/api/green/updateosscallbacksetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateOssCallbackSettingWithChan(request *UpdateOssCallbackSettingRequest) (<-chan *UpdateOssCallbackSettingResponse, <-chan error) {
	responseChan := make(chan *UpdateOssCallbackSettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateOssCallbackSetting(request)
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

// UpdateOssCallbackSettingWithCallback invokes the green.UpdateOssCallbackSetting API asynchronously
// api document: https://help.aliyun.com/api/green/updateosscallbacksetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateOssCallbackSettingWithCallback(request *UpdateOssCallbackSettingRequest, callback func(response *UpdateOssCallbackSettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateOssCallbackSettingResponse
		var err error
		defer close(result)
		response, err = client.UpdateOssCallbackSetting(request)
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

// UpdateOssCallbackSettingRequest is the request struct for api UpdateOssCallbackSetting
type UpdateOssCallbackSettingRequest struct {
	*requests.RpcRequest
	ScanCallbackSuggestions string           `position:"Query" name:"ScanCallbackSuggestions"`
	SourceIp                string           `position:"Query" name:"SourceIp"`
	CallbackSeed            string           `position:"Query" name:"CallbackSeed"`
	ServiceModules          string           `position:"Query" name:"ServiceModules"`
	AuditCallback           requests.Boolean `position:"Query" name:"AuditCallback"`
	ScanCallback            requests.Boolean `position:"Query" name:"ScanCallback"`
	CallbackUrl             string           `position:"Query" name:"CallbackUrl"`
}

// UpdateOssCallbackSettingResponse is the response struct for api UpdateOssCallbackSetting
type UpdateOssCallbackSettingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateOssCallbackSettingRequest creates a request to invoke UpdateOssCallbackSetting API
func CreateUpdateOssCallbackSettingRequest() (request *UpdateOssCallbackSettingRequest) {
	request = &UpdateOssCallbackSettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "UpdateOssCallbackSetting", "green", "openAPI")
	return
}

// CreateUpdateOssCallbackSettingResponse creates a response to parse from UpdateOssCallbackSetting response
func CreateUpdateOssCallbackSettingResponse() (response *UpdateOssCallbackSettingResponse) {
	response = &UpdateOssCallbackSettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
