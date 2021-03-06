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

func (client *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
	response = CreateDeleteLoadBalancerResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteLoadBalancerWithChan(request *DeleteLoadBalancerRequest) (<-chan *DeleteLoadBalancerResponse, <-chan error) {
	responseChan := make(chan *DeleteLoadBalancerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLoadBalancer(request)
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

func (client *Client) DeleteLoadBalancerWithCallback(request *DeleteLoadBalancerRequest, callback func(response *DeleteLoadBalancerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLoadBalancerResponse
		var err error
		defer close(result)
		response, err = client.DeleteLoadBalancer(request)
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

type DeleteLoadBalancerRequest struct {
	*requests.RpcRequest
	Tags                 string           `position:"Query" name:"Tags"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DeleteLoadBalancerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteLoadBalancerRequest() (request *DeleteLoadBalancerRequest) {
	request = &DeleteLoadBalancerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DeleteLoadBalancer", "slb", "openAPI")
	return
}

func CreateDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
	response = &DeleteLoadBalancerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
