package cr_ee

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

// DeleteChartRepository invokes the cr.DeleteChartRepository API synchronously
// api document: https://help.aliyun.com/api/cr/deletechartrepository.html
func (client *Client) DeleteChartRepository(request *DeleteChartRepositoryRequest) (response *DeleteChartRepositoryResponse, err error) {
	response = CreateDeleteChartRepositoryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteChartRepositoryWithChan invokes the cr.DeleteChartRepository API asynchronously
// api document: https://help.aliyun.com/api/cr/deletechartrepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteChartRepositoryWithChan(request *DeleteChartRepositoryRequest) (<-chan *DeleteChartRepositoryResponse, <-chan error) {
	responseChan := make(chan *DeleteChartRepositoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteChartRepository(request)
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

// DeleteChartRepositoryWithCallback invokes the cr.DeleteChartRepository API asynchronously
// api document: https://help.aliyun.com/api/cr/deletechartrepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteChartRepositoryWithCallback(request *DeleteChartRepositoryRequest, callback func(response *DeleteChartRepositoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteChartRepositoryResponse
		var err error
		defer close(result)
		response, err = client.DeleteChartRepository(request)
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

// DeleteChartRepositoryRequest is the request struct for api DeleteChartRepository
type DeleteChartRepositoryRequest struct {
	*requests.RpcRequest
	InstanceId        string `position:"Query" name:"InstanceId"`
	RepoNamespaceName string `position:"Query" name:"RepoNamespaceName"`
	RepoName          string `position:"Query" name:"RepoName"`
}

// DeleteChartRepositoryResponse is the response struct for api DeleteChartRepository
type DeleteChartRepositoryResponse struct {
	*responses.BaseResponse
	DeleteChartRepositoryIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                           string `json:"Code" xml:"Code"`
	RequestId                      string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteChartRepositoryRequest creates a request to invoke DeleteChartRepository API
func CreateDeleteChartRepositoryRequest() (request *DeleteChartRepositoryRequest) {
	request = &DeleteChartRepositoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "DeleteChartRepository", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteChartRepositoryResponse creates a response to parse from DeleteChartRepository response
func CreateDeleteChartRepositoryResponse() (response *DeleteChartRepositoryResponse) {
	response = &DeleteChartRepositoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
