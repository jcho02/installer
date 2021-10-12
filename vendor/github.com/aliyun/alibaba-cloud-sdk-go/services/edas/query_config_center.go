package edas

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

// QueryConfigCenter invokes the edas.QueryConfigCenter API synchronously
func (client *Client) QueryConfigCenter(request *QueryConfigCenterRequest) (response *QueryConfigCenterResponse, err error) {
	response = CreateQueryConfigCenterResponse()
	err = client.DoAction(request, response)
	return
}

// QueryConfigCenterWithChan invokes the edas.QueryConfigCenter API asynchronously
func (client *Client) QueryConfigCenterWithChan(request *QueryConfigCenterRequest) (<-chan *QueryConfigCenterResponse, <-chan error) {
	responseChan := make(chan *QueryConfigCenterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryConfigCenter(request)
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

// QueryConfigCenterWithCallback invokes the edas.QueryConfigCenter API asynchronously
func (client *Client) QueryConfigCenterWithCallback(request *QueryConfigCenterRequest, callback func(response *QueryConfigCenterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryConfigCenterResponse
		var err error
		defer close(result)
		response, err = client.QueryConfigCenter(request)
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

// QueryConfigCenterRequest is the request struct for api QueryConfigCenter
type QueryConfigCenterRequest struct {
	*requests.RoaRequest
	DataId          string `position:"Query" name:"DataId"`
	LogicalRegionId string `position:"Query" name:"LogicalRegionId"`
	Group           string `position:"Query" name:"Group"`
}

// QueryConfigCenterResponse is the response struct for api QueryConfigCenter
type QueryConfigCenterResponse struct {
	*responses.BaseResponse
	Code             int              `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	ConfigCenterInfo ConfigCenterInfo `json:"configCenterInfo" xml:"configCenterInfo"`
}

// CreateQueryConfigCenterRequest creates a request to invoke QueryConfigCenter API
func CreateQueryConfigCenterRequest() (request *QueryConfigCenterRequest) {
	request = &QueryConfigCenterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "QueryConfigCenter", "/pop/v5/configCenter", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryConfigCenterResponse creates a response to parse from QueryConfigCenter response
func CreateQueryConfigCenterResponse() (response *QueryConfigCenterResponse) {
	response = &QueryConfigCenterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
