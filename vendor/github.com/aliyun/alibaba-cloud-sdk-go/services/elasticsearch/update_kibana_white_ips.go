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

// UpdateKibanaWhiteIps invokes the elasticsearch.UpdateKibanaWhiteIps API synchronously
func (client *Client) UpdateKibanaWhiteIps(request *UpdateKibanaWhiteIpsRequest) (response *UpdateKibanaWhiteIpsResponse, err error) {
	response = CreateUpdateKibanaWhiteIpsResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateKibanaWhiteIpsWithChan invokes the elasticsearch.UpdateKibanaWhiteIps API asynchronously
func (client *Client) UpdateKibanaWhiteIpsWithChan(request *UpdateKibanaWhiteIpsRequest) (<-chan *UpdateKibanaWhiteIpsResponse, <-chan error) {
	responseChan := make(chan *UpdateKibanaWhiteIpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateKibanaWhiteIps(request)
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

// UpdateKibanaWhiteIpsWithCallback invokes the elasticsearch.UpdateKibanaWhiteIps API asynchronously
func (client *Client) UpdateKibanaWhiteIpsWithCallback(request *UpdateKibanaWhiteIpsRequest, callback func(response *UpdateKibanaWhiteIpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateKibanaWhiteIpsResponse
		var err error
		defer close(result)
		response, err = client.UpdateKibanaWhiteIps(request)
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

// UpdateKibanaWhiteIpsRequest is the request struct for api UpdateKibanaWhiteIps
type UpdateKibanaWhiteIpsRequest struct {
	*requests.RoaRequest
	ModifyMode  string `position:"Query" name:"modifyMode"`
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// UpdateKibanaWhiteIpsResponse is the response struct for api UpdateKibanaWhiteIps
type UpdateKibanaWhiteIpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUpdateKibanaWhiteIpsRequest creates a request to invoke UpdateKibanaWhiteIps API
func CreateUpdateKibanaWhiteIpsRequest() (request *UpdateKibanaWhiteIpsRequest) {
	request = &UpdateKibanaWhiteIpsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "UpdateKibanaWhiteIps", "/openapi/instances/[InstanceId]/kibana-white-ips", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateKibanaWhiteIpsResponse creates a response to parse from UpdateKibanaWhiteIps response
func CreateUpdateKibanaWhiteIpsResponse() (response *UpdateKibanaWhiteIpsResponse) {
	response = &UpdateKibanaWhiteIpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
