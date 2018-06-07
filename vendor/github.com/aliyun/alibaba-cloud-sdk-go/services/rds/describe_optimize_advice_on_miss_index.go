package rds

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

// DescribeOptimizeAdviceOnMissIndex invokes the rds.DescribeOptimizeAdviceOnMissIndex API synchronously
// api document: https://help.aliyun.com/api/rds/describeoptimizeadviceonmissindex.html
func (client *Client) DescribeOptimizeAdviceOnMissIndex(request *DescribeOptimizeAdviceOnMissIndexRequest) (response *DescribeOptimizeAdviceOnMissIndexResponse, err error) {
	response = CreateDescribeOptimizeAdviceOnMissIndexResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOptimizeAdviceOnMissIndexWithChan invokes the rds.DescribeOptimizeAdviceOnMissIndex API asynchronously
// api document: https://help.aliyun.com/api/rds/describeoptimizeadviceonmissindex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOptimizeAdviceOnMissIndexWithChan(request *DescribeOptimizeAdviceOnMissIndexRequest) (<-chan *DescribeOptimizeAdviceOnMissIndexResponse, <-chan error) {
	responseChan := make(chan *DescribeOptimizeAdviceOnMissIndexResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOptimizeAdviceOnMissIndex(request)
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

// DescribeOptimizeAdviceOnMissIndexWithCallback invokes the rds.DescribeOptimizeAdviceOnMissIndex API asynchronously
// api document: https://help.aliyun.com/api/rds/describeoptimizeadviceonmissindex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOptimizeAdviceOnMissIndexWithCallback(request *DescribeOptimizeAdviceOnMissIndexRequest, callback func(response *DescribeOptimizeAdviceOnMissIndexResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOptimizeAdviceOnMissIndexResponse
		var err error
		defer close(result)
		response, err = client.DescribeOptimizeAdviceOnMissIndex(request)
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

// DescribeOptimizeAdviceOnMissIndexRequest is the request struct for api DescribeOptimizeAdviceOnMissIndex
type DescribeOptimizeAdviceOnMissIndexRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// DescribeOptimizeAdviceOnMissIndexResponse is the response struct for api DescribeOptimizeAdviceOnMissIndex
type DescribeOptimizeAdviceOnMissIndexResponse struct {
	*responses.BaseResponse
	RequestId         string                                   `json:"RequestId" xml:"RequestId"`
	DBInstanceId      string                                   `json:"DBInstanceId" xml:"DBInstanceId"`
	TotalRecordsCount int                                      `json:"TotalRecordsCount" xml:"TotalRecordsCount"`
	PageNumber        int                                      `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount   int                                      `json:"PageRecordCount" xml:"PageRecordCount"`
	Items             ItemsInDescribeOptimizeAdviceOnMissIndex `json:"Items" xml:"Items"`
}

// CreateDescribeOptimizeAdviceOnMissIndexRequest creates a request to invoke DescribeOptimizeAdviceOnMissIndex API
func CreateDescribeOptimizeAdviceOnMissIndexRequest() (request *DescribeOptimizeAdviceOnMissIndexRequest) {
	request = &DescribeOptimizeAdviceOnMissIndexRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceOnMissIndex", "rds", "openAPI")
	return
}

// CreateDescribeOptimizeAdviceOnMissIndexResponse creates a response to parse from DescribeOptimizeAdviceOnMissIndex response
func CreateDescribeOptimizeAdviceOnMissIndexResponse() (response *DescribeOptimizeAdviceOnMissIndexResponse) {
	response = &DescribeOptimizeAdviceOnMissIndexResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}