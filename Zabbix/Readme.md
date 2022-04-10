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
### 3 参数说明

#### 3-1 全局参数说明

```go
	Url      string `mapstructure:"url" json:"url" yaml:"url"`  //ZabbixAPI接口地址
	Username string `mapstructure:"username" json:"username" yaml:"username"` //Zabbix用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` //Zabbix密码
```
#### 3-2 请求入参说明
```go

```

### 3方法API（可调用方法）
```go

//获取所有主机列表
GetHostAll()


```

### 4. 可直接调用接口

    获取主机列表接口： /zabbix/getHostAll [get] 已配置swagger注释

### 添加api SQL语句

```sql
    INSERT INTO `你的数据库名字`.`sys_apis`(`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (100, '2022-03-28 14:04:19.194', '2022-03-28 14:04:19.194', NULL, '/zabbix/getHostAll', 'Zabbix主机列表', 'Zabbix', 'GET');

```
