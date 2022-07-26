package slb

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

// DescribeRuleAttribute invokes the slb.DescribeRuleAttribute API synchronously
func (client *Client) DescribeRuleAttribute(request *DescribeRuleAttributeRequest) (response *DescribeRuleAttributeResponse, err error) {
	response = CreateDescribeRuleAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRuleAttributeWithChan invokes the slb.DescribeRuleAttribute API asynchronously
func (client *Client) DescribeRuleAttributeWithChan(request *DescribeRuleAttributeRequest) (<-chan *DescribeRuleAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeRuleAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRuleAttribute(request)
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

// DescribeRuleAttributeWithCallback invokes the slb.DescribeRuleAttribute API asynchronously
func (client *Client) DescribeRuleAttributeWithCallback(request *DescribeRuleAttributeRequest, callback func(response *DescribeRuleAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRuleAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeRuleAttribute(request)
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

// DescribeRuleAttributeRequest is the request struct for api DescribeRuleAttribute
type DescribeRuleAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	RuleId               string           `position:"Query" name:"RuleId"`
}

// DescribeRuleAttributeResponse is the response struct for api DescribeRuleAttribute
type DescribeRuleAttributeResponse struct {
	*responses.BaseResponse
	VServerGroupId         string `json:"VServerGroupId" xml:"VServerGroupId"`
	Cookie                 string `json:"Cookie" xml:"Cookie"`
	LoadBalancerId         string `json:"LoadBalancerId" xml:"LoadBalancerId"`
	RuleId                 string `json:"RuleId" xml:"RuleId"`
	ServiceManagedMode     string `json:"ServiceManagedMode" xml:"ServiceManagedMode"`
	RequestId              string `json:"RequestId" xml:"RequestId"`
	HealthCheckConnectPort int    `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	HealthCheckTimeout     int    `json:"HealthCheckTimeout" xml:"HealthCheckTimeout"`
	CookieTimeout          int    `json:"CookieTimeout" xml:"CookieTimeout"`
	HealthCheckDomain      string `json:"HealthCheckDomain" xml:"HealthCheckDomain"`
	UnhealthyThreshold     int    `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	HealthCheckHttpCode    string `json:"HealthCheckHttpCode" xml:"HealthCheckHttpCode"`
	Domain                 string `json:"Domain" xml:"Domain"`
	ListenerPort           string `json:"ListenerPort" xml:"ListenerPort"`
	Url                    string `json:"Url" xml:"Url"`
	HealthCheckInterval    int    `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	HealthCheckURI         string `json:"HealthCheckURI" xml:"HealthCheckURI"`
	RuleName               string `json:"RuleName" xml:"RuleName"`
	StickySessionType      string `json:"StickySessionType" xml:"StickySessionType"`
	Scheduler              string `json:"Scheduler" xml:"Scheduler"`
	ListenerSync           string `json:"ListenerSync" xml:"ListenerSync"`
	HealthyThreshold       int    `json:"HealthyThreshold" xml:"HealthyThreshold"`
	StickySession          string `json:"StickySession" xml:"StickySession"`
	HealthCheck            string `json:"HealthCheck" xml:"HealthCheck"`
}

// CreateDescribeRuleAttributeRequest creates a request to invoke DescribeRuleAttribute API
func CreateDescribeRuleAttributeRequest() (request *DescribeRuleAttributeRequest) {
	request = &DescribeRuleAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeRuleAttribute", "Slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRuleAttributeResponse creates a response to parse from DescribeRuleAttribute response
func CreateDescribeRuleAttributeResponse() (response *DescribeRuleAttributeResponse) {
	response = &DescribeRuleAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
