package cs

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

// DescribeClusterLogs invokes the cs.DescribeClusterLogs API synchronously
func (client *Client) DescribeClusterLogs(request *DescribeClusterLogsRequest) (response *DescribeClusterLogsResponse, err error) {
	response = CreateDescribeClusterLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterLogsWithChan invokes the cs.DescribeClusterLogs API asynchronously
func (client *Client) DescribeClusterLogsWithChan(request *DescribeClusterLogsRequest) (<-chan *DescribeClusterLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterLogs(request)
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

// DescribeClusterLogsWithCallback invokes the cs.DescribeClusterLogs API asynchronously
func (client *Client) DescribeClusterLogsWithCallback(request *DescribeClusterLogsRequest, callback func(response *DescribeClusterLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterLogs(request)
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

// DescribeClusterLogsRequest is the request struct for api DescribeClusterLogs
type DescribeClusterLogsRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

// DescribeClusterLogsResponse is the response struct for api DescribeClusterLogs
type DescribeClusterLogsResponse struct {
	*responses.BaseResponse
}

// CreateDescribeClusterLogsRequest creates a request to invoke DescribeClusterLogs API
func CreateDescribeClusterLogsRequest() (request *DescribeClusterLogsRequest) {
	request = &DescribeClusterLogsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterLogs", "/clusters/[ClusterId]/logs", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeClusterLogsResponse creates a response to parse from DescribeClusterLogs response
func CreateDescribeClusterLogsResponse() (response *DescribeClusterLogsResponse) {
	response = &DescribeClusterLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
