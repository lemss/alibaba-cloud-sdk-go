package domain_intl

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

func (client *Client) SaveBatchTaskForCreatingOrderRedeem(request *SaveBatchTaskForCreatingOrderRedeemRequest) (response *SaveBatchTaskForCreatingOrderRedeemResponse, err error) {
	response = CreateSaveBatchTaskForCreatingOrderRedeemResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SaveBatchTaskForCreatingOrderRedeemWithChan(request *SaveBatchTaskForCreatingOrderRedeemRequest) (<-chan *SaveBatchTaskForCreatingOrderRedeemResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForCreatingOrderRedeemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForCreatingOrderRedeem(request)
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

func (client *Client) SaveBatchTaskForCreatingOrderRedeemWithCallback(request *SaveBatchTaskForCreatingOrderRedeemRequest, callback func(response *SaveBatchTaskForCreatingOrderRedeemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForCreatingOrderRedeemResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForCreatingOrderRedeem(request)
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

type SaveBatchTaskForCreatingOrderRedeemRequest struct {
	*requests.RpcRequest
	OrderRedeemParam *[]SaveBatchTaskForCreatingOrderRedeemOrderRedeemParam `position:"Query" name:"OrderRedeemParam"  type:"Repeated"`
	UserClientIp     string                                                 `position:"Query" name:"UserClientIp"`
	Lang             string                                                 `position:"Query" name:"Lang"`
}

type SaveBatchTaskForCreatingOrderRedeemOrderRedeemParam struct {
	DomainName            string `name:"DomainName"`
	CurrentExpirationDate string `name:"CurrentExpirationDate"`
}

type SaveBatchTaskForCreatingOrderRedeemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

func CreateSaveBatchTaskForCreatingOrderRedeemRequest() (request *SaveBatchTaskForCreatingOrderRedeemRequest) {
	request = &SaveBatchTaskForCreatingOrderRedeemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForCreatingOrderRedeem", "", "")
	return
}

func CreateSaveBatchTaskForCreatingOrderRedeemResponse() (response *SaveBatchTaskForCreatingOrderRedeemResponse) {
	response = &SaveBatchTaskForCreatingOrderRedeemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
