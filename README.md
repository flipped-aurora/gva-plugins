## GIN-VUE-ADMIN插件开发规范

### 请勿提交任何具有危险性的代码

### 详细介绍

插件分类（pr时请标注）
- 前端功能类插件
- 前端美化类插件
- 后端工具性插件
- 后端中间件
- 业务性插件 (完整的后端业务链条，配备前端vue静态文件页面)
- 其他

### 开发规范

插件需按照gin-vue-admin基本目录结构书写 插件前端部分如果需要则在views、api、utils等需要到对应功能的目录下分别创建同插件名的文件例如  必须配备readme 使用说明

```
  web
    --src
      --api
        --plugs.js
      --view
        --plugsTable.vue
        --plugsForm.vue
      --utils
        --plugs.js
```

后端开发需要在plugin下需创建等同于gin-vue-admin后端目录的结构并安放你开发的插件代码  必须配备readme 使用说明

```
  server
    --plugin
      --你的插件名字
        --api
        --你需要的功能包 对应gva的目录结构和功能结构即可
```

```
后端提供的api需要按照规范填写 swagger注解
示例：
  // @Tags System
  // @Summary 发送测试邮件
  // @Security ApiKeyAuth
  // @Produce  application/json
  // @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
  // @Router /email/emailTest [post]
  func (s *EmailApi) EmailTest(c *gin.Context) {
    if err := service.ServiceGroupApp.EmailTest(); err != nil {
      global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
      response.FailWithMessage("发送失败", c)
    } else {
      response.OkWithData("发送成功", c)
    }
  }
```
所有可直接调用的方法需要提供详细出入参说明

  //@author: [maplepie](https://github.com/maplepie)
  //@function: EmailTest
  //@description: 发送邮件测试
  //@return: err error
  //@params to string 	 收件人
  //@params subject string   标题（主题）
  //@params body  string 	 邮件内容

  func (e *EmailService) SendEmail(to, subject, body string) (err error) {
    err = utils.Email(to, subject, body)
    return err
  }

```
后端提供的api需要按照规范填写 swagger注解
示例：
  // @Tags System
  // @Summary 发送测试邮件
  // @Security ApiKeyAuth
  // @Produce  application/json
  // @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
  // @Router /email/emailTest [post]
  func (s *EmailApi) EmailTest(c *gin.Context) {
    if err := service.ServiceGroupApp.EmailTest(); err != nil {
      global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
      response.FailWithMessage("发送失败", c)
    } else {
      response.OkWithData("发送成功", c)
    }
  }
```

### 插件提供的readme中需要对自身开发的工具进行最详细的描述
#### 示例

```
  ## GVA 邮件发送功能插件
#### 开发者：GIN-VUE-ADMIN 官方

### 使用步骤

#### 1. 前往GVA主程序下的initialize/router.go 在Routers 方法最末尾按照你需要的及安全模式添加本插件
    例：
    本插件可以采用gva的配置文件 也可以直接写死内容作为配置 建议为gva添加配置文件结构 然后将配置传入
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
		))

    同样也可以再传入时写死

    PluginInit(PrivateGroup, email.CreateEmailPlug(
    "a@qq.com",
    "b@qq.com",
    "smtp.qq.com",
    "global.GVA_CONFIG.Email.Secret",
    "登录密钥",
    465,
    true,
    ))

### 2. 配置说明

#### 2-1 全局配置结构体说明
    //其中 Form 和 Secret 通常来说就是用户名和密码

    type Email struct {
	    To       string  // 收件人:多个以英文逗号分隔 例：a@qq.com b@qq.com 正式开发中请把此项目作为参数使用 此处配置主要用于发送错误监控邮件
	    From     string  // 发件人  你自己要发邮件的邮箱
	    Host     string  // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
	    Secret   string  // 密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥
	    Nickname string  // 昵称    发件人昵称 自定义即可 可以不填
	    Port     int     // 端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465
	    IsSSL    bool    // 是否SSL   是否开启SSL
    }
#### 2-2 入参结构说明
    //其中 Form 和 Secret 通常来说就是用户名和密码

    type Email struct {
        To      string `json:"to"`      // 邮件发送给谁
        Subject string `json:"subject"` // 邮件标题
        Body    string `json:"body"`    // 邮件内容
    }


### 3. 方法API

    utils.EmailTest(邮件标题，邮件主体) 发送测试邮件
    例:utils.EmailTest("测试邮件"，"测试邮件")
    utils.ErrorToEmail(邮件标题,邮件主体) 错误监控
    例:utils.ErrorToEmail("测试邮件"，"测试邮件")
    utils.Email(目标邮箱多个的话用逗号分隔，邮件标题，邮件主体) 发送测试邮件
    例:utils.Email(”a.qq.com,b.qq.com“,"测试邮件"，"测试邮件")

### 4. 可直接调用的接口

    测试接口： /email/emailTest [post] 已配置swagger

    发送邮件接口接口： /email/emailSend [post] 已配置swagger
    入参：
    type Email struct {
        To      string `json:"to"`      // 邮件发送给谁
        Subject string `json:"subject"` // 邮件标题
        Body    string `json:"body"`    // 邮件内容
    }
   
### api插入语句

    ```sql
        INSERT INTO `你的数据库名字`.`sys_apis`(`created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES ( '2021-08-25 23:09:12', '2021-08-25 23:09:12', NULL, '/email/emailTest', '发送测试邮件', 'email', 'POST');
        INSERT INTO `你的数据库名字`.`sys_apis`(`created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES ( '2021-08-28 14:20:27', '2021-08-28 14:20:27', NULL, '/email/sendEmail', '发送邮件', 'email', 'POST');
    ```

```


### 加入开发者需求

本仓库为gin-vue-admin插件仓库，插件请提交至gin-vue-admin项目的gva-pliugs分支，审核测试通过后，将有官方人员转移添加至此仓库，或自行调整为单独的包向此仓库pr。
