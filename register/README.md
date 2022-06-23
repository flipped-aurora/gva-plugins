## GVA 普通用户注册插件

#### 开发者：XRSec

本插件用于普通用户注册

### 1. 使用场景

- 开放用户注册

### 2. 配置说明

下载插件解压至`server/plugin/register`

在`server/initialize/plugin.go` 文件中注册插件

```go
func InstallPlugin(PublicGroup *gin.RouterGroup, PrivateGroup *gin.RouterGroup) {
    // 888 为普通用户ID
    PluginInit(PublicGroup, register.CreateRegisterPlug("888"))
}
```

### 3 参数说明

#### 3-1 全局参数说明

> 请参考 system_user Register

#### 3-2 请求入参说明

> 请参考 system_user Register

#### 3.3 vue示例

```vue
<!--login/index.vue-->
<el-form-item>
<el-button
    type="primary"
    style="width: 38%"
    size="large"
    @click="checkInit"
>前往初始化
</el-button>
<el-button
    type="primary"
    size="large"
    style="width: 38%; margin-left: 8%"
    @click="submitForm"
>
  <div v-if="loginType.status">注 册</div>
  <div v-if="!loginType.status">登 录</div>
</el-button>
<el-switch
    v-model="loginType.status"
    style="width: 13%; margin-left: 3%"/>
</el-form-item>
...

<script setup>
const loginFormData = reactive({
  username: 'admin',
  password: '123456',
  captcha: '',
  captchaId: '',
})
const loginType = reactive({
  status: false,
})
...
const userStore = useUserStore()
const login = async() => {
  return await userStore.LoginIn(loginFormData)
}
const register = async() => {
  return await userStore.Register(loginFormData)
}
const submitForm = () => {
  loginForm.value.validate(async(v) => {
    if (v) {
      let flag
      if (loginType.status) {
        flag = await register()
      } else {
        flag = await login()
      }
      if (!flag) {
        loginVerify()
        loginType.status = false
      }
    } else {
      ElMessage({
        type: 'error',
        message: '请正确填写登录信息',
        showClose: true,
      })
      loginVerify()
      return false
    }
  })
}
</script>
```

```js
  // web/src/pinia/modules/user.js
  // 注册
import { login, getUserInfo, setSelfInfo, userRegister } from '@/api/user'
...
const Register = async(loginInfo) => {
    loadingInstance.value = ElLoading.service({
        fullscreen: true,
        text: '注册中，请稍候...',
    })
    try {
        const res = await userRegister(loginInfo)
        if (res.code === 0) {
            setUserInfo(res.data.user)
            setToken(res.data.token)
            const routerStore = useRouterStore()
            await routerStore.SetAsyncRouter()
            const asyncRouters = routerStore.asyncRouters
            asyncRouters.forEach((asyncRouter) => {
                router.addRoute(asyncRouter)
            })
            router.push({ name: userInfo.value.authority.defaultRouter })
            return true
        }
    } catch (e) {
        loadingInstance.value.close()
    }
    loadingInstance.value.close()
}
...

return {
    Register,
}
```

```js
// web/src/api/user.js
// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /register [post]
export const userRegister = (data) => {
    return service({
        url: '/register',
        method: 'post',
        data: data,
    })
}
```