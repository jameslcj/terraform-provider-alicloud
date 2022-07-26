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

// QuerySavingsPlansDeductLog invokes the bssopenapi.QuerySavingsPlansDeductLog API synchronously
func (client *Client) QuerySavingsPlansDeductLog(request *QuerySavingsPlansDeductLogRequest) (response *QuerySavingsPlansDeductLogResponse, err error) {
	response = CreateQuerySavingsPlansDeductLogResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySavingsPlansDeductLogWithChan invokes the bssopenapi.QuerySavingsPlansDeductLog API asynchronously
func (client *Client) QuerySavingsPlansDeductLogWithChan(request *QuerySavingsPlansDeductLogRequest) (<-chan *QuerySavingsPlansDeductLogResponse, <-chan error) {
	responseChan := make(chan *QuerySavingsPlansDeductLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySavingsPlansDeductLog(request)
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

// QuerySavingsPlansDeductLogWithCallback invokes the bssopenapi.QuerySavingsPlansDeductLog API asynchronously
func (client *Client) QuerySavingsPlansDeductLogWithCallback(request *QuerySavingsPlansDeductLogRequest, callback func(response *QuerySavingsPlansDeductLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySavingsPlansDeductLogResponse
		var err error
		defer close(result)
		response, err = client.QuerySavingsPlansDeductLog(request)
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

// QuerySavingsPlansDeductLogRequest is the request struct for api QuerySavingsPlansDeductLog
type QuerySavingsPlansDeductLogRequest struct {
	*requests.RpcRequest
	EndTime      string           `position:"Query" name:"EndTime"`
	StartTime    string           `position:"Query" name:"StartTime"`
	Locale       string           `position:"Query" name:"Locale"`
	PageNum      requests.Integer `position:"Query" name:"PageNum"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	InstanceType string           `position:"Query" name:"InstanceType"`
}

// QuerySavingsPlansDeductLogResponse is the response struct for api QuerySavingsPlansDeductLog
type QuerySavingsPlansDeductLogResponse struct {
	*responses.BaseResponse
	Code      string                           `json:"Code" xml:"Code"`
	Message   string                           `json:"Message" xml:"Message"`
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	Success   bool                             `json:"Success" xml:"Success"`
	Data      DataInQuerySavingsPlansDeductLog `json:"Data" xml:"Data"`
}

// CreateQuerySavingsPlansDeductLogRequest creates a request to invoke QuerySavingsPlansDeductLog API
func CreateQuerySavingsPlansDeductLogRequest() (request *QuerySavingsPlansDeductLogRequest) {
	request = &QuerySavingsPlansDeductLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QuerySavingsPlansDeductLog", "", "")
	request.Method = requests.POST
	return
}

// CreateQuerySavingsPlansDeductLogResponse creates a response to parse from QuerySavingsPlansDeductLog response
func CreateQuerySavingsPlansDeductLogResponse() (response *QuerySavingsPlansDeductLogResponse) {
	response = &QuerySavingsPlansDeductLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
