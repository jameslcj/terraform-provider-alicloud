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

// DescribeKubernetesVersionMetadata invokes the cs.DescribeKubernetesVersionMetadata API synchronously
func (client *Client) DescribeKubernetesVersionMetadata(request *DescribeKubernetesVersionMetadataRequest) (response *DescribeKubernetesVersionMetadataResponse, err error) {
	response = CreateDescribeKubernetesVersionMetadataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeKubernetesVersionMetadataWithChan invokes the cs.DescribeKubernetesVersionMetadata API asynchronously
func (client *Client) DescribeKubernetesVersionMetadataWithChan(request *DescribeKubernetesVersionMetadataRequest) (<-chan *DescribeKubernetesVersionMetadataResponse, <-chan error) {
	responseChan := make(chan *DescribeKubernetesVersionMetadataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeKubernetesVersionMetadata(request)
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

// DescribeKubernetesVersionMetadataWithCallback invokes the cs.DescribeKubernetesVersionMetadata API asynchronously
func (client *Client) DescribeKubernetesVersionMetadataWithCallback(request *DescribeKubernetesVersionMetadataRequest, callback func(response *DescribeKubernetesVersionMetadataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeKubernetesVersionMetadataResponse
		var err error
		defer close(result)
		response, err = client.DescribeKubernetesVersionMetadata(request)
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

// DescribeKubernetesVersionMetadataRequest is the request struct for api DescribeKubernetesVersionMetadata
type DescribeKubernetesVersionMetadataRequest struct {
	*requests.RoaRequest
	ClusterType       string `position:"Query" name:"ClusterType"`
	KubernetesVersion string `position:"Query" name:"KubernetesVersion"`
	Profile           string `position:"Query" name:"Profile"`
	Region            string `position:"Query" name:"Region"`
}

// DescribeKubernetesVersionMetadataResponse is the response struct for api DescribeKubernetesVersionMetadata
type DescribeKubernetesVersionMetadataResponse struct {
	*responses.BaseResponse
}

// CreateDescribeKubernetesVersionMetadataRequest creates a request to invoke DescribeKubernetesVersionMetadata API
func CreateDescribeKubernetesVersionMetadataRequest() (request *DescribeKubernetesVersionMetadataRequest) {
	request = &DescribeKubernetesVersionMetadataRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeKubernetesVersionMetadata", "/api/v1/metadata/versions", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeKubernetesVersionMetadataResponse creates a response to parse from DescribeKubernetesVersionMetadata response
func CreateDescribeKubernetesVersionMetadataResponse() (response *DescribeKubernetesVersionMetadataResponse) {
	response = &DescribeKubernetesVersionMetadataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}