package dnspod_tencentcloud_sdk_api_ddnspro

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

func GetDNSPODRecordList_TencloudSDK(sercetId, secretKey, sub_domain string) {
	credential := common.NewCredential(
		sercetId,
		secretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := dnspod.NewClient(credential, "", cpf)

	request := dnspod.NewDescribeRecordListRequest()
	request.Domain = common.StringPtr(sub_domain)
	response, err := client.DescribeRecordList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(response.ToJsonString())
	var tempMap map[string]interface{}
	err = json.Unmarshal([]byte(response.ToJsonString()), &tempMap)
	if err != nil {
		panic(err)
	}
	println(tempMap)
}
