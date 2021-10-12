package ddoscoo

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

// DescribePortConnsList invokes the ddoscoo.DescribePortConnsList API synchronously
func (client *Client) DescribePortConnsList(request *DescribePortConnsListRequest) (response *DescribePortConnsListResponse, err error) {
	response = CreateDescribePortConnsListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePortConnsListWithChan invokes the ddoscoo.DescribePortConnsList API asynchronously
func (client *Client) DescribePortConnsListWithChan(request *DescribePortConnsListRequest) (<-chan *DescribePortConnsListResponse, <-chan error) {
	responseChan := make(chan *DescribePortConnsListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePortConnsList(request)
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

// DescribePortConnsListWithCallback invokes the ddoscoo.DescribePortConnsList API asynchronously
func (client *Client) DescribePortConnsListWithCallback(request *DescribePortConnsListRequest, callback func(response *DescribePortConnsListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePortConnsListResponse
		var err error
		defer close(result)
		response, err = client.DescribePortConnsList(request)
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

// DescribePortConnsListRequest is the request struct for api DescribePortConnsList
type DescribePortConnsListRequest struct {
	*requests.RpcRequest
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	InstanceIds     *[]string        `position:"Query" name:"InstanceIds"  type:"Repeated"`
	Port            string           `position:"Query" name:"Port"`
	Interval        requests.Integer `position:"Query" name:"Interval"`
}

// DescribePortConnsListResponse is the response struct for api DescribePortConnsList
type DescribePortConnsListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ConnsList []Conn `json:"ConnsList" xml:"ConnsList"`
}

// CreateDescribePortConnsListRequest creates a request to invoke DescribePortConnsList API
func CreateDescribePortConnsListRequest() (request *DescribePortConnsListRequest) {
	request = &DescribePortConnsListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribePortConnsList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePortConnsListResponse creates a response to parse from DescribePortConnsList response
func CreateDescribePortConnsListResponse() (response *DescribePortConnsListResponse) {
	response = &DescribePortConnsListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
