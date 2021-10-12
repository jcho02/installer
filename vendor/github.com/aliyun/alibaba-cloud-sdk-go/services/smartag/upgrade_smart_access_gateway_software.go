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

// UpgradeSmartAccessGatewaySoftware invokes the smartag.UpgradeSmartAccessGatewaySoftware API synchronously
func (client *Client) UpgradeSmartAccessGatewaySoftware(request *UpgradeSmartAccessGatewaySoftwareRequest) (response *UpgradeSmartAccessGatewaySoftwareResponse, err error) {
	response = CreateUpgradeSmartAccessGatewaySoftwareResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeSmartAccessGatewaySoftwareWithChan invokes the smartag.UpgradeSmartAccessGatewaySoftware API asynchronously
func (client *Client) UpgradeSmartAccessGatewaySoftwareWithChan(request *UpgradeSmartAccessGatewaySoftwareRequest) (<-chan *UpgradeSmartAccessGatewaySoftwareResponse, <-chan error) {
	responseChan := make(chan *UpgradeSmartAccessGatewaySoftwareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeSmartAccessGatewaySoftware(request)
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

// UpgradeSmartAccessGatewaySoftwareWithCallback invokes the smartag.UpgradeSmartAccessGatewaySoftware API asynchronously
func (client *Client) UpgradeSmartAccessGatewaySoftwareWithCallback(request *UpgradeSmartAccessGatewaySoftwareRequest, callback func(response *UpgradeSmartAccessGatewaySoftwareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeSmartAccessGatewaySoftwareResponse
		var err error
		defer close(result)
		response, err = client.UpgradeSmartAccessGatewaySoftware(request)
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

// UpgradeSmartAccessGatewaySoftwareRequest is the request struct for api UpgradeSmartAccessGatewaySoftware
type UpgradeSmartAccessGatewaySoftwareRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query"`
	UserCount            requests.Integer `position:"Query"`
	AutoPay              requests.Boolean `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
	SmartAGId            string           `position:"Query"`
	DataPlan             requests.Integer `position:"Query"`
}

// UpgradeSmartAccessGatewaySoftwareResponse is the response struct for api UpgradeSmartAccessGatewaySoftware
type UpgradeSmartAccessGatewaySoftwareResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateUpgradeSmartAccessGatewaySoftwareRequest creates a request to invoke UpgradeSmartAccessGatewaySoftware API
func CreateUpgradeSmartAccessGatewaySoftwareRequest() (request *UpgradeSmartAccessGatewaySoftwareRequest) {
	request = &UpgradeSmartAccessGatewaySoftwareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "UpgradeSmartAccessGatewaySoftware", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpgradeSmartAccessGatewaySoftwareResponse creates a response to parse from UpgradeSmartAccessGatewaySoftware response
func CreateUpgradeSmartAccessGatewaySoftwareResponse() (response *UpgradeSmartAccessGatewaySoftwareResponse) {
	response = &UpgradeSmartAccessGatewaySoftwareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
