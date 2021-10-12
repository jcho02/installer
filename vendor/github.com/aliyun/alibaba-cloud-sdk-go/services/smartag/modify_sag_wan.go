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

// ModifySagWan invokes the smartag.ModifySagWan API synchronously
func (client *Client) ModifySagWan(request *ModifySagWanRequest) (response *ModifySagWanResponse, err error) {
	response = CreateModifySagWanResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySagWanWithChan invokes the smartag.ModifySagWan API asynchronously
func (client *Client) ModifySagWanWithChan(request *ModifySagWanRequest) (<-chan *ModifySagWanResponse, <-chan error) {
	responseChan := make(chan *ModifySagWanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySagWan(request)
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

// ModifySagWanWithCallback invokes the smartag.ModifySagWan API asynchronously
func (client *Client) ModifySagWanWithCallback(request *ModifySagWanRequest, callback func(response *ModifySagWanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySagWanResponse
		var err error
		defer close(result)
		response, err = client.ModifySagWan(request)
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

// ModifySagWanRequest is the request struct for api ModifySagWan
type ModifySagWanRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query"`
	ISP                  string           `position:"Query"`
	Password             string           `position:"Query"`
	Vlan                 string           `position:"Query"`
	Mask                 string           `position:"Query"`
	StartIp              string           `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	Bandwidth            requests.Integer `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	IP                   string           `position:"Query"`
	Weight               requests.Integer `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
	IPType               string           `position:"Query"`
	Priority             requests.Integer `position:"Query"`
	SourceIps            string           `position:"Query"`
	SmartAGId            string           `position:"Query"`
	SmartAGSn            string           `position:"Query"`
	PortName             string           `position:"Query"`
	StopIp               string           `position:"Query"`
	Gateway              string           `position:"Query"`
	Username             string           `position:"Query"`
}

// ModifySagWanResponse is the response struct for api ModifySagWan
type ModifySagWanResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySagWanRequest creates a request to invoke ModifySagWan API
func CreateModifySagWanRequest() (request *ModifySagWanRequest) {
	request = &ModifySagWanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifySagWan", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySagWanResponse creates a response to parse from ModifySagWan response
func CreateModifySagWanResponse() (response *ModifySagWanResponse) {
	response = &ModifySagWanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
