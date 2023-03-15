// This file is auto-generated, don't edit it. Thanks.
// aliyun doc https://help.aliyun.com/document_detail/437673.html
// aliyun debug https://next.api.aliyun.com/api/alikafka/2019-09-16/CreateTopic?lang=GO
package admin

import (
	"os"

	alikafka20190916 "github.com/alibabacloud-go/alikafka-20190916/v3/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *alikafka20190916.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("alikafka.cn-shanghai.aliyuncs.com")
	_result = &alikafka20190916.Client{}
	_result, _err = alikafka20190916.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if _err != nil {
		return _err
	}

	createTopicRequest := &alikafka20190916.CreateTopicRequest{
		InstanceId: tea.String("my-instanceId"),
		Topic:      tea.String("my-topic"),
		Remark:     tea.String("my-remark"),
		RegionId:   tea.String("my-regionId"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.CreateTopicWithOptions(createTopicRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
