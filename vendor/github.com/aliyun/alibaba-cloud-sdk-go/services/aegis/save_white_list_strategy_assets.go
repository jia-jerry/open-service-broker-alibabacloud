package aegis

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

// SaveWhiteListStrategyAssets invokes the aegis.SaveWhiteListStrategyAssets API synchronously
// api document: https://help.aliyun.com/api/aegis/savewhiteliststrategyassets.html
func (client *Client) SaveWhiteListStrategyAssets(request *SaveWhiteListStrategyAssetsRequest) (response *SaveWhiteListStrategyAssetsResponse, err error) {
	response = CreateSaveWhiteListStrategyAssetsResponse()
	err = client.DoAction(request, response)
	return
}

// SaveWhiteListStrategyAssetsWithChan invokes the aegis.SaveWhiteListStrategyAssets API asynchronously
// api document: https://help.aliyun.com/api/aegis/savewhiteliststrategyassets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveWhiteListStrategyAssetsWithChan(request *SaveWhiteListStrategyAssetsRequest) (<-chan *SaveWhiteListStrategyAssetsResponse, <-chan error) {
	responseChan := make(chan *SaveWhiteListStrategyAssetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveWhiteListStrategyAssets(request)
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

// SaveWhiteListStrategyAssetsWithCallback invokes the aegis.SaveWhiteListStrategyAssets API asynchronously
// api document: https://help.aliyun.com/api/aegis/savewhiteliststrategyassets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveWhiteListStrategyAssetsWithCallback(request *SaveWhiteListStrategyAssetsRequest, callback func(response *SaveWhiteListStrategyAssetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveWhiteListStrategyAssetsResponse
		var err error
		defer close(result)
		response, err = client.SaveWhiteListStrategyAssets(request)
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

// SaveWhiteListStrategyAssetsRequest is the request struct for api SaveWhiteListStrategyAssets
type SaveWhiteListStrategyAssetsRequest struct {
	*requests.RpcRequest
	Operations   string           `position:"Query" name:"Operations"`
	RelationType requests.Integer `position:"Query" name:"RelationType"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	StrategyId   requests.Integer `position:"Query" name:"StrategyId"`
	Lang         string           `position:"Query" name:"Lang"`
}

// SaveWhiteListStrategyAssetsResponse is the response struct for api SaveWhiteListStrategyAssets
type SaveWhiteListStrategyAssetsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSaveWhiteListStrategyAssetsRequest creates a request to invoke SaveWhiteListStrategyAssets API
func CreateSaveWhiteListStrategyAssetsRequest() (request *SaveWhiteListStrategyAssetsRequest) {
	request = &SaveWhiteListStrategyAssetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "SaveWhiteListStrategyAssets", "vipaegis", "openAPI")
	return
}

// CreateSaveWhiteListStrategyAssetsResponse creates a response to parse from SaveWhiteListStrategyAssets response
func CreateSaveWhiteListStrategyAssetsResponse() (response *SaveWhiteListStrategyAssetsResponse) {
	response = &SaveWhiteListStrategyAssetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
