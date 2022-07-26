package dds

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

// DescribeDBInstanceSSL invokes the dds.DescribeDBInstanceSSL API synchronously
func (client *Client) DescribeDBInstanceSSL(request *DescribeDBInstanceSSLRequest) (response *DescribeDBInstanceSSLResponse, err error) {
	response = CreateDescribeDBInstanceSSLResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceSSLWithChan invokes the dds.DescribeDBInstanceSSL API asynchronously
func (client *Client) DescribeDBInstanceSSLWithChan(request *DescribeDBInstanceSSLRequest) (<-chan *DescribeDBInstanceSSLResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceSSLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceSSL(request)
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

// DescribeDBInstanceSSLWithCallback invokes the dds.DescribeDBInstanceSSL API asynchronously
func (client *Client) DescribeDBInstanceSSLWithCallback(request *DescribeDBInstanceSSLRequest, callback func(response *DescribeDBInstanceSSLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceSSLResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceSSL(request)
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

// DescribeDBInstanceSSLRequest is the request struct for api DescribeDBInstanceSSL
type DescribeDBInstanceSSLRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBInstanceSSLResponse is the response struct for api DescribeDBInstanceSSL
type DescribeDBInstanceSSLResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	SSLExpiredTime string `json:"SSLExpiredTime" xml:"SSLExpiredTime"`
	CertCommonName string `json:"CertCommonName" xml:"CertCommonName"`
	SSLStatus      string `json:"SSLStatus" xml:"SSLStatus"`
}

// CreateDescribeDBInstanceSSLRequest creates a request to invoke DescribeDBInstanceSSL API
func CreateDescribeDBInstanceSSLRequest() (request *DescribeDBInstanceSSLRequest) {
	request = &DescribeDBInstanceSSLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeDBInstanceSSL", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstanceSSLResponse creates a response to parse from DescribeDBInstanceSSL response
func CreateDescribeDBInstanceSSLResponse() (response *DescribeDBInstanceSSLResponse) {
	response = &DescribeDBInstanceSSLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
