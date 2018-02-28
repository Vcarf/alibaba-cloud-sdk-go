package teslamaxcompute

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

func (client *Client) GetProjectInstance(request *GetProjectInstanceRequest) (response *GetProjectInstanceResponse, err error) {
	response = CreateGetProjectInstanceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetProjectInstanceWithChan(request *GetProjectInstanceRequest) (<-chan *GetProjectInstanceResponse, <-chan error) {
	responseChan := make(chan *GetProjectInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProjectInstance(request)
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

func (client *Client) GetProjectInstanceWithCallback(request *GetProjectInstanceRequest, callback func(response *GetProjectInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProjectInstanceResponse
		var err error
		defer close(result)
		response, err = client.GetProjectInstance(request)
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

type GetProjectInstanceRequest struct {
	*requests.RpcRequest
	Project  string           `position:"Query" name:"Project"`
	PageNum  requests.Integer `position:"Query" name:"PageNum"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	Status   string           `position:"Query" name:"Status"`
	Region   string           `position:"Query" name:"Region"`
}

type GetProjectInstanceResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

func CreateGetProjectInstanceRequest() (request *GetProjectInstanceRequest) {
	request = &GetProjectInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetProjectInstance", "", "")
	return
}

func CreateGetProjectInstanceResponse() (response *GetProjectInstanceResponse) {
	response = &GetProjectInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
