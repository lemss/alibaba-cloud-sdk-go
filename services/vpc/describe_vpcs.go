package vpc

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

func (client *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
	response = CreateDescribeVpcsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeVpcsWithChan(request *DescribeVpcsRequest) (<-chan *DescribeVpcsResponse, <-chan error) {
	responseChan := make(chan *DescribeVpcsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpcs(request)
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

func (client *Client) DescribeVpcsWithCallback(request *DescribeVpcsRequest, callback func(response *DescribeVpcsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpcsResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpcs(request)
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

type DescribeVpcsRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	IsDefault            requests.Boolean `position:"Query" name:"IsDefault"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	VpcName              string           `position:"Query" name:"VpcName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VpcId                string           `position:"Query" name:"VpcId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeVpcsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Vpcs       struct {
		Vpc []struct {
			VpcId        string `json:"VpcId" xml:"VpcId"`
			RegionId     string `json:"RegionId" xml:"RegionId"`
			Status       string `json:"Status" xml:"Status"`
			VpcName      string `json:"VpcName" xml:"VpcName"`
			CreationTime string `json:"CreationTime" xml:"CreationTime"`
			CidrBlock    string `json:"CidrBlock" xml:"CidrBlock"`
			VRouterId    string `json:"VRouterId" xml:"VRouterId"`
			Description  string `json:"Description" xml:"Description"`
			IsDefault    bool   `json:"IsDefault" xml:"IsDefault"`
			VSwitchIds   struct {
				VSwitchId []string `json:"VSwitchId" xml:"VSwitchId"`
			} `json:"VSwitchIds" xml:"VSwitchIds"`
			UserCidrs struct {
				UserCidr []string `json:"UserCidr" xml:"UserCidr"`
			} `json:"UserCidrs" xml:"UserCidrs"`
			NatGatewayIds struct {
				NatGatewayIds []string `json:"NatGatewayIds" xml:"NatGatewayIds"`
			} `json:"NatGatewayIds" xml:"NatGatewayIds"`
			RouterTableIds struct {
				RouterTableIds []string `json:"RouterTableIds" xml:"RouterTableIds"`
			} `json:"RouterTableIds" xml:"RouterTableIds"`
		} `json:"Vpc" xml:"Vpc"`
	} `json:"Vpcs" xml:"Vpcs"`
}

func CreateDescribeVpcsRequest() (request *DescribeVpcsRequest) {
	request = &DescribeVpcsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpcs", "vpc", "openAPI")
	return
}

func CreateDescribeVpcsResponse() (response *DescribeVpcsResponse) {
	response = &DescribeVpcsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
