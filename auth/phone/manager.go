package phone

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

type Config struct {
	SecretID	string	`json:"secret_id"`
	SecretKey	string	`json:"secret_key"`
	TemplateID	string	`json:"template_id"`
	SmsSdkAppid	string	`json:"sms_sdk_appid"`
	Sign		string	`json:"sign"`
}

type Manager struct {
	config	*Config
	handler	*sms.Client
}

func New(config *Config) (*Manager, error) {
	credential := common.NewCredential (
		config.SecretID,
		config.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, err := sms.NewClient(credential, "", cpf)
	if err != nil {
		return nil, err
	}
	return &Manager{
		config: config,
		handler: client,
	}, nil
}

func (m *Manager) SendCode(phone string) (string, error) {
	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(
		[]string{
			fmt.Sprintf("+86%s",phone),
		},
	)

	code := fmt.Sprintf("%04d", rand.Intn(10000))
	request.TemplateParamSet = common.StringPtrs(
		[]string{
			code,
		},
	)
	request.TemplateID = common.StringPtr(m.config.TemplateID)
	request.SmsSdkAppid = common.StringPtr(m.config.SmsSdkAppid)
	request.Sign = common.StringPtr("康养网")

	_, err := m.handler.SendSms(request)
	return code, err
}
