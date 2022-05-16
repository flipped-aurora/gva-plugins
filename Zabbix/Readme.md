## GVA Zabbix插件

本插件用于集成Zabbix

### 1. 使用场景

- 当需要集成Zabbix监控时使用

### 2. 配置说明

在`plugin/notify/global/global.go` 文件中配置启用Zabbix插件

```go
//  在gin-vue-admin 主程序的initialize中的plugin的InstallPlugin 函数中写入如下代码
   PluginInit(PublicGroup, Zabbix.CreateZabbixPlugin(
		global.GVA_CONFIG.Zabbix.Url,
		global.GVA_CONFIG.Zabbix.Username,
		global.GVA_CONFIG.Zabbix.Password,
	))
}
```

在config.yaml文件中配置Zabbix插件登录信息

	zabbix:
	  url: 'http://localhost/zabbix/api_jsonrpc.php'
	  username: Admin
	  password: Zabbix
在全局config目录下增加Zabbix.go及修改config.go文件

	Zabbix.go
	type Zabbix struct {
		Url      string `mapstructure:"url" json:"url" yaml:"url"`
		Username string `mapstructure:"username" json:"username" yaml:"username"`
		Password string `mapstructure:"password" json:"password" yaml:"password"`
	}
	config.go
	Zabbix  Zabbix  `mapstructure:"zabbix" json:"zabbix" yaml:"zabbix"` //增加在结构体后面

### 3 参数说明

#### 3-1 全局参数说明

```go
	Url      string `mapstructure:"url" json:"url" yaml:"url"`  //ZabbixAPI接口地址
	Username string `mapstructure:"username" json:"username" yaml:"username"` //Zabbix用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` //Zabbix密码
```
#### 3-2 请求入参说明
```go
/zabbix/getHostAll //请求为GET无需设置参数
/zabbix/getTiggerList //请求为GET无需设置参数
/zabbix/getwebcode //请求为POST入参如下
{
    "url":"http://www.baidu.com,http://google.com" //支持批量查询,以英文逗号分割
}
```

### 3方法API（可调用方法）
```go
GetHost() //获取所有主机列表
GetAlertTiggerList() //获取所有正在告警状态的触发器
WebMonitor(urls string) //获取url参数中提供的网址StatusCode状态是否为200 OK
```

### 4. 可直接调用接口

    //获取所有主机列表
    /zabbix/getHostAll   [get] 已配置swagger注释
    //获取所有正在告警状态的触发器 status:0为触发器启用状态,status:1为触发器禁用状态(可忽略的触发器告警)
    /zabbix/getTiggerList  [get] 已配置swagger注释
    //获取url参数中提供的网址StatusCode状态是否为200 OK
    /zabbix/getwebcode  [post] 已配置swagger注释
    

### 添加api SQL语句

```sql
    INSERT INTO `你的数据库名字`.`sys_apis`(`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (104, '2022-04-12 10:33:05.000', '2022-04-12 10:33:05.000', NULL, '/zabbix/getHostAll', 'Zabbix主机列表', 'Zabbix', 'GET');
INSERT INTO `你的数据库名字`.`sys_apis`(`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (106, '2022-04-12 15:03:31.302', '2022-04-12 15:03:31.302', NULL, '/zabbix/getwebcode', '监控网页状态组件', 'Zabbix', 'POST');
INSERT INTO `你的数据库名字`.`sys_apis`(`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (108, '2022-04-13 16:23:07.714', '2022-04-13 16:23:07.714', NULL, '/zabbix/getTiggerList', '获取正在触发的告警列表', 'Zabbix', 'GET');
```
