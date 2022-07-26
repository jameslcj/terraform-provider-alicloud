package drds

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

// DescribeBackupSets invokes the drds.DescribeBackupSets API synchronously
func (client *Client) DescribeBackupSets(request *DescribeBackupSetsRequest) (response *DescribeBackupSetsResponse, err error) {
	response = CreateDescribeBackupSetsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupSetsWithChan invokes the drds.DescribeBackupSets API asynchronously
func (client *Client) DescribeBackupSetsWithChan(request *DescribeBackupSetsRequest) (<-chan *DescribeBackupSetsResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupSetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupSets(request)
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

// DescribeBackupSetsWithCallback invokes the drds.DescribeBackupSets API asynchronously
func (client *Client) DescribeBackupSetsWithCallback(request *DescribeBackupSetsRequest, callback func(response *DescribeBackupSetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupSetsResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupSets(request)
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

// DescribeBackupSetsRequest is the request struct for api DescribeBackupSets
type DescribeBackupSetsRequest struct {
	*requests.RpcRequest
	EndTime        string `position:"Query" name:"EndTime"`
	StartTime      string `position:"Query" name:"StartTime"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// DescribeBackupSetsResponse is the response struct for api DescribeBackupSets
type DescribeBackupSetsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Success    bool       `json:"Success" xml:"Success"`
	BackupSets BackupSets `json:"BackupSets" xml:"BackupSets"`
}

// CreateDescribeBackupSetsRequest creates a request to invoke DescribeBackupSets API
func CreateDescribeBackupSetsRequest() (request *DescribeBackupSetsRequest) {
	request = &DescribeBackupSetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeBackupSets", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBackupSetsResponse creates a response to parse from DescribeBackupSets response
func CreateDescribeBackupSetsResponse() (response *DescribeBackupSetsResponse) {
	response = &DescribeBackupSetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
