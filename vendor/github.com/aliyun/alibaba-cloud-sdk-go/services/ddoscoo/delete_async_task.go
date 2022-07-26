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

// DeleteAsyncTask invokes the ddoscoo.DeleteAsyncTask API synchronously
func (client *Client) DeleteAsyncTask(request *DeleteAsyncTaskRequest) (response *DeleteAsyncTaskResponse, err error) {
	response = CreateDeleteAsyncTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAsyncTaskWithChan invokes the ddoscoo.DeleteAsyncTask API asynchronously
func (client *Client) DeleteAsyncTaskWithChan(request *DeleteAsyncTaskRequest) (<-chan *DeleteAsyncTaskResponse, <-chan error) {
	responseChan := make(chan *DeleteAsyncTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAsyncTask(request)
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

// DeleteAsyncTaskWithCallback invokes the ddoscoo.DeleteAsyncTask API asynchronously
func (client *Client) DeleteAsyncTaskWithCallback(request *DeleteAsyncTaskRequest, callback func(response *DeleteAsyncTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAsyncTaskResponse
		var err error
		defer close(result)
		response, err = client.DeleteAsyncTask(request)
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

// DeleteAsyncTaskRequest is the request struct for api DeleteAsyncTask
type DeleteAsyncTaskRequest struct {
	*requests.RpcRequest
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Lang            string           `position:"Query" name:"Lang"`
	TaskId          requests.Integer `position:"Query" name:"TaskId"`
}

// DeleteAsyncTaskResponse is the response struct for api DeleteAsyncTask
type DeleteAsyncTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteAsyncTaskRequest creates a request to invoke DeleteAsyncTask API
func CreateDeleteAsyncTaskRequest() (request *DeleteAsyncTaskRequest) {
	request = &DeleteAsyncTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "DeleteAsyncTask", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteAsyncTaskResponse creates a response to parse from DeleteAsyncTask response
func CreateDeleteAsyncTaskResponse() (response *DeleteAsyncTaskResponse) {
	response = &DeleteAsyncTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
