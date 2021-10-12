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

// DescribeSagWan4G invokes the smartag.DescribeSagWan4G API synchronously
func (client *Client) DescribeSagWan4G(request *DescribeSagWan4GRequest) (response *DescribeSagWan4GResponse, err error) {
	response = CreateDescribeSagWan4GResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSagWan4GWithChan invokes the smartag.DescribeSagWan4G API asynchronously
func (client *Client) DescribeSagWan4GWithChan(request *DescribeSagWan4GRequest) (<-chan *DescribeSagWan4GResponse, <-chan error) {
	responseChan := make(chan *DescribeSagWan4GResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSagWan4G(request)
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

// DescribeSagWan4GWithCallback invokes the smartag.DescribeSagWan4G API asynchronously
func (client *Client) DescribeSagWan4GWithCallback(request *DescribeSagWan4GRequest, callback func(response *DescribeSagWan4GResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSagWan4GResponse
		var err error
		defer close(result)
		response, err = client.DescribeSagWan4G(request)
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

// DescribeSagWan4GRequest is the request struct for api DescribeSagWan4G
type DescribeSagWan4GRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
	SmartAGId            string           `position:"Query"`
	SmartAGSn            string           `position:"Query"`
}

// DescribeSagWan4GResponse is the response struct for api DescribeSagWan4G
type DescribeSagWan4GResponse struct {
	*responses.BaseResponse
	Status       string `json:"Status" xml:"Status"`
	TrafficState string `json:"TrafficState" xml:"TrafficState"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Priority     int    `json:"Priority" xml:"Priority"`
	Ip           string `json:"Ip" xml:"Ip"`
	Mac          string `json:"Mac" xml:"Mac"`
	Strength     string `json:"Strength" xml:"Strength"`
}

// CreateDescribeSagWan4GRequest creates a request to invoke DescribeSagWan4G API
func CreateDescribeSagWan4GRequest() (request *DescribeSagWan4GRequest) {
	request = &DescribeSagWan4GRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSagWan4G", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSagWan4GResponse creates a response to parse from DescribeSagWan4G response
func CreateDescribeSagWan4GResponse() (response *DescribeSagWan4GResponse) {
	response = &DescribeSagWan4GResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
