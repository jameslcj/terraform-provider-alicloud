package emr

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

// MetastoreDescribeKafkaConsumerGroup invokes the emr.MetastoreDescribeKafkaConsumerGroup API synchronously
// api document: https://help.aliyun.com/api/emr/metastoredescribekafkaconsumergroup.html
func (client *Client) MetastoreDescribeKafkaConsumerGroup(request *MetastoreDescribeKafkaConsumerGroupRequest) (response *MetastoreDescribeKafkaConsumerGroupResponse, err error) {
	response = CreateMetastoreDescribeKafkaConsumerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// MetastoreDescribeKafkaConsumerGroupWithChan invokes the emr.MetastoreDescribeKafkaConsumerGroup API asynchronously
// api document: https://help.aliyun.com/api/emr/metastoredescribekafkaconsumergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreDescribeKafkaConsumerGroupWithChan(request *MetastoreDescribeKafkaConsumerGroupRequest) (<-chan *MetastoreDescribeKafkaConsumerGroupResponse, <-chan error) {
	responseChan := make(chan *MetastoreDescribeKafkaConsumerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MetastoreDescribeKafkaConsumerGroup(request)
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

// MetastoreDescribeKafkaConsumerGroupWithCallback invokes the emr.MetastoreDescribeKafkaConsumerGroup API asynchronously
// api document: https://help.aliyun.com/api/emr/metastoredescribekafkaconsumergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreDescribeKafkaConsumerGroupWithCallback(request *MetastoreDescribeKafkaConsumerGroupRequest, callback func(response *MetastoreDescribeKafkaConsumerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MetastoreDescribeKafkaConsumerGroupResponse
		var err error
		defer close(result)
		response, err = client.MetastoreDescribeKafkaConsumerGroup(request)
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

// MetastoreDescribeKafkaConsumerGroupRequest is the request struct for api MetastoreDescribeKafkaConsumerGroup
type MetastoreDescribeKafkaConsumerGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TopicId         string           `position:"Query" name:"TopicId"`
	ConsumerGroupId string           `position:"Query" name:"ConsumerGroupId"`
}

// MetastoreDescribeKafkaConsumerGroupResponse is the response struct for api MetastoreDescribeKafkaConsumerGroup
type MetastoreDescribeKafkaConsumerGroupResponse struct {
	*responses.BaseResponse
	RequestId           string                                             `json:"RequestId" xml:"RequestId"`
	Id                  string                                             `json:"Id" xml:"Id"`
	DataSourceId        string                                             `json:"DataSourceId" xml:"DataSourceId"`
	ConsumerGroup       string                                             `json:"ConsumerGroup" xml:"ConsumerGroup"`
	ConsumerType        string                                             `json:"ConsumerType" xml:"ConsumerType"`
	TotalLag            int64                                              `json:"TotalLag" xml:"TotalLag"`
	PartitionProportion string                                             `json:"PartitionProportion" xml:"PartitionProportion"`
	GmtCreate           int64                                              `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified         int64                                              `json:"GmtModified" xml:"GmtModified"`
	PartitionList       PartitionListInMetastoreDescribeKafkaConsumerGroup `json:"PartitionList" xml:"PartitionList"`
}

// CreateMetastoreDescribeKafkaConsumerGroupRequest creates a request to invoke MetastoreDescribeKafkaConsumerGroup API
func CreateMetastoreDescribeKafkaConsumerGroupRequest() (request *MetastoreDescribeKafkaConsumerGroupRequest) {
	request = &MetastoreDescribeKafkaConsumerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "MetastoreDescribeKafkaConsumerGroup", "emr", "openAPI")
	return
}

// CreateMetastoreDescribeKafkaConsumerGroupResponse creates a response to parse from MetastoreDescribeKafkaConsumerGroup response
func CreateMetastoreDescribeKafkaConsumerGroupResponse() (response *MetastoreDescribeKafkaConsumerGroupResponse) {
	response = &MetastoreDescribeKafkaConsumerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}