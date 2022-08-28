package dnspod_tenSDK

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"log"
)

func GetDNSPODRecordList_TencloudSDK(sercetId, secretKey, domain string, subdomain string) dnspod.DescribeRecordListResponse {
	credential := common.NewCredential(
		sercetId,
		secretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := dnspod.NewClient(credential, "", cpf)

	request := dnspod.NewDescribeRecordListRequest()
	request.Domain = common.StringPtr(domain)
	request.Subdomain = common.StringPtr(subdomain)
	response, err := client.DescribeRecordList(request)
	if err != nil {
		log.Fatal(err.Error())
	}
	return *response
}
