## GVA 腾讯云短信发送功能插件

#### 开发者：shubo6

### 使用步骤

#### 1. 前往GVA主程序下的initialize/router.go 在Routers 方法最末尾按照你需要的及安全模式添加本插件

    例：
    本插件可以采用gva的配置文件 也可以直接写死内容作为配置 建议为gva添加配置文件结构 然后将配置传入
		PluginInit(PublicGroup, sms.CreateTencentSmsPlug("短信的SecretId", "短信的SecretKey", "短信的 SdkAppId", "短信的 SignName"))

### 2. 配置说明

#### 2-1 全局配置结构体说明

```go

    type TencentSMS struct {
        SecretId  string `mapstructure:"secret-id" json:"secret-id" yaml:"secret-id"`// 短信的SecretId
        SecretKey string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`// 短信的SecretKey
        SdkAppId  string `mapstructure:"sdk-app-id" json:"sdk-app-id" yaml:"sdk-app-id"`//短信所属的APPID
        SignName  string `mapstructure:"sign-name" json:"sign-name" yaml:"sign-name"`// 短信的 SignName
    }
	
```



#### 2-2 入参结构说明
```go
    type TencentModel struct {
        Phones        []string `json:"phones"` // 手机号
        TemplateId    string   `json:"templateId"` //模板ID,需要在腾讯云控制台提交审核
        TemplateParam []string `json:"templateParam"` //模板参数,按顺序的模板参数
    }

```



### 3. 方法API

// 无

### 4. 可直接调用的接口

    发送邮件接口接口： /tencentSms/sendSms [post] 已配置swagger
    入参：
    type TencentModel struct {
        Phones        []string `json:"phones"` // 手机号
        TemplateId    string   `json:"templateId"` //模板ID,需要在腾讯云控制台提交审核
        TemplateParam []string `json:"templateParam"` //模板参数,按顺序的模板参数
    }
   
