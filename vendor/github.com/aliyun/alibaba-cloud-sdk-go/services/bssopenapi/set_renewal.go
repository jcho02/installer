package bssopenapi

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

// SetRenewal invokes the bssopenapi.SetRenewal API synchronously
func (client *Client) SetRenewal(request *SetRenewalRequest) (response *SetRenewalResponse, err error) {
	response = CreateSetRenewalResponse()
	err = client.DoAction(request, response)
	return
}

// SetRenewalWithChan invokes the bssopenapi.SetRenewal API asynchronously
func (client *Client) SetRenewalWithChan(request *SetRenewalRequest) (<-chan *SetRenewalResponse, <-chan error) {
	responseChan := make(chan *SetRenewalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetRenewal(request)
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

// SetRenewalWithCallback invokes the bssopenapi.SetRenewal API asynchronously
func (client *Client) SetRenewalWithCallback(request *SetRenewalRequest, callback func(response *SetRenewalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetRenewalResponse
		var err error
		defer close(result)
		response, err = client.SetRenewal(request)
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

// SetRenewalRequest is the request struct for api SetRenewal
type SetRenewalRequest struct {
	*requests.RpcRequest
	ProductCode       string           `position:"Query" name:"ProductCode"`
	SubscriptionType  string           `position:"Query" name:"SubscriptionType"`
	RenewalPeriod     requests.Integer `position:"Query" name:"RenewalPeriod"`
	OwnerId           requests.Integer `position:"Query" name:"OwnerId"`
	ProductType       string           `position:"Query" name:"ProductType"`
	InstanceIDs       string           `position:"Query" name:"InstanceIDs"`
	RenewalStatus     string           `position:"Query" name:"RenewalStatus"`
	RenewalPeriodUnit string           `position:"Query" name:"RenewalPeriodUnit"`
}

// SetRenewalResponse is the response struct for api SetRenewal
type SetRenewalResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateSetRenewalRequest creates a request to invoke SetRenewal API
func CreateSetRenewalRequest() (request *SetRenewalRequest) {
	request = &SetRenewalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "SetRenewal", "", "")
	request.Method = requests.POST
	return
}

// CreateSetRenewalResponse creates a response to parse from SetRenewal response
func CreateSetRenewalResponse() (response *SetRenewalResponse) {
	response = &SetRenewalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
