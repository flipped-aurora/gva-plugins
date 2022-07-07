## GVA ElasticSearch基础操作功能插件
#### 开发者：fengyikang888

### 使用步骤

#### 1. 前往GVA主程序下的initialize/router.go 在Routers 方法最末尾按照你需要的及安全模式添加本插件
    例：
    1、本插件可以采用gva的配置文件 也可以直接写死内容作为配置 建议为gva添加配置文件结构 然后将配置传入
	PluginInit(PrivateGroup, es.CreateEsPlug(
		global.GVA_CONFIG.Elasticsearch.Host,
		global.GVA_CONFIG.Elasticsearch.Port,
		global.GVA_CONFIG.Elasticsearch.User,
		global.GVA_CONFIG.Elasticsearch.Pass,
	))

    2、同样也需要在GVA主程序下config/config.go下添加配置
    import es_config "github.com/flipped-aurora/gin-vue-admin/server/plugin/es/config"
    type Server struct {
    //ES
	Elasticsearch es_config.Elasticsearch `mapstructure:"elasticsearch" json:"elasticsearch" yaml:"elasticsearch"`
    }
    3、在GVA主程序下的core/server.go中添加配置    
    // 保证文本顺序输出
    // In order to ensure that the text order output can be deleted
	es_initialize.LoadService() // Elasticsearch连接服务地址
	

### 2. 配置说明

#### 2-1 全局配置结构体说明
    //其中 User 和 Pass 就是用户名和密码

    type Elasticsearch struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"` // 哪个数据库
	Port string `mapstructure:"port" json:"port" yaml:"port"` // 服务器地址:端口
	User string `mapstructure:"user" json:"user" yaml:"user"` // 用户名
	Pass string `mapstructure:"pass" json:"pass" yaml:"pass"` // 密码
    }
#### 2-2 入参结构说明
    //其中 Employee结构体自己对应需要存储的数据

    type Employee struct {
	Type           int      `json:"type"`
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	BrandName      string   `json:"brandName"`
	OneClass       int      `json:"oneClass"`
	TwoClass       int      `json:"twoClass"`
	ThreeClass     int      `json:"threeClass"`
	BrandId        int      `json:"brandId"`
	SealId         int      `json:"sealId"`
	SealName       string   `json:"sealName"`
	OneClassName   string   `json:"oneClassName"`
	TwoClassName   string   `json:"twoClassName"`
	ThreeClassName string   `json:"threeClassName"`
	Price          string   `json:"price"`
	Img            string   `json:"img"`
	IsShow         int      `json:"isShow"`
	About          string   `json:"about"`
	Interests      []string `json:"interests"`
    }


### 3. 方法API

    可以参考plugin/es/api/es.go里的写法


### 4. 可直接调用的接口

    新增接口： /es/esAdd [post] 已配置config.yaml
    删除接口： /es/delEs [delete] 
    修改接口： /es/updateEs [put] 
    根据ID查询接口： /es/getEsId [get] 
    全局查询接口： /es/getEsQuery [get] 
    模糊查询接口： /es/getEsMultipleQuery [get]
    更新经纬度接口： /es/updateEsLocation [put]
    附近的人查询接口： /es/getEsNearby [get]

## 鸣谢

> 以下排名不分先后

[gin-vue-admin 极速开发的全栈开发基础平台](https://www.gin-vue-admin.com/)

[Elasticsearch组件库](https://github.com/olivere/elastic)



   
