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

// UpdateIndexTemplate invokes the elasticsearch.UpdateIndexTemplate API synchronously
func (client *Client) UpdateIndexTemplate(request *UpdateIndexTemplateRequest) (response *UpdateIndexTemplateResponse, err error) {
	response = CreateUpdateIndexTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateIndexTemplateWithChan invokes the elasticsearch.UpdateIndexTemplate API asynchronously
func (client *Client) UpdateIndexTemplateWithChan(request *UpdateIndexTemplateRequest) (<-chan *UpdateIndexTemplateResponse, <-chan error) {
	responseChan := make(chan *UpdateIndexTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateIndexTemplate(request)
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

// UpdateIndexTemplateWithCallback invokes the elasticsearch.UpdateIndexTemplate API asynchronously
func (client *Client) UpdateIndexTemplateWithCallback(request *UpdateIndexTemplateRequest, callback func(response *UpdateIndexTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateIndexTemplateResponse
		var err error
		defer close(result)
		response, err = client.UpdateIndexTemplate(request)
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

// UpdateIndexTemplateRequest is the request struct for api UpdateIndexTemplate
type UpdateIndexTemplateRequest struct {
	*requests.RoaRequest
	InstanceId    string `position:"Path" name:"InstanceId"`
	ClientToken   string `position:"Query" name:"ClientToken"`
	IndexTemplate string `position:"Path" name:"IndexTemplate"`
}

// UpdateIndexTemplateResponse is the response struct for api UpdateIndexTemplate
type UpdateIndexTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateUpdateIndexTemplateRequest creates a request to invoke UpdateIndexTemplate API
func CreateUpdateIndexTemplateRequest() (request *UpdateIndexTemplateRequest) {
	request = &UpdateIndexTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "UpdateIndexTemplate", "/openapi/instances/[InstanceId]/index-templates/[IndexTemplate]", "elasticsearch", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateIndexTemplateResponse creates a response to parse from UpdateIndexTemplate response
func CreateUpdateIndexTemplateResponse() (response *UpdateIndexTemplateResponse) {
	response = &UpdateIndexTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
