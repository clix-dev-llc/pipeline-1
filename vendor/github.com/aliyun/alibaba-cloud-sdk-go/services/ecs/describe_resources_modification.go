package ecs

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

// DescribeResourcesModification invokes the ecs.DescribeResourcesModification API synchronously
// api document: https://help.aliyun.com/api/ecs/describeresourcesmodification.html
func (client *Client) DescribeResourcesModification(request *DescribeResourcesModificationRequest) (response *DescribeResourcesModificationResponse, err error) {
	response = CreateDescribeResourcesModificationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourcesModificationWithChan invokes the ecs.DescribeResourcesModification API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeresourcesmodification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourcesModificationWithChan(request *DescribeResourcesModificationRequest) (<-chan *DescribeResourcesModificationResponse, <-chan error) {
	responseChan := make(chan *DescribeResourcesModificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourcesModification(request)
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

// DescribeResourcesModificationWithCallback invokes the ecs.DescribeResourcesModification API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeresourcesmodification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourcesModificationWithCallback(request *DescribeResourcesModificationRequest, callback func(response *DescribeResourcesModificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourcesModificationResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourcesModification(request)
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

// DescribeResourcesModificationRequest is the request struct for api DescribeResourcesModification
type DescribeResourcesModificationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Memory               requests.Integer `position:"Query" name:"Memory"`
	Cores                requests.Integer `position:"Query" name:"Cores"`
	MigrateAcrossZone    requests.Boolean `position:"Query" name:"MigrateAcrossZone"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	ResourceId           string           `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OperationType        string           `position:"Query" name:"OperationType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DestinationResource  string           `position:"Query" name:"DestinationResource"`
}

// DescribeResourcesModificationResponse is the response struct for api DescribeResourcesModification
type DescribeResourcesModificationResponse struct {
	*responses.BaseResponse
	RequestId      string                                        `json:"RequestId" xml:"RequestId"`
	AvailableZones AvailableZonesInDescribeResourcesModification `json:"AvailableZones" xml:"AvailableZones"`
}

// CreateDescribeResourcesModificationRequest creates a request to invoke DescribeResourcesModification API
func CreateDescribeResourcesModificationRequest() (request *DescribeResourcesModificationRequest) {
	request = &DescribeResourcesModificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeResourcesModification", "ecs", "openAPI")
	return
}

// CreateDescribeResourcesModificationResponse creates a response to parse from DescribeResourcesModification response
func CreateDescribeResourcesModificationResponse() (response *DescribeResourcesModificationResponse) {
	response = &DescribeResourcesModificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
