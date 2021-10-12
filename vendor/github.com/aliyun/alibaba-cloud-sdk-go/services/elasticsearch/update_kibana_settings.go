package elasticsearch

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

// UpdateKibanaSettings invokes the elasticsearch.UpdateKibanaSettings API synchronously
func (client *Client) UpdateKibanaSettings(request *UpdateKibanaSettingsRequest) (response *UpdateKibanaSettingsResponse, err error) {
	response = CreateUpdateKibanaSettingsResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateKibanaSettingsWithChan invokes the elasticsearch.UpdateKibanaSettings API asynchronously
func (client *Client) UpdateKibanaSettingsWithChan(request *UpdateKibanaSettingsRequest) (<-chan *UpdateKibanaSettingsResponse, <-chan error) {
	responseChan := make(chan *UpdateKibanaSettingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateKibanaSettings(request)
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

// UpdateKibanaSettingsWithCallback invokes the elasticsearch.UpdateKibanaSettings API asynchronously
func (client *Client) UpdateKibanaSettingsWithCallback(request *UpdateKibanaSettingsRequest, callback func(response *UpdateKibanaSettingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateKibanaSettingsResponse
		var err error
		defer close(result)
		response, err = client.UpdateKibanaSettings(request)
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

// UpdateKibanaSettingsRequest is the request struct for api UpdateKibanaSettings
type UpdateKibanaSettingsRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// UpdateKibanaSettingsResponse is the response struct for api UpdateKibanaSettings
type UpdateKibanaSettingsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateUpdateKibanaSettingsRequest creates a request to invoke UpdateKibanaSettings API
func CreateUpdateKibanaSettingsRequest() (request *UpdateKibanaSettingsRequest) {
	request = &UpdateKibanaSettingsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "UpdateKibanaSettings", "/openapi/instances/[InstanceId]/actions/update-kibana-settings", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateKibanaSettingsResponse creates a response to parse from UpdateKibanaSettings response
func CreateUpdateKibanaSettingsResponse() (response *UpdateKibanaSettingsResponse) {
	response = &UpdateKibanaSettingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
